package main

import (
	"flag"
	"github.com/artemzavg/avito-tech-backend-ta/internal/app/backend"
	"log"
	"os"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/json/config.json", "path to config file")
}

func main() {
	flag.Parse()

	configFile, err := os.OpenFile(configPath, os.O_RDONLY, 0755)
	if err != nil {
		log.Fatal("Wrong config path")
	}

	config := backend.NewConfig(configFile)
	server := backend.New(config)

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
