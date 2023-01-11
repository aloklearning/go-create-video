package source

import "github.com/google/uuid"

func AllVideos() *Video {
	return &Video{
		ID: uuid.New().String(),
	}
}
