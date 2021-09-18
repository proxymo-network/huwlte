package main

import (
	"bytes"
	"fmt"

	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"
)

var (
	rawCmd = &cli.Command{
		Name:  "raw",
		Usage: "Raw request",
		Subcommands: []*cli.Command{
			rawGetCmd,
		},
	}

	rawGetCmd = &cli.Command{
		Name:      "get",
		Usage:     "GET request",
		ArgsUsage: "<path>",
		Action: func(c *cli.Context) error {
			if !c.Args().Present() {
				return xerrors.New("path is required")
			}

			ctx, cancel := newCtx(c)
			defer cancel()

			client, cancel, err := newClient(ctx, c)
			if err != nil {
				return xerrors.Errorf("failed to create client: %w", err)
			}
			defer cancel()

			buf := bytes.Buffer{}

			if err := client.Get(ctx, c.Args().First(), &buf); err != nil {
				return xerrors.Errorf("get raw: %w", err)
			}

			fmt.Println(buf.String())

			return nil
		},
	}
)
