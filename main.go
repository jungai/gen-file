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
				Name:    "editorconfig",
				Aliases: []string{"e"},
				Usage:   "generate editorconfig config file (.editorconfig)",
				Action: func(c *cli.Context) error {
					if c.Args().Len() < 2 {
						fmt.Println("required <config type> <template> <path>")

						return nil
					}
					value := c.Args().Slice()

					template := value[0]
					templatePath := fmt.Sprintf("./templates/editorconfig/%v", template)
					targetPath := fmt.Sprintf("%s/.editorconfig", value[1])

					content, err := os.ReadFile(templatePath)

					if err != nil {
						fmt.Println("editorconfig template not found !!")

						return nil
					}

					err = os.WriteFile(targetPath, content, 0644)

					if err != nil {
						fmt.Println("cannot write file")

						return nil
					}

					return nil
				},
			},
			{
				// TODO: add action
				Name:    "prettier",
				Aliases: []string{"p"},
				Usage:   "generate prettier config file (.prettier)",
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
