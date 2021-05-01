package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/dankokin/fivegen-assignment/models"
	"github.com/dankokin/fivegen-assignment/services"
)

// Worker structure implements the logic for deleting old files
type Worker struct {
	ExpirationInterval time.Duration
	WorkersQuantity uint
	SleepTime time.Duration

	Db services.DataStore
}

func NewWorker(epxTime time.Duration, quantity uint,
		sleep time.Duration, db services.DataStore) *Worker {
	return &Worker{
		ExpirationInterval: epxTime,
		WorkersQuantity:    quantity,
		SleepTime:          sleep,
		Db:                 db,
	}
}

func (w *Worker) DeleteExpiredFiles() {
	for {
		var wg sync.WaitGroup
		taskChannel := make(chan string, 256)

		for i := uint(0); i < w.WorkersQuantity; i++ {
			wg.Add(1)
			go w.launchHelper(&wg, taskChannel)
		}

		go w.Db.AllFilesRecords(taskChannel)

		wg.Wait()
		time.Sleep(w.SleepTime)
	}
}

func (w *Worker) launchHelper(wg *sync.WaitGroup, taskChannel chan string) {
	defer wg.Done()
	for fileRecord := range taskChannel {
		var file models.File
		err := json.NewDecoder(bytes.NewReader([]byte(fileRecord))).Decode(&file)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			if time.Unix(file.CreatedAt, 0).Add(w.ExpirationInterval).Unix() < time.Now().Unix() {
				w.Db.DeleteRecord(file.ShortUrl)
				fmt.Println(os.Remove(file.HashedName).Error())
			}
		}
	}
}

