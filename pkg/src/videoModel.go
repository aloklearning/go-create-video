package source

import "time"

type Error struct {
	ErrorMessage string `json:"error"`
}

type Video struct {
	ID          string       `json:"video_id"`
	URL         string       `json:"video_url"`
	METADATA    Metadata     `json:"meta_data"`
	ANNOTATIONS []Annotation `json:"annotations"`
}

type Metadata struct {
	AUTHOR     string    `json:"author"`
	NAME       string    `json:"video_name"`
	CREATEDAT  time.Time `json:"created_at"`
	MODIFIEDAT time.Time `json:"modified_at"`
	DURATION   int       `json:"total_duration"`
}

type Annotation struct {
	ID              string   `json:"annotation_id"`
	STARTTIME       int      `json:"start_time"`
	ENDTIME         int      `json:"end_time"`
	TYPE            string   `json:"type"`
	ANNOTATION      string   `json:"annotation"`
	ADDITIONALNOTES []string `json:"additional_notes"`
}
