// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: calculator.proto

package proto

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

// CalculatorServiceClient is the client API for CalculatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalculatorServiceClient interface {
	Calculator(ctx context.Context, in *CalculatorRequest, opts ...grpc.CallOption) (*CalculatorResponse, error)
	CalculatorPrimes(ctx context.Context, in *CalculatorRequest, opts ...grpc.CallOption) (CalculatorService_CalculatorPrimesClient, error)
	CalculatorAvg(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_CalculatorAvgClient, error)
}

type calculatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculatorServiceClient(cc grpc.ClientConnInterface) CalculatorServiceClient {
	return &calculatorServiceClient{cc}
}

func (c *calculatorServiceClient) Calculator(ctx context.Context, in *CalculatorRequest, opts ...grpc.CallOption) (*CalculatorResponse, error) {
	out := new(CalculatorResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalculatorService/Calculator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) CalculatorPrimes(ctx context.Context, in *CalculatorRequest, opts ...grpc.CallOption) (CalculatorService_CalculatorPrimesClient, error) {
	stream, err := c.cc.NewStream(ctx, &CalculatorService_ServiceDesc.Streams[0], "/calculator.CalculatorService/CalculatorPrimes", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceCalculatorPrimesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalculatorService_CalculatorPrimesClient interface {
	Recv() (*CalculatorResponse, error)
	grpc.ClientStream
}

type calculatorServiceCalculatorPrimesClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceCalculatorPrimesClient) Recv() (*CalculatorResponse, error) {
	m := new(CalculatorResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) CalculatorAvg(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_CalculatorAvgClient, error) {
	stream, err := c.cc.NewStream(ctx, &CalculatorService_ServiceDesc.Streams[1], "/calculator.CalculatorService/CalculatorAvg", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceCalculatorAvgClient{stream}
	return x, nil
}

type CalculatorService_CalculatorAvgClient interface {
	Send(*CalculatorRequest) error
	CloseAndRecv() (*CalculatorAvgResponse, error)
	grpc.ClientStream
}

type calculatorServiceCalculatorAvgClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceCalculatorAvgClient) Send(m *CalculatorRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorServiceCalculatorAvgClient) CloseAndRecv() (*CalculatorAvgResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(CalculatorAvgResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalculatorServiceServer is the server API for CalculatorService service.
// All implementations must embed UnimplementedCalculatorServiceServer
// for forward compatibility
type CalculatorServiceServer interface {
	Calculator(context.Context, *CalculatorRequest) (*CalculatorResponse, error)
	CalculatorPrimes(*CalculatorRequest, CalculatorService_CalculatorPrimesServer) error
	CalculatorAvg(CalculatorService_CalculatorAvgServer) error
	mustEmbedUnimplementedCalculatorServiceServer()
}

// UnimplementedCalculatorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCalculatorServiceServer struct {
}

func (UnimplementedCalculatorServiceServer) Calculator(context.Context, *CalculatorRequest) (*CalculatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Calculator not implemented")
}
func (UnimplementedCalculatorServiceServer) CalculatorPrimes(*CalculatorRequest, CalculatorService_CalculatorPrimesServer) error {
	return status.Errorf(codes.Unimplemented, "method CalculatorPrimes not implemented")
}
func (UnimplementedCalculatorServiceServer) CalculatorAvg(CalculatorService_CalculatorAvgServer) error {
	return status.Errorf(codes.Unimplemented, "method CalculatorAvg not implemented")
}
func (UnimplementedCalculatorServiceServer) mustEmbedUnimplementedCalculatorServiceServer() {}

// UnsafeCalculatorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalculatorServiceServer will
// result in compilation errors.
type UnsafeCalculatorServiceServer interface {
	mustEmbedUnimplementedCalculatorServiceServer()
}

func RegisterCalculatorServiceServer(s grpc.ServiceRegistrar, srv CalculatorServiceServer) {
	s.RegisterService(&CalculatorService_ServiceDesc, srv)
}

func _CalculatorService_Calculator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Calculator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculatorService/Calculator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Calculator(ctx, req.(*CalculatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_CalculatorPrimes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CalculatorRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalculatorServiceServer).CalculatorPrimes(m, &calculatorServiceCalculatorPrimesServer{stream})
}

type CalculatorService_CalculatorPrimesServer interface {
	Send(*CalculatorResponse) error
	grpc.ServerStream
}

type calculatorServiceCalculatorPrimesServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceCalculatorPrimesServer) Send(m *CalculatorResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CalculatorService_CalculatorAvg_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).CalculatorAvg(&calculatorServiceCalculatorAvgServer{stream})
}

type CalculatorService_CalculatorAvgServer interface {
	SendAndClose(*CalculatorAvgResponse) error
	Recv() (*CalculatorRequest, error)
	grpc.ServerStream
}

type calculatorServiceCalculatorAvgServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceCalculatorAvgServer) SendAndClose(m *CalculatorAvgResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorServiceCalculatorAvgServer) Recv() (*CalculatorRequest, error) {
	m := new(CalculatorRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalculatorService_ServiceDesc is the grpc.ServiceDesc for CalculatorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CalculatorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.CalculatorService",
	HandlerType: (*CalculatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Calculator",
			Handler:    _CalculatorService_Calculator_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CalculatorPrimes",
			Handler:       _CalculatorService_CalculatorPrimes_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "CalculatorAvg",
			Handler:       _CalculatorService_CalculatorAvg_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "calculator.proto",
}
