// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: proto/v1/dictionary_item.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	DictionaryItemService_Show_FullMethodName   = "/proto.v1.DictionaryItemService/Show"
	DictionaryItemService_Create_FullMethodName = "/proto.v1.DictionaryItemService/Create"
	DictionaryItemService_Update_FullMethodName = "/proto.v1.DictionaryItemService/Update"
	DictionaryItemService_Delete_FullMethodName = "/proto.v1.DictionaryItemService/Delete"
)

// DictionaryItemServiceClient is the client API for DictionaryItemService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DictionaryItemServiceClient interface {
	Show(ctx context.Context, in *DictionaryItemShowRequest, opts ...grpc.CallOption) (*DictionaryItemShowResponse, error)
	Create(ctx context.Context, in *DictionaryItemCreateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Update(ctx context.Context, in *DictionaryItemUpdateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Delete(ctx context.Context, in *DictionaryItemDeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type dictionaryItemServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDictionaryItemServiceClient(cc grpc.ClientConnInterface) DictionaryItemServiceClient {
	return &dictionaryItemServiceClient{cc}
}

func (c *dictionaryItemServiceClient) Show(ctx context.Context, in *DictionaryItemShowRequest, opts ...grpc.CallOption) (*DictionaryItemShowResponse, error) {
	out := new(DictionaryItemShowResponse)
	err := c.cc.Invoke(ctx, DictionaryItemService_Show_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictionaryItemServiceClient) Create(ctx context.Context, in *DictionaryItemCreateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, DictionaryItemService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictionaryItemServiceClient) Update(ctx context.Context, in *DictionaryItemUpdateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, DictionaryItemService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictionaryItemServiceClient) Delete(ctx context.Context, in *DictionaryItemDeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, DictionaryItemService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DictionaryItemServiceServer is the server API for DictionaryItemService service.
// All implementations must embed UnimplementedDictionaryItemServiceServer
// for forward compatibility
type DictionaryItemServiceServer interface {
	Show(context.Context, *DictionaryItemShowRequest) (*DictionaryItemShowResponse, error)
	Create(context.Context, *DictionaryItemCreateRequest) (*emptypb.Empty, error)
	Update(context.Context, *DictionaryItemUpdateRequest) (*emptypb.Empty, error)
	Delete(context.Context, *DictionaryItemDeleteRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedDictionaryItemServiceServer()
}

// UnimplementedDictionaryItemServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDictionaryItemServiceServer struct {
}

func (UnimplementedDictionaryItemServiceServer) Show(context.Context, *DictionaryItemShowRequest) (*DictionaryItemShowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Show not implemented")
}
func (UnimplementedDictionaryItemServiceServer) Create(context.Context, *DictionaryItemCreateRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedDictionaryItemServiceServer) Update(context.Context, *DictionaryItemUpdateRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedDictionaryItemServiceServer) Delete(context.Context, *DictionaryItemDeleteRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedDictionaryItemServiceServer) mustEmbedUnimplementedDictionaryItemServiceServer() {}

// UnsafeDictionaryItemServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DictionaryItemServiceServer will
// result in compilation errors.
type UnsafeDictionaryItemServiceServer interface {
	mustEmbedUnimplementedDictionaryItemServiceServer()
}

func RegisterDictionaryItemServiceServer(s grpc.ServiceRegistrar, srv DictionaryItemServiceServer) {
	s.RegisterService(&DictionaryItemService_ServiceDesc, srv)
}

func _DictionaryItemService_Show_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DictionaryItemShowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DictionaryItemServiceServer).Show(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DictionaryItemService_Show_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DictionaryItemServiceServer).Show(ctx, req.(*DictionaryItemShowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DictionaryItemService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DictionaryItemCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DictionaryItemServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DictionaryItemService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DictionaryItemServiceServer).Create(ctx, req.(*DictionaryItemCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DictionaryItemService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DictionaryItemUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DictionaryItemServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DictionaryItemService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DictionaryItemServiceServer).Update(ctx, req.(*DictionaryItemUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DictionaryItemService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DictionaryItemDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DictionaryItemServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DictionaryItemService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DictionaryItemServiceServer).Delete(ctx, req.(*DictionaryItemDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DictionaryItemService_ServiceDesc is the grpc.ServiceDesc for DictionaryItemService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DictionaryItemService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.v1.DictionaryItemService",
	HandlerType: (*DictionaryItemServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Show",
			Handler:    _DictionaryItemService_Show_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _DictionaryItemService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _DictionaryItemService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _DictionaryItemService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/v1/dictionary_item.proto",
}
