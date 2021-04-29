package files

import (
	"io/ioutil"
	"os"
)

func SaveFile(file []byte, path string, errChan chan error) {
	if _, err := os.Stat(path); err == nil {
		errChan <- nil
	} else if os.IsNotExist(err) {
		CreateFile(file, path, errChan)
	} else {
		errChan <- err
	}
}

func CreateFile(file []byte, path string, errChan chan error) {
	err := ioutil.WriteFile(path, file, 0777)
	if err != nil {
		errChan <- err
		return
	}

	errChan <- nil
}
