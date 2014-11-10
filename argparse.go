package main

import (
	"github.com/codegangsta/cli"
	"os"
	"strconv"
	"strings"
)

func argParse(command_name string, c *cli.Context) (fnames string, neurons []Neuron, group_num int) {
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
		neurons, group_num = genNeuronIndex(c.Int("N"), c.Bool("each"))
	} else if argc == 2 {
		fname = args.Get(1)
		neurons, group_num = parseNeuronIndex(args.Get(0), c.Bool("each"))
	}
	return fname, neurons, group_num
}

func genNeuronIndex(N int, is_each bool) (neurons []Neuron, group_num int) {
	for i := 0; i < N; i++ {
		if is_each {
			neurons = append(neurons, Neuron{
				id:       i,
				x:        0.0,
				x2:       0.0,
				group_id: i,
			})
		} else {
			neurons = append(neurons, Neuron{
				id:       i,
				x:        0.0,
				x2:       0.0,
				group_id: 0,
			})
		}
	}
	if is_each {
		group_num = N
	} else {
		group_num = 1
	}
	return neurons, group_num
}

func parseNeuronIndex(s string, is_each bool) (neurons []Neuron, group_num int) {
	group_by := strings.Split(s, "%")

	for group_id, group := range group_by {
		for _, tmp := range strings.Split(group, ",") {
			if strings.Contains(tmp, "-") {
				tmp2 := strings.Split(tmp, "-")
				lower, _ := strconv.Atoi(tmp2[0])
				upper, _ := strconv.Atoi(tmp2[1])
				for i := lower; i <= upper; i++ {
					if is_each {
						neurons = append(neurons, Neuron{
							id:       i,
							x:        0.0,
							x2:       0.0,
							group_id: i,
						})
					} else {
						neurons = append(neurons, Neuron{
							id:       i,
							x:        0.0,
							x2:       0.0,
							group_id: group_id,
						})
					}
				}
			} else {
				id, _ := strconv.Atoi(tmp)
				if is_each {
					neurons = append(neurons, Neuron{
						id:       id,
						x:        0.0,
						x2:       0.0,
						group_id: id,
					})
				} else {
					neurons = append(neurons, Neuron{
						id:       id,
						x:        0.0,
						x2:       0.0,
						group_id: group_id,
					})
				}
			}
		}
	}
	if is_each {
		group_num = len(neurons)
	} else {
		group_num = len(group_by)
	}

	return neurons, group_num
}
