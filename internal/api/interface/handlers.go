package apiInterface

import "github.com/gofiber/fiber/v2"

type IHandler interface {
	Process() fiber.Handler
}
