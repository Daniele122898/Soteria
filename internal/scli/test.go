package scli

import "github.com/urfave/cli/v2"

func (clic *Scli) RegisterTestCommands() {
	clic.app.Commands = append(clic.app.Commands, &cli.Command{
		Name: "test",
		Aliases: []string{"t"},
		Category: "testing",
		Usage: "test the system",
		UsageText: "tests the system test",
		Description: "Why is there so much description",
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "extended", Aliases: []string{"ext"}},
		},
		Subcommands: []*cli.Command{
			&cli.Command{
				Name: "subtest",
				Aliases: []string{"subt"},
				Action: func(context *cli.Context) error {
					clic.log.Info("Ran sub test command")
					return nil
				},
			},
		},
		Action: func(context *cli.Context) error {
			clic.log.Info("Ran actual test command")
			clic.log.Info("flags %v", context.FlagNames())
			if context.NArg() > 0 {
				clic.log.Info("Arg 0: %v", context.Args().Get(0))
			}
			clic.log.Info("string: %v", context.String("extended"))
			return nil
		},
	})
}
