package contracts

import (
	"context"
	"fmt"
)

var ErrNoItemProcessed = fmt.Errorf("no item processed")

type Runner interface {
	Run(context.Context) error
}
