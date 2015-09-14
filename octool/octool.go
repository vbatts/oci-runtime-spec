package main

import (
	"os"

	"github.com/codegangsta/cli"
)

var (
	registeredCommands = []cli.Command{
		{
			Name:  "validate",
			Usage: "validation of sorts",
			Subcommands: []cli.Command{
				{
					Name:   "config",
					Usage:  "validate the config.json",
					Action: cmdValidateConfig,
				},
				{
					Name:   "env",
					Usage:  "validate the environment of a container runtime",
					Action: cmdValidateEnv,
				},
			},
		},
	}
)

func main() {
	app := cli.NewApp()
	app.Name = os.Args[0]
	app.Usage = "OCI/specs utility"
	app.Flags = []cli.Flag{}
	app.Commands = registeredCommands
	app.Run(os.Args)
}
