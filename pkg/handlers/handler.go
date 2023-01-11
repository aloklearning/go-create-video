package handlers

import (
	"encoding/json"
	source "go-create-video/pkg/source"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAllVideos(w http.ResponseWriter, r *http.Request) {
	videos := source.AllVideos()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(videos)
}

func CreateVideo(w http.ResponseWriter, r *http.Request) {
	var videoData source.Video

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewDecoder(r.Body).Decode(&videoData)

	newVideoList, errorMessage := source.Create(videoData)
	if errorMessage != "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(source.Error{ErrorMessage: errorMessage})

		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newVideoList)
}

func GetAllAnnotations(w http.ResponseWriter, r *http.Request) {
	//var annotations []source.Annotation
	urlParam := mux.Vars(r)

	w.Header().Set("Content-Type", "application/json")
	annotations, errorMessage := source.AllAnnotations(urlParam["url"])
	if errorMessage != "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(source.Error{ErrorMessage: errorMessage})

		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(annotations)
}
