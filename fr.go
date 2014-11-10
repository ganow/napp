package main

import (
	"github.com/codegangsta/cli"
)

func doFr(c *cli.Context) {
	fname, neurons, group_num := argParse("fr", c)
	logStatus("file name:\t%s\n", fname)
	logStatus("neurons:\t%s\n", neurons)
	logStatus("group num:\t%d\n", group_num)
}
