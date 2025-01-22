// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: api_server.proto

package api

import (
	context "context"
	types "github.com/hyle-team/tss-svc/internal/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	API_SubmitWithdrawal_FullMethodName = "/api.API/SubmitWithdrawal"
	API_CheckWithdrawal_FullMethodName  = "/api.API/CheckWithdrawal"
)

// APIClient is the client API for API service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type APIClient interface {
	SubmitWithdrawal(ctx context.Context, in *types.DepositIdentifier, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CheckWithdrawal(ctx context.Context, in *types.DepositIdentifier, opts ...grpc.CallOption) (*CheckWithdrawalResponse, error)
}

type aPIClient struct {
	cc grpc.ClientConnInterface
}

func NewAPIClient(cc grpc.ClientConnInterface) APIClient {
	return &aPIClient{cc}
}

func (c *aPIClient) SubmitWithdrawal(ctx context.Context, in *types.DepositIdentifier, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, API_SubmitWithdrawal_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) CheckWithdrawal(ctx context.Context, in *types.DepositIdentifier, opts ...grpc.CallOption) (*CheckWithdrawalResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CheckWithdrawalResponse)
	err := c.cc.Invoke(ctx, API_CheckWithdrawal_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// APIServer is the server API for API service.
// All implementations should embed UnimplementedAPIServer
// for forward compatibility.
type APIServer interface {
	SubmitWithdrawal(context.Context, *types.DepositIdentifier) (*emptypb.Empty, error)
	CheckWithdrawal(context.Context, *types.DepositIdentifier) (*CheckWithdrawalResponse, error)
}

// UnimplementedAPIServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAPIServer struct{}

func (UnimplementedAPIServer) SubmitWithdrawal(context.Context, *types.DepositIdentifier) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitWithdrawal not implemented")
}
func (UnimplementedAPIServer) CheckWithdrawal(context.Context, *types.DepositIdentifier) (*CheckWithdrawalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckWithdrawal not implemented")
}
func (UnimplementedAPIServer) testEmbeddedByValue() {}

// UnsafeAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to APIServer will
// result in compilation errors.
type UnsafeAPIServer interface {
	mustEmbedUnimplementedAPIServer()
}

func RegisterAPIServer(s grpc.ServiceRegistrar, srv APIServer) {
	// If the following call pancis, it indicates UnimplementedAPIServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&API_ServiceDesc, srv)
}

func _API_SubmitWithdrawal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.DepositIdentifier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).SubmitWithdrawal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: API_SubmitWithdrawal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).SubmitWithdrawal(ctx, req.(*types.DepositIdentifier))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_CheckWithdrawal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.DepositIdentifier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).CheckWithdrawal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: API_CheckWithdrawal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).CheckWithdrawal(ctx, req.(*types.DepositIdentifier))
	}
	return interceptor(ctx, in, info, handler)
}

// API_ServiceDesc is the grpc.ServiceDesc for API service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var API_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.API",
	HandlerType: (*APIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitWithdrawal",
			Handler:    _API_SubmitWithdrawal_Handler,
		},
		{
			MethodName: "CheckWithdrawal",
			Handler:    _API_CheckWithdrawal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api_server.proto",
}
