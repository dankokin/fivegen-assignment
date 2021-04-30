package models

import "github.com/dankokin/fivegen-assignment/utils"

type File struct {
	CreatedAt    int64  `json:"created_at"`
	HashedName   string `json:"hashed_name"`
	OriginalName string `json:"original_name"`
	ShortUrl     string `json:"short_url"`
}

type ShortUrlResponse struct {
	ShortUrl string `json:"short_url"`
}

func MakeResponse(addr, url string) ShortUrlResponse {
	return ShortUrlResponse{
		ShortUrl: utils.ConcatenateStrings("http://", addr, "/", url),
	}
}
