package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "greet",
		Usage: "A simple CLI app with urfave/cli",
		Commands: []*cli.Command{
			{
				Name:  "hello",
				Usage: "Prints a greeting",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "name",
						Usage: "Your name",
					},
				},
				Action: func(c *cli.Context) error {
					name := c.String("name")
					if name == "" {
						name = "World"
					}
					fmt.Printf("Hello, %s!\n", name)
					return nil
				},
			},
		},
	}

	// Run the app
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
