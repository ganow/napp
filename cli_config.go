package main

import (
	"github.com/codegangsta/cli"
)

func setupCli() {
	cli.CommandHelpTemplate = `NAME:
   {{.Name}}

USAGE:
   {{.Usage}}

DESCRIPTION:
   {{.Description}}

OPTIONS:
   {{range .Flags}}{{.}}
   {{end}}
`
}
