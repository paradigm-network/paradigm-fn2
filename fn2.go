package main

import (
	"log"
	"os"

	"github.com/urfave/cli"

	"github.com/paradigm-network/paradigm-fn2/server"
)

func main() {
	app := cli.NewApp()
	app.Name = "fn2"
	app.Usage = "make function as a service"
	app.Version = "0.1"

	app.Commands = []cli.Command{
		{
			Name:  "serve",
			Usage: "start fn2 server on current host",
			Action: func(c *cli.Context) error {
				return server.Start(true)
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
