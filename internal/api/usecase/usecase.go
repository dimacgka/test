package usecase

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	apiInterface "test/internal/api/interface"
	"test/pkg/batch"
)

type UC struct {
	batchService batch.Service
}

func NewUC(batchService batch.Service) apiInterface.IUseCase {
	return &UC{
		batchService: batchService,
	}
}

func (u *UC) Process(ctx *fiber.Ctx, params batch.Batch) error {
	context := ctx.Context()

	if err := u.batchService.Process(context, params); err != nil {
		if err == batch.ErrBlocked {
			return errors.New("service is blocked")
		} else {
			return errors.New("internal error")
		}
	}

	return nil
}
