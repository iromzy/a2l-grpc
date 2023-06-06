// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.1
// source: protobufs/API.proto

package a2l

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
	A2L_GetTree_FullMethodName = "/A2L/GetTree"
)

// A2LClient is the client API for A2L service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type A2LClient interface {
	GetTree(ctx context.Context, in *GetTreeRequest, opts ...grpc.CallOption) (*RootNodeType, error)
}

type a2LClient struct {
	cc grpc.ClientConnInterface
}

func NewA2LClient(cc grpc.ClientConnInterface) A2LClient {
	return &a2LClient{cc}
}

func (c *a2LClient) GetTree(ctx context.Context, in *GetTreeRequest, opts ...grpc.CallOption) (*RootNodeType, error) {
	out := new(RootNodeType)
	err := c.cc.Invoke(ctx, A2L_GetTree_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// A2LServer is the server API for A2L service.
// All implementations must embed UnimplementedA2LServer
// for forward compatibility
type A2LServer interface {
	GetTree(context.Context, *GetTreeRequest) (*RootNodeType, error)
	mustEmbedUnimplementedA2LServer()
}

// UnimplementedA2LServer must be embedded to have forward compatible implementations.
type UnimplementedA2LServer struct {
}

func (UnimplementedA2LServer) GetTree(context.Context, *GetTreeRequest) (*RootNodeType, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTree not implemented")
}
func (UnimplementedA2LServer) mustEmbedUnimplementedA2LServer() {}

// UnsafeA2LServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to A2LServer will
// result in compilation errors.
type UnsafeA2LServer interface {
	mustEmbedUnimplementedA2LServer()
}

func RegisterA2LServer(s grpc.ServiceRegistrar, srv A2LServer) {
	s.RegisterService(&A2L_ServiceDesc, srv)
}

func _A2L_GetTree_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTreeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(A2LServer).GetTree(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: A2L_GetTree_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(A2LServer).GetTree(ctx, req.(*GetTreeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// A2L_ServiceDesc is the grpc.ServiceDesc for A2L service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var A2L_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "A2L",
	HandlerType: (*A2LServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTree",
			Handler:    _A2L_GetTree_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobufs/API.proto",
}