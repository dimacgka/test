package httpServer

import (
	"fmt"
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
	if err := s.MapHandlers(s.fiber, s.logger); err != nil {
		s.logger.Warnf("Cannot map handlers: ", err)
		panic(err)
	}

	s.logger.Infof("Start server on port: %s:%s", s.cfg.BaseConfig.System.Host, s.cfg.BaseConfig.System.Port)
	if err := s.fiber.Listen(fmt.Sprintf("%s:%s", s.cfg.BaseConfig.System.Host, s.cfg.BaseConfig.System.Port)); err != nil {
		s.logger.Warnf("Error starting Server: ", err)
		panic(err)
	}

	return nil
}
