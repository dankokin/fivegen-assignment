package handlers

import (
	"net/http"

	"github.com/dankokin/fivegen-assignment/utils"
)

func (u *Uploader) ServeFile(w http.ResponseWriter, r *http.Request) {
	if len(r.RequestURI) <= 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	shortUrl := r.RequestURI[1:]
	if shortUrl == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fileInfo := u.Db.DownloadFileName(shortUrl)
	if fileInfo == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	path := utils.ConcatenateStrings(u.DataPath, "/", fileInfo.HashedName)

	value := utils.ConcatenateStrings("attachment; filename=", fileInfo.OriginalName)
	w.Header().Set("Content-Disposition", value)
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	http.ServeFile(w, r, path)
}
