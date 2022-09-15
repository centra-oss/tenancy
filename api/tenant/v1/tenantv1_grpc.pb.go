// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: tenant/v1/tenantv1.proto

package tenantv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TenantServiceClient is the client API for TenantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TenantServiceClient interface {
	// ListTenants lists all the tenants a given User is associated with
	ListTenants(ctx context.Context, in *ListTenantsRequest, opts ...grpc.CallOption) (*ListTenantsResponse, error)
}

type tenantServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTenantServiceClient(cc grpc.ClientConnInterface) TenantServiceClient {
	return &tenantServiceClient{cc}
}

func (c *tenantServiceClient) ListTenants(ctx context.Context, in *ListTenantsRequest, opts ...grpc.CallOption) (*ListTenantsResponse, error) {
	out := new(ListTenantsResponse)
	err := c.cc.Invoke(ctx, "/tenant.v1.TenantService/ListTenants", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TenantServiceServer is the server API for TenantService service.
// All implementations should embed UnimplementedTenantServiceServer
// for forward compatibility
type TenantServiceServer interface {
	// ListTenants lists all the tenants a given User is associated with
	ListTenants(context.Context, *ListTenantsRequest) (*ListTenantsResponse, error)
}

// UnimplementedTenantServiceServer should be embedded to have forward compatible implementations.
type UnimplementedTenantServiceServer struct {
}

func (UnimplementedTenantServiceServer) ListTenants(context.Context, *ListTenantsRequest) (*ListTenantsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTenants not implemented")
}

// UnsafeTenantServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TenantServiceServer will
// result in compilation errors.
type UnsafeTenantServiceServer interface {
	mustEmbedUnimplementedTenantServiceServer()
}

func RegisterTenantServiceServer(s grpc.ServiceRegistrar, srv TenantServiceServer) {
	s.RegisterService(&TenantService_ServiceDesc, srv)
}

func _TenantService_ListTenants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTenantsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServiceServer).ListTenants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tenant.v1.TenantService/ListTenants",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServiceServer).ListTenants(ctx, req.(*ListTenantsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TenantService_ServiceDesc is the grpc.ServiceDesc for TenantService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TenantService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tenant.v1.TenantService",
	HandlerType: (*TenantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListTenants",
			Handler:    _TenantService_ListTenants_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tenant/v1/tenantv1.proto",
}