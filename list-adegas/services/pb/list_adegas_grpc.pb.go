// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: list_adegas.proto

package pb

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
	ListAndGetAdegasService_ListAdegas_FullMethodName   = "/ListAndGetAdegasService/ListAdegas"
	ListAndGetAdegasService_GetAdegaById_FullMethodName = "/ListAndGetAdegasService/GetAdegaById"
)

// ListAndGetAdegasServiceClient is the client API for ListAndGetAdegasService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ListAndGetAdegasServiceClient interface {
	ListAdegas(ctx context.Context, in *ListAdegasFilter, opts ...grpc.CallOption) (grpc.ServerStreamingClient[AdegaResponse], error)
	GetAdegaById(ctx context.Context, in *GetAdegasByIdFilter, opts ...grpc.CallOption) (*AdegaResponse, error)
}

type listAndGetAdegasServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewListAndGetAdegasServiceClient(cc grpc.ClientConnInterface) ListAndGetAdegasServiceClient {
	return &listAndGetAdegasServiceClient{cc}
}

func (c *listAndGetAdegasServiceClient) ListAdegas(ctx context.Context, in *ListAdegasFilter, opts ...grpc.CallOption) (grpc.ServerStreamingClient[AdegaResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ListAndGetAdegasService_ServiceDesc.Streams[0], ListAndGetAdegasService_ListAdegas_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[ListAdegasFilter, AdegaResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ListAndGetAdegasService_ListAdegasClient = grpc.ServerStreamingClient[AdegaResponse]

func (c *listAndGetAdegasServiceClient) GetAdegaById(ctx context.Context, in *GetAdegasByIdFilter, opts ...grpc.CallOption) (*AdegaResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AdegaResponse)
	err := c.cc.Invoke(ctx, ListAndGetAdegasService_GetAdegaById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ListAndGetAdegasServiceServer is the server API for ListAndGetAdegasService service.
// All implementations must embed UnimplementedListAndGetAdegasServiceServer
// for forward compatibility.
type ListAndGetAdegasServiceServer interface {
	ListAdegas(*ListAdegasFilter, grpc.ServerStreamingServer[AdegaResponse]) error
	GetAdegaById(context.Context, *GetAdegasByIdFilter) (*AdegaResponse, error)
	mustEmbedUnimplementedListAndGetAdegasServiceServer()
}

// UnimplementedListAndGetAdegasServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedListAndGetAdegasServiceServer struct{}

func (UnimplementedListAndGetAdegasServiceServer) ListAdegas(*ListAdegasFilter, grpc.ServerStreamingServer[AdegaResponse]) error {
	return status.Errorf(codes.Unimplemented, "method ListAdegas not implemented")
}
func (UnimplementedListAndGetAdegasServiceServer) GetAdegaById(context.Context, *GetAdegasByIdFilter) (*AdegaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAdegaById not implemented")
}
func (UnimplementedListAndGetAdegasServiceServer) mustEmbedUnimplementedListAndGetAdegasServiceServer() {
}
func (UnimplementedListAndGetAdegasServiceServer) testEmbeddedByValue() {}

// UnsafeListAndGetAdegasServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ListAndGetAdegasServiceServer will
// result in compilation errors.
type UnsafeListAndGetAdegasServiceServer interface {
	mustEmbedUnimplementedListAndGetAdegasServiceServer()
}

func RegisterListAndGetAdegasServiceServer(s grpc.ServiceRegistrar, srv ListAndGetAdegasServiceServer) {
	// If the following call pancis, it indicates UnimplementedListAndGetAdegasServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ListAndGetAdegasService_ServiceDesc, srv)
}

func _ListAndGetAdegasService_ListAdegas_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListAdegasFilter)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ListAndGetAdegasServiceServer).ListAdegas(m, &grpc.GenericServerStream[ListAdegasFilter, AdegaResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ListAndGetAdegasService_ListAdegasServer = grpc.ServerStreamingServer[AdegaResponse]

func _ListAndGetAdegasService_GetAdegaById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAdegasByIdFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListAndGetAdegasServiceServer).GetAdegaById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ListAndGetAdegasService_GetAdegaById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListAndGetAdegasServiceServer).GetAdegaById(ctx, req.(*GetAdegasByIdFilter))
	}
	return interceptor(ctx, in, info, handler)
}

// ListAndGetAdegasService_ServiceDesc is the grpc.ServiceDesc for ListAndGetAdegasService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ListAndGetAdegasService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ListAndGetAdegasService",
	HandlerType: (*ListAndGetAdegasServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAdegaById",
			Handler:    _ListAndGetAdegasService_GetAdegaById_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListAdegas",
			Handler:       _ListAndGetAdegasService_ListAdegas_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "list_adegas.proto",
}
