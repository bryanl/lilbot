// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1alpha1

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

// LilbotPluginServiceClient is the client API for LilbotPluginService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LilbotPluginServiceClient interface {
	// Invoke a command.
	Invoke(ctx context.Context, in *InvokeRequest, opts ...grpc.CallOption) (*InvokeResponse, error)
}

type lilbotPluginServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLilbotPluginServiceClient(cc grpc.ClientConnInterface) LilbotPluginServiceClient {
	return &lilbotPluginServiceClient{cc}
}

func (c *lilbotPluginServiceClient) Invoke(ctx context.Context, in *InvokeRequest, opts ...grpc.CallOption) (*InvokeResponse, error) {
	out := new(InvokeResponse)
	err := c.cc.Invoke(ctx, "/plugin.v1alpha1.LilbotPluginService/Invoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LilbotPluginServiceServer is the server API for LilbotPluginService service.
// All implementations must embed UnimplementedLilbotPluginServiceServer
// for forward compatibility
type LilbotPluginServiceServer interface {
	// Invoke a command.
	Invoke(context.Context, *InvokeRequest) (*InvokeResponse, error)
	mustEmbedUnimplementedLilbotPluginServiceServer()
}

// UnimplementedLilbotPluginServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLilbotPluginServiceServer struct {
}

func (UnimplementedLilbotPluginServiceServer) Invoke(context.Context, *InvokeRequest) (*InvokeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Invoke not implemented")
}
func (UnimplementedLilbotPluginServiceServer) mustEmbedUnimplementedLilbotPluginServiceServer() {}

// UnsafeLilbotPluginServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LilbotPluginServiceServer will
// result in compilation errors.
type UnsafeLilbotPluginServiceServer interface {
	mustEmbedUnimplementedLilbotPluginServiceServer()
}

func RegisterLilbotPluginServiceServer(s grpc.ServiceRegistrar, srv LilbotPluginServiceServer) {
	s.RegisterService(&LilbotPluginService_ServiceDesc, srv)
}

func _LilbotPluginService_Invoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvokeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LilbotPluginServiceServer).Invoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/plugin.v1alpha1.LilbotPluginService/Invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LilbotPluginServiceServer).Invoke(ctx, req.(*InvokeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LilbotPluginService_ServiceDesc is the grpc.ServiceDesc for LilbotPluginService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LilbotPluginService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "plugin.v1alpha1.LilbotPluginService",
	HandlerType: (*LilbotPluginServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Invoke",
			Handler:    _LilbotPluginService_Invoke_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "plugin/v1alpha1/plugin.proto",
}
