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
	"time"

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
		os.Getenv("NGINX_ADDRESS")+":"+os.Getenv("NGINX_PORT"), "ServerAddress of web-server")

	flag.Parse()

	var templates = template.Must(template.ParseFiles(mainPagePath))
	rdb := services.NewRedisDataStore(os.Getenv("REDIS_URL"), os.Getenv("REDIS_PASSWORD"), 0, context.Background())
	u := handlers.CreateUploader(rdb, conf.MaxFileSize, templates, mainPagePath, dataPath, serverAddr)

	worker := handlers.NewWorker(
		time.Duration(conf.ExpirationTimeInDays)*time.Hour*24,
		conf.WorkersQuantity, time.Duration(conf.WorkerTimeoutInDays)*time.Hour*24, rdb, dataPath)

	fmt.Println(worker)
	go worker.DeleteExpiredFiles()

	http.HandleFunc("/main", u.MainPageHandler)
	http.HandleFunc("/api/upload", u.UploadFileHandler)
	http.HandleFunc("/", u.ServeFileHandler)

	fmt.Println("starting server at :" + conf.ServerPort)
	log.Fatal(http.ListenAndServe(":"+conf.ServerPort, nil))
}
