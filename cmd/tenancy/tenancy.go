package main

import (
	"os"

	"go.h4n.io/centra/component-base/cli"
	"go.h4n.io/centra/tenancy/cmd/tenancy/app"
)

var version = `v0.0.0`

func main() {
    cmd := app.NewTenancyCommand(version)
    code := cli.Run(cmd)
    os.Exit(code)
}

