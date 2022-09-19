package main

import (
	"github.com/urfave/cli/v2"
	"testtask/model"
)

type Config struct {
	FilePath     string
	Filter       *model.Filter
	amountWorker int
}

func NewConfig() ([]cli.Flag, *Config) {
	cfg := Config{
		amountWorker: 10,
		Filter:       &model.Filter{},
	}

	flags := []cli.Flag{
		&cli.StringFlag{
			Name:        "filePath",
			Aliases:     []string{"f"},
			Usage:       "path to source file",
			Destination: &cfg.FilePath,
			Required:    true,
		},
		&cli.StringFlag{
			Name:        "genres",
			Aliases:     []string{"g"},
			Usage:       "provide genre for find",
			Destination: &cfg.Filter.Genre,
		},
	}

	return flags, &cfg
}
