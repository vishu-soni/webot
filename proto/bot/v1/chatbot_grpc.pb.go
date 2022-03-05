// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: proto/bot/v1/chatbot.proto

package v1

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

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	Chit(ctx context.Context, in *ChitRequest, opts ...grpc.CallOption) (*ChitResponse, error)
	Chat(ctx context.Context, in *ChatRequest, opts ...grpc.CallOption) (Service_ChatClient, error)
	InitiateChat(ctx context.Context, in *InitiateChatRequest, opts ...grpc.CallOption) (*IntiateChatResponse, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) Chit(ctx context.Context, in *ChitRequest, opts ...grpc.CallOption) (*ChitResponse, error) {
	out := new(ChitResponse)
	err := c.cc.Invoke(ctx, "/proto.bot.v1.Service/Chit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) Chat(ctx context.Context, in *ChatRequest, opts ...grpc.CallOption) (Service_ChatClient, error) {
	stream, err := c.cc.NewStream(ctx, &Service_ServiceDesc.Streams[0], "/proto.bot.v1.Service/Chat", opts...)
	if err != nil {
		return nil, err
	}
	x := &serviceChatClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Service_ChatClient interface {
	Recv() (*ChatResponse, error)
	grpc.ClientStream
}

type serviceChatClient struct {
	grpc.ClientStream
}

func (x *serviceChatClient) Recv() (*ChatResponse, error) {
	m := new(ChatResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *serviceClient) InitiateChat(ctx context.Context, in *InitiateChatRequest, opts ...grpc.CallOption) (*IntiateChatResponse, error) {
	out := new(IntiateChatResponse)
	err := c.cc.Invoke(ctx, "/proto.bot.v1.Service/InitiateChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	Chit(context.Context, *ChitRequest) (*ChitResponse, error)
	Chat(*ChatRequest, Service_ChatServer) error
	InitiateChat(context.Context, *InitiateChatRequest) (*IntiateChatResponse, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) Chit(context.Context, *ChitRequest) (*ChitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Chit not implemented")
}
func (UnimplementedServiceServer) Chat(*ChatRequest, Service_ChatServer) error {
	return status.Errorf(codes.Unimplemented, "method Chat not implemented")
}
func (UnimplementedServiceServer) InitiateChat(context.Context, *InitiateChatRequest) (*IntiateChatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitiateChat not implemented")
}
func (UnimplementedServiceServer) mustEmbedUnimplementedServiceServer() {}

// UnsafeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceServer will
// result in compilation errors.
type UnsafeServiceServer interface {
	mustEmbedUnimplementedServiceServer()
}

func RegisterServiceServer(s grpc.ServiceRegistrar, srv ServiceServer) {
	s.RegisterService(&Service_ServiceDesc, srv)
}

func _Service_Chit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Chit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.bot.v1.Service/Chit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Chit(ctx, req.(*ChitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_Chat_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ChatRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ServiceServer).Chat(m, &serviceChatServer{stream})
}

type Service_ChatServer interface {
	Send(*ChatResponse) error
	grpc.ServerStream
}

type serviceChatServer struct {
	grpc.ServerStream
}

func (x *serviceChatServer) Send(m *ChatResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Service_InitiateChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitiateChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).InitiateChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.bot.v1.Service/InitiateChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).InitiateChat(ctx, req.(*InitiateChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.bot.v1.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Chit",
			Handler:    _Service_Chit_Handler,
		},
		{
			MethodName: "InitiateChat",
			Handler:    _Service_InitiateChat_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Chat",
			Handler:       _Service_Chat_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/bot/v1/chatbot.proto",
}