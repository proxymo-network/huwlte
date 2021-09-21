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
			deviceInfoBasicCmd,
			deviceInfo,
			deviceReboot,
		},
	}

	deviceReboot = &cli.Command{
		Name:  "reboot",
		Usage: "reboot device",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "wait",
				Usage: "wait for device to come back online",
			},
		},
		Action: func(c *cli.Context) error {
			ctx, cancel := newCtx(c)
			defer cancel()

			client, cancel, err := newClient(ctx, c)
			if err != nil {
				return xerrors.Errorf("failed to create client: %w", err)
			}
			defer cancel()

			if err := client.Device.Control(ctx, 1); err != nil {
				return xerrors.Errorf("send request: %w", err)
			}

			return nil
		},
	}

	deviceInfoBasicCmd = &cli.Command{
		Name:    "basic-information",
		Aliases: []string{"basic-info"},
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

	deviceInfo = &cli.Command{
		Name:    "information",
		Aliases: []string{"info"},
		Usage:   "get full information of device",
		Action: func(c *cli.Context) error {
			ctx, cancel := newCtx(c)
			defer cancel()

			client, cancel, err := newClient(ctx, c)
			if err != nil {
				return xerrors.Errorf("failed to create client: %w", err)
			}
			defer cancel()

			info, err := client.Device.Information(ctx)
			if err != nil {
				return xerrors.Errorf("failed to get basic information: %w", err)
			}

			return pretty(info)
		},
	}
)
