package app

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"go.h4n.io/centra/component-base/logs"
	"go.h4n.io/centra/tenancy/handlers/v1/tenants"
)

var (
	ginLogger gin.HandlerFunc
)

func NewTenancyCommand(version string) *cobra.Command {
	cmd := &cobra.Command{
		Short:        `tenancy`,
		Version:      version,
		SilenceUsage: true, // hide usage on error

		PersistentPreRun: func(_ *cobra.Command, _ []string) {
			loggerConfig := gin.LoggerConfig{
				Formatter: func(params gin.LogFormatterParams) string {
					b := strings.Builder{}
					b.WriteString(fmt.Sprintf("%s - ", params.ClientIP))
					b.WriteString(fmt.Sprintf("%v ", params.Latency))
					b.WriteString(fmt.Sprintf(`"%s %s" `, params.Method, params.Path))
					b.WriteString(fmt.Sprintf("%v ", params.StatusCode))
					b.WriteString(fmt.Sprintf("%v ", params.BodySize))
					b.WriteString(fmt.Sprintf(`"%v" `, params.Request.UserAgent()))

					return b.String()
				},
				Output:    logs.KlogWriter{},
				SkipPaths: []string{},
			}

			ginLogger = gin.LoggerWithConfig(loggerConfig)

            gin.SetMode(gin.ReleaseMode)
		},

		RunE: func(_ *cobra.Command, _ []string) error {
			log.Println("starting tenancy service...")

			r := gin.New()
			r.Use(ginLogger)
			r.Use(gin.Recovery())

			r.GET("/v1/tenants", tenants.TenantsIndex)

			s := http.Server{
				Addr:           `:8080`,
				Handler:        r,
				ReadTimeout:    10 * time.Second,
				WriteTimeout:   10 * time.Second,
				MaxHeaderBytes: 1 << 20,
			}

			return s.ListenAndServe()
		},
	}

	return cmd
}
