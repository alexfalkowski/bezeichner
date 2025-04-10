package main

import (
	"github.com/alexfalkowski/bezeichner/internal/cmd"
	sc "github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/env"
)

func main() {
	command().ExitOnError()
}

func command() *sc.Command {
	command := sc.New(env.NewName(), env.NewVersion())

	cmd.RegisterServer(command)

	return command
}
