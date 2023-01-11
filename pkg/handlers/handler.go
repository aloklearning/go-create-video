package handlers

import (
	"encoding/json"
	source "go-create-video/pkg/source"
	"net/http"
)

func GetAllVideos(w http.ResponseWriter, r *http.Request) {
	videos := source.AllVideos()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(videos)
}

func CreateVideo(w http.ResponseWriter, r *http.Request) {
	var videoData source.Video

	_ = json.NewDecoder(r.Body).Decode(&videoData)

	newVideoList := source.Create(videoData)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newVideoList)
}
