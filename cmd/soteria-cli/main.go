package main

import (
	"Soteria/internal/scli"
	"Soteria/pkg/log"
	"fmt"
	"github.com/urfave/cli/v2"
)


func main() {
	log, err := log.CreateNamed(log.Trace, "CLI")
	if err != nil {
		fmt.Println("Failed to initialize logger")
		panic(err)
	}
	log.Trace("Initialized logger")

	app := cli.NewApp()
	app.EnableBashCompletion = true

	scli := scli.CreateClic(log, app)
	scli.RegisterCreateActions()
}

