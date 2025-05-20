package main

import (
	"context"

	"github.com/alexfalkowski/bezeichner/internal/cmd"
	"github.com/alexfalkowski/go-service/v2/cli"
)

var app = cli.NewApplication(func(command cli.Commander) {
	cmd.RegisterServer(command)
})

func main() {
	app.ExitOnError(context.Background())
}
