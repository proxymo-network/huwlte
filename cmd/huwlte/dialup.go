package main

import (
	"fmt"

	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"
)

var (
	dialupCmd = &cli.Command{
		Name:  "dialup",
		Usage: "contains dialup related commands",
		Subcommands: []*cli.Command{
			dialupMobileSwitch,
		},
	}

	dialupMobileSwitch = &cli.Command{
		Name:    "mobile-switch",
		Aliases: []string{"ms"},
		Usage:   "get or change status of mobile connection",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "on",
				Usage: "turn on mobile switch",
			},
			&cli.BoolFlag{
				Name:  "off",
				Usage: "turn off mobile switch",
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

			switch {
			case c.Bool("on"):
				if err := client.Dialup.SetMobileSwitch(ctx, true); err != nil {
					return xerrors.Errorf("set mobile switch: %w", err)
				}
			case c.Bool("off"):
				if err := client.Dialup.SetMobileSwitch(ctx, false); err != nil {
					return xerrors.Errorf("set mobile switch: %w", err)
				}
			default:
				enabled, err := client.Dialup.MobileSwitch(ctx)
				if err != nil {
					return xerrors.Errorf("failed to get basic information: %w", err)
				}
				printSwitch(c, enabled)
			}

			return nil
		},
	}
)

func printSwitch(c *cli.Context, v bool) {
	if c.Bool("no-emoji") {
		printSwitchNoEmoji(c, v)
	} else {
		printSwitchEmoji(c, v)
	}
}

func printSwitchNoEmoji(c *cli.Context, v bool) {
	if v {
		fmt.Println("ON")
	} else {
		fmt.Println("OFF")
	}
}

func printSwitchEmoji(c *cli.Context, v bool) {
	if v {
		fmt.Println("✅ ON")
	} else {
		fmt.Println("🔴 OFF")
	}
}
