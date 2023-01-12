package main

import (
	"fmt"
	"log"
	"net/http"

	db "go-create-video/pkg/db"
	routes "go-create-video/pkg/handlers"

	"github.com/gorilla/mux"
)

func main() {

	// Connecting to the DB
	vidoeDB, status := db.DBConnect()
	if status != "Successfully Connected to the DB" {
		fmt.Printf("%s", status)
		return
	}

	defer vidoeDB.Close()

	err := db.CreateTables(vidoeDB)
	if err != "" {
		fmt.Printf("%s", err)
		return
	}

	router := mux.NewRouter()
	routerHandler := &routes.RouterHandler{Database: vidoeDB}

	// Route Handlers / Endpoints
	router.HandleFunc("/api/v1/videos", routerHandler.GetAllVideos).Methods("GET")
	router.HandleFunc("/api/v1/createVideo", routerHandler.CreateVideo).Methods("POST")
	router.HandleFunc("/api/v1/annotations", routerHandler.GetAllAnnotations).Methods("GET")
	router.HandleFunc("/api/v1/updateAdditionalNotes", routerHandler.UpdateAnnotationAdditionalNotes).Methods("PUT")
	router.HandleFunc("/api/v1/updateAnnotation/{video_url}/{type}", routerHandler.UpdateAnnotation).Methods("PUT")
	// router.HandleFunc("/api/v1/deleteAnnotation", routes.DeleteAnnotation).Methods("DELETE")
	router.HandleFunc("/api/v1/deleteVideo", routerHandler.DeleteVideo).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
