package source

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func AllVideos(db *sql.DB) (*[]Video, string) {
	var videos []Video

	rows, err := db.Query("Select * FROM videos")
	if err != nil {
		return nil, fmt.Sprintf("%v", err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var video Video
		var metadataJSON []byte
		var annotationsJSON []byte

		err := rows.Scan(&video.ID, &video.URL, &metadataJSON, &annotationsJSON)
		if err != nil {
			return nil, fmt.Sprintf("%v", err.Error())
		}

		json.Unmarshal(metadataJSON, &video.METADATA)
		json.Unmarshal(annotationsJSON, &video.ANNOTATIONS)

		videos = append(videos, video)
	}

	return &videos, ""
}

func Create(db *sql.DB, videoData Video) (*[]Video, string) {
	videoData.ID = uuid.NewString()
	videoData.METADATA.CREATEDAT = time.Now()
	videoData.METADATA.MODIFIEDAT = time.Now()

	// Adding UUID to each annotation from backend
	for _, annotation := range videoData.ANNOTATIONS {
		annotation.ID = uuid.NewString()
	}

	// To be able to insert the data into the table for the below structs
	metadataJSON, _ := json.Marshal(videoData.METADATA)
	annotationsJSON, _ := json.Marshal(videoData.ANNOTATIONS)

	if videoData.ANNOTATIONS[0].ENDTIME > videoData.METADATA.DURATION {
		errMessage := fmt.Sprintf("Your annotations end time %d is out of bounds of duration of the video %d",
			videoData.ANNOTATIONS[0].ENDTIME, videoData.METADATA.DURATION)
		return nil, errMessage
	}

	// Getting all the video data for validation
	finalVideoData, errorMessage := AllVideos(db)
	if errorMessage != "" {
		return nil, errorMessage
	}

	for _, data := range *finalVideoData {
		// Securing duplicate entries
		if videoData.URL == data.URL {
			return nil, fmt.Sprintf("Data already exists in the table for the URL: %s", data.URL)
		}

		_, err := db.Exec("INSERT INTO videos (video_id, video_url, metadata, categories) VALUES (?, ?, ?, ?)", videoData.ID, videoData.URL, string(metadataJSON), string(annotationsJSON))
		if err != nil {
			return nil, err.Error()
		}
	}

	return finalVideoData, ""
}

func AllAnnotations(db *sql.DB, videoURL string) ([]Annotation, string) {
	videoRecords, errorMessage := AllVideos(db)
	if errorMessage != "" {
		return nil, errorMessage
	}

	for _, videoRecord := range *videoRecords {
		if videoURL == videoRecord.URL {
			return videoRecord.ANNOTATIONS, ""
		}
	}

	return nil, fmt.Sprintf("No video exists with searched URL: '%s' to show the annotations details", videoURL)
}

// func AddAdditionalNotes(videoURL, annotationType, notes string) (*Video, string) {
// 	if notes != "" || len(notes) > 0 {
// 		for _, video := range videos {
// 			if videoURL == video.URL {
// 				for index, annotation := range video.ANNOTATIONS {
// 					if annotationType == annotation.TYPE {
// 						annotation.ADDITIONALNOTES = append(annotation.ADDITIONALNOTES, notes)

// 						video.ANNOTATIONS[index].ADDITIONALNOTES = annotation.ADDITIONALNOTES
// 						return &video, ""
// 					}
// 				}

// 				return nil, "No annotation type was found to add additional notes"
// 			}
// 		}

// 		return nil, "No video exists to show the annotations details"
// 	}

// 	return nil, "Empty notes were passed. Please add and try again"
// }

// func UpdateAnnotationDetails(videoURL, annotationType string, newAnnotation Annotation) (*Video, string) {
// 	for _, video := range videos {
// 		if videoURL == video.URL {
// 			for index, annotation := range video.ANNOTATIONS {
// 				if annotationType == annotation.TYPE {
// 					video.ANNOTATIONS[index] = newAnnotation
// 					return &video, ""
// 				}
// 			}

// 			return nil, "No annotation type was found to update the annotation"
// 		}
// 	}

// 	return nil, "No video exists to show the annotations details"
// }

// func DeleteAnnotationData(videoURL, annotationType string) (*Video, string) {
// 	for videoIndex, video := range videos {
// 		if videoURL == video.URL {
// 			for index, annotation := range video.ANNOTATIONS {
// 				if annotationType == annotation.TYPE {
// 					video.ANNOTATIONS = append(video.ANNOTATIONS[:index], video.ANNOTATIONS[index+1:]...)

// 					// To keep the updated element in the main data
// 					videos[videoIndex] = video
// 					return &video, ""
// 				}
// 			}

// 			return nil, "No annotation type was found to delete the annotation"
// 		}
// 	}

// 	return nil, "No video exists to show the annotations details"
// }

// func DeleteVideoData(videoURL string) (*[]Video, string) {
// 	for index, video := range videos {
// 		if videoURL == video.URL {
// 			videos = append(videos[:index], videos[index+1:]...)

// 			return &videos, ""
// 		}
// 	}

// 	return nil, "No video found to be deleted from the data"
// }
