package worker

import (
	"context"
	"fmt"

	"github.com/dptsi/its-go-worker/contracts"
	"github.com/samber/do"
)

type Worker struct {
	ctx context.Context
	i   *do.Injector
}

func NewWorker(ctx context.Context, i *do.Injector) contracts.Worker {
	return &Worker{
		ctx: ctx,
		i:   i,
	}
}

type Provider[T any] func(w contracts.Worker) (T, error)

func Bind[T any](w contracts.Worker, provider Provider[T]) {
	do.Provide(w.Injector(), func(i *do.Injector) (T, error) {
		return provider(w)
	})
}

func MustMake[T any](w contracts.Worker) T {
	instance, err := do.Invoke[T](w.Injector())
	if err != nil {
		panic(fmt.Errorf("error when creating object: %w", err))
	}
	return instance
}

func Make[T any](w contracts.Worker) (T, error) {
	return do.Invoke[T](w.Injector())
}

func (w *Worker) Context() context.Context {
	return w.ctx
}

func (w *Worker) Services() []string {
	return w.i.ListProvidedServices()
}

func (w *Worker) Injector() *do.Injector {
	return w.i
}

func (w *Worker) Shutdown() error {
	return w.i.Shutdown()
}
