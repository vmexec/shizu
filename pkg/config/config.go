package config

import (
	"encoding/json"
	"io/ioutil"
)

// config values from json, port and test
type Config struct {
	Port int `json:"port"`
}

func LoadConfigFromFile(filePath string) (*Config, error) {
	// Read the JSON file
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON data into a Config struct
	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
