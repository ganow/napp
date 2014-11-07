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
	Usage: "calculate mean firing rate from rst file",
	Description: `
`,
	Action: doFr,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "separator, s",
			Value: "\t",
			Usage: "value separator for rst file",
		},
		cli.Float64Flag{
			Name:  "from, f",
			Value: 0,
			Usage: "aggregate values from #[ms]",
		},
		cli.Float64Flag{
			Name:  "to, t",
			Value: -1,
			Usage: "aggregate values to #[ms]. -1 means 'until end'",
		},
		cli.Float64Flag{
			Name:  "binsize, b",
			Value: -1,
			Usage: "aggregate values for each #[ms]. -1 means gather them altogether",
		},
	},
}

var commandFf = cli.Command{
	Name:  "ff",
	Usage: "",
	Description: `
`,
	Action: doFf,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "separator, s",
			Value: "\t",
			Usage: "value separator for rst file",
		},
		cli.Float64Flag{
			Name:  "from, f",
			Value: 0,
			Usage: "aggregate values from #[ms]",
		},
		cli.Float64Flag{
			Name:  "to, t",
			Value: -1,
			Usage: "aggregate values to #[ms]. -1 means 'until end'",
		},
		cli.Float64Flag{
			Name:  "binsize, b",
			Value: -1,
			Usage: "aggregate values for each #[ms]. -1 means gather them altogether",
		},
	},
}

var commandCorr = cli.Command{
	Name:  "corr",
	Usage: "",
	Description: `
`,
	Action: doCorr,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "separator, s",
			Value: "\t",
			Usage: "value separator for rst file",
		},
		cli.Float64Flag{
			Name:  "from, f",
			Value: 0,
			Usage: "aggregate values from #[ms]",
		},
		cli.Float64Flag{
			Name:  "to, t",
			Value: -1,
			Usage: "aggregate values to #[ms]. -1 means 'until end'",
		},
		cli.Float64Flag{
			Name:  "binsize, b",
			Value: -1,
			Usage: "aggregate values for each #[ms]. -1 means gather them altogether",
		},
		cli.IntFlag{
			Name:  "aggregate-by, a",
			Value: 1,
			Usage: "spike count correlations between # and rest args",
		},
	},
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
