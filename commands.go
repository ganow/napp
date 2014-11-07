package main

import (
	"github.com/codegangsta/cli"
	"log"
	"os"
)

var Commands = []cli.Command{
	commandFr,
	commandFf,
	commandCorr,
}

var commandFr = cli.Command{
	Name:  "fr",
	Usage: "",
	Description: `
`,
	Action: doFr,
}

var commandFf = cli.Command{
	Name:  "ff",
	Usage: "",
	Description: `
`,
	Action: doFf,
}

var commandCorr = cli.Command{
	Name:  "corr",
	Usage: "",
	Description: `
`,
	Action: doCorr,
}

func debug(v ...interface{}) {
	if os.Getenv("DEBUG") != "" {
		log.Println(v...)
	}
}

func assert(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func doFr(c *cli.Context) {
}

func doFf(c *cli.Context) {
}

func doCorr(c *cli.Context) {
}
