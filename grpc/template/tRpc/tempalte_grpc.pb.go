// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package tRpc

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

// TemplateRpcServiceClient is the client API for TemplateRpcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TemplateRpcServiceClient interface {
	AddTemplate(ctx context.Context, in *TemplateRpcAdd, opts ...grpc.CallOption) (*Success, error)
	List(ctx context.Context, in *PageList, opts ...grpc.CallOption) (*Templates, error)
}

type templateRpcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTemplateRpcServiceClient(cc grpc.ClientConnInterface) TemplateRpcServiceClient {
	return &templateRpcServiceClient{cc}
}

func (c *templateRpcServiceClient) AddTemplate(ctx context.Context, in *TemplateRpcAdd, opts ...grpc.CallOption) (*Success, error) {
	out := new(Success)
	err := c.cc.Invoke(ctx, "/TemplateRpcService/AddTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *templateRpcServiceClient) List(ctx context.Context, in *PageList, opts ...grpc.CallOption) (*Templates, error) {
	out := new(Templates)
	err := c.cc.Invoke(ctx, "/TemplateRpcService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TemplateRpcServiceServer is the server API for TemplateRpcService service.
// All implementations must embed UnimplementedTemplateRpcServiceServer
// for forward compatibility
type TemplateRpcServiceServer interface {
	AddTemplate(context.Context, *TemplateRpcAdd) (*Success, error)
	List(context.Context, *PageList) (*Templates, error)
	mustEmbedUnimplementedTemplateRpcServiceServer()
}

// UnimplementedTemplateRpcServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTemplateRpcServiceServer struct {
}

func (UnimplementedTemplateRpcServiceServer) AddTemplate(context.Context, *TemplateRpcAdd) (*Success, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTemplate not implemented")
}
func (UnimplementedTemplateRpcServiceServer) List(context.Context, *PageList) (*Templates, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedTemplateRpcServiceServer) mustEmbedUnimplementedTemplateRpcServiceServer() {}

// UnsafeTemplateRpcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TemplateRpcServiceServer will
// result in compilation errors.
type UnsafeTemplateRpcServiceServer interface {
	mustEmbedUnimplementedTemplateRpcServiceServer()
}

func RegisterTemplateRpcServiceServer(s grpc.ServiceRegistrar, srv TemplateRpcServiceServer) {
	s.RegisterService(&TemplateRpcService_ServiceDesc, srv)
}

func _TemplateRpcService_AddTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TemplateRpcAdd)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateRpcServiceServer).AddTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TemplateRpcService/AddTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateRpcServiceServer).AddTemplate(ctx, req.(*TemplateRpcAdd))
	}
	return interceptor(ctx, in, info, handler)
}

func _TemplateRpcService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PageList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateRpcServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TemplateRpcService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateRpcServiceServer).List(ctx, req.(*PageList))
	}
	return interceptor(ctx, in, info, handler)
}

// TemplateRpcService_ServiceDesc is the grpc.ServiceDesc for TemplateRpcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TemplateRpcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "TemplateRpcService",
	HandlerType: (*TemplateRpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddTemplate",
			Handler:    _TemplateRpcService_AddTemplate_Handler,
		},
		{
			MethodName: "List",
			Handler:    _TemplateRpcService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tempalte.proto",
}
