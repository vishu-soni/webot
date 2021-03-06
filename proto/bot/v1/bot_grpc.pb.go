// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: proto/bot/v1/bot.proto

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

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatServiceClient interface {
	GetResponse(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error)
	JoinChannel(ctx context.Context, in *Channel, opts ...grpc.CallOption) (ChatService_JoinChannelClient, error)
	SendMessage(ctx context.Context, opts ...grpc.CallOption) (ChatService_SendMessageClient, error)
	ChitChat(ctx context.Context, opts ...grpc.CallOption) (ChatService_ChitChatClient, error)
}

type chatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServiceClient(cc grpc.ClientConnInterface) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) GetResponse(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	out := new(QueryResponse)
	err := c.cc.Invoke(ctx, "/proto.bot.v1.ChatService/getResponse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) JoinChannel(ctx context.Context, in *Channel, opts ...grpc.CallOption) (ChatService_JoinChannelClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChatService_ServiceDesc.Streams[0], "/proto.bot.v1.ChatService/JoinChannel", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceJoinChannelClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ChatService_JoinChannelClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type chatServiceJoinChannelClient struct {
	grpc.ClientStream
}

func (x *chatServiceJoinChannelClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatServiceClient) SendMessage(ctx context.Context, opts ...grpc.CallOption) (ChatService_SendMessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChatService_ServiceDesc.Streams[1], "/proto.bot.v1.ChatService/SendMessage", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceSendMessageClient{stream}
	return x, nil
}

type ChatService_SendMessageClient interface {
	Send(*Message) error
	CloseAndRecv() (*MessageAck, error)
	grpc.ClientStream
}

type chatServiceSendMessageClient struct {
	grpc.ClientStream
}

func (x *chatServiceSendMessageClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatServiceSendMessageClient) CloseAndRecv() (*MessageAck, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(MessageAck)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatServiceClient) ChitChat(ctx context.Context, opts ...grpc.CallOption) (ChatService_ChitChatClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChatService_ServiceDesc.Streams[2], "/proto.bot.v1.ChatService/ChitChat", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceChitChatClient{stream}
	return x, nil
}

type ChatService_ChitChatClient interface {
	Send(*Chit) error
	CloseAndRecv() (*Chat, error)
	grpc.ClientStream
}

type chatServiceChitChatClient struct {
	grpc.ClientStream
}

func (x *chatServiceChitChatClient) Send(m *Chit) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatServiceChitChatClient) CloseAndRecv() (*Chat, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Chat)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatServiceServer is the server API for ChatService service.
// All implementations must embed UnimplementedChatServiceServer
// for forward compatibility
type ChatServiceServer interface {
	GetResponse(context.Context, *QueryRequest) (*QueryResponse, error)
	JoinChannel(*Channel, ChatService_JoinChannelServer) error
	SendMessage(ChatService_SendMessageServer) error
	ChitChat(ChatService_ChitChatServer) error
	mustEmbedUnimplementedChatServiceServer()
}

// UnimplementedChatServiceServer must be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (UnimplementedChatServiceServer) GetResponse(context.Context, *QueryRequest) (*QueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetResponse not implemented")
}
func (UnimplementedChatServiceServer) JoinChannel(*Channel, ChatService_JoinChannelServer) error {
	return status.Errorf(codes.Unimplemented, "method JoinChannel not implemented")
}
func (UnimplementedChatServiceServer) SendMessage(ChatService_SendMessageServer) error {
	return status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedChatServiceServer) ChitChat(ChatService_ChitChatServer) error {
	return status.Errorf(codes.Unimplemented, "method ChitChat not implemented")
}
func (UnimplementedChatServiceServer) mustEmbedUnimplementedChatServiceServer() {}

// UnsafeChatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServiceServer will
// result in compilation errors.
type UnsafeChatServiceServer interface {
	mustEmbedUnimplementedChatServiceServer()
}

func RegisterChatServiceServer(s grpc.ServiceRegistrar, srv ChatServiceServer) {
	s.RegisterService(&ChatService_ServiceDesc, srv)
}

func _ChatService_GetResponse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).GetResponse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.bot.v1.ChatService/getResponse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).GetResponse(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_JoinChannel_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Channel)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatServiceServer).JoinChannel(m, &chatServiceJoinChannelServer{stream})
}

type ChatService_JoinChannelServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type chatServiceJoinChannelServer struct {
	grpc.ServerStream
}

func (x *chatServiceJoinChannelServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func _ChatService_SendMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServiceServer).SendMessage(&chatServiceSendMessageServer{stream})
}

type ChatService_SendMessageServer interface {
	SendAndClose(*MessageAck) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type chatServiceSendMessageServer struct {
	grpc.ServerStream
}

func (x *chatServiceSendMessageServer) SendAndClose(m *MessageAck) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatServiceSendMessageServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ChatService_ChitChat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServiceServer).ChitChat(&chatServiceChitChatServer{stream})
}

type ChatService_ChitChatServer interface {
	SendAndClose(*Chat) error
	Recv() (*Chit, error)
	grpc.ServerStream
}

type chatServiceChitChatServer struct {
	grpc.ServerStream
}

func (x *chatServiceChitChatServer) SendAndClose(m *Chat) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatServiceChitChatServer) Recv() (*Chit, error) {
	m := new(Chit)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatService_ServiceDesc is the grpc.ServiceDesc for ChatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.bot.v1.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getResponse",
			Handler:    _ChatService_GetResponse_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "JoinChannel",
			Handler:       _ChatService_JoinChannel_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SendMessage",
			Handler:       _ChatService_SendMessage_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "ChitChat",
			Handler:       _ChatService_ChitChat_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "proto/bot/v1/bot.proto",
}
