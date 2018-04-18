package commands

import (
	"path"

	"fmt"

	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

func init() {
	register(cli.Command{
		Name:   "execute",
		Usage:  "upload and execute the given script",
		Action: Execute,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "name",
				Usage: "name of the nexus script",
			},
			cli.StringFlag{
				Name:  "payload",
				Usage: "payload argument for execution",
			},
			cli.StringFlag{
				Name:  "file-payload",
				Usage: "path to file which is used as payload argument for execution",
			},
		},
	})
}

// Execute upload and executes the given script
func Execute(context *cli.Context) error {
	manager, err := createManager(context)
	if err != nil {
		return err
	}

	filename := context.Args().First()
	if filename == "" {
		return cli.NewExitError("filename argument is required", 1)
	}

	name := context.String("name")
	if name == "" {
		name = strings.TrimSuffix(path.Base(filename), filepath.Ext(filename))
	}

	script, err := manager.CreateFromFile(name, filename)
	if err != nil {
		return errors.Wrapf(err, "failed to create script %s from %s", name, filename)
	}

	var result string
	filePayload := context.String("file-payload")
	payload := context.String("payload")

	if filePayload != "" && payload != "" {
		return cli.NewExitError("only one of the parameters can be used: payload or file-payload", 2)
	}

	if filePayload != "" {
		result, err = script.ExecuteWithFilePayload(filePayload)
		if err != nil {
			return errors.Wrapf(err, "execution of script %s with filePayload %s failed", filename, filePayload)
		}
	} else if payload != "" {
		result, err = script.ExecuteWithStringPayload(payload)
		if err != nil {
			return errors.Wrapf(err, "execution of script %s failed", filename)
		}
	} else {
		result, err = script.ExecuteWithoutPayload()
		if err != nil {
			return errors.Wrapf(err, "execution of script %s failed", filename)
		}
	}

	fmt.Println(result)

	return nil
}
