package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/dankokin/fivegen-assignment/models"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/dankokin/fivegen-assignment/handlers"
	"github.com/dankokin/fivegen-assignment/services"
)

func main() {
	conf, err := models.InitConfigFile("settings")
	if err != nil {
		log.Fatal(err.Error())
	}

	var mainPagePath string
	flag.StringVar(&mainPagePath, "main", "static/main_page.html",
		"Path to main page of web-server")

	var dataPath string
	flag.StringVar(&dataPath, "data", "data", "Path to stored files")

	var serverAddr string
	flag.StringVar(&serverAddr, "addr",
		conf.Address + ":" + conf.Port, "Address of web-server")

	flag.Parse()

	var templates = template.Must(template.ParseFiles(mainPagePath))
	rdb := services.NewRedisDataStore(os.Getenv("REDIS_URL"), os.Getenv("REDIS_PASSWORD"), 0, context.Background())
	u := handlers.CreateUploader(rdb, conf.MaxFileSize, templates, mainPagePath, dataPath, serverAddr)

	http.HandleFunc("/main", u.MainPageHandler)
	http.HandleFunc("/api/upload", u.UploadFileHandler)
	http.HandleFunc("/", u.ServeFile)

	fmt.Printf("starting server at :%s port", conf.Port)
	log.Fatal(http.ListenAndServe(":" + conf.Port, nil))
}
