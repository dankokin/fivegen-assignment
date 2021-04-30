package models

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Address     string `json:"address"`
	Port        string `json:"port"`
	MaxFileSize uint   `json:"max_file_size"`
}

func InitConfigFile(cnfFile string) (*Config, error) {
	jsonFile, err := ioutil.ReadFile(cnfFile + ".json")

	var cnf Config
	jErr := json.Unmarshal(jsonFile, &cnf)

	if err != nil || jErr != nil {
		return nil, err
	}

	return &cnf, nil
}
