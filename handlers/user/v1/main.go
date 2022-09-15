package v1

import (
	"context"
	"time"

	userv1 "go.h4n.io/centra/tenancy/api/user/v1"
	"k8s.io/klog/v2"
)

type UserService struct {
    *userv1.UnimplementedUserServiceServer
}

func (us UserService) ListUsers(_ context.Context, r *userv1.ListUsersRequest) (*userv1.ListUsersResponse, error) {
    klog.Info("call ListUsers")
    resp := userv1.ListUsersResponse{
        Users: []*userv1.User{
            {
                Id: `a0a7dbb9-e688-4b3b-a301-d57263aa1d22`,
                DisplayName: `Hayden Young`,
                Email: `hayden@hbjy.dev`,
                CreatedAt: time.Now().Format(time.RFC3339),
                UpdatedAt: time.Now().Format(time.RFC3339),
            },
            {
                Id: `a0a7dbb9-e688-4b3b-a301-d57263aa1d22`,
                DisplayName: `John Doe`,
                Email: `j.doe@centra.app`,
                CreatedAt: time.Now().Format(time.RFC3339),
                UpdatedAt: time.Now().Format(time.RFC3339),
            },
        },
    }

    return &resp, nil
}

func (s UserService) CreateUser(_ context.Context, r *userv1.CreateUserRequest) (*userv1.CreateUserResponse, error) {
    klog.Info("call CreateUser")
    resp := userv1.CreateUserResponse{
        User: &userv1.User{
            Id: `a0a7dbb9-e688-4b3b-a301-d57263aa1d22`,
            DisplayName: `Hayden Young`,
            Email: r.Email,
            CreatedAt: time.Now().Format(time.RFC3339),
            UpdatedAt: time.Now().Format(time.RFC3339),
        },
    }

    return &resp, nil
}

