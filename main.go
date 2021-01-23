package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "template-go-cli"
	app.Usage = "This is cli template with go lang urfave"
	app.Version = "0.0.1"

	app.Action = func(context *cli.Context) error {
		fmt.Println(context.Args().Get(0))
		return nil
	}

	app.Run(os.Args)
}
