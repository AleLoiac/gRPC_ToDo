// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.22.0--rc3
// source: ToDopb/todo.proto

package ToDopb

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

// ToDoServiceClient is the client API for ToDoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ToDoServiceClient interface {
	CreateToDo(ctx context.Context, in *NewToDo, opts ...grpc.CallOption) (*ToDoResponse, error)
	ListToDos(ctx context.Context, in *Empty, opts ...grpc.CallOption) (ToDoService_ListToDosClient, error)
	CheckUncheck(ctx context.Context, in *ToDoId, opts ...grpc.CallOption) (*ToDoResponse, error)
	DeleteToDo(ctx context.Context, in *ToDoId, opts ...grpc.CallOption) (*Empty, error)
}

type toDoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewToDoServiceClient(cc grpc.ClientConnInterface) ToDoServiceClient {
	return &toDoServiceClient{cc}
}

func (c *toDoServiceClient) CreateToDo(ctx context.Context, in *NewToDo, opts ...grpc.CallOption) (*ToDoResponse, error) {
	out := new(ToDoResponse)
	err := c.cc.Invoke(ctx, "/todo.ToDoService/CreateToDo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toDoServiceClient) ListToDos(ctx context.Context, in *Empty, opts ...grpc.CallOption) (ToDoService_ListToDosClient, error) {
	stream, err := c.cc.NewStream(ctx, &ToDoService_ServiceDesc.Streams[0], "/todo.ToDoService/ListToDos", opts...)
	if err != nil {
		return nil, err
	}
	x := &toDoServiceListToDosClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ToDoService_ListToDosClient interface {
	Recv() (*ToDoResponse, error)
	grpc.ClientStream
}

type toDoServiceListToDosClient struct {
	grpc.ClientStream
}

func (x *toDoServiceListToDosClient) Recv() (*ToDoResponse, error) {
	m := new(ToDoResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *toDoServiceClient) CheckUncheck(ctx context.Context, in *ToDoId, opts ...grpc.CallOption) (*ToDoResponse, error) {
	out := new(ToDoResponse)
	err := c.cc.Invoke(ctx, "/todo.ToDoService/CheckUncheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toDoServiceClient) DeleteToDo(ctx context.Context, in *ToDoId, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/todo.ToDoService/DeleteToDo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ToDoServiceServer is the server API for ToDoService service.
// All implementations must embed UnimplementedToDoServiceServer
// for forward compatibility
type ToDoServiceServer interface {
	CreateToDo(context.Context, *NewToDo) (*ToDoResponse, error)
	ListToDos(*Empty, ToDoService_ListToDosServer) error
	CheckUncheck(context.Context, *ToDoId) (*ToDoResponse, error)
	DeleteToDo(context.Context, *ToDoId) (*Empty, error)
	mustEmbedUnimplementedToDoServiceServer()
}

// UnimplementedToDoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedToDoServiceServer struct {
}

func (UnimplementedToDoServiceServer) CreateToDo(context.Context, *NewToDo) (*ToDoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateToDo not implemented")
}
func (UnimplementedToDoServiceServer) ListToDos(*Empty, ToDoService_ListToDosServer) error {
	return status.Errorf(codes.Unimplemented, "method ListToDos not implemented")
}
func (UnimplementedToDoServiceServer) CheckUncheck(context.Context, *ToDoId) (*ToDoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckUncheck not implemented")
}
func (UnimplementedToDoServiceServer) DeleteToDo(context.Context, *ToDoId) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteToDo not implemented")
}
func (UnimplementedToDoServiceServer) mustEmbedUnimplementedToDoServiceServer() {}

// UnsafeToDoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ToDoServiceServer will
// result in compilation errors.
type UnsafeToDoServiceServer interface {
	mustEmbedUnimplementedToDoServiceServer()
}

func RegisterToDoServiceServer(s grpc.ServiceRegistrar, srv ToDoServiceServer) {
	s.RegisterService(&ToDoService_ServiceDesc, srv)
}

func _ToDoService_CreateToDo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewToDo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToDoServiceServer).CreateToDo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.ToDoService/CreateToDo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToDoServiceServer).CreateToDo(ctx, req.(*NewToDo))
	}
	return interceptor(ctx, in, info, handler)
}

func _ToDoService_ListToDos_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ToDoServiceServer).ListToDos(m, &toDoServiceListToDosServer{stream})
}

type ToDoService_ListToDosServer interface {
	Send(*ToDoResponse) error
	grpc.ServerStream
}

type toDoServiceListToDosServer struct {
	grpc.ServerStream
}

func (x *toDoServiceListToDosServer) Send(m *ToDoResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ToDoService_CheckUncheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ToDoId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToDoServiceServer).CheckUncheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.ToDoService/CheckUncheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToDoServiceServer).CheckUncheck(ctx, req.(*ToDoId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ToDoService_DeleteToDo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ToDoId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToDoServiceServer).DeleteToDo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.ToDoService/DeleteToDo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToDoServiceServer).DeleteToDo(ctx, req.(*ToDoId))
	}
	return interceptor(ctx, in, info, handler)
}

// ToDoService_ServiceDesc is the grpc.ServiceDesc for ToDoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ToDoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "todo.ToDoService",
	HandlerType: (*ToDoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateToDo",
			Handler:    _ToDoService_CreateToDo_Handler,
		},
		{
			MethodName: "CheckUncheck",
			Handler:    _ToDoService_CheckUncheck_Handler,
		},
		{
			MethodName: "DeleteToDo",
			Handler:    _ToDoService_DeleteToDo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListToDos",
			Handler:       _ToDoService_ListToDos_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "ToDopb/todo.proto",
}
