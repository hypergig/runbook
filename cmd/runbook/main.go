package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/hypergig/runbook/internal/modules/steps"
	"github.com/urfave/cli/v3"
	"gopkg.in/yaml.v3"
)

func main() {

	var runbookPath string
	cmd := &cli.Command{
		Name:  "runbook",
		Usage: "Run a runbook",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			raw, err := os.ReadFile(runbookPath)
			if err != nil {
				return err
			}

			var runbook steps.Step
			if err := yaml.Unmarshal(raw, &runbook); err != nil {
				return err
			}

			return runbook.Run(ctx)
		},
		ArgsUsage: "[runbook.yaml]",
		Arguments: []cli.Argument{
			&cli.StringArg{
				Name:        "runbook file",
				Value:       "runbook.yaml",
				UsageText:   "the location of the runbook",
				Max:         1,
				Destination: &runbookPath,
			},
		},
	}

	ctx := context.Background()

	if err := cmd.Run(ctx, os.Args); err != nil {
		slog.ErrorContext(ctx, "error", slog.Any("err", err))
	}
}
