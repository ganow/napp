package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	setupCli()

	app := cli.NewApp()
	app.Name = "napp"
	app.Version = Version
	app.Usage = "neural activity processing pipeline"
	app.Author = "Yoshihiro Nagano"
	app.Email = "y.nagano.92@gmail.com"
	app.Commands = Commands

	app.Run(os.Args)
}
