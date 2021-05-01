package handlers

import (
	"net/http"

	"github.com/dankokin/fivegen-assignment/utils"
)

// ServeFileHandler is function that provides the ability to download files
func (u *Uploader) ServeFileHandler(w http.ResponseWriter, r *http.Request) {
	// Checking for an empty uri
	if len(r.RequestURI) <= 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Take everything except 1 character, which is a slash
	shortUrl := r.RequestURI[1:]
	if shortUrl == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Load data about the file from the database and check if it is there
	fileInfo := u.Db.DownloadFileInfo(shortUrl)
	if fileInfo == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Give the file
	path := utils.ConcatenateStrings(u.DataPath, "/", fileInfo.HashedName)
	headerValue := utils.ConcatenateStrings("attachment; filename=", fileInfo.OriginalName)

	w.Header().Set("Content-Disposition", headerValue)
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))

	http.ServeFile(w, r, path)
}
