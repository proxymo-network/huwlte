package main

import (
	"github.com/proxymo-network/huwlte"
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"
)

var (
	smsCmd = &cli.Command{
		Name:  "sms",
		Usage: "contains SMS related commands",
		Subcommands: []*cli.Command{
			smsCount,
			smsList,
		},
	}

	smsList = &cli.Command{
		Name:  "list",
		Usage: "list sms messages",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "page-index",
				Aliases: []string{"p"},
				Usage:   "page index",
				Value:   1,
			},
			&cli.IntFlag{
				Name:    "read-count",
				Aliases: []string{"r"},
				Usage:   "read count",
				Value:   20,
			},
			&cli.IntFlag{
				Name:    "box-type",
				Aliases: []string{"b"},
				Usage:   "box type",
				Value:   1,
			},
			&cli.IntFlag{
				Name:    "sort-type",
				Aliases: []string{"s"},
				Usage:   "sort type",
				Value:   0,
			},
			&cli.IntFlag{
				Name:    "ascending",
				Aliases: []string{"a"},
				Usage:   "ascending",
				Value:   0,
			},
			&cli.BoolFlag{
				Name:    "unread-preferred",
				Aliases: []string{"u"},
				Usage:   "unread preferred",
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

			var unreadPreffer int
			if c.Bool("unread-preferred") {
				unreadPreffer = 1
			}

			info, err := client.SMS.List(ctx, huwlte.SMSListOptions{
				PageIndex:       c.Int("page-index"),
				ReadCount:       c.Int("read-count"),
				BoxType:         c.Int("box-type"),
				SortType:        c.Int("sort-type"),
				Ascending:       c.Int("ascending"),
				UnreadPreferred: unreadPreffer,
			})
			if err != nil {
				return xerrors.Errorf("failed to get sms count: %w", err)
			}

			return pretty(info)
		},
	}

	smsCount = &cli.Command{
		Name:  "count",
		Usage: "get sms counters",
		Action: func(c *cli.Context) error {
			ctx, cancel := newCtx(c)
			defer cancel()

			client, cancel, err := newClient(ctx, c)
			if err != nil {
				return xerrors.Errorf("failed to create client: %w", err)
			}
			defer cancel()

			info, err := client.SMS.Count(ctx)
			if err != nil {
				return xerrors.Errorf("failed to get sms count: %w", err)
			}

			return pretty(info)
		},
	}
)
