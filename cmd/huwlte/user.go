package main

import (
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"
)

var (
	userCmd = &cli.Command{
		Name:  "user",
		Usage: "contains auth related commands",
		Subcommands: []*cli.Command{
			userLoginCmd,
			userStateLoginCmd,
		},
	}

	userLoginCmd = &cli.Command{
		Name:  "login",
		Usage: "login to the modem ui",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "username",
				Usage: "modem admin username",
				Value: "admin",
			},
			&cli.StringFlag{
				Name:     "password",
				Usage:    "modem admin password",
				Required: true,
			},
			&cli.BoolFlag{
				Name:  "relogin",
				Usage: "force relogin if session exists",
				Value: false,
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

			if err := client.User.Login(ctx, c.String("username"), c.String("password"), false); err != nil {
				return xerrors.Errorf("failed to login: %w", err)
			}

			return nil
		},
	}

	userStateLoginCmd = &cli.Command{
		Name:    "state-login",
		Aliases: []string{"state"},
		Usage:   "get current login state",
		Action: func(c *cli.Context) error {
			ctx, cancel := newCtx(c)
			defer cancel()

			client, cancel, err := newClient(ctx, c)
			if err != nil {
				return xerrors.Errorf("failed to create client: %w", err)
			}
			defer cancel()

			info, err := client.User.StateLogin(ctx)
			if err != nil {
				return xerrors.Errorf("failed to get basic information: %w", err)
			}

			return pretty(info)
		},
	}
)
