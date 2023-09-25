package main

import (
	"log"
	"test/config"
	"test/internal/httpServer"
)

func main() {
	viperInstance, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Cannot load config. Error: {%s}", err.Error())
	}

	cfg, err := config.ParseConfig(viperInstance)
	if err != nil {
		log.Fatalf("Cannot parse config. Error: {%s}", err.Error())
	}

	if err = httpServer.NewServer(cfg).Run(); err != nil {
		log.Panic(err)
	}
}
