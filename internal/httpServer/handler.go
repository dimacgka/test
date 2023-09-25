package httpServer

import (
	"github.com/gofiber/fiber/v2"
	apiHttp "test/internal/api/delivery/http"
	"test/internal/api/usecase"
	"test/pkg/batch"
	"time"
)

func (s *Server) MapHandlers(app *fiber.App) error {
	apiGroup := app.Group("api")

	maxItems := uint64(10)
	interval := 2 * time.Second

	batchService := batch.NewService(maxItems, interval)

	uc := usecase.NewUC(*batchService)
	handlers := apiHttp.NewHandler(uc)

	apiHttp.MapApiRoutes(apiGroup, handlers)

	return nil
}
