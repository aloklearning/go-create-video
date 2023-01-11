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
	for index, _ := range videoData.ANNOTATIONS {
		videoData.ANNOTATIONS[index].ID = uuid.NewString()
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
	// finalVideoData, errorMessage := AllVideos(db)
	// if errorMessage != "" {
	// 	return nil, errorMessage
	// }

	// for _, data := range *finalVideoData {
	// 	// Securing duplicate entries
	// 	if videoData.URL == data.URL {
	// 		return nil, fmt.Sprintf("Data already exists in the table for the URL: %s", data.URL)
	// 	}

	// 	_, err := db.Exec("INSERT INTO videos (video_id, video_url, metadata, category) VALUES (?, ?, ?, ?)", videoData.ID, videoData.URL, string(metadataJSON), string(annotationsJSON))
	// 	if err != nil {
	// 		return nil, err.Error()
	// 	}
	// }

	_, err := db.Exec("INSERT INTO videos (video_id, video_url, metadata, annotations) VALUES (?, ?, ?, ?)", videoData.ID, videoData.URL, string(metadataJSON), string(annotationsJSON))
	if err != nil {
		return nil, err.Error()
	}

	finalVideoData, _ := AllVideos(db)
	return finalVideoData, ""
}

func AllAnnotations(db *sql.DB, videoURL string) ([]Annotation, string) {
	var video Video
	var metadataJSON []byte
	var annotationsJSON []byte

	err := db.QueryRow("SELECT video_id, video_url, metadata, annotations FROM videos WHERE video_url = ?", videoURL).
		Scan(&video.ID, &video.URL, &metadataJSON, &annotationsJSON)
	if err != nil {
		fmt.Print(err.Error())
		return nil, fmt.Sprintf("No such record found in the data for the URL: '%s'", videoURL)
	}

	json.Unmarshal(metadataJSON, &video.METADATA)
	json.Unmarshal(annotationsJSON, &video.ANNOTATIONS)

	return video.ANNOTATIONS, ""
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

func DeleteCompleteVideoData(db *sql.DB, videoURL string) (string, string) {
	// DELETE Meta Data
	_, metadataError := db.Exec("DELETE FROM metadata WHERE author in (SELECT metadata.author FROM metadata INNER JOIN videos ON metadata.author = videos.metadata WHERE video_url = ?)", videoURL)
	if metadataError != nil {
		return "", fmt.Sprintf("Metadata deletion error %s", metadataError.Error())
	}

	// DELETE Annotations
	_, annnotationError := db.Exec("DELETE FROM annotations WHERE type in (SELECT annotations.type FROM annotations INNER JOIN videos ON annotations.type = videos.annotations WHERE video_url = ?)", videoURL)
	if annnotationError != nil {
		return "", fmt.Sprintf("Annotation deletion error %s", annnotationError.Error())
	}

	// DELETE Videos
	_, videosError := db.Exec("DELETE FROM videos WHERE video_url = ?", videoURL)
	if videosError != nil {
		return "", fmt.Sprintf("Video deletion error %s", videosError.Error())
	}

	return "{Success: Video deleted successfully}", ""
}
