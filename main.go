package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	allConfigs, err := os.ReadDir("./templates")

	if err != nil {
		log.Fatalf("cannot read directory")
	}

	app := &cli.App{
		Commands: []*cli.Command{
			{
				// TODO: add action
				Name:    "prettier",
				Aliases: []string{"p"},
				Usage:   "generate preitter config file (.prettier)",
				Action: func(c *cli.Context) error {
					fmt.Println(c.Args().First())

					return nil
				},
			},
			{
				// TODO: add action
				Name:    "editorconfig",
				Aliases: []string{"p"},
				Usage:   "generate editorconfig config file (.editorconfig)",
			},
			{
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "show all available extension",
				Action: func(c *cli.Context) error {
					for i := 0; i < len(allConfigs); i++ {
						fmt.Println(allConfigs[i].Name())
					}

					return nil
				},
			},
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
