package main

import (
	"os"

	"go.h4n.io/centra/component-base/cli"
	"go.h4n.io/centra/tenancy/cmd/tenancy/app"
)

func main() {
    cmd := app.NewTenancyCommand()
    code := cli.Run(cmd)
    os.Exit(code)
}

