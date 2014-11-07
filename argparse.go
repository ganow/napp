package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func argParse(command_name string, c *cli.Context) (fname string, neuron_index []int) {
	args := c.Args()
	argc := len(args)

	if argc > 2 {
		logError("too many arguments: %s\n", args)
		cli.ShowCommandHelp(c, "fr")
		os.Exit(-1)
	} else if argc == 0 {
		cli.ShowCommandHelp(c, "fr")
		os.Exit(-1)
	} else if argc == 1 {
		fname = args.Get(0)
	} else if argc == 2 {
		fname = args.Get(1)
	}
	neuron_index = []int{1, 2, 3}
	return fname, neuron_index
}
