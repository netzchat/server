// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: core/v1/channel.proto

package corev1

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

// ChannelServiceClient is the client API for ChannelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChannelServiceClient interface {
	List(ctx context.Context, in *ListChannelsRequest, opts ...grpc.CallOption) (*ListChannelsResponse, error)
	Get(ctx context.Context, in *GetChannelRequest, opts ...grpc.CallOption) (*Channel, error)
	Create(ctx context.Context, in *CreateChannelRequest, opts ...grpc.CallOption) (*Channel, error)
	Update(ctx context.Context, in *UpdateChannelRequest, opts ...grpc.CallOption) (*Channel, error)
	Delete(ctx context.Context, in *DeleteChannelRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type channelServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChannelServiceClient(cc grpc.ClientConnInterface) ChannelServiceClient {
	return &channelServiceClient{cc}
}

func (c *channelServiceClient) List(ctx context.Context, in *ListChannelsRequest, opts ...grpc.CallOption) (*ListChannelsResponse, error) {
	out := new(ListChannelsResponse)
	err := c.cc.Invoke(ctx, "/core.v1.ChannelService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelServiceClient) Get(ctx context.Context, in *GetChannelRequest, opts ...grpc.CallOption) (*Channel, error) {
	out := new(Channel)
	err := c.cc.Invoke(ctx, "/core.v1.ChannelService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelServiceClient) Create(ctx context.Context, in *CreateChannelRequest, opts ...grpc.CallOption) (*Channel, error) {
	out := new(Channel)
	err := c.cc.Invoke(ctx, "/core.v1.ChannelService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelServiceClient) Update(ctx context.Context, in *UpdateChannelRequest, opts ...grpc.CallOption) (*Channel, error) {
	out := new(Channel)
	err := c.cc.Invoke(ctx, "/core.v1.ChannelService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelServiceClient) Delete(ctx context.Context, in *DeleteChannelRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/core.v1.ChannelService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChannelServiceServer is the server API for ChannelService service.
// All implementations must embed UnimplementedChannelServiceServer
// for forward compatibility
type ChannelServiceServer interface {
	List(context.Context, *ListChannelsRequest) (*ListChannelsResponse, error)
	Get(context.Context, *GetChannelRequest) (*Channel, error)
	Create(context.Context, *CreateChannelRequest) (*Channel, error)
	Update(context.Context, *UpdateChannelRequest) (*Channel, error)
	Delete(context.Context, *DeleteChannelRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedChannelServiceServer()
}

// UnimplementedChannelServiceServer must be embedded to have forward compatible implementations.
type UnimplementedChannelServiceServer struct {
}

func (UnimplementedChannelServiceServer) List(context.Context, *ListChannelsRequest) (*ListChannelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedChannelServiceServer) Get(context.Context, *GetChannelRequest) (*Channel, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedChannelServiceServer) Create(context.Context, *CreateChannelRequest) (*Channel, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedChannelServiceServer) Update(context.Context, *UpdateChannelRequest) (*Channel, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedChannelServiceServer) Delete(context.Context, *DeleteChannelRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedChannelServiceServer) mustEmbedUnimplementedChannelServiceServer() {}

// UnsafeChannelServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChannelServiceServer will
// result in compilation errors.
type UnsafeChannelServiceServer interface {
	mustEmbedUnimplementedChannelServiceServer()
}

func RegisterChannelServiceServer(s grpc.ServiceRegistrar, srv ChannelServiceServer) {
	s.RegisterService(&ChannelService_ServiceDesc, srv)
}

func _ChannelService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListChannelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.v1.ChannelService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelServiceServer).List(ctx, req.(*ListChannelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChannelService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.v1.ChannelService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelServiceServer).Get(ctx, req.(*GetChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChannelService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.v1.ChannelService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelServiceServer).Create(ctx, req.(*CreateChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChannelService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.v1.ChannelService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelServiceServer).Update(ctx, req.(*UpdateChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChannelService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.v1.ChannelService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelServiceServer).Delete(ctx, req.(*DeleteChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChannelService_ServiceDesc is the grpc.ServiceDesc for ChannelService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChannelService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "core.v1.ChannelService",
	HandlerType: (*ChannelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _ChannelService_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ChannelService_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _ChannelService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ChannelService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ChannelService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "core/v1/channel.proto",
}