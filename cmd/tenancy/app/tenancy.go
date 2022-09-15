package app

import (
	"log"
	"net/http"
	"time"

	"github.com/spf13/cobra"
	"go.h4n.io/centra/tenancy/util/notfoundhandler"
	"go.h4n.io/centra/component-base/healthz"
)

func NewTenancyCommand(version string) *cobra.Command {
	cmd := &cobra.Command{
		Short: `tenancy`,
        Version: version,
        SilenceUsage: true, // hide usage on error
		RunE: func(_ *cobra.Command, _ []string) error {
			log.Println("starting tenancy service...")

            m := http.NewServeMux()

            healthz.SetUp("ready for connections")

            notfound := notfoundhandler.New()

            m.HandleFunc("/-/healthz", healthz.Handler)

            m.Handle("/", notfound)

            s := http.Server{
                Addr: `:8080`,
                Handler: m,
                ReadTimeout: 10 * time.Second,
                WriteTimeout: 10 * time.Second,
                MaxHeaderBytes: 1 << 20,
            }
            
			return s.ListenAndServe()
		},
	}

	return cmd
}

