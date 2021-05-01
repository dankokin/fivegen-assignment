package models

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	ServerAddress string `json:"server_address"`
	ServerPort    string `json:"server_port"`
	MaxFileSize   uint   `json:"max_file_size"`

	ExpirationTimeInDays int `json:"expiration_time_in_days"`
	WorkersQuantity uint `json:"workers_quantity"`
	WorkerTimeoutInDays int `json:"worker_timeout_in_days"`

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
