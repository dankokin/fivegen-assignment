package services

import "github.com/dankokin/fivegen-assignment/models"

type DataStore interface {
	UploadFileName(file *models.File, errChan chan error)
	DownloadFileName(file models.File, errChan chan error)
	IsExists(key string, fileDataHash string) bool
}
