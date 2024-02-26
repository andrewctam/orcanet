// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.2
// source: market.proto

package market

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

// MarketServiceClient is the client API for MarketService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MarketServiceClient interface {
	RequestQuery(ctx context.Context, opts ...grpc.CallOption) (MarketService_RequestQueryClient, error)
}

type marketServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMarketServiceClient(cc grpc.ClientConnInterface) MarketServiceClient {
	return &marketServiceClient{cc}
}

func (c *marketServiceClient) RequestQuery(ctx context.Context, opts ...grpc.CallOption) (MarketService_RequestQueryClient, error) {
	stream, err := c.cc.NewStream(ctx, &MarketService_ServiceDesc.Streams[0], "/market.MarketService/RequestQuery", opts...)
	if err != nil {
		return nil, err
	}
	x := &marketServiceRequestQueryClient{stream}
	return x, nil
}

type MarketService_RequestQueryClient interface {
	Send(*Request) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type marketServiceRequestQueryClient struct {
	grpc.ClientStream
}

func (x *marketServiceRequestQueryClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *marketServiceRequestQueryClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MarketServiceServer is the server API for MarketService service.
// All implementations must embed UnimplementedMarketServiceServer
// for forward compatibility
type MarketServiceServer interface {
	RequestQuery(MarketService_RequestQueryServer) error
	mustEmbedUnimplementedMarketServiceServer()
}

// UnimplementedMarketServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMarketServiceServer struct {
}

func (UnimplementedMarketServiceServer) RequestQuery(MarketService_RequestQueryServer) error {
	return status.Errorf(codes.Unimplemented, "method RequestQuery not implemented")
}
func (UnimplementedMarketServiceServer) mustEmbedUnimplementedMarketServiceServer() {}

// UnsafeMarketServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MarketServiceServer will
// result in compilation errors.
type UnsafeMarketServiceServer interface {
	mustEmbedUnimplementedMarketServiceServer()
}

func RegisterMarketServiceServer(s grpc.ServiceRegistrar, srv MarketServiceServer) {
	s.RegisterService(&MarketService_ServiceDesc, srv)
}

func _MarketService_RequestQuery_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MarketServiceServer).RequestQuery(&marketServiceRequestQueryServer{stream})
}

type MarketService_RequestQueryServer interface {
	Send(*Response) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type marketServiceRequestQueryServer struct {
	grpc.ServerStream
}

func (x *marketServiceRequestQueryServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *marketServiceRequestQueryServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MarketService_ServiceDesc is the grpc.ServiceDesc for MarketService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MarketService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "market.MarketService",
	HandlerType: (*MarketServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RequestQuery",
			Handler:       _MarketService_RequestQuery_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "market.proto",
}
