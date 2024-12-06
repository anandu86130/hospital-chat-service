// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.3
// source: chat.proto

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ChatService_Connect_FullMethodName        = "/pbC.ChatService/Connect"
	ChatService_FetchHistory_FullMethodName   = "/pbC.ChatService/FetchHistory"
	ChatService_StartVideoCall_FullMethodName = "/pbC.ChatService/StartVideoCall"
	ChatService_FetchVideoCall_FullMethodName = "/pbC.ChatService/FetchVideoCall"
)

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatServiceClient interface {
	Connect(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[Message, Message], error)
	FetchHistory(ctx context.Context, in *ChatID, opts ...grpc.CallOption) (*ChatHistory, error)
	StartVideoCall(ctx context.Context, in *VideoCallRequest, opts ...grpc.CallOption) (*VideoCallResponse, error)
	FetchVideoCall(ctx context.Context, in *ChatID, opts ...grpc.CallOption) (*ChatHistory, error)
}

type chatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServiceClient(cc grpc.ClientConnInterface) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) Connect(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[Message, Message], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ChatService_ServiceDesc.Streams[0], ChatService_Connect_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[Message, Message]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ChatService_ConnectClient = grpc.BidiStreamingClient[Message, Message]

func (c *chatServiceClient) FetchHistory(ctx context.Context, in *ChatID, opts ...grpc.CallOption) (*ChatHistory, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ChatHistory)
	err := c.cc.Invoke(ctx, ChatService_FetchHistory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) StartVideoCall(ctx context.Context, in *VideoCallRequest, opts ...grpc.CallOption) (*VideoCallResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VideoCallResponse)
	err := c.cc.Invoke(ctx, ChatService_StartVideoCall_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) FetchVideoCall(ctx context.Context, in *ChatID, opts ...grpc.CallOption) (*ChatHistory, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ChatHistory)
	err := c.cc.Invoke(ctx, ChatService_FetchVideoCall_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServiceServer is the server API for ChatService service.
// All implementations must embed UnimplementedChatServiceServer
// for forward compatibility.
type ChatServiceServer interface {
	Connect(grpc.BidiStreamingServer[Message, Message]) error
	FetchHistory(context.Context, *ChatID) (*ChatHistory, error)
	StartVideoCall(context.Context, *VideoCallRequest) (*VideoCallResponse, error)
	FetchVideoCall(context.Context, *ChatID) (*ChatHistory, error)
	mustEmbedUnimplementedChatServiceServer()
}

// UnimplementedChatServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedChatServiceServer struct{}

func (UnimplementedChatServiceServer) Connect(grpc.BidiStreamingServer[Message, Message]) error {
	return status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (UnimplementedChatServiceServer) FetchHistory(context.Context, *ChatID) (*ChatHistory, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchHistory not implemented")
}
func (UnimplementedChatServiceServer) StartVideoCall(context.Context, *VideoCallRequest) (*VideoCallResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartVideoCall not implemented")
}
func (UnimplementedChatServiceServer) FetchVideoCall(context.Context, *ChatID) (*ChatHistory, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchVideoCall not implemented")
}
func (UnimplementedChatServiceServer) mustEmbedUnimplementedChatServiceServer() {}
func (UnimplementedChatServiceServer) testEmbeddedByValue()                     {}

// UnsafeChatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServiceServer will
// result in compilation errors.
type UnsafeChatServiceServer interface {
	mustEmbedUnimplementedChatServiceServer()
}

func RegisterChatServiceServer(s grpc.ServiceRegistrar, srv ChatServiceServer) {
	// If the following call pancis, it indicates UnimplementedChatServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ChatService_ServiceDesc, srv)
}

func _ChatService_Connect_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServiceServer).Connect(&grpc.GenericServerStream[Message, Message]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ChatService_ConnectServer = grpc.BidiStreamingServer[Message, Message]

func _ChatService_FetchHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).FetchHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_FetchHistory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).FetchHistory(ctx, req.(*ChatID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_StartVideoCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VideoCallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).StartVideoCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_StartVideoCall_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).StartVideoCall(ctx, req.(*VideoCallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_FetchVideoCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).FetchVideoCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_FetchVideoCall_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).FetchVideoCall(ctx, req.(*ChatID))
	}
	return interceptor(ctx, in, info, handler)
}

// ChatService_ServiceDesc is the grpc.ServiceDesc for ChatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pbC.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchHistory",
			Handler:    _ChatService_FetchHistory_Handler,
		},
		{
			MethodName: "StartVideoCall",
			Handler:    _ChatService_StartVideoCall_Handler,
		},
		{
			MethodName: "FetchVideoCall",
			Handler:    _ChatService_FetchVideoCall_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Connect",
			Handler:       _ChatService_Connect_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "chat.proto",
}
