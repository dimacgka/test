package usecase

import (
	"errors"
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

func (u *UC) Process(ctx fiber.Context, params batch.Batch) error {
	if err := service.Process(ctx, params); err != nil {
		if err == batch.ErrBlocked {
			return errors.New("service is blocked")
		} else {
			return errors.New("internal error")
		}
	}
	return nil
}
