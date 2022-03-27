package backend

import (
	"encoding/json"
	"io"
	"log"
)

// Config ...
type Config struct {
	DbConnectionString string `json:"DbConnectionString"`
	BindAddr           string `json:"BindAddress"`
}

// NewConfig ...
func NewConfig(reader io.Reader) *Config {
	config := &Config{}

	encoder := json.NewDecoder(reader)
	err := encoder.Decode(config)
	if err != nil {
		log.Fatal(err)
	}

	return config
}
