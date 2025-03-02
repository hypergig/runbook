package steps

import (
	"context"

	"github.com/hypergig/runbook/internal/modules"
)

type Steps []*Step

var _ modules.Module = (Steps)(nil)

func (s Steps) Run(ctx context.Context) error {
	for _, step := range s {
		if err := step.Run(ctx); err != nil {
			return err
		}
	}
	return nil
}
