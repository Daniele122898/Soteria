package scli

import (
	"Soteria/pkg/log"

	"github.com/urfave/cli/v2"
)

type Scli struct {
	log log.Log
	app *cli.App
}

// Create CLI Context
func CreateClic(log log.Log, app *cli.App) *Scli {
	return &Scli{
		log: log,
		app: app,
	}
}

func (clic *Scli) RegisterCreateCommands() {
	//clic.app.Commands = append(clic.app.Commands, &cli.Command{})
}
