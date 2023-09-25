package apiInterface

type IHandler interface {
	Process() fiber.Handler
}
