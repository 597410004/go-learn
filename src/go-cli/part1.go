package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

//运行方式: go run part1.go -help
/*NAME:
greet - fight the loneliness!

USAGE:
greet [global options] command [command options] [arguments...]

VERSION:
0.0.0

COMMANDS:
help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS
--version Shows version information*/
func main() {
	app := cli.NewApp()
	app.Name = "greet"
	app.Usage = "fight the loneliness!"
	app.Action = func(c *cli.Context) error {
		fmt.Println("Hello friend!")
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
