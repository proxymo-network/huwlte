package main

import (
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"
)

var (
	monitoringCmd = &cli.Command{
		Name:  "monitoring",
		Usage: "contains monitoring related commands",
		Subcommands: []*cli.Command{
			monitoringStatusCmd,
		},
	}

	monitoringStatusCmd = &cli.Command{
		Name:  "status",
		Usage: "get status of connection",
		Action: func(c *cli.Context) error {
			ctx, cancel := newCtx(c)
			defer cancel()

			client, cancel, err := newClient(ctx, c)
			if err != nil {
				return xerrors.Errorf("failed to create client: %w", err)
			}
			defer cancel()

			info, err := client.Monitoring.Status(ctx)
			if err != nil {
				return xerrors.Errorf("failed to get basic information: %w", err)
			}

			return pretty(info)
		},
	}
)
