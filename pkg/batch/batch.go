package batch

import (
	"context"
	"errors"
	"time"
)

// ErrBlocked сообщает о блокировке сервиса.
var ErrBlocked = errors.New("blocked")

// Batch is a batch of items.
type Batch []Item

// Item is some abstract item.
type Item struct{}

// Service определяет внешний сервис, который может обрабатывать батчи объектов.
type Service struct {
	maxItems   uint64
	interval   time.Duration
	processing chan bool
}

// NewService создает новый экземпляр сервиса с заданными лимитами.
func NewService(maxItems uint64, interval time.Duration) *Service {
	return &Service{
		maxItems:   maxItems,
		interval:   interval,
		processing: make(chan bool, 1),
	}
}

// GetLimits возвращает текущие лимиты сервиса.
func (s *Service) GetLimits() (uint64, time.Duration) {
	return s.maxItems, s.interval
}

// Process выполняет обработку батча объектов.
func (s *Service) Process(ctx context.Context, batch Batch) error {
	select {
	case s.processing <- true:
		defer func() { <-s.processing }()
		if uint64(len(batch)) > s.maxItems {
			return ErrBlocked
		}
		// Здесь должен быть код для обработки батча.
		// Например, можно просто заснуть на указанный интервал времени.
		time.Sleep(s.interval)
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
