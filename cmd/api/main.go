package main

import (
	"github.com/rs/zerolog/pkgerrors"
	"log"
	"test/config"
	"test/internal/httpServer"
)

func main() {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	viperInstance, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Cannot load config. Error: {%s}", err.Error())
	}

	cfg, err := config.ParseConfig(viperInstance)
	if err != nil {
		log.Fatalf("Cannot parse config. Error: {%s}", err.Error())
	}

	s := httpServer.NewServer(cfg)
	if err = s.Run(); err != nil {
		log.Printf("%v", err)
	}
}
