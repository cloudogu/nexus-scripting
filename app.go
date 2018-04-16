package main

import (
	"log"
	"os"

	"github.com/cloudogu/nexus-scripting/commands"
	"github.com/urfave/cli"
)

var (
	// Version of the application
	Version string
)

func main() {
	app := cli.NewApp()
	app.Name = "nexus-scripting"
	app.Usage = "manage nexus 3 scripts"
	app.Version = Version
	app.Commands = commands.Commands
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "url",
			EnvVar: "NEXUS_URL",
			Usage:  "url to nexus server (with context path)",
		},
		cli.StringFlag{
			Name:   "username",
			EnvVar: "NEXUS_USER",
			Usage:  "username for nexus authentication",
		},
		cli.StringFlag{
			Name:   "password",
			EnvVar: "NEXUS_PASSWORD",
			Usage:  "password for nexus authentication",
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
