package modules

import (
	"context"
)

type Module interface {
	Run(ctx context.Context) error
}
