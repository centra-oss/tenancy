package main

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"go.h4n.io/centra/component-base/cli"
)

func NewTenancyCommand() *cobra.Command {
    cmd := &cobra.Command{
        Short: `tenancy`,
        RunE: func(cmd *cobra.Command, args []string) error {
            log.Println("Testing tenancy CLI logging")

            return nil
        },
    }

    return cmd
}

func main() {
    cmd := NewTenancyCommand()
    code := cli.Run(cmd)
    os.Exit(code)
}

