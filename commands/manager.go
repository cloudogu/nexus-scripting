package commands

import (
	"github.com/cloudogu/nexus-scripting/manager"
	"github.com/urfave/cli"
)

func createManager(c *cli.Context) (*manager.Manager, error) {
	url := c.GlobalString("url")
	if url == "" {
		return nil, cli.NewExitError("url is required", 1)
	}

	username := c.GlobalString("username")
	if username == "" {
		return nil, cli.NewExitError("username is required", 1)
	}

	password := c.GlobalString("password")
	if password == "" {
		return nil, cli.NewExitError("password is required", 1)
	}

	return manager.New(url, username, password), nil
}
