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
