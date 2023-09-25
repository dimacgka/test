package http

import (
	"fmt"
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

		var Ip = "192.254.25.65"

		fmt.Println(Ip)

		if err := ctx.BodyParser(&params); err != nil {
			return err
		}

		err := h.uc.Process(ctx, params)
		if err != nil {
			return err
		}

		return ctx.JSON(response)
	}
}
