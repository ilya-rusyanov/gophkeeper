// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: gophkeeper.proto

package proto

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
	Gophkeeper_Register_FullMethodName = "/gophkeeper.Gophkeeper/Register"
	Gophkeeper_LogIn_FullMethodName    = "/gophkeeper.Gophkeeper/LogIn"
	Gophkeeper_Store_FullMethodName    = "/gophkeeper.Gophkeeper/Store"
	Gophkeeper_List_FullMethodName     = "/gophkeeper.Gophkeeper/List"
	Gophkeeper_Show_FullMethodName     = "/gophkeeper.Gophkeeper/Show"
)

// GophkeeperClient is the client API for Gophkeeper service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GophkeeperClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*Empty, error)
	LogIn(ctx context.Context, in *LogInRequest, opts ...grpc.CallOption) (*Empty, error)
	Store(ctx context.Context, in *StoreRequest, opts ...grpc.CallOption) (*Empty, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Show(ctx context.Context, in *ShowRequest, opts ...grpc.CallOption) (*ShowResponse, error)
}

type gophkeeperClient struct {
	cc grpc.ClientConnInterface
}

func NewGophkeeperClient(cc grpc.ClientConnInterface) GophkeeperClient {
	return &gophkeeperClient{cc}
}

func (c *gophkeeperClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Gophkeeper_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophkeeperClient) LogIn(ctx context.Context, in *LogInRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Gophkeeper_LogIn_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophkeeperClient) Store(ctx context.Context, in *StoreRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Gophkeeper_Store_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophkeeperClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, Gophkeeper_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophkeeperClient) Show(ctx context.Context, in *ShowRequest, opts ...grpc.CallOption) (*ShowResponse, error) {
	out := new(ShowResponse)
	err := c.cc.Invoke(ctx, Gophkeeper_Show_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GophkeeperServer is the server API for Gophkeeper service.
// All implementations must embed UnimplementedGophkeeperServer
// for forward compatibility
type GophkeeperServer interface {
	Register(context.Context, *RegisterRequest) (*Empty, error)
	LogIn(context.Context, *LogInRequest) (*Empty, error)
	Store(context.Context, *StoreRequest) (*Empty, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
	Show(context.Context, *ShowRequest) (*ShowResponse, error)
	mustEmbedUnimplementedGophkeeperServer()
}

// UnimplementedGophkeeperServer must be embedded to have forward compatible implementations.
type UnimplementedGophkeeperServer struct {
}

func (UnimplementedGophkeeperServer) Register(context.Context, *RegisterRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedGophkeeperServer) LogIn(context.Context, *LogInRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogIn not implemented")
}
func (UnimplementedGophkeeperServer) Store(context.Context, *StoreRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Store not implemented")
}
func (UnimplementedGophkeeperServer) List(context.Context, *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedGophkeeperServer) Show(context.Context, *ShowRequest) (*ShowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Show not implemented")
}
func (UnimplementedGophkeeperServer) mustEmbedUnimplementedGophkeeperServer() {}

// UnsafeGophkeeperServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GophkeeperServer will
// result in compilation errors.
type UnsafeGophkeeperServer interface {
	mustEmbedUnimplementedGophkeeperServer()
}

func RegisterGophkeeperServer(s grpc.ServiceRegistrar, srv GophkeeperServer) {
	s.RegisterService(&Gophkeeper_ServiceDesc, srv)
}

func _Gophkeeper_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophkeeperServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gophkeeper_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophkeeperServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gophkeeper_LogIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophkeeperServer).LogIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gophkeeper_LogIn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophkeeperServer).LogIn(ctx, req.(*LogInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gophkeeper_Store_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophkeeperServer).Store(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gophkeeper_Store_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophkeeperServer).Store(ctx, req.(*StoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gophkeeper_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophkeeperServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gophkeeper_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophkeeperServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gophkeeper_Show_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophkeeperServer).Show(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gophkeeper_Show_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophkeeperServer).Show(ctx, req.(*ShowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gophkeeper_ServiceDesc is the grpc.ServiceDesc for Gophkeeper service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gophkeeper_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gophkeeper.Gophkeeper",
	HandlerType: (*GophkeeperServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Gophkeeper_Register_Handler,
		},
		{
			MethodName: "LogIn",
			Handler:    _Gophkeeper_LogIn_Handler,
		},
		{
			MethodName: "Store",
			Handler:    _Gophkeeper_Store_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Gophkeeper_List_Handler,
		},
		{
			MethodName: "Show",
			Handler:    _Gophkeeper_Show_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gophkeeper.proto",
}
