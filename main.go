package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"

	"github.com/egsam98/voting/cli/cmd/kafkaproducer"
)

func main() {
	app := cli.App{
		Commands: cli.Commands{
			kafkaproducer.Cmd,
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Printf("error: %+v\n", err)
		os.Exit(1)
	}
}
