package main

import (
	"runtime/debug"

	"github.com/alexfalkowski/bezeichner/internal/cmd"
	sc "github.com/alexfalkowski/go-service/cmd"
)

func main() {
	command().ExitOnError()
}

func command() *sc.Command {
	info, _ := debug.ReadBuildInfo()

	command := sc.New(info.Main.Version)
	command.RegisterInput(command.Root(), "env:BEZEICHNER_CONFIG_FILE")

	cmd.RegisterServer(command)

	return command
}
