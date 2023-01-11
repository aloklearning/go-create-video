package source

import (
	"github.com/google/uuid"
)

var videos []Video

func AllVideos() *[]Video {
	return &videos
}

func Create(videoData Video) *[]Video {
	videoData.ID = uuid.NewString()

	for _, video := range videos {
		if video.URL == videoData.URL {
			return &videos
		}
	}

	videos = append(videos, videoData)

	return &videos
}
