package main

import (
	"fmt"
	"github.com/urfave/cli/v3"
	"log"
	"os"
)

func main() {
	app := &cli.Command{
		Name: "tasks",
		Commands: []*cli.Command{
			{
				Name:  "list",
				Usage: "List all tasks",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    "completed",
						Aliases: []string{"c"},
						Usage:   "Show completed tasks",
					},
				},
				Action: func(c *cli.Context) error {
					if c.Bool("completed") {
						fmt.Println("Completed tasks:")
					} else {
						fmt.Println("Pending tasks:")
					}
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
