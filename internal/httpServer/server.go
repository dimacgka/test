package httpServer

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"test/config"
)

// Server struct
type Server struct {
	fiber *fiber.App
	cfg   *config.Config
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		fiber: fiber.New(fiber.Config{DisableStartupMessage: true}),
		cfg:   cfg,
	}
}

func (s *Server) Run() error {
	if err := s.MapHandlers(s.fiber); err != nil {
		panic(err)
	}

	log.Printf("Start server on port: %s:%s", s.cfg.BaseConfig.System.Host, s.cfg.BaseConfig.System.Port)

	if err := s.fiber.Listen(fmt.Sprintf("%s:%s", s.cfg.BaseConfig.System.Host, s.cfg.BaseConfig.System.Port)); err != nil {
		panic(err)
	}

	return nil
}
