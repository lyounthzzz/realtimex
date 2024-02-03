// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: api/msg/msg.proto

package msg

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

// MsgServiceClient is the client API for MsgService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgServiceClient interface {
	// 连接
	Connect(ctx context.Context, in *ConnectReq, opts ...grpc.CallOption) (*ConnectAck, error)
	// 订阅主题
	Subscribe(ctx context.Context, in *SubscribeReq, opts ...grpc.CallOption) (*SubscribeAck, error)
	// 取消订阅主题
	Unsubscribe(ctx context.Context, in *UnsubscribeReq, opts ...grpc.CallOption) (*UnsubscribeAck, error)
	// 发布消息
	Publish(ctx context.Context, in *PublishReq, opts ...grpc.CallOption) (*PublishAck, error)
	// 心跳
	Ping(ctx context.Context, in *PingReq, opts ...grpc.CallOption) (*PingAck, error)
	// 断开连接
	Disconnect(ctx context.Context, in *DisconnectReq, opts ...grpc.CallOption) (*DisconnectAck, error)
}

type msgServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgServiceClient(cc grpc.ClientConnInterface) MsgServiceClient {
	return &msgServiceClient{cc}
}

func (c *msgServiceClient) Connect(ctx context.Context, in *ConnectReq, opts ...grpc.CallOption) (*ConnectAck, error) {
	out := new(ConnectAck)
	err := c.cc.Invoke(ctx, "/realtimex.msg.MsgService/Connect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) Subscribe(ctx context.Context, in *SubscribeReq, opts ...grpc.CallOption) (*SubscribeAck, error) {
	out := new(SubscribeAck)
	err := c.cc.Invoke(ctx, "/realtimex.msg.MsgService/Subscribe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) Unsubscribe(ctx context.Context, in *UnsubscribeReq, opts ...grpc.CallOption) (*UnsubscribeAck, error) {
	out := new(UnsubscribeAck)
	err := c.cc.Invoke(ctx, "/realtimex.msg.MsgService/Unsubscribe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) Publish(ctx context.Context, in *PublishReq, opts ...grpc.CallOption) (*PublishAck, error) {
	out := new(PublishAck)
	err := c.cc.Invoke(ctx, "/realtimex.msg.MsgService/Publish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) Ping(ctx context.Context, in *PingReq, opts ...grpc.CallOption) (*PingAck, error) {
	out := new(PingAck)
	err := c.cc.Invoke(ctx, "/realtimex.msg.MsgService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) Disconnect(ctx context.Context, in *DisconnectReq, opts ...grpc.CallOption) (*DisconnectAck, error) {
	out := new(DisconnectAck)
	err := c.cc.Invoke(ctx, "/realtimex.msg.MsgService/Disconnect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServiceServer is the server API for MsgService service.
// All implementations must embed UnimplementedMsgServiceServer
// for forward compatibility
type MsgServiceServer interface {
	// 连接
	Connect(context.Context, *ConnectReq) (*ConnectAck, error)
	// 订阅主题
	Subscribe(context.Context, *SubscribeReq) (*SubscribeAck, error)
	// 取消订阅主题
	Unsubscribe(context.Context, *UnsubscribeReq) (*UnsubscribeAck, error)
	// 发布消息
	Publish(context.Context, *PublishReq) (*PublishAck, error)
	// 心跳
	Ping(context.Context, *PingReq) (*PingAck, error)
	// 断开连接
	Disconnect(context.Context, *DisconnectReq) (*DisconnectAck, error)
	mustEmbedUnimplementedMsgServiceServer()
}

// UnimplementedMsgServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServiceServer struct {
}

func (UnimplementedMsgServiceServer) Connect(context.Context, *ConnectReq) (*ConnectAck, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (UnimplementedMsgServiceServer) Subscribe(context.Context, *SubscribeReq) (*SubscribeAck, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedMsgServiceServer) Unsubscribe(context.Context, *UnsubscribeReq) (*UnsubscribeAck, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unsubscribe not implemented")
}
func (UnimplementedMsgServiceServer) Publish(context.Context, *PublishReq) (*PublishAck, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (UnimplementedMsgServiceServer) Ping(context.Context, *PingReq) (*PingAck, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedMsgServiceServer) Disconnect(context.Context, *DisconnectReq) (*DisconnectAck, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Disconnect not implemented")
}
func (UnimplementedMsgServiceServer) mustEmbedUnimplementedMsgServiceServer() {}

// UnsafeMsgServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServiceServer will
// result in compilation errors.
type UnsafeMsgServiceServer interface {
	mustEmbedUnimplementedMsgServiceServer()
}

func RegisterMsgServiceServer(s grpc.ServiceRegistrar, srv MsgServiceServer) {
	s.RegisterService(&MsgService_ServiceDesc, srv)
}

func _MsgService_Connect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).Connect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/realtimex.msg.MsgService/Connect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).Connect(ctx, req.(*ConnectReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_Subscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscribeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).Subscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/realtimex.msg.MsgService/Subscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).Subscribe(ctx, req.(*SubscribeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_Unsubscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnsubscribeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).Unsubscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/realtimex.msg.MsgService/Unsubscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).Unsubscribe(ctx, req.(*UnsubscribeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/realtimex.msg.MsgService/Publish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).Publish(ctx, req.(*PublishReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/realtimex.msg.MsgService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).Ping(ctx, req.(*PingReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_Disconnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisconnectReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).Disconnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/realtimex.msg.MsgService/Disconnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).Disconnect(ctx, req.(*DisconnectReq))
	}
	return interceptor(ctx, in, info, handler)
}

// MsgService_ServiceDesc is the grpc.ServiceDesc for MsgService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MsgService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "realtimex.msg.MsgService",
	HandlerType: (*MsgServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Connect",
			Handler:    _MsgService_Connect_Handler,
		},
		{
			MethodName: "Subscribe",
			Handler:    _MsgService_Subscribe_Handler,
		},
		{
			MethodName: "Unsubscribe",
			Handler:    _MsgService_Unsubscribe_Handler,
		},
		{
			MethodName: "Publish",
			Handler:    _MsgService_Publish_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _MsgService_Ping_Handler,
		},
		{
			MethodName: "Disconnect",
			Handler:    _MsgService_Disconnect_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/msg/msg.proto",
}

// PushServiceClient is the client API for PushService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PushServiceClient interface {
	// 广播消息
	Broadcast(ctx context.Context, in *PushMsg, opts ...grpc.CallOption) (*BroadcastResp, error)
	// 推送消息
	Push(ctx context.Context, in *PushMsg, opts ...grpc.CallOption) (*PushMsgResp, error)
}

type pushServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPushServiceClient(cc grpc.ClientConnInterface) PushServiceClient {
	return &pushServiceClient{cc}
}

func (c *pushServiceClient) Broadcast(ctx context.Context, in *PushMsg, opts ...grpc.CallOption) (*BroadcastResp, error) {
	out := new(BroadcastResp)
	err := c.cc.Invoke(ctx, "/realtimex.msg.PushService/Broadcast", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pushServiceClient) Push(ctx context.Context, in *PushMsg, opts ...grpc.CallOption) (*PushMsgResp, error) {
	out := new(PushMsgResp)
	err := c.cc.Invoke(ctx, "/realtimex.msg.PushService/Push", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PushServiceServer is the server API for PushService service.
// All implementations must embed UnimplementedPushServiceServer
// for forward compatibility
type PushServiceServer interface {
	// 广播消息
	Broadcast(context.Context, *PushMsg) (*BroadcastResp, error)
	// 推送消息
	Push(context.Context, *PushMsg) (*PushMsgResp, error)
	mustEmbedUnimplementedPushServiceServer()
}

// UnimplementedPushServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPushServiceServer struct {
}

func (UnimplementedPushServiceServer) Broadcast(context.Context, *PushMsg) (*BroadcastResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Broadcast not implemented")
}
func (UnimplementedPushServiceServer) Push(context.Context, *PushMsg) (*PushMsgResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Push not implemented")
}
func (UnimplementedPushServiceServer) mustEmbedUnimplementedPushServiceServer() {}

// UnsafePushServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PushServiceServer will
// result in compilation errors.
type UnsafePushServiceServer interface {
	mustEmbedUnimplementedPushServiceServer()
}

func RegisterPushServiceServer(s grpc.ServiceRegistrar, srv PushServiceServer) {
	s.RegisterService(&PushService_ServiceDesc, srv)
}

func _PushService_Broadcast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PushServiceServer).Broadcast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/realtimex.msg.PushService/Broadcast",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PushServiceServer).Broadcast(ctx, req.(*PushMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _PushService_Push_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PushServiceServer).Push(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/realtimex.msg.PushService/Push",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PushServiceServer).Push(ctx, req.(*PushMsg))
	}
	return interceptor(ctx, in, info, handler)
}

// PushService_ServiceDesc is the grpc.ServiceDesc for PushService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PushService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "realtimex.msg.PushService",
	HandlerType: (*PushServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Broadcast",
			Handler:    _PushService_Broadcast_Handler,
		},
		{
			MethodName: "Push",
			Handler:    _PushService_Push_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/msg/msg.proto",
}
