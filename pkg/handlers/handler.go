package handlers

import (
	"encoding/json"
	source "go-create-video/pkg/source"
	"net/http"
)

func GetAllVideos(w http.ResponseWriter, r *http.Request) {
	videos := source.AllVideos()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(videos)
}
