package main

import (
	"github.com/compico/shopsite/cmd/commands"
	_ "github.com/compico/shopsite/cmd/commands/http"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := cli.App{
		Commands: commands.Commands,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
