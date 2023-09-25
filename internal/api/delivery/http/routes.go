package http

func MapApiRoutes(router fiber.Router, h tokenInterface.IHandler) {
	router.Post("/process", h.Process())
}
