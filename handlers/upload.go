package handlers

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"hash/crc32"
	"io/ioutil"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/dankokin/fivegen-assignment/files"
	"github.com/dankokin/fivegen-assignment/models"
	"github.com/dankokin/fivegen-assignment/services"
	"github.com/dankokin/fivegen-assignment/utils"
)

type Uploader struct {
	Db services.DataStore

	FileMaxSize  uint `json:"file_max_size"`
	TemplateFile *template.Template

	WebServerAddress string
	MainPagePath     string
	DataPath         string
}

func CreateUploader(db services.DataStore, maxSize uint, tmp *template.Template,
	mainPath, dataPath, addr string) *Uploader {
	return &Uploader{
		Db:               db,
		FileMaxSize:      maxSize,
		TemplateFile:     tmp,
		MainPagePath:     mainPath,
		DataPath:         dataPath,
		WebServerAddress: addr,
	}
}

func (u *Uploader) Hash(data []byte) string {
	hasher := md5.New()
	hasher.Write(data)
	return hex.EncodeToString(hasher.Sum(nil))
}


func (u *Uploader) NewShortURL(fileDataHash string) string {
	crcH := crc32.ChecksumIEEE([]byte(fileDataHash))
	dataHash := strconv.FormatUint(uint64(crcH), 36)
	// Continue the loop until the hash function returns a unique value
	for i := uint64(0); u.Db.IsExists(dataHash, fileDataHash); i++ {
		fmt.Println(i)
		salt := strconv.FormatUint(i, 10)
		fileDataHash = utils.ConcatenateStrings(dataHash, salt)
		crcH = crc32.ChecksumIEEE([]byte(fileDataHash))
		dataHash = strconv.FormatUint(uint64(crcH), 36)
	}
	return dataHash
}

func (u *Uploader) MainPageHandler(w http.ResponseWriter, r *http.Request) {
	err := u.TemplateFile.Execute(w, u.MainPagePath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// UploadFileHandler uploads files to the our server
func (u *Uploader) UploadFileHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(int64(u.FileMaxSize) << 20)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	file, fileHeader, err := r.FormFile("my_file")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer file.Close()

	rawFile, err := ioutil.ReadAll(file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fileDataHash := u.Hash(rawFile)

	path := utils.ConcatenateStrings(
		u.DataPath,
		"/",
		fileDataHash)

	errChan := make(chan error, 2)
	go files.SaveFile(rawFile, path, errChan)

	fileModel := models.File{
		CreatedAt:    time.Now().Unix(),
		HashedName:   fileDataHash,
		OriginalName: fileHeader.Filename,
		ShortUrl:     u.NewShortURL(fileDataHash),
	}

	go u.Db.UploadFileInfo(&fileModel, errChan)

	for i := 0; i < 2; i++ {
		err = <-errChan
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	json.NewEncoder(w).Encode(models.MakeResponse(u.WebServerAddress, fileModel.ShortUrl))
}
