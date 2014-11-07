package main

import (
	"github.com/codegangsta/cli"
)

func setupCli() {
	cli.CommandHelpTemplate = `NAME:
   {{.Name}} - {{.Usage}}

USAGE:
   napp {{.Name}} [command options] [neuron index] file_name

DESCRIPTION:
   {{.Description}}

OPTIONS:
   {{range .Flags}}{{.}}
   {{end}}
`
}
