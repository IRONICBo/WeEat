// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.19.4
// source: asr_service.proto

package asrservice

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
	ASRService_StreamRecognize_FullMethodName = "/asr.ASRService/StreamRecognize"
	ASRService_SyncRecognize_FullMethodName   = "/asr.ASRService/SyncRecognize"
)

// ASRServiceClient is the client API for ASRService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Export asr service for client
type ASRServiceClient interface {
	// ASR with stream response method
	StreamRecognize(ctx context.Context, in *ASRRequest, opts ...grpc.CallOption) (ASRService_StreamRecognizeClient, error)
	// ASR with sync response method
	SyncRecognize(ctx context.Context, in *ASRRequest, opts ...grpc.CallOption) (*ASRResponse, error)
}

type aSRServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewASRServiceClient(cc grpc.ClientConnInterface) ASRServiceClient {
	return &aSRServiceClient{cc}
}

func (c *aSRServiceClient) StreamRecognize(ctx context.Context, in *ASRRequest, opts ...grpc.CallOption) (ASRService_StreamRecognizeClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ASRService_ServiceDesc.Streams[0], ASRService_StreamRecognize_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &aSRServiceStreamRecognizeClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ASRService_StreamRecognizeClient interface {
	Recv() (*ASRResponse, error)
	grpc.ClientStream
}

type aSRServiceStreamRecognizeClient struct {
	grpc.ClientStream
}

func (x *aSRServiceStreamRecognizeClient) Recv() (*ASRResponse, error) {
	m := new(ASRResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *aSRServiceClient) SyncRecognize(ctx context.Context, in *ASRRequest, opts ...grpc.CallOption) (*ASRResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ASRResponse)
	err := c.cc.Invoke(ctx, ASRService_SyncRecognize_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ASRServiceServer is the server API for ASRService service.
// All implementations must embed UnimplementedASRServiceServer
// for forward compatibility
//
// Export asr service for client
type ASRServiceServer interface {
	// ASR with stream response method
	StreamRecognize(*ASRRequest, ASRService_StreamRecognizeServer) error
	// ASR with sync response method
	SyncRecognize(context.Context, *ASRRequest) (*ASRResponse, error)
	mustEmbedUnimplementedASRServiceServer()
}

// UnimplementedASRServiceServer must be embedded to have forward compatible implementations.
type UnimplementedASRServiceServer struct {
}

func (UnimplementedASRServiceServer) StreamRecognize(*ASRRequest, ASRService_StreamRecognizeServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamRecognize not implemented")
}
func (UnimplementedASRServiceServer) SyncRecognize(context.Context, *ASRRequest) (*ASRResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncRecognize not implemented")
}
func (UnimplementedASRServiceServer) mustEmbedUnimplementedASRServiceServer() {}

// UnsafeASRServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ASRServiceServer will
// result in compilation errors.
type UnsafeASRServiceServer interface {
	mustEmbedUnimplementedASRServiceServer()
}

func RegisterASRServiceServer(s grpc.ServiceRegistrar, srv ASRServiceServer) {
	s.RegisterService(&ASRService_ServiceDesc, srv)
}

func _ASRService_StreamRecognize_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ASRRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ASRServiceServer).StreamRecognize(m, &aSRServiceStreamRecognizeServer{ServerStream: stream})
}

type ASRService_StreamRecognizeServer interface {
	Send(*ASRResponse) error
	grpc.ServerStream
}

type aSRServiceStreamRecognizeServer struct {
	grpc.ServerStream
}

func (x *aSRServiceStreamRecognizeServer) Send(m *ASRResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ASRService_SyncRecognize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ASRRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ASRServiceServer).SyncRecognize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ASRService_SyncRecognize_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ASRServiceServer).SyncRecognize(ctx, req.(*ASRRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ASRService_ServiceDesc is the grpc.ServiceDesc for ASRService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ASRService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "asr.ASRService",
	HandlerType: (*ASRServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SyncRecognize",
			Handler:    _ASRService_SyncRecognize_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamRecognize",
			Handler:       _ASRService_StreamRecognize_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "asr_service.proto",
}