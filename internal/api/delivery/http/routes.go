package http

import (
	"github.com/gofiber/fiber/v2"
	apiInterface "test/internal/api/interface"
)

func MapApiRoutes(router fiber.Router, h apiInterface.IHandler) {
	router.Post("/process", h.Process())
}
