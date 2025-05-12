package main

import (
	"context"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v3"
)

func main() {
	app := &cli.Command{
		Name:  "mysql-connectivity",
		Usage: "MySQL connectivity check utility",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "user",
				Destination: &mysql_user,
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "pass",
				Destination: &mysql_pass,
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "addr",
				Destination: &mysql_addr,
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "dbname",
				Destination: &mysql_dbname,
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "command",
				Destination: &mysql_command,
				Required:    true,
			},
		},
		Action: func(context.Context, *cli.Command) error {
			execute()
			return nil
		},
	}

	err := app.Run(context.Background(), os.Args)
	if err != nil {
		log.Fatal(err.Error())
	}
}
