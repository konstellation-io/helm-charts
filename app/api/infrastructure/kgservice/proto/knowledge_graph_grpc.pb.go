// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package kgpb

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

// KGServiceClient is the client API for KGService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KGServiceClient interface {
	GetGraph(ctx context.Context, in *GetGraphReq, opts ...grpc.CallOption) (*GetGraphRes, error)
	GetItem(ctx context.Context, in *GetItemReq, opts ...grpc.CallOption) (*GetItemRes, error)
	GetDescriptionQuality(ctx context.Context, in *DescriptionQualityReq, opts ...grpc.CallOption) (*DescriptionQualityRes, error)
}

type kGServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKGServiceClient(cc grpc.ClientConnInterface) KGServiceClient {
	return &kGServiceClient{cc}
}

func (c *kGServiceClient) GetGraph(ctx context.Context, in *GetGraphReq, opts ...grpc.CallOption) (*GetGraphRes, error) {
	out := new(GetGraphRes)
	err := c.cc.Invoke(ctx, "/kg.KGService/GetGraph", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kGServiceClient) GetItem(ctx context.Context, in *GetItemReq, opts ...grpc.CallOption) (*GetItemRes, error) {
	out := new(GetItemRes)
	err := c.cc.Invoke(ctx, "/kg.KGService/GetItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kGServiceClient) GetDescriptionQuality(ctx context.Context, in *DescriptionQualityReq, opts ...grpc.CallOption) (*DescriptionQualityRes, error) {
	out := new(DescriptionQualityRes)
	err := c.cc.Invoke(ctx, "/kg.KGService/GetDescriptionQuality", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KGServiceServer is the server API for KGService service.
// All implementations must embed UnimplementedKGServiceServer
// for forward compatibility
type KGServiceServer interface {
	GetGraph(context.Context, *GetGraphReq) (*GetGraphRes, error)
	GetItem(context.Context, *GetItemReq) (*GetItemRes, error)
	GetDescriptionQuality(context.Context, *DescriptionQualityReq) (*DescriptionQualityRes, error)
	mustEmbedUnimplementedKGServiceServer()
}

// UnimplementedKGServiceServer must be embedded to have forward compatible implementations.
type UnimplementedKGServiceServer struct {
}

func (UnimplementedKGServiceServer) GetGraph(context.Context, *GetGraphReq) (*GetGraphRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGraph not implemented")
}
func (UnimplementedKGServiceServer) GetItem(context.Context, *GetItemReq) (*GetItemRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetItem not implemented")
}
func (UnimplementedKGServiceServer) GetDescriptionQuality(context.Context, *DescriptionQualityReq) (*DescriptionQualityRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDescriptionQuality not implemented")
}
func (UnimplementedKGServiceServer) mustEmbedUnimplementedKGServiceServer() {}

// UnsafeKGServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KGServiceServer will
// result in compilation errors.
type UnsafeKGServiceServer interface {
	mustEmbedUnimplementedKGServiceServer()
}

func RegisterKGServiceServer(s grpc.ServiceRegistrar, srv KGServiceServer) {
	s.RegisterService(&KGService_ServiceDesc, srv)
}

func _KGService_GetGraph_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGraphReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KGServiceServer).GetGraph(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kg.KGService/GetGraph",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KGServiceServer).GetGraph(ctx, req.(*GetGraphReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _KGService_GetItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetItemReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KGServiceServer).GetItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kg.KGService/GetItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KGServiceServer).GetItem(ctx, req.(*GetItemReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _KGService_GetDescriptionQuality_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescriptionQualityReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KGServiceServer).GetDescriptionQuality(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kg.KGService/GetDescriptionQuality",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KGServiceServer).GetDescriptionQuality(ctx, req.(*DescriptionQualityReq))
	}
	return interceptor(ctx, in, info, handler)
}

// KGService_ServiceDesc is the grpc.ServiceDesc for KGService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KGService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kg.KGService",
	HandlerType: (*KGServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGraph",
			Handler:    _KGService_GetGraph_Handler,
		},
		{
			MethodName: "GetItem",
			Handler:    _KGService_GetItem_Handler,
		},
		{
			MethodName: "GetDescriptionQuality",
			Handler:    _KGService_GetDescriptionQuality_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/knowledge_graph.proto",
}
