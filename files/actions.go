package files

import (
	"io"
	"os"
)

func SaveFile(file io.Reader, path string, errChan chan error) {
	if _, err := os.Stat(path); err == nil {
		errChan <- nil
	} else if os.IsNotExist(err) {
		CreateFile(file, path, errChan)
	} else {
		errChan <- err
	}
}

func CreateFile(file io.Reader, path string, errChan chan error) {
	systemFile, err := os.Create(path)
	if err != nil {
		errChan <- err
		return
	}
	defer systemFile.Close()

	_, err = io.Copy(systemFile, file)
	if err != nil {
		errChan <- err
		return
	}

	errChan <- nil
}
