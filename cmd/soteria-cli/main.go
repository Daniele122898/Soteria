package main

import (
	"Soteria/internal/scli"
	"Soteria/pkg/log"
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	log, err := log.CreateNamed(log.Trace, "CLI")
	if err != nil {
		fmt.Println("Failed to initialize logger")
		panic(err)
	}

	app := &cli.App{
		Name:    "Soteria cli",
		Version: "v0.0.1",
		Authors: []*cli.Author{
			&cli.Author{
				Name:  "Serenity",
				Email: "support@argonaut.pw",
			},
		},
		Copyright: "(c) 2022 Serenity at Argonaut Developments",
		Usage:     "Cli to interface with Soteria",
	}
	app.EnableBashCompletion = true

	scli := scli.CreateClic(log, app)
	scli.RegisterTestCommands()
	scli.RegisterCreateCommands()

	app.Run(os.Args)
}
