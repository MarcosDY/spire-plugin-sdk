// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package test

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

// SomeHostServiceClient is the client API for SomeHostService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SomeHostServiceClient interface {
	HostServiceEcho(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error)
}

type someHostServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSomeHostServiceClient(cc grpc.ClientConnInterface) SomeHostServiceClient {
	return &someHostServiceClient{cc}
}

func (c *someHostServiceClient) HostServiceEcho(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error) {
	out := new(EchoResponse)
	err := c.cc.Invoke(ctx, "/test.SomeHostService/HostServiceEcho", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SomeHostServiceServer is the server API for SomeHostService service.
// All implementations must embed UnimplementedSomeHostServiceServer
// for forward compatibility
type SomeHostServiceServer interface {
	HostServiceEcho(context.Context, *EchoRequest) (*EchoResponse, error)
	mustEmbedUnimplementedSomeHostServiceServer()
}

// UnimplementedSomeHostServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSomeHostServiceServer struct {
}

func (UnimplementedSomeHostServiceServer) HostServiceEcho(context.Context, *EchoRequest) (*EchoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HostServiceEcho not implemented")
}
func (UnimplementedSomeHostServiceServer) mustEmbedUnimplementedSomeHostServiceServer() {}

// UnsafeSomeHostServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SomeHostServiceServer will
// result in compilation errors.
type UnsafeSomeHostServiceServer interface {
	mustEmbedUnimplementedSomeHostServiceServer()
}

func RegisterSomeHostServiceServer(s grpc.ServiceRegistrar, srv SomeHostServiceServer) {
	s.RegisterService(&SomeHostService_ServiceDesc, srv)
}

func _SomeHostService_HostServiceEcho_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SomeHostServiceServer).HostServiceEcho(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.SomeHostService/HostServiceEcho",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SomeHostServiceServer).HostServiceEcho(ctx, req.(*EchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SomeHostService_ServiceDesc is the grpc.ServiceDesc for SomeHostService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SomeHostService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "test.SomeHostService",
	HandlerType: (*SomeHostServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HostServiceEcho",
			Handler:    _SomeHostService_HostServiceEcho_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test/somehostservice.proto",
}