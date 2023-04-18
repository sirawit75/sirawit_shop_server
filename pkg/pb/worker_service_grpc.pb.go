// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: worker_service.proto

package pb

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

// WorkerServiceClient is the client API for WorkerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WorkerServiceClient interface {
	InsertLoginTimestamp(ctx context.Context, in *LoginTimestampReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SendWelcomeEmail(ctx context.Context, in *SendWelcomeEmailReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CheckRediness(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type workerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkerServiceClient(cc grpc.ClientConnInterface) WorkerServiceClient {
	return &workerServiceClient{cc}
}

func (c *workerServiceClient) InsertLoginTimestamp(ctx context.Context, in *LoginTimestampReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.WorkerService/InsertLoginTimestamp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerServiceClient) SendWelcomeEmail(ctx context.Context, in *SendWelcomeEmailReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.WorkerService/SendWelcomeEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerServiceClient) CheckRediness(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.WorkerService/CheckRediness", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkerServiceServer is the server API for WorkerService service.
// All implementations must embed UnimplementedWorkerServiceServer
// for forward compatibility
type WorkerServiceServer interface {
	InsertLoginTimestamp(context.Context, *LoginTimestampReq) (*emptypb.Empty, error)
	SendWelcomeEmail(context.Context, *SendWelcomeEmailReq) (*emptypb.Empty, error)
	CheckRediness(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	mustEmbedUnimplementedWorkerServiceServer()
}

// UnimplementedWorkerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWorkerServiceServer struct {
}

func (UnimplementedWorkerServiceServer) InsertLoginTimestamp(context.Context, *LoginTimestampReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertLoginTimestamp not implemented")
}
func (UnimplementedWorkerServiceServer) SendWelcomeEmail(context.Context, *SendWelcomeEmailReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendWelcomeEmail not implemented")
}
func (UnimplementedWorkerServiceServer) CheckRediness(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckRediness not implemented")
}
func (UnimplementedWorkerServiceServer) mustEmbedUnimplementedWorkerServiceServer() {}

// UnsafeWorkerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WorkerServiceServer will
// result in compilation errors.
type UnsafeWorkerServiceServer interface {
	mustEmbedUnimplementedWorkerServiceServer()
}

func RegisterWorkerServiceServer(s grpc.ServiceRegistrar, srv WorkerServiceServer) {
	s.RegisterService(&WorkerService_ServiceDesc, srv)
}

func _WorkerService_InsertLoginTimestamp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginTimestampReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServiceServer).InsertLoginTimestamp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.WorkerService/InsertLoginTimestamp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServiceServer).InsertLoginTimestamp(ctx, req.(*LoginTimestampReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkerService_SendWelcomeEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendWelcomeEmailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServiceServer).SendWelcomeEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.WorkerService/SendWelcomeEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServiceServer).SendWelcomeEmail(ctx, req.(*SendWelcomeEmailReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkerService_CheckRediness_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServiceServer).CheckRediness(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.WorkerService/CheckRediness",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServiceServer).CheckRediness(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// WorkerService_ServiceDesc is the grpc.ServiceDesc for WorkerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WorkerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.WorkerService",
	HandlerType: (*WorkerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InsertLoginTimestamp",
			Handler:    _WorkerService_InsertLoginTimestamp_Handler,
		},
		{
			MethodName: "SendWelcomeEmail",
			Handler:    _WorkerService_SendWelcomeEmail_Handler,
		},
		{
			MethodName: "CheckRediness",
			Handler:    _WorkerService_CheckRediness_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "worker_service.proto",
}
