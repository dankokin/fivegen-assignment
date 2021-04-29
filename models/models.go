package models

type File struct {
	CreatedAt    int64
	HashedName   string
	OriginalName string
	ShortUrl     string `json:"short_url"`
}
