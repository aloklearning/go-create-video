package handlers

import (
	"encoding/json"
	source "go-create-video/pkg/src"
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
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	annotations, errorMessage := source.AllAnnotations(r.FormValue("video_url"))
	if errorMessage != "" {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(source.Error{ErrorMessage: errorMessage})

		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(annotations)
}

func UpdateAnnotationAdditionalNotes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	updatedVideoData, errorMessage := source.AddAdditionalNotes(
		r.FormValue("video_url"),
		r.FormValue("type"),
		r.FormValue("notes"),
	)

	if errorMessage != "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(source.Error{ErrorMessage: errorMessage})

		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(updatedVideoData)
}

func UpdateAnnotation(w http.ResponseWriter, r *http.Request) {
	var annotationDetails source.Annotation
	w.Header().Set("Content-Type", "application/json")

	paramData := mux.Vars(r)

	_ = json.NewDecoder(r.Body).Decode(&annotationDetails)

	updatedVideoData, errorMessage := source.UpdateAnnotationDetails(
		paramData["video_url"],
		paramData["type"],
		annotationDetails,
	)

	if errorMessage != "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(source.Error{ErrorMessage: errorMessage})

		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedVideoData)
}

func DeleteAnnotation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	updatedVideoData, errorMessage := source.DeleteAnnotationData(
		r.FormValue("video_url"),
		r.FormValue("type"),
	)

	if errorMessage != "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(source.Error{ErrorMessage: errorMessage})

		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedVideoData)
}

func DeleteVideo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	updatedVideoData, errorMessage := source.DeleteVideoData(r.FormValue("video_url"))

	if errorMessage != "" {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(source.Error{ErrorMessage: errorMessage})

		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedVideoData)
}
