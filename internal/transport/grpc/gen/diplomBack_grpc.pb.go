// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: diplomBack.proto

package mainBack

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

const (
	MainBack_Login_FullMethodName    = "/mainBack.mainBack/Login"
	MainBack_GetRouts_FullMethodName = "/mainBack.mainBack/GetRouts"
)

// MainBackClient is the client API for MainBack service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MainBackClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	GetRouts(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*RequestsCount, error)
}

type mainBackClient struct {
	cc grpc.ClientConnInterface
}

func NewMainBackClient(cc grpc.ClientConnInterface) MainBackClient {
	return &mainBackClient{cc}
}

func (c *mainBackClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, MainBack_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mainBackClient) GetRouts(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*RequestsCount, error) {
	out := new(RequestsCount)
	err := c.cc.Invoke(ctx, MainBack_GetRouts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MainBackServer is the server API for MainBack service.
// All implementations must embed UnimplementedMainBackServer
// for forward compatibility
type MainBackServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	GetRouts(context.Context, *EmptyRequest) (*RequestsCount, error)
	mustEmbedUnimplementedMainBackServer()
}

// UnimplementedMainBackServer must be embedded to have forward compatible implementations.
type UnimplementedMainBackServer struct {
}

func (UnimplementedMainBackServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedMainBackServer) GetRouts(context.Context, *EmptyRequest) (*RequestsCount, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRouts not implemented")
}
func (UnimplementedMainBackServer) mustEmbedUnimplementedMainBackServer() {}

// UnsafeMainBackServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MainBackServer will
// result in compilation errors.
type UnsafeMainBackServer interface {
	mustEmbedUnimplementedMainBackServer()
}

func RegisterMainBackServer(s grpc.ServiceRegistrar, srv MainBackServer) {
	s.RegisterService(&MainBack_ServiceDesc, srv)
}

func _MainBack_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainBackServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MainBack_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainBackServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MainBack_GetRouts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainBackServer).GetRouts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MainBack_GetRouts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainBackServer).GetRouts(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MainBack_ServiceDesc is the grpc.ServiceDesc for MainBack service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MainBack_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mainBack.mainBack",
	HandlerType: (*MainBackServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _MainBack_Login_Handler,
		},
		{
			MethodName: "GetRouts",
			Handler:    _MainBack_GetRouts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "diplomBack.proto",
}
