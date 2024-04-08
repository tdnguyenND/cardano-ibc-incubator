// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: tx-cardano/config.proto

package tx_cardano

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
	ConfigService_UpdatePathConfig_FullMethodName = "/config.ConfigService/UpdatePathConfig"
	ConfigService_ShowPathConfig_FullMethodName   = "/config.ConfigService/ShowPathConfig"
)

// ConfigServiceClient is the client API for ConfigService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConfigServiceClient interface {
	UpdatePathConfig(ctx context.Context, in *UpdatePathConfigRequest, opts ...grpc.CallOption) (*UpdatePathConfigResponse, error)
	ShowPathConfig(ctx context.Context, in *ShowPathConfigRequest, opts ...grpc.CallOption) (*ShowPathConfigResponse, error)
}

type configServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConfigServiceClient(cc grpc.ClientConnInterface) ConfigServiceClient {
	return &configServiceClient{cc}
}

func (c *configServiceClient) UpdatePathConfig(ctx context.Context, in *UpdatePathConfigRequest, opts ...grpc.CallOption) (*UpdatePathConfigResponse, error) {
	out := new(UpdatePathConfigResponse)
	err := c.cc.Invoke(ctx, ConfigService_UpdatePathConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServiceClient) ShowPathConfig(ctx context.Context, in *ShowPathConfigRequest, opts ...grpc.CallOption) (*ShowPathConfigResponse, error) {
	out := new(ShowPathConfigResponse)
	err := c.cc.Invoke(ctx, ConfigService_ShowPathConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConfigServiceServer is the server API for ConfigService service.
// All implementations must embed UnimplementedConfigServiceServer
// for forward compatibility
type ConfigServiceServer interface {
	UpdatePathConfig(context.Context, *UpdatePathConfigRequest) (*UpdatePathConfigResponse, error)
	ShowPathConfig(context.Context, *ShowPathConfigRequest) (*ShowPathConfigResponse, error)
	mustEmbedUnimplementedConfigServiceServer()
}

// UnimplementedConfigServiceServer must be embedded to have forward compatible implementations.
type UnimplementedConfigServiceServer struct {
}

func (UnimplementedConfigServiceServer) UpdatePathConfig(context.Context, *UpdatePathConfigRequest) (*UpdatePathConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePathConfig not implemented")
}
func (UnimplementedConfigServiceServer) ShowPathConfig(context.Context, *ShowPathConfigRequest) (*ShowPathConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowPathConfig not implemented")
}
func (UnimplementedConfigServiceServer) mustEmbedUnimplementedConfigServiceServer() {}

// UnsafeConfigServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConfigServiceServer will
// result in compilation errors.
type UnsafeConfigServiceServer interface {
	mustEmbedUnimplementedConfigServiceServer()
}

func RegisterConfigServiceServer(s grpc.ServiceRegistrar, srv ConfigServiceServer) {
	s.RegisterService(&ConfigService_ServiceDesc, srv)
}

func _ConfigService_UpdatePathConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePathConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).UpdatePathConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConfigService_UpdatePathConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).UpdatePathConfig(ctx, req.(*UpdatePathConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigService_ShowPathConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShowPathConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).ShowPathConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConfigService_ShowPathConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).ShowPathConfig(ctx, req.(*ShowPathConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ConfigService_ServiceDesc is the grpc.ServiceDesc for ConfigService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConfigService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "config.ConfigService",
	HandlerType: (*ConfigServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdatePathConfig",
			Handler:    _ConfigService_UpdatePathConfig_Handler,
		},
		{
			MethodName: "ShowPathConfig",
			Handler:    _ConfigService_ShowPathConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tx-cardano/config.proto",
}
