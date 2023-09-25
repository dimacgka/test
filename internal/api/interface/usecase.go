package apiInterface

import "test/pkg/batch"

type IUseCase interface {
	Process(ctx fiber.Context, params batch.Batch) error
}
