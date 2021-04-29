package main

import (
	"flag"
	"fmt"
	"net/http"
	"text/template"
	"time"

	"github.com/dankokin/fivegen-assignment/handlers"
	"github.com/dankokin/fivegen-assignment/models"
)

type mockDb struct {
}

func (db *mockDb) UploadFileName(file *models.File, errChan chan error) {
	errChan <- nil
	return
}

func (db *mockDb) DownloadFileName(url string) *models.File {
	return &models.File{
		CreatedAt:    time.Now().Unix(),
		HashedName:   "602f5b7eeff60a54f482a6f9e5df343a",
		OriginalName: "Лр1.doc",
		ShortUrl:     "aaaaaa",
	}
}

func (db *mockDb) IsExists(key string, fileDataHash string) bool {
	return false
}

func main() {
	var mainPagePath string
	flag.StringVar(&mainPagePath, "main", "static/main_page.html", "Path to main page of web-server")
	var dataPath string
	flag.StringVar(&dataPath, "data", "data", "Path to stored files")
	var templates = template.Must(template.ParseFiles(mainPagePath))
	flag.Parse()

	db := &mockDb{}
	u := handlers.CreateUploader(db, 15, templates, mainPagePath, dataPath)

	http.HandleFunc("/main", u.MainPageHandler)
	http.HandleFunc("/api/upload", u.UploadFileHandler)
	http.HandleFunc("/", u.ServeFile)

	fmt.Println("starting server at :8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		return
	}
}
