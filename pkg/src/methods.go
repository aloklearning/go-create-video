package source

import (
	"fmt"

	"github.com/google/uuid"
)

var videos []Video

func AllVideos() *[]Video {
	return &videos
}

func Create(videoData Video) (*[]Video, string) {
	videoData.ID = uuid.NewString()

	if videoData.ANNOTATIONS[0].ENDTIME > videoData.METADATA.DURATION {
		errMessage := fmt.Sprintf("Your annotations end time %d is out of bounds of duration of the video %d",
			videoData.ANNOTATIONS[0].ENDTIME, videoData.METADATA.DURATION)
		return nil, errMessage
	}

	for _, video := range videos {
		if video.URL == videoData.URL {
			return &videos, ""
		}
	}

	videos = append(videos, videoData)

	return &videos, ""
}

func AllAnnotations(videoURL string) ([]Annotation, string) {
	for _, video := range videos {
		if videoURL == video.URL {
			return video.ANNOTATIONS, ""
		}
	}

	return nil, "No video exists to show the annotations details"
}

func AddAdditionalNotes(videoURL, annotationType, notes string) (*Video, string) {
	if notes != "" || len(notes) > 0 {
		for _, video := range videos {
			if videoURL == video.URL {
				for index, annotation := range video.ANNOTATIONS {
					if annotationType == annotation.TYPE {
						annotation.ADDITIONALNOTES = append(annotation.ADDITIONALNOTES, notes)

						video.ANNOTATIONS[index].ADDITIONALNOTES = annotation.ADDITIONALNOTES
						return &video, ""
					}
				}

				return nil, "No annotation type was found to add additional notes"
			}
		}

		return nil, "No video exists to show the annotations details"
	}

	return nil, "Empty notes were passed. Please add and try again"
}

func UpdateAnnotationDetails(videoURL, annotationType string, newAnnotation Annotation) (*Video, string) {
	for _, video := range videos {
		if videoURL == video.URL {
			for index, annotation := range video.ANNOTATIONS {
				if annotationType == annotation.TYPE {
					video.ANNOTATIONS[index] = newAnnotation
					return &video, ""
				}
			}

			return nil, "No annotation type was found to update the annotation"
		}
	}

	return nil, "No video exists to show the annotations details"
}

func DeleteAnnotationData(videoURL, annotationType string) (*Video, string) {
	for videoIndex, video := range videos {
		if videoURL == video.URL {
			for index, annotation := range video.ANNOTATIONS {
				if annotationType == annotation.TYPE {
					video.ANNOTATIONS = append(video.ANNOTATIONS[:index], video.ANNOTATIONS[index+1:]...)

					// To keep the updated element in the main data
					videos[videoIndex] = video
					return &video, ""
				}
			}

			return nil, "No annotation type was found to delete the annotation"
		}
	}

	return nil, "No video exists to show the annotations details"
}

func DeleteVideoData(videoURL string) (*[]Video, string) {
	for index, video := range videos {
		if videoURL == video.URL {
			videos = append(videos[:index], videos[index+1:]...)

			return &videos, ""
		}
	}

	return nil, "No video found to be deleted from the data"
}
