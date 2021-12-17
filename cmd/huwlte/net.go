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
			netCurrentPLMN,
		},
	}

	netCurrentPLMN = &cli.Command{
		Name:    "current-plmn",
		Aliases: []string{"plmn"},
		Usage:   "get current carrier info",
		Action: func(c *cli.Context) error {
			ctx, cancel := newCtx(c)
			defer cancel()

			client, cancel, err := newClient(ctx, c)
			if err != nil {
				return xerrors.Errorf("failed to create client: %w", err)
			}
			defer cancel()

			info, err := client.Net.CurrentPLMN(ctx)
			if err != nil {
				return xerrors.Errorf("failed to get basic information: %w", err)
			}

			return pretty(info)
		},
	}

	netModeCmd = &cli.Command{
		Name:    "mode",
		Aliases: []string{"mode"},
		Usage:   "configure connection mode",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "mode",
				Usage: "set network mode",
			},
			&cli.StringFlag{
				Name:  "band",
				Usage: "set network band",
			},
		},
		Action: func(c *cli.Context) error {
			mode, band := c.String("mode"), c.String("band")

			ctx, cancel := newCtx(c)
			defer cancel()

			client, cancel, err := newClient(ctx, c)
			if err != nil {
				return xerrors.Errorf("failed to create client: %w", err)
			}
			defer cancel()

			if mode == "" && band == "" {
				modem, err := client.Net.Mode(ctx)
				if err != nil {
					return xerrors.Errorf("failed to get network mode: %w", err)
				}

				return pretty(modem)
			} else {
				netMode, err := huwlte.ParseNetworkMode(mode)
				if err != nil {
					return xerrors.Errorf("failed to parse network mode: %w", err)
				}

				netBand, err := huwlte.ParseNetworkBand(band)
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
			}

			return nil
		},
	}
)
