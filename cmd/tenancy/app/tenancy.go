package app

import (
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/go-grpc-middleware/providers/openmetrics/v2"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
	tenantpbv1 "go.h4n.io/centra/tenancy/api/tenant/v1"
	userpbv1 "go.h4n.io/centra/tenancy/api/user/v1"
	tenantv1 "go.h4n.io/centra/tenancy/handlers/tenant/v1"
	userv1 "go.h4n.io/centra/tenancy/handlers/user/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"k8s.io/klog/v2"
)

func NewTenancyCommand(version string) *cobra.Command {
	cmd := &cobra.Command{
		Short:        `tenancy`,
		Version:      version,
		SilenceUsage: true, // hide usage on error

		RunE: func(_ *cobra.Command, _ []string) error {
			log.Println("starting tenancy service...")

            // Open port 8080
            l, err := net.Listen("tcp", ":8080")
            if err != nil {
                return err
            }

            // Configure gRPC server
            serverMetrics := metrics.NewServerMetrics()
            s := grpc.NewServer(
                grpc.ChainStreamInterceptor(
                    metrics.StreamServerInterceptor(serverMetrics),
                ),
                grpc.ChainUnaryInterceptor(
                    metrics.UnaryServerInterceptor(serverMetrics),
                ),
            )

            // Register services
            userpbv1.RegisterUserServiceServer(s, userv1.UserService{})
            tenantpbv1.RegisterTenantServiceServer(s, tenantv1.TenantService{})

            // Configure schema reflection
            reflection.Register(s)

            go func() {
                s := http.Server{
                    Addr: ":8081",
                    Handler: promhttp.Handler(),
                }
                if err := s.ListenAndServe(); err != nil {
                    klog.Fatal(err)
                }
            }()

            return s.Serve(l)
		},
	}

	return cmd
}
