package main

import (
	"context"

	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"
)

func newCtx(c *cli.Context) (context.Context, context.CancelFunc) {
	ctx := context.Background()
	return context.WithTimeout(ctx, c.Duration("timeout"))
}

var (
	deviceCmd = &cli.Command{
		Name:  "device",
		Usage: "contains device related commands",
		Subcommands: []*cli.Command{
			deviceInfoCmd,
		},
	}

	deviceInfoCmd = &cli.Command{
		Name:    "basic-information",
		Aliases: []string{"info"},
		Usage:   "get basic information of device",
		Action: func(c *cli.Context) error {
			ctx, cancel := newCtx(c)
			defer cancel()

			client, cancel, err := newClient(ctx, c)
			if err != nil {
				return xerrors.Errorf("failed to create client: %w", err)
			}
			defer cancel()

			info, err := client.Device.BasicInformation(ctx)
			if err != nil {
				return xerrors.Errorf("failed to get basic information: %w", err)
			}

			return pretty(info)
		},
	}
)
