package v1

import (
	"context"
	"time"

	tenantv1 "go.h4n.io/centra/tenancy/api/tenant/v1"
	"k8s.io/klog/v2"
)

type TenantService struct {
    *tenantv1.UnimplementedTenantServiceServer
}

func (us TenantService) ListTenants(_ context.Context, r *tenantv1.ListTenantsRequest) (*tenantv1.ListTenantsResponse, error) {
    klog.Info("call ListTenants")
    resp := tenantv1.ListTenantsResponse{
        Tenants: []*tenantv1.Tenant{
            {
                Id: `a0a7dbb9-e688-4b3b-a301-d57263aa1d22`,
                DisplayName: `Centra Internal`,
                CreatedAt: time.Now().Format(time.RFC3339),
                UpdatedAt: time.Now().Format(time.RFC3339),
            },
        },
    }

    return &resp, nil
}

