package migrate

import (
	"embed"
	"github.com/compico/shopsite/internal/cmd/commands"
	_ "github.com/compico/shopsite/internal/database/migrations"
	"github.com/pressly/goose/v3"
	"github.com/urfave/cli/v2"
)

//go:embed migrations/*.go
var migrations embed.FS

func init() {
	commands.RegisterCommand(
		&cli.Command{
			Name: "migrate",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "config",
					Value:   "configs/default.yaml",
					Usage:   "Path to config",
					Aliases: []string{"c"},
				},
			},
			Action: func(ctx *cli.Context) error {
				migrator, err := InitializeMigrator(
					ctx.String("config"),
					[]string{"multiStatements=true"},
				)
				if err != nil {
					return nil
				}

				goose.SetBaseFS(migrations)
				err = goose.SetDialect(migrator.conf.GetDialect())
				if err != nil {
					return err
				}

				return goose.RunContext(
					ctx.Context,
					ctx.Args().First(),
					migrator.connection,
					"migrations",
					ctx.Args().Slice()[1:]...,
				)
			},
		},
	)
}
