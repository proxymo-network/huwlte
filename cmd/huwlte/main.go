package main

import (
	"context"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/anexia-it/go-human"
	"github.com/proxymo-network/huwlte"
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"
)

func newClient(c *cli.Context) (*huwlte.Client, error) {
	doer := http.DefaultClient

	proxy := c.String("proxy")
	if proxy != "" {
		proxyURL, err := url.Parse(proxy)
		if err != nil {
			return nil, xerrors.Errorf("failed to parse proxy URL: %w", err)
		}

		doer = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxyURL),
			},
		}
	}

	return huwlte.NewClient(c.String("url"), huwlte.WithDoer(doer)), nil
}

func pretty(v interface{}) error {
	encoder, err := human.NewEncoder(os.Stdout)
	if err != nil {
		return xerrors.Errorf("failed to create encoder: %w", err)
	}
	return encoder.Encode(v)
}

func main() {
	ctx := context.Background()

	app := &cli.App{
		Name:  "huwlte",
		Usage: "simple CLI for manage Huawei LTE dongle",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "url",
				Usage:   "URL of Huawei LTE dongle",
				EnvVars: []string{"HUWLTE_URL"},
				Value:   "http://192.168.8.1",
			},
			&cli.StringFlag{
				Name:    "proxy",
				Aliases: []string{"x"},
				Usage:   "URL of proxy server",
				EnvVars: []string{"HUWLTE_PROXY"},
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "device",
				Usage: "contains device related commands",
				Subcommands: []*cli.Command{
					{
						Name:    "basic-information",
						Aliases: []string{"info"},
						Usage:   "get basic information of device",
						Action: func(c *cli.Context) error {
							client, err := newClient(c)
							if err != nil {
								return xerrors.Errorf("failed to create client: %w", err)
							}

							info, err := client.Device.BasicInformation(ctx)
							if err != nil {
								return xerrors.Errorf("failed to get basic information: %w", err)
							}

							return pretty(info)
						},
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
