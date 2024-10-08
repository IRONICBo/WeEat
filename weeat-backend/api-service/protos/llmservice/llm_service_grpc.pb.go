// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.19.4
// source: llm_service.proto

package llmservice

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	LLMService_StreamChat_FullMethodName = "/chat.LLMService/StreamChat"
	LLMService_SyncChat_FullMethodName   = "/chat.LLMService/SyncChat"
)

// LLMServiceClient is the client API for LLMService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Exporty llm service for client
type LLMServiceClient interface {
	// Chat with stream response method
	StreamChat(ctx context.Context, in *ChatMessageRequest, opts ...grpc.CallOption) (LLMService_StreamChatClient, error)
	// Chat with sync response method
	SyncChat(ctx context.Context, in *ChatMessageRequest, opts ...grpc.CallOption) (*ChatMessageResponse, error)
}

type lLMServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLLMServiceClient(cc grpc.ClientConnInterface) LLMServiceClient {
	return &lLMServiceClient{cc}
}

func (c *lLMServiceClient) StreamChat(ctx context.Context, in *ChatMessageRequest, opts ...grpc.CallOption) (LLMService_StreamChatClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &LLMService_ServiceDesc.Streams[0], LLMService_StreamChat_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &lLMServiceStreamChatClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LLMService_StreamChatClient interface {
	Recv() (*ChatMessageResponse, error)
	grpc.ClientStream
}

type lLMServiceStreamChatClient struct {
	grpc.ClientStream
}

func (x *lLMServiceStreamChatClient) Recv() (*ChatMessageResponse, error) {
	m := new(ChatMessageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *lLMServiceClient) SyncChat(ctx context.Context, in *ChatMessageRequest, opts ...grpc.CallOption) (*ChatMessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ChatMessageResponse)
	err := c.cc.Invoke(ctx, LLMService_SyncChat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LLMServiceServer is the server API for LLMService service.
// All implementations must embed UnimplementedLLMServiceServer
// for forward compatibility
//
// Exporty llm service for client
type LLMServiceServer interface {
	// Chat with stream response method
	StreamChat(*ChatMessageRequest, LLMService_StreamChatServer) error
	// Chat with sync response method
	SyncChat(context.Context, *ChatMessageRequest) (*ChatMessageResponse, error)
	mustEmbedUnimplementedLLMServiceServer()
}

// UnimplementedLLMServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLLMServiceServer struct {
}

func (UnimplementedLLMServiceServer) StreamChat(*ChatMessageRequest, LLMService_StreamChatServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamChat not implemented")
}
func (UnimplementedLLMServiceServer) SyncChat(context.Context, *ChatMessageRequest) (*ChatMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncChat not implemented")
}
func (UnimplementedLLMServiceServer) mustEmbedUnimplementedLLMServiceServer() {}

// UnsafeLLMServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LLMServiceServer will
// result in compilation errors.
type UnsafeLLMServiceServer interface {
	mustEmbedUnimplementedLLMServiceServer()
}

func RegisterLLMServiceServer(s grpc.ServiceRegistrar, srv LLMServiceServer) {
	s.RegisterService(&LLMService_ServiceDesc, srv)
}

func _LLMService_StreamChat_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ChatMessageRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LLMServiceServer).StreamChat(m, &lLMServiceStreamChatServer{ServerStream: stream})
}

type LLMService_StreamChatServer interface {
	Send(*ChatMessageResponse) error
	grpc.ServerStream
}

type lLMServiceStreamChatServer struct {
	grpc.ServerStream
}

func (x *lLMServiceStreamChatServer) Send(m *ChatMessageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _LLMService_SyncChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LLMServiceServer).SyncChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LLMService_SyncChat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LLMServiceServer).SyncChat(ctx, req.(*ChatMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LLMService_ServiceDesc is the grpc.ServiceDesc for LLMService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LLMService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chat.LLMService",
	HandlerType: (*LLMServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SyncChat",
			Handler:    _LLMService_SyncChat_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamChat",
			Handler:       _LLMService_StreamChat_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "llm_service.proto",
}
