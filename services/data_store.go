package services

import "github.com/dankokin/fivegen-assignment/models"

type DataStore interface {
	UploadFileInfo(file *models.File, errChan chan error)
	DownloadFileInfo(url string) *models.File
	IsExists(key string, fileDataHash string) bool
	AllFilesRecords(chan string)
	DeleteRecord(string)
}
