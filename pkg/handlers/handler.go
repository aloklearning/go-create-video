package handlers

import (
	"database/sql"
	"encoding/json"
	source "go-create-video/pkg/src"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type RouterHandler struct {
	Database *sql.DB
}

func (handler *RouterHandler) GetAllVideos(w http.ResponseWriter, r *http.Request) {
	videos, errorMessage := source.AllVideos(handler.Database)
	if errorMessage != "" {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(source.Error{ErrorMessage: errorMessage})

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(videos)
}

func (handler *RouterHandler) CreateVideo(w http.ResponseWriter, r *http.Request) {
	var videoData source.Video

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewDecoder(r.Body).Decode(&videoData)

	newVideoList, errorMessage := source.Create(handler.Database, videoData)
	if errorMessage != "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(source.Error{ErrorMessage: errorMessage})

		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newVideoList)
}

func (handler *RouterHandler) GetAllAnnotations(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	annotations, errorMessage := source.AllAnnotations(handler.Database, r.FormValue("video_url"))
	if errorMessage != "" {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(source.Error{ErrorMessage: errorMessage})

		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(annotations)
}

// func UpdateAnnotationAdditionalNotes(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	err := r.ParseMultipartForm(32 << 20)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	updatedVideoData, errorMessage := source.AddAdditionalNotes(
// 		r.FormValue("video_url"),
// 		r.FormValue("type"),
// 		r.FormValue("notes"),
// 	)

// 	if errorMessage != "" {
// 		w.WriteHeader(http.StatusBadRequest)
// 		json.NewEncoder(w).Encode(source.Error{ErrorMessage: errorMessage})

// 		return
// 	}

// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode(updatedVideoData)
// }

// func UpdateAnnotation(w http.ResponseWriter, r *http.Request) {
// 	var annotationDetails source.Annotation
// 	w.Header().Set("Content-Type", "application/json")

// 	paramData := mux.Vars(r)

// 	_ = json.NewDecoder(r.Body).Decode(&annotationDetails)

// 	updatedVideoData, errorMessage := source.UpdateAnnotationDetails(
// 		paramData["video_url"],
// 		paramData["type"],
// 		annotationDetails,
// 	)

// 	if errorMessage != "" {
// 		w.WriteHeader(http.StatusBadRequest)
// 		json.NewEncoder(w).Encode(source.Error{ErrorMessage: errorMessage})

// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(updatedVideoData)
// }

// func DeleteAnnotation(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	err := r.ParseMultipartForm(32 << 20)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	updatedVideoData, errorMessage := source.DeleteAnnotationData(
// 		r.FormValue("video_url"),
// 		r.FormValue("type"),
// 	)

// 	if errorMessage != "" {
// 		w.WriteHeader(http.StatusBadRequest)
// 		json.NewEncoder(w).Encode(source.Error{ErrorMessage: errorMessage})

// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(updatedVideoData)
// }

// func DeleteVideo(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	err := r.ParseMultipartForm(32 << 20)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	updatedVideoData, errorMessage := source.DeleteVideoData(r.FormValue("video_url"))

// 	if errorMessage != "" {
// 		w.WriteHeader(http.StatusNotFound)
// 		json.NewEncoder(w).Encode(source.Error{ErrorMessage: errorMessage})

// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(updatedVideoData)
// }
