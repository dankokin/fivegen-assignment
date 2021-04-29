package services

import "github.com/dankokin/fivegen-assignment/models"

type DataStore interface {
	UploadFileName(file *models.File, errChan chan error)
	DownloadFileName(url string) *models.File
	IsExists(key string, fileDataHash string) bool
}
