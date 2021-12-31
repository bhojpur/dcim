// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// DcimServiceClient is the client API for DcimService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DcimServiceClient interface {
	// StartLocalRack starts a Rack on the Bhojpur.NET Platform directly.
	// The incoming requests are expected in the following order:
	//   1. metadata
	//   2. all bytes constituting the dcim/config.yaml
	//   3. all bytes constituting the Rack YAML that will be executed (that the config.yaml points to)
	//   4. all bytes constituting the gzipped Bhojpur.NET Platform application tar stream
	//   5. the Bhojpur.NET Platform application tar stream done marker
	StartLocalRack(ctx context.Context, opts ...grpc.CallOption) (DcimService_StartLocalRackClient, error)
	// StartFromPreviousRack starts a new Rack based on a previous one.
	// If the previous Rack does not have the can-replay condition set this call will result in an error.
	StartFromPreviousRack(ctx context.Context, in *StartFromPreviousRackRequest, opts ...grpc.CallOption) (*StartRackResponse, error)
	// StartRackRequest starts a new Rack based on its specification.
	StartRack(ctx context.Context, in *StartRackRequest, opts ...grpc.CallOption) (*StartRackResponse, error)
	// Searches for Rack(s) known to this instance
	ListRacks(ctx context.Context, in *ListRacksRequest, opts ...grpc.CallOption) (*ListRacksResponse, error)
	// Subscribe listens to new Rack(s) updates
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (DcimService_SubscribeClient, error)
	// GetRack retrieves details of a single Rack
	GetRack(ctx context.Context, in *GetRackRequest, opts ...grpc.CallOption) (*GetRackResponse, error)
	// Listen listens to Rack updates and log output of a running Rack
	Listen(ctx context.Context, in *ListenRequest, opts ...grpc.CallOption) (DcimService_ListenClient, error)
	// StopRack stops a currently running Rack
	StopRack(ctx context.Context, in *StopRackRequest, opts ...grpc.CallOption) (*StopRackResponse, error)
}

type dcimServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDcimServiceClient(cc grpc.ClientConnInterface) DcimServiceClient {
	return &dcimServiceClient{cc}
}

func (c *dcimServiceClient) StartLocalRack(ctx context.Context, opts ...grpc.CallOption) (DcimService_StartLocalRackClient, error) {
	stream, err := c.cc.NewStream(ctx, &DcimService_ServiceDesc.Streams[0], "/v1.DcimService/StartLocalRack", opts...)
	if err != nil {
		return nil, err
	}
	x := &dcimServiceStartLocalRackClient{stream}
	return x, nil
}

type DcimService_StartLocalRackClient interface {
	Send(*StartLocalRackRequest) error
	CloseAndRecv() (*StartRackResponse, error)
	grpc.ClientStream
}

type dcimServiceStartLocalRackClient struct {
	grpc.ClientStream
}

func (x *dcimServiceStartLocalRackClient) Send(m *StartLocalRackRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dcimServiceStartLocalRackClient) CloseAndRecv() (*StartRackResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StartRackResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dcimServiceClient) StartFromPreviousRack(ctx context.Context, in *StartFromPreviousRackRequest, opts ...grpc.CallOption) (*StartRackResponse, error) {
	out := new(StartRackResponse)
	err := c.cc.Invoke(ctx, "/v1.DcimService/StartFromPreviousRack", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dcimServiceClient) StartRack(ctx context.Context, in *StartRackRequest, opts ...grpc.CallOption) (*StartRackResponse, error) {
	out := new(StartRackResponse)
	err := c.cc.Invoke(ctx, "/v1.DcimService/StartRack", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dcimServiceClient) ListRacks(ctx context.Context, in *ListRacksRequest, opts ...grpc.CallOption) (*ListRacksResponse, error) {
	out := new(ListRacksResponse)
	err := c.cc.Invoke(ctx, "/v1.DcimService/ListRacks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dcimServiceClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (DcimService_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &DcimService_ServiceDesc.Streams[1], "/v1.DcimService/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &dcimServiceSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DcimService_SubscribeClient interface {
	Recv() (*SubscribeResponse, error)
	grpc.ClientStream
}

type dcimServiceSubscribeClient struct {
	grpc.ClientStream
}

func (x *dcimServiceSubscribeClient) Recv() (*SubscribeResponse, error) {
	m := new(SubscribeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dcimServiceClient) GetRack(ctx context.Context, in *GetRackRequest, opts ...grpc.CallOption) (*GetRackResponse, error) {
	out := new(GetRackResponse)
	err := c.cc.Invoke(ctx, "/v1.DcimService/GetRack", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dcimServiceClient) Listen(ctx context.Context, in *ListenRequest, opts ...grpc.CallOption) (DcimService_ListenClient, error) {
	stream, err := c.cc.NewStream(ctx, &DcimService_ServiceDesc.Streams[2], "/v1.DcimService/Listen", opts...)
	if err != nil {
		return nil, err
	}
	x := &dcimServiceListenClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DcimService_ListenClient interface {
	Recv() (*ListenResponse, error)
	grpc.ClientStream
}

type dcimServiceListenClient struct {
	grpc.ClientStream
}

func (x *dcimServiceListenClient) Recv() (*ListenResponse, error) {
	m := new(ListenResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dcimServiceClient) StopRack(ctx context.Context, in *StopRackRequest, opts ...grpc.CallOption) (*StopRackResponse, error) {
	out := new(StopRackResponse)
	err := c.cc.Invoke(ctx, "/v1.DcimService/StopRack", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DcimServiceServer is the server API for DcimService service.
// All implementations must embed UnimplementedDcimServiceServer
// for forward compatibility
type DcimServiceServer interface {
	// StartLocalRack starts a Rack on the Bhojpur.NET Platform directly.
	// The incoming requests are expected in the following order:
	//   1. metadata
	//   2. all bytes constituting the dcim/config.yaml
	//   3. all bytes constituting the Rack YAML that will be executed (that the config.yaml points to)
	//   4. all bytes constituting the gzipped Bhojpur.NET Platform application tar stream
	//   5. the Bhojpur.NET Platform application tar stream done marker
	StartLocalRack(DcimService_StartLocalRackServer) error
	// StartFromPreviousRack starts a new Rack based on a previous one.
	// If the previous Rack does not have the can-replay condition set this call will result in an error.
	StartFromPreviousRack(context.Context, *StartFromPreviousRackRequest) (*StartRackResponse, error)
	// StartRackRequest starts a new Rack based on its specification.
	StartRack(context.Context, *StartRackRequest) (*StartRackResponse, error)
	// Searches for Rack(s) known to this instance
	ListRacks(context.Context, *ListRacksRequest) (*ListRacksResponse, error)
	// Subscribe listens to new Rack(s) updates
	Subscribe(*SubscribeRequest, DcimService_SubscribeServer) error
	// GetRack retrieves details of a single Rack
	GetRack(context.Context, *GetRackRequest) (*GetRackResponse, error)
	// Listen listens to Rack updates and log output of a running Rack
	Listen(*ListenRequest, DcimService_ListenServer) error
	// StopRack stops a currently running Rack
	StopRack(context.Context, *StopRackRequest) (*StopRackResponse, error)
	mustEmbedUnimplementedDcimServiceServer()
}

// UnimplementedDcimServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDcimServiceServer struct {
}

func (UnimplementedDcimServiceServer) StartLocalRack(DcimService_StartLocalRackServer) error {
	return status.Errorf(codes.Unimplemented, "method StartLocalRack not implemented")
}
func (UnimplementedDcimServiceServer) StartFromPreviousRack(context.Context, *StartFromPreviousRackRequest) (*StartRackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartFromPreviousRack not implemented")
}
func (UnimplementedDcimServiceServer) StartRack(context.Context, *StartRackRequest) (*StartRackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartRack not implemented")
}
func (UnimplementedDcimServiceServer) ListRacks(context.Context, *ListRacksRequest) (*ListRacksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRacks not implemented")
}
func (UnimplementedDcimServiceServer) Subscribe(*SubscribeRequest, DcimService_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedDcimServiceServer) GetRack(context.Context, *GetRackRequest) (*GetRackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRack not implemented")
}
func (UnimplementedDcimServiceServer) Listen(*ListenRequest, DcimService_ListenServer) error {
	return status.Errorf(codes.Unimplemented, "method Listen not implemented")
}
func (UnimplementedDcimServiceServer) StopRack(context.Context, *StopRackRequest) (*StopRackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopRack not implemented")
}
func (UnimplementedDcimServiceServer) mustEmbedUnimplementedDcimServiceServer() {}

// UnsafeDcimServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DcimServiceServer will
// result in compilation errors.
type UnsafeDcimServiceServer interface {
	mustEmbedUnimplementedDcimServiceServer()
}

func RegisterDcimServiceServer(s grpc.ServiceRegistrar, srv DcimServiceServer) {
	s.RegisterService(&DcimService_ServiceDesc, srv)
}

func _DcimService_StartLocalRack_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DcimServiceServer).StartLocalRack(&dcimServiceStartLocalRackServer{stream})
}

type DcimService_StartLocalRackServer interface {
	SendAndClose(*StartRackResponse) error
	Recv() (*StartLocalRackRequest, error)
	grpc.ServerStream
}

type dcimServiceStartLocalRackServer struct {
	grpc.ServerStream
}

func (x *dcimServiceStartLocalRackServer) SendAndClose(m *StartRackResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dcimServiceStartLocalRackServer) Recv() (*StartLocalRackRequest, error) {
	m := new(StartLocalRackRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _DcimService_StartFromPreviousRack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartFromPreviousRackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DcimServiceServer).StartFromPreviousRack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.DcimService/StartFromPreviousRack",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DcimServiceServer).StartFromPreviousRack(ctx, req.(*StartFromPreviousRackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DcimService_StartRack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartRackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DcimServiceServer).StartRack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.DcimService/StartRack",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DcimServiceServer).StartRack(ctx, req.(*StartRackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DcimService_ListRacks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRacksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DcimServiceServer).ListRacks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.DcimService/ListRacks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DcimServiceServer).ListRacks(ctx, req.(*ListRacksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DcimService_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DcimServiceServer).Subscribe(m, &dcimServiceSubscribeServer{stream})
}

type DcimService_SubscribeServer interface {
	Send(*SubscribeResponse) error
	grpc.ServerStream
}

type dcimServiceSubscribeServer struct {
	grpc.ServerStream
}

func (x *dcimServiceSubscribeServer) Send(m *SubscribeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _DcimService_GetRack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DcimServiceServer).GetRack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.DcimService/GetRack",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DcimServiceServer).GetRack(ctx, req.(*GetRackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DcimService_Listen_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListenRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DcimServiceServer).Listen(m, &dcimServiceListenServer{stream})
}

type DcimService_ListenServer interface {
	Send(*ListenResponse) error
	grpc.ServerStream
}

type dcimServiceListenServer struct {
	grpc.ServerStream
}

func (x *dcimServiceListenServer) Send(m *ListenResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _DcimService_StopRack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopRackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DcimServiceServer).StopRack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.DcimService/StopRack",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DcimServiceServer).StopRack(ctx, req.(*StopRackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DcimService_ServiceDesc is the grpc.ServiceDesc for DcimService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DcimService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.DcimService",
	HandlerType: (*DcimServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartFromPreviousRack",
			Handler:    _DcimService_StartFromPreviousRack_Handler,
		},
		{
			MethodName: "StartRack",
			Handler:    _DcimService_StartRack_Handler,
		},
		{
			MethodName: "ListRacks",
			Handler:    _DcimService_ListRacks_Handler,
		},
		{
			MethodName: "GetRack",
			Handler:    _DcimService_GetRack_Handler,
		},
		{
			MethodName: "StopRack",
			Handler:    _DcimService_StopRack_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StartLocalRack",
			Handler:       _DcimService_StartLocalRack_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Subscribe",
			Handler:       _DcimService_Subscribe_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Listen",
			Handler:       _DcimService_Listen_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "dcim.proto",
}
