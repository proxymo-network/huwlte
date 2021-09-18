package main

import (
	"context"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"time"

	"github.com/anexia-it/go-human"
	"github.com/proxymo-network/huwlte"
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"
)

func newClient(ctx context.Context, c *cli.Context) (*huwlte.Client, context.CancelFunc, error) {
	doer := http.DefaultClient

	proxy := c.String("proxy")
	if proxy != "" {
		proxyURL, err := url.Parse(proxy)
		if err != nil {
			return nil, nil, xerrors.Errorf("failed to parse proxy URL: %w", err)
		}

		doer = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxyURL),
			},
		}
	}

	opts := []huwlte.ClientOpt{
		huwlte.WithDoer(doer),
	}

	if session := c.String("session"); session != "" {
		storage := huwlte.LocalSessionStorage{
			Dir: c.String("session-dir"),
		}

		opts = append(opts, huwlte.WithStorage(session, &storage))
	}

	client := huwlte.NewClient(c.String("url"), opts...)

	if err := client.LoadSession(ctx); err != nil {
		return nil, nil, xerrors.Errorf("load session: %w", err)
	}

	return client, func() {
		if err := client.SaveSession(ctx); err != nil {
			log.Printf("failed to save session: %v", err)
		}
	}, nil
}

func pretty(v interface{}) error {
	encoder, err := human.NewEncoder(os.Stdout)
	if err != nil {
		return xerrors.Errorf("failed to create encoder: %w", err)
	}
	return encoder.Encode(v)
}

func main() {

	cacheDir, err := os.UserCacheDir()
	if err != nil {
		log.Fatal(err)
	}

	sessionDir := path.Join(cacheDir, "huwlte")

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
			&cli.DurationFlag{
				Name:    "timeout",
				Aliases: []string{"t"},
				Usage:   "timeout for HTTP requests",
				EnvVars: []string{"HUWLTE_TIMEOUT"},
				Value:   30 * time.Second,
			},
			&cli.StringFlag{
				Name:    "session",
				Aliases: []string{"s"},
				Usage:   "session name",
			},
			&cli.PathFlag{
				Name:    "session-dir",
				Usage:   "store session in a directory",
				EnvVars: []string{"HUWLTE_SESSION_DIR"},
				Value:   sessionDir,
			},
			&cli.BoolFlag{
				Name:    "reset-session",
				Aliases: []string{"r"},
				Usage:   "erase all session data",
			},
			&cli.BoolFlag{
				Name:  "no-emoji",
				Usage: "disable emoji",
			},
		},
		Commands: []*cli.Command{
			userCmd,
			deviceCmd,
			monitoringCmd,
			dialupCmd,
			rawCmd,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
