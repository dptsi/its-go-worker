package contracts

import (
	"context"

	"github.com/samber/do"
)

type Worker interface {
	Context() context.Context
	Services() []string
	Injector() *do.Injector
	Shutdown() error
}
