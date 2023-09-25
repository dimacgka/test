package batch

import (
	"context"
	"testing"
	"time"
)

func TestService_GetLimits(t *testing.T) {
	maxItems := uint64(10)
	interval := 2 * time.Second

	service := NewService(maxItems, interval)

	n, p := service.GetLimits()

	if n != maxItems {
		t.Errorf("Expected max items: %d, got: %d", maxItems, n)
	}

	if p != interval {
		t.Errorf("Expected interval: %v, got: %v", interval, p)
	}
}

func TestService_Process(t *testing.T) {
	maxItems := uint64(10)
	interval := 2 * time.Second

	service := NewService(maxItems, interval)

	t.Run("Process with valid batch size", func(t *testing.T) {
		batch := make(Batch, 5)
		err := service.Process(context.Background(), batch)

		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}
	})

	t.Run("Process with blocked batch size", func(t *testing.T) {
		batch := make(Batch, 15)
		err := service.Process(context.Background(), batch)

		if err != ErrBlocked {
			t.Errorf("Expected ErrBlocked, got: %v", err)
		}
	})

	t.Run("Process with context timeout", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()

		batch := make(Batch, 5)
		err := service.Process(ctx, batch)

		if err != context.DeadlineExceeded {
			t.Errorf("Expected context.DeadlineExceeded, got: %v", err)
		}
	})
}
