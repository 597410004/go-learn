package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
)

//自定义lang命令
func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "lang",
			Value: "chinese",
			Usage: "language for the greeting",
		},
	}

	app.Action = func(c *cli.Context) error {
		name := "二狗"
		if c.NArg() > 0 {
			name = c.Args().Get(0)
		}
		if c.String("lang") == "chinese" {
			fmt.Println("你好", name)
		} else {
			fmt.Println("侬好", name)
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
