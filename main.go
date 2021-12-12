package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "gen-file",
		Usage: "generate file for you!!",
		Action: func(c *cli.Context) error {
			fmt.Println("I will do this trust me!")
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
