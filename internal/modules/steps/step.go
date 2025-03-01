package steps

import (
	"context"
	"errors"
	"log/slog"

	"github.com/hypergig/runbook/internal/modules"
	"github.com/hypergig/runbook/internal/modules/exec"
)

type Step struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`

	// all possible modules
	Exec  *exec.Exec `json:"exec,omitempty"`
	Steps *Steps     `json:"steps,omitempty"`
}

var _ modules.Module = (*Step)(nil)

func (s *Step) getModule() (modules.Module, error) {
	switch {
	case s.Exec != nil:
		return s.Exec, nil
	case s.Steps != nil:
		return s.Steps, nil
	default:
		return nil, errors.New("missing module in step")
	}
}

func (s *Step) Run(ctx context.Context) error {
	module, err := s.getModule()
	if err != nil {
		return err
	}

	slog.InfoContext(ctx, "starting step", slog.String("name", s.Name), slog.String("description", s.Description))

	err = module.Run(ctx)

	if err != nil {
		slog.InfoContext(ctx, "step errored out", slog.Any("error", err))
		return err
	}
	slog.InfoContext(ctx, "finished step", slog.String("name", s.Name))

	return nil
}
