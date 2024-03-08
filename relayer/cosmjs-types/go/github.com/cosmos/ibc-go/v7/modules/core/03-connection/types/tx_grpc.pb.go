// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package types

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

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// ConnectionOpenInit defines a rpc handler method for MsgConnectionOpenInit.
	ConnectionOpenInit(ctx context.Context, in *MsgConnectionOpenInit, opts ...grpc.CallOption) (*MsgConnectionOpenInitResponse, error)
	// ConnectionOpenTry defines a rpc handler method for MsgConnectionOpenTry.
	ConnectionOpenTry(ctx context.Context, in *MsgConnectionOpenTry, opts ...grpc.CallOption) (*MsgConnectionOpenTryResponse, error)
	// ConnectionOpenAck defines a rpc handler method for MsgConnectionOpenAck.
	ConnectionOpenAck(ctx context.Context, in *MsgConnectionOpenAck, opts ...grpc.CallOption) (*MsgConnectionOpenAckResponse, error)
	// ConnectionOpenConfirm defines a rpc handler method for
	// MsgConnectionOpenConfirm.
	ConnectionOpenConfirm(ctx context.Context, in *MsgConnectionOpenConfirm, opts ...grpc.CallOption) (*MsgConnectionOpenConfirmResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) ConnectionOpenInit(ctx context.Context, in *MsgConnectionOpenInit, opts ...grpc.CallOption) (*MsgConnectionOpenInitResponse, error) {
	out := new(MsgConnectionOpenInitResponse)
	err := c.cc.Invoke(ctx, "/ibc.core.connection.v1.Msg/ConnectionOpenInit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ConnectionOpenTry(ctx context.Context, in *MsgConnectionOpenTry, opts ...grpc.CallOption) (*MsgConnectionOpenTryResponse, error) {
	out := new(MsgConnectionOpenTryResponse)
	err := c.cc.Invoke(ctx, "/ibc.core.connection.v1.Msg/ConnectionOpenTry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ConnectionOpenAck(ctx context.Context, in *MsgConnectionOpenAck, opts ...grpc.CallOption) (*MsgConnectionOpenAckResponse, error) {
	out := new(MsgConnectionOpenAckResponse)
	err := c.cc.Invoke(ctx, "/ibc.core.connection.v1.Msg/ConnectionOpenAck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ConnectionOpenConfirm(ctx context.Context, in *MsgConnectionOpenConfirm, opts ...grpc.CallOption) (*MsgConnectionOpenConfirmResponse, error) {
	out := new(MsgConnectionOpenConfirmResponse)
	err := c.cc.Invoke(ctx, "/ibc.core.connection.v1.Msg/ConnectionOpenConfirm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// ConnectionOpenInit defines a rpc handler method for MsgConnectionOpenInit.
	ConnectionOpenInit(context.Context, *MsgConnectionOpenInit) (*MsgConnectionOpenInitResponse, error)
	// ConnectionOpenTry defines a rpc handler method for MsgConnectionOpenTry.
	ConnectionOpenTry(context.Context, *MsgConnectionOpenTry) (*MsgConnectionOpenTryResponse, error)
	// ConnectionOpenAck defines a rpc handler method for MsgConnectionOpenAck.
	ConnectionOpenAck(context.Context, *MsgConnectionOpenAck) (*MsgConnectionOpenAckResponse, error)
	// ConnectionOpenConfirm defines a rpc handler method for
	// MsgConnectionOpenConfirm.
	ConnectionOpenConfirm(context.Context, *MsgConnectionOpenConfirm) (*MsgConnectionOpenConfirmResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) ConnectionOpenInit(context.Context, *MsgConnectionOpenInit) (*MsgConnectionOpenInitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConnectionOpenInit not implemented")
}
func (UnimplementedMsgServer) ConnectionOpenTry(context.Context, *MsgConnectionOpenTry) (*MsgConnectionOpenTryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConnectionOpenTry not implemented")
}
func (UnimplementedMsgServer) ConnectionOpenAck(context.Context, *MsgConnectionOpenAck) (*MsgConnectionOpenAckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConnectionOpenAck not implemented")
}
func (UnimplementedMsgServer) ConnectionOpenConfirm(context.Context, *MsgConnectionOpenConfirm) (*MsgConnectionOpenConfirmResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConnectionOpenConfirm not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_ConnectionOpenInit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgConnectionOpenInit)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ConnectionOpenInit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.core.connection.v1.Msg/ConnectionOpenInit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ConnectionOpenInit(ctx, req.(*MsgConnectionOpenInit))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ConnectionOpenTry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgConnectionOpenTry)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ConnectionOpenTry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.core.connection.v1.Msg/ConnectionOpenTry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ConnectionOpenTry(ctx, req.(*MsgConnectionOpenTry))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ConnectionOpenAck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgConnectionOpenAck)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ConnectionOpenAck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.core.connection.v1.Msg/ConnectionOpenAck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ConnectionOpenAck(ctx, req.(*MsgConnectionOpenAck))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ConnectionOpenConfirm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgConnectionOpenConfirm)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ConnectionOpenConfirm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.core.connection.v1.Msg/ConnectionOpenConfirm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ConnectionOpenConfirm(ctx, req.(*MsgConnectionOpenConfirm))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ibc.core.connection.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ConnectionOpenInit",
			Handler:    _Msg_ConnectionOpenInit_Handler,
		},
		{
			MethodName: "ConnectionOpenTry",
			Handler:    _Msg_ConnectionOpenTry_Handler,
		},
		{
			MethodName: "ConnectionOpenAck",
			Handler:    _Msg_ConnectionOpenAck_Handler,
		},
		{
			MethodName: "ConnectionOpenConfirm",
			Handler:    _Msg_ConnectionOpenConfirm_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ibc/core/connection/v1/tx.proto",
}
