// protoc --proto_path=./api --go_out=./pkg/api/ --go_opt=paths=source_relative  --go-grpc_out ./pkg/api/ --go-grpc_opt=paths=source_relative ./api/*.proto
// protoc --proto_path=./api --go_out=./pkg/api/ --go_opt=paths=source_relative  --go-grpc_out ./pkg/api --go-grpc_opt=paths=source_relative --grpc-gateway_out ./pkg/api --grpc-gateway_opt paths=source_relative --openapi_out ./api/v1/ --openapi_opt=title="TodoAppService API" ./api/v1/service.proto

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.0
// source: v1/service.proto

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
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	TodoAppServiceV1_CreateList_FullMethodName = "/todoapp.v1.TodoAppServiceV1/CreateList"
	TodoAppServiceV1_DeleteList_FullMethodName = "/todoapp.v1.TodoAppServiceV1/DeleteList"
	TodoAppServiceV1_GetList_FullMethodName    = "/todoapp.v1.TodoAppServiceV1/GetList"
	TodoAppServiceV1_CreateItem_FullMethodName = "/todoapp.v1.TodoAppServiceV1/CreateItem"
	TodoAppServiceV1_DeleteItem_FullMethodName = "/todoapp.v1.TodoAppServiceV1/DeleteItem"
)

// TodoAppServiceV1Client is the client API for TodoAppServiceV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TodoAppServiceV1Client interface {
	// List
	CreateList(ctx context.Context, in *CreateListRequest, opts ...grpc.CallOption) (*CreateListResponse, error)
	DeleteList(ctx context.Context, in *DeleteListRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetListResponse, error)
	// Item
	CreateItem(ctx context.Context, in *CreateItemRequest, opts ...grpc.CallOption) (*CreateItemResponse, error)
	DeleteItem(ctx context.Context, in *DeleteItemRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type todoAppServiceV1Client struct {
	cc grpc.ClientConnInterface
}

func NewTodoAppServiceV1Client(cc grpc.ClientConnInterface) TodoAppServiceV1Client {
	return &todoAppServiceV1Client{cc}
}

func (c *todoAppServiceV1Client) CreateList(ctx context.Context, in *CreateListRequest, opts ...grpc.CallOption) (*CreateListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateListResponse)
	err := c.cc.Invoke(ctx, TodoAppServiceV1_CreateList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoAppServiceV1Client) DeleteList(ctx context.Context, in *DeleteListRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, TodoAppServiceV1_DeleteList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoAppServiceV1Client) GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetListResponse)
	err := c.cc.Invoke(ctx, TodoAppServiceV1_GetList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoAppServiceV1Client) CreateItem(ctx context.Context, in *CreateItemRequest, opts ...grpc.CallOption) (*CreateItemResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateItemResponse)
	err := c.cc.Invoke(ctx, TodoAppServiceV1_CreateItem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoAppServiceV1Client) DeleteItem(ctx context.Context, in *DeleteItemRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, TodoAppServiceV1_DeleteItem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoAppServiceV1Server is the server API for TodoAppServiceV1 service.
// All implementations must embed UnimplementedTodoAppServiceV1Server
// for forward compatibility.
type TodoAppServiceV1Server interface {
	// List
	CreateList(context.Context, *CreateListRequest) (*CreateListResponse, error)
	DeleteList(context.Context, *DeleteListRequest) (*emptypb.Empty, error)
	GetList(context.Context, *GetListRequest) (*GetListResponse, error)
	// Item
	CreateItem(context.Context, *CreateItemRequest) (*CreateItemResponse, error)
	DeleteItem(context.Context, *DeleteItemRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedTodoAppServiceV1Server()
}

// UnimplementedTodoAppServiceV1Server must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTodoAppServiceV1Server struct{}

func (UnimplementedTodoAppServiceV1Server) CreateList(context.Context, *CreateListRequest) (*CreateListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateList not implemented")
}
func (UnimplementedTodoAppServiceV1Server) DeleteList(context.Context, *DeleteListRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteList not implemented")
}
func (UnimplementedTodoAppServiceV1Server) GetList(context.Context, *GetListRequest) (*GetListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedTodoAppServiceV1Server) CreateItem(context.Context, *CreateItemRequest) (*CreateItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateItem not implemented")
}
func (UnimplementedTodoAppServiceV1Server) DeleteItem(context.Context, *DeleteItemRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteItem not implemented")
}
func (UnimplementedTodoAppServiceV1Server) mustEmbedUnimplementedTodoAppServiceV1Server() {}
func (UnimplementedTodoAppServiceV1Server) testEmbeddedByValue()                          {}

// UnsafeTodoAppServiceV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TodoAppServiceV1Server will
// result in compilation errors.
type UnsafeTodoAppServiceV1Server interface {
	mustEmbedUnimplementedTodoAppServiceV1Server()
}

func RegisterTodoAppServiceV1Server(s grpc.ServiceRegistrar, srv TodoAppServiceV1Server) {
	// If the following call pancis, it indicates UnimplementedTodoAppServiceV1Server was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TodoAppServiceV1_ServiceDesc, srv)
}

func _TodoAppServiceV1_CreateList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoAppServiceV1Server).CreateList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TodoAppServiceV1_CreateList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoAppServiceV1Server).CreateList(ctx, req.(*CreateListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoAppServiceV1_DeleteList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoAppServiceV1Server).DeleteList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TodoAppServiceV1_DeleteList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoAppServiceV1Server).DeleteList(ctx, req.(*DeleteListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoAppServiceV1_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoAppServiceV1Server).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TodoAppServiceV1_GetList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoAppServiceV1Server).GetList(ctx, req.(*GetListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoAppServiceV1_CreateItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoAppServiceV1Server).CreateItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TodoAppServiceV1_CreateItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoAppServiceV1Server).CreateItem(ctx, req.(*CreateItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoAppServiceV1_DeleteItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoAppServiceV1Server).DeleteItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TodoAppServiceV1_DeleteItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoAppServiceV1Server).DeleteItem(ctx, req.(*DeleteItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TodoAppServiceV1_ServiceDesc is the grpc.ServiceDesc for TodoAppServiceV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TodoAppServiceV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "todoapp.v1.TodoAppServiceV1",
	HandlerType: (*TodoAppServiceV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateList",
			Handler:    _TodoAppServiceV1_CreateList_Handler,
		},
		{
			MethodName: "DeleteList",
			Handler:    _TodoAppServiceV1_DeleteList_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _TodoAppServiceV1_GetList_Handler,
		},
		{
			MethodName: "CreateItem",
			Handler:    _TodoAppServiceV1_CreateItem_Handler,
		},
		{
			MethodName: "DeleteItem",
			Handler:    _TodoAppServiceV1_DeleteItem_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/service.proto",
}
