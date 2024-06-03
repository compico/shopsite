package http

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/compico/shopsite/cmd/commands"
	"github.com/urfave/cli/v2"
)

func init() {
	commands.RegisterCommand(
		&cli.Command{
			Name: "http",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "config",
					Value:   "configs/default.yaml",
					Usage:   "Path to config",
					Aliases: []string{"c"},
				},
			},
			Action: func(context *cli.Context) error {
				app, err := InitializeApp(context.String("config"))
				if err != nil {
					return err
				}

				go app.server.Start()

				sigCh := make(chan os.Signal, 1)
				signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
				<-sigCh

				app.server.Stop()
				return nil
			},
		},
	)
}
