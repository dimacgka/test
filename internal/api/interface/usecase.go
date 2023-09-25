package apiInterface

import (
	"github.com/gofiber/fiber/v2"
	"test/pkg/batch"
)

type IUseCase interface {
	Process(ctx *fiber.Ctx, params batch.Batch) error
}
