package models

import (
	"encoding/json"
	"os"
)

type Config struct {
	Database string `json:"database" bson:"database"`
}

func LoadConfiguration() (Config, error) {
	var config Config
	configFile, err := os.Open("config.json")
	defer configFile.Close()

	if err != nil {
		return config, err
	}
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)

	return config, err
}
