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

	// Preventing duplicate entry
	// Getting all the video data for validation
	previousVideoData, errorMessage := AllVideos(db)
	if errorMessage != "" {
		return nil, errorMessage
	}

	if len(*previousVideoData) > 0 {
		for _, video := range *previousVideoData {
			if videoData.URL == video.URL {
				return nil, fmt.Sprintf("Duplicate record for the URL '%s' found", videoData.URL)
			}
		}

		_, err := db.Exec("INSERT INTO videos (video_id, video_url, metadata, annotations) VALUES (?, ?, ?, ?)", videoData.ID, videoData.URL, string(metadataJSON), string(annotationsJSON))
		if err != nil {
			return nil, err.Error()
		}
	}

	// If there were not data present in the record. Simply add the item
	_, err := db.Exec("INSERT INTO videos (video_id, video_url, metadata, annotations) VALUES (?, ?, ?, ?)", videoData.ID, videoData.URL, string(metadataJSON), string(annotationsJSON))
	if err != nil {
		return nil, err.Error()
	}

	// Getting all the video data for validation
	finalVideoData, errorMessage := AllVideos(db)
	if errorMessage != "" {
		return nil, errorMessage
	}

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

func AddAdditionalNotes(db *sql.DB, videoURL, annotationType, notes string) (*Video, string) {
	var updatedVideoItem Video
	var metadataJSON []byte
	var annotationsJSON []byte

	if notes != "" || len(notes) > 0 {
		currentVideoList, err := AllVideos(db)
		if err != "" {
			return nil, err
		}

		for _, video := range *currentVideoList {
			if videoURL == video.URL {
				for index, annotation := range video.ANNOTATIONS {
					if annotationType == annotation.TYPE {
						annotation.ADDITIONALNOTES = append(annotation.ADDITIONALNOTES, notes)

						// Workaround for updating the record with PRIMARY KEY
						// Adding the whole Annotations again to the video item
						video.ANNOTATIONS[index].ADDITIONALNOTES = annotation.ADDITIONALNOTES
						video.METADATA.MODIFIEDAT = time.Now()

						jsonMetadata, _ := json.Marshal(video.METADATA)
						jsonAnnotations, _ := json.Marshal(video.ANNOTATIONS)

						_, updateError := db.Exec("UPDATE videos SET metadata = ?, annotations = ? WHERE video_url = ?", string(jsonMetadata), string(jsonAnnotations), videoURL)
						if updateError != nil {
							return nil, updateError.Error()
						}

						queryError := db.QueryRow("SELECT * FROM videos WHERE video_url = ?", videoURL).
							Scan(&updatedVideoItem.ID, &updatedVideoItem.URL, &metadataJSON, &annotationsJSON)
						if queryError != nil {
							return nil, queryError.Error()
						}

						// Adding the updated data to the current updated Video item
						json.Unmarshal(metadataJSON, &updatedVideoItem.METADATA)
						json.Unmarshal(annotationsJSON, &updatedVideoItem.ANNOTATIONS)

						return &updatedVideoItem, ""
					}
				}

				return nil, fmt.Sprintf("No annotation with type '%s' was found to add additional notes", annotationType)
			}
		}

		return nil, fmt.Sprintf("No video with the URL '%s' exists to show the annotations details", videoURL)
	}

	return nil, "Empty notes were passed. Please add and try again"
}

// Assuming we will recieve the Annotation Item from body always
func UpdateAnnotationDetails(db *sql.DB, videoURL, annotationType string, newAnnotation Annotation) (*Video, string) {
	var updatedVideoItem Video
	var metadataJSON []byte
	var annotationsJSON []byte

	currentVideoList, err := AllVideos(db)
	if err != "" {
		return nil, err
	}

	for _, video := range *currentVideoList {
		if videoURL == video.URL {
			for index, annotation := range video.ANNOTATIONS {
				if annotationType == annotation.TYPE {
					// Maintaing a copy of the ID
					annotationID := video.ANNOTATIONS[index].ID

					video.ANNOTATIONS[index] = newAnnotation
					video.ANNOTATIONS[index].ID = annotationID
					video.METADATA.MODIFIEDAT = time.Now() // modifications timing capture and updated

					jsonMetadata, _ := json.Marshal(video.METADATA)
					jsonAnnotations, _ := json.Marshal(video.ANNOTATIONS)

					_, updateError := db.Exec("UPDATE videos SET metadata = ?, annotations = ? WHERE video_url = ?", string(jsonMetadata), string(jsonAnnotations), videoURL)
					if updateError != nil {
						return nil, updateError.Error()
					}

					queryError := db.QueryRow("SELECT * FROM videos WHERE video_url = ?", videoURL).
						Scan(&updatedVideoItem.ID, &updatedVideoItem.URL, &metadataJSON, &annotationsJSON)
					if queryError != nil {
						return nil, queryError.Error()
					}

					// Adding the updated data to the current updated Video item
					json.Unmarshal(metadataJSON, &updatedVideoItem.METADATA)
					json.Unmarshal(annotationsJSON, &updatedVideoItem.ANNOTATIONS)

					return &updatedVideoItem, ""
				}
			}

			return nil, fmt.Sprintf("No annotation with type '%s' was found to update the annotation", annotationType)
		}
	}

	return nil, "No video exists to show the annotations details"
}

func DeleteAnnotationData(db *sql.DB, videoURL, annotationType string) (*Video, string) {
	var updatedVideoItem Video
	var metadataJSON []byte
	var annotationsJSON []byte

	currentVideoList, err := AllVideos(db)
	if err != "" {
		return nil, err
	}

	for _, video := range *currentVideoList {
		if videoURL == video.URL {
			for index, annotation := range video.ANNOTATIONS {
				if annotationType == annotation.TYPE {
					video.ANNOTATIONS = append(video.ANNOTATIONS[:index], video.ANNOTATIONS[index+1:]...)
					video.METADATA.MODIFIEDAT = time.Now() // modifications timing capture and updated

					jsonMetadata, _ := json.Marshal(video.METADATA)
					jsonAnnotations, _ := json.Marshal(video.ANNOTATIONS)

					_, updateError := db.Exec("UPDATE videos SET metadata = ?, annotations = ? WHERE video_url = ?", string(jsonMetadata), string(jsonAnnotations), videoURL)
					if updateError != nil {
						return nil, updateError.Error()
					}

					queryError := db.QueryRow("SELECT * FROM videos WHERE video_url = ?", videoURL).
						Scan(&updatedVideoItem.ID, &updatedVideoItem.URL, &metadataJSON, &annotationsJSON)
					if queryError != nil {
						return nil, queryError.Error()
					}

					// Adding the updated data to the current updated Video item
					json.Unmarshal(metadataJSON, &updatedVideoItem.METADATA)
					json.Unmarshal(annotationsJSON, &updatedVideoItem.ANNOTATIONS)

					return &updatedVideoItem, ""
				}
			}

			return nil, fmt.Sprintf("No annotation with type '%s' was found to delete the annotation", annotationType)
		}
	}

	return nil, "No video exists to show the annotations details"
}

func DeleteCompleteVideoData(db *sql.DB, videoURL string) (string, string) {
	// DELETE Videos
	_, videosError := db.Exec("DELETE FROM videos WHERE video_url = ?", videoURL)
	if videosError != nil {
		return "", fmt.Sprintf("Video deletion error %s", videosError.Error())
	}

	return "{Success: Video deleted successfully}", ""
}
