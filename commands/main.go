package commands

import (
	"github.com/urfave/cli"
)

var Commands []cli.Command

func register(command cli.Command) {
	Commands = append(Commands, command)
}
