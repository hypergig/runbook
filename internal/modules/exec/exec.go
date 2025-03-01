package exec

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"strings"

	"github.com/hypergig/runbook/internal/modules"
)

type Exec struct {
	Env  []string `json:"env,omitempty"`
	Cmd  string   `json:"cmd,omitempty"`
	Args []string `json:"args,omitempty"`
}

var _ modules.Module = (*Exec)(nil)

type logWritter struct {
	ctx   context.Context
	level slog.Level
}

func (e *Exec) parseEnv() ([]string, error) {
	r := []string{}

	for _, env := range e.Env {
		if strings.Contains(env, "=") {
			r = append(r, env)
			continue
		}

		value, ok := os.LookupEnv(env)
		if !ok {
			return nil, fmt.Errorf("no variable named %s in environment", env)
		}
		r = append(r, fmt.Sprintf("%s=%s", env, value))
	}

	return r, nil
}

func (e *Exec) Run(ctx context.Context) error {
	cmd := exec.CommandContext(ctx, e.Cmd, e.Args...)
	cmd.Stdout = &logWritter{ctx, slog.LevelInfo}
	cmd.Stderr = &logWritter{ctx, slog.LevelError}

	var err error
	cmd.Env, err = e.parseEnv()
	if err != nil {
		return err
	}

	return cmd.Run()
}
