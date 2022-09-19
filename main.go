package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	flags, cfg := NewConfig()
	app := &cli.App{
		Name:  "boom",
		Usage: "make an explosive entrance",
		Flags: flags,
		Action: func(ctx *cli.Context) error {
			return Run(ctx, cfg)
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
