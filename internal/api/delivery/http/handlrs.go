package http

import (
	"github.com/gofiber/fiber/v2"
	apiInterface "test/internal/api/interface"
	"test/pkg/batch"
)

type Handler struct {
	uc apiInterface.IUseCase
}

func NewHandler(uc apiInterface.IUseCase) apiInterface.IHandler {
	return &Handler{
		uc: uc,
	}
}

func (h *Handler) Process() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params batch.Batch

		if err := ctx.BodyParser(&params); err != nil {
			return err
		}

		err := h.uc.Process(ctx, params)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "OK",
		})
	}
}
