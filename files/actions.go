package files

import (
	"io/ioutil"
	"os"
)

// SaveFile function saves the file to the specified path on server
func SaveFile(file []byte, path string, errorChan chan error) {
	if _, err := os.Stat(path); err == nil {
		errorChan <- nil
	} else if os.IsNotExist(err) {
		WriteFile(file, path, errorChan)
	} else {
		errorChan <- err
	}
}

// WriteFile function copies the contents of the file to a new one to the
// specified path on server
func WriteFile(file []byte, path string, errChan chan error) {
	err := ioutil.WriteFile(path, file, 0777)
	if err != nil {
		errChan <- err
		return
	}

	errChan <- nil
}
