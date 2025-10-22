package rwtesting

import (
	"context"
	"time"
)

// Exemplo de função testável com injeção de dependência
type Worker interface {
	Do(ctx context.Context) error
}

type RealWorker struct{}

func (r *RealWorker) Do(ctx context.Context) error {
	select {
	case <-time.After(10 * time.Millisecond):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// Test helper que usa contexto e timeout
func RunWorker(w Worker) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return w.Do(ctx)
}
