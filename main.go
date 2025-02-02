package main

import (
	"os"

	"github.com/alexfalkowski/bezeichner/cmd"
	sc "github.com/alexfalkowski/go-service/cmd"
)

func main() {
	if err := command().Run(); err != nil {
		os.Exit(1)
	}
}

func command() *sc.Command {
	c := sc.New(cmd.Version)
	c.RegisterInput(c.Root(), "env:BEZEICHNER_CONFIG_FILE")
	c.AddServer("server", "Start bezeichner server", cmd.ServerOptions...)

	return c
}
