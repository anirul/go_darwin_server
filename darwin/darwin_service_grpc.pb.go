// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: darwin_service.proto

package darwin

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
	DarwinService_Update_FullMethodName = "/darwin.DarwinService/Update"
	DarwinService_Push_FullMethodName   = "/darwin.DarwinService/Push"
)

// DarwinServiceClient is the client API for DarwinService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DarwinServiceClient interface {
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (DarwinService_UpdateClient, error)
	Push(ctx context.Context, in *PushRequest, opts ...grpc.CallOption) (*PushResponse, error)
}

type darwinServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDarwinServiceClient(cc grpc.ClientConnInterface) DarwinServiceClient {
	return &darwinServiceClient{cc}
}

func (c *darwinServiceClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (DarwinService_UpdateClient, error) {
	stream, err := c.cc.NewStream(ctx, &DarwinService_ServiceDesc.Streams[0], DarwinService_Update_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &darwinServiceUpdateClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DarwinService_UpdateClient interface {
	Recv() (*UpdateResponse, error)
	grpc.ClientStream
}

type darwinServiceUpdateClient struct {
	grpc.ClientStream
}

func (x *darwinServiceUpdateClient) Recv() (*UpdateResponse, error) {
	m := new(UpdateResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *darwinServiceClient) Push(ctx context.Context, in *PushRequest, opts ...grpc.CallOption) (*PushResponse, error) {
	out := new(PushResponse)
	err := c.cc.Invoke(ctx, DarwinService_Push_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DarwinServiceServer is the server API for DarwinService service.
// All implementations must embed UnimplementedDarwinServiceServer
// for forward compatibility
type DarwinServiceServer interface {
	Update(*UpdateRequest, DarwinService_UpdateServer) error
	Push(context.Context, *PushRequest) (*PushResponse, error)
	mustEmbedUnimplementedDarwinServiceServer()
}

// UnimplementedDarwinServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDarwinServiceServer struct {
}

func (UnimplementedDarwinServiceServer) Update(*UpdateRequest, DarwinService_UpdateServer) error {
	return status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedDarwinServiceServer) Push(context.Context, *PushRequest) (*PushResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Push not implemented")
}
func (UnimplementedDarwinServiceServer) mustEmbedUnimplementedDarwinServiceServer() {}

// UnsafeDarwinServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DarwinServiceServer will
// result in compilation errors.
type UnsafeDarwinServiceServer interface {
	mustEmbedUnimplementedDarwinServiceServer()
}

func RegisterDarwinServiceServer(s grpc.ServiceRegistrar, srv DarwinServiceServer) {
	s.RegisterService(&DarwinService_ServiceDesc, srv)
}

func _DarwinService_Update_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(UpdateRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DarwinServiceServer).Update(m, &darwinServiceUpdateServer{stream})
}

type DarwinService_UpdateServer interface {
	Send(*UpdateResponse) error
	grpc.ServerStream
}

type darwinServiceUpdateServer struct {
	grpc.ServerStream
}

func (x *darwinServiceUpdateServer) Send(m *UpdateResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _DarwinService_Push_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DarwinServiceServer).Push(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DarwinService_Push_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DarwinServiceServer).Push(ctx, req.(*PushRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DarwinService_ServiceDesc is the grpc.ServiceDesc for DarwinService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DarwinService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "darwin.DarwinService",
	HandlerType: (*DarwinServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Push",
			Handler:    _DarwinService_Push_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Update",
			Handler:       _DarwinService_Update_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "darwin_service.proto",
}