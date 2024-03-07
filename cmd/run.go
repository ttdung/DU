package cmd

import (
	"context"
	"fmt"
	"github.com/ttdung/du/config"
	"github.com/ttdung/du/internal"
	cli "github.com/urfave/cli/v2"
	"log"
	"path/filepath"
)

const (
	flagConfig = "config"
)

func Run(args []string) error {
	cliApp := &cli.App{
		Name:                 filepath.Base(args[0]),
		Usage:                "DU",
		Version:              "v0.0.1",
		Copyright:            "??",
		EnableBashCompletion: true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    flagConfig,
				Value:   "./config.json",
				Usage:   "The config file to load from",
				EnvVars: []string{"CONFIG_FILE"},
			},
		},
		Action: func(ctx *cli.Context) error {
			if args := ctx.Args(); args.Len() > 0 {
				return fmt.Errorf("unexpected arguments: %q", args.Get(0))
			}

			// Prepare FileConfig
			configPath := ctx.String(flagConfig)
			cfg, err := config.LoadConfigFromFile(configPath)
			if err != nil {
				log.Printf("failed to load config from file %v: %v", configPath, err)
				tmpCfg := config.DefaultConfig()
				cfg = &tmpCfg
			}

			mainApp, err := internal.NewApp(cfg)
			if err != nil {
				return err
			}

			mainApp.Start(context.Background())

			return nil
		},
	}

	err := cliApp.Run(args)
	if err != nil {
		return err
	}

	return nil
}
