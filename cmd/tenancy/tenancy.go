package main

import (
	"os"

	"go.h4n.io/centra/component-base/cli"
	"go.h4n.io/centra/tenancy/cmd/tenancy/app"
)

// @title Centra Tenancy Service
// @version 1.0.0
// @description The multitenancy service of Centra, holding tenant and user information.
// @host localhost:8080
// @license.name Apache 2.0
// @accept json
// @produce json

var version = `v0.0.0`

func main() {
    cmd := app.NewTenancyCommand(version)
    code := cli.Run(cmd)
    os.Exit(code)
}

