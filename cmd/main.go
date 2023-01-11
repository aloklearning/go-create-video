package main

import (
	"log"
	"net/http"

	routes "go-create-video/pkg/handlers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Route Handlers / Endpoints
	router.HandleFunc("/api/v1/videos", routes.GetAllVideos).Methods("GET")
	router.HandleFunc("/api/v1/createVideo", routes.CreateVideo).Methods("POST")
	router.HandleFunc("/api/v1/annotations/{url}", routes.GetAllAnnotations).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
