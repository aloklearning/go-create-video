package source

import (
	"github.com/google/uuid"
)

var videos []Video

func AllVideos() *[]Video {
	videoList := append(videos, Video{
		ID:  uuid.NewString(),
		URL: "Random Video URL",
		METADATA: Metadata{
			AUTHOR: "Alok",
		},
		ANNOTATIONS: []Annotation{
			{
				TYPE:            "None",
				ANNOTATION:      "Some Annotation",
				ADDITIONALNOTES: "Additional Notes",
			},
		},
	})

	return &videoList
}

func Create(videoData Video) *[]Video {
	videoData.ID = uuid.NewString()

	videos := append(videos, videoData)

	return &videos
}
