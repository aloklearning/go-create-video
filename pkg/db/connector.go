package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// Standard practise to not return error as String
// Or just return nil when something happens in GO
func DBConnect() *sql.DB {
	db, err := sql.Open("sqlite3", "../pkg/db/videos.db")
	if err != nil {
		return nil
	}

	return db
}

func CreateTables(db *sql.DB) string {
	// Create Metadata Table
	_, metadataError := db.Exec("CREATE TABLE IF NOT EXISTS metadata (author TEXT, video_name TEXT, created_at TEXT, modified_at TEXT, total_duration INTEGER)")
	if metadataError != nil {
		return fmt.Sprintf("Metadata Creation Table Error:%v", metadataError.Error())
	}

	// Create Categories Table
	_, annotationError := db.Exec("CREATE TABLE IF NOT EXISTS annotations (annotation_id TEXT, start_time INTEGER, end_time INTEGER, type TEXT PRIMARY KEY, annotation TEXT, additional_notes TEXT)")
	if annotationError != nil {
		return fmt.Sprintf("Annotation Creation Table Error: %v", annotationError.Error())
	}

	// Create Videos Table
	_, videosError := db.Exec("CREATE TABLE IF NOT EXISTS videos (video_id TEXT PRIMARY KEY, video_url TEXT, metadata TEXT, annotations TEXT)")
	if videosError != nil {
		return fmt.Sprintf("Video Table Creation Table Error: %v", videosError.Error())
	}

	fmt.Printf("Table Created Successfully\n")
	return ""
}

// In case we are required to drop the table
func DropTable(db *sql.DB) string {
	_, videosError := db.Exec("DROP TABLE IF EXISTS videos")
	if videosError != nil {
		return fmt.Sprintf("Video Table Removal Error: %v", videosError.Error())
	}

	fmt.Printf("Table Dropped Successfully\n")
	return ""
}
