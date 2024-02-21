// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: market/market.proto

package market

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

const (
	Market_RequestFile_FullMethodName   = "/market.Market/RequestFile"
	Market_CheckRequests_FullMethodName = "/market.Market/CheckRequests"
	Market_CheckHolders_FullMethodName  = "/market.Market/CheckHolders"
	Market_RegisterFile_FullMethodName  = "/market.Market/RegisterFile"
)

// MarketClient is the client API for Market service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MarketClient interface {
	RequestFile(ctx context.Context, in *FileRequest, opts ...grpc.CallOption) (*FileResponse, error)
	CheckRequests(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*Requests, error)
	CheckHolders(ctx context.Context, in *CheckHolder, opts ...grpc.CallOption) (*ListReply, error)
	RegisterFile(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type marketClient struct {
	cc grpc.ClientConnInterface
}

func NewMarketClient(cc grpc.ClientConnInterface) MarketClient {
	return &marketClient{cc}
}

func (c *marketClient) RequestFile(ctx context.Context, in *FileRequest, opts ...grpc.CallOption) (*FileResponse, error) {
	out := new(FileResponse)
	err := c.cc.Invoke(ctx, Market_RequestFile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketClient) CheckRequests(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*Requests, error) {
	out := new(Requests)
	err := c.cc.Invoke(ctx, Market_CheckRequests_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketClient) CheckHolders(ctx context.Context, in *CheckHolder, opts ...grpc.CallOption) (*ListReply, error) {
	out := new(ListReply)
	err := c.cc.Invoke(ctx, Market_CheckHolders_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketClient) RegisterFile(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Market_RegisterFile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MarketServer is the server API for Market service.
// All implementations must embed UnimplementedMarketServer
// for forward compatibility
type MarketServer interface {
	RequestFile(context.Context, *FileRequest) (*FileResponse, error)
	CheckRequests(context.Context, *CheckRequest) (*Requests, error)
	CheckHolders(context.Context, *CheckHolder) (*ListReply, error)
	RegisterFile(context.Context, *RegisterRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedMarketServer()
}

// UnimplementedMarketServer must be embedded to have forward compatible implementations.
type UnimplementedMarketServer struct {
}

func (UnimplementedMarketServer) RequestFile(context.Context, *FileRequest) (*FileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestFile not implemented")
}
func (UnimplementedMarketServer) CheckRequests(context.Context, *CheckRequest) (*Requests, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckRequests not implemented")
}
func (UnimplementedMarketServer) CheckHolders(context.Context, *CheckHolder) (*ListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckHolders not implemented")
}
func (UnimplementedMarketServer) RegisterFile(context.Context, *RegisterRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterFile not implemented")
}
func (UnimplementedMarketServer) mustEmbedUnimplementedMarketServer() {}

// UnsafeMarketServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MarketServer will
// result in compilation errors.
type UnsafeMarketServer interface {
	mustEmbedUnimplementedMarketServer()
}

func RegisterMarketServer(s grpc.ServiceRegistrar, srv MarketServer) {
	s.RegisterService(&Market_ServiceDesc, srv)
}

func _Market_RequestFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).RequestFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Market_RequestFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).RequestFile(ctx, req.(*FileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Market_CheckRequests_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).CheckRequests(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Market_CheckRequests_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).CheckRequests(ctx, req.(*CheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Market_CheckHolders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckHolder)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).CheckHolders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Market_CheckHolders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).CheckHolders(ctx, req.(*CheckHolder))
	}
	return interceptor(ctx, in, info, handler)
}

func _Market_RegisterFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).RegisterFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Market_RegisterFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).RegisterFile(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Market_ServiceDesc is the grpc.ServiceDesc for Market service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Market_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "market.Market",
	HandlerType: (*MarketServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestFile",
			Handler:    _Market_RequestFile_Handler,
		},
		{
			MethodName: "CheckRequests",
			Handler:    _Market_CheckRequests_Handler,
		},
		{
			MethodName: "CheckHolders",
			Handler:    _Market_CheckHolders_Handler,
		},
		{
			MethodName: "RegisterFile",
			Handler:    _Market_RegisterFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "market/market.proto",
}
