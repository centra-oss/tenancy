package app

import (
	"log"

	"github.com/spf13/cobra"
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
