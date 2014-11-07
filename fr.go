package main

import (
	"github.com/codegangsta/cli"
)

func doFr(c *cli.Context) {
	fname, _ := argParse("fr", c)
	logStatus("file name:\t%s\n", fname)
}
