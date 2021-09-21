package main

import (
	"github.com/proxymo-network/huwlte"
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"
)

var (
	netCmd = &cli.Command{
		Name:  "net",
		Usage: "contains mobile network related commands",
		Subcommands: []*cli.Command{
			netModeCmd,
		},
	}

	netModeCmd = &cli.Command{
		Name:    "mode",
		Aliases: []string{"mode"},
		Usage:   "configure connection mode",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "mode",
				Usage:    "set network mode",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "band",
				Usage:    "set network band",
				Required: true,
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

			netMode, err := huwlte.ParseNetworkMode(c.String("mode"))
			if err != nil {
				return xerrors.Errorf("failed to parse network mode: %w", err)
			}

			netBand, err := huwlte.ParseNetworkBand(c.String("band"))
			if err != nil {
				return xerrors.Errorf("failed to parse network band: %w", err)
			}

			if err := client.Net.SetMode(ctx, huwlte.NetMode{
				NetworkMode: netMode,
				NetworkBand: netBand,
				LTEBand:     huwlte.LTEBandAll,
			}); err != nil {
				return xerrors.Errorf("failed to set network mode: %w", err)
			}

			return nil
		},
	}
)
