package commands

import (
	"github.com/cloudogu/nexus-scripting/manager"
	"github.com/urfave/cli"
)

func createManager(context *cli.Context) (*manager.Manager, error) {
	url := context.GlobalString("url")
	if url == "" {
		return nil, cli.NewExitError("url is required", 1)
	}

	username := context.GlobalString("username")
	if username == "" {
		return nil, cli.NewExitError("username is required", 1)
	}

	password := context.GlobalString("password")
	if password == "" {
		return nil, cli.NewExitError("password is required", 1)
	}

	return manager.New(url, username, password), nil
}
