// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: orchard.proto

package rpc

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
	Controller_Watch_FullMethodName       = "/Controller/Watch"
	Controller_PortForward_FullMethodName = "/Controller/PortForward"
)

// ControllerClient is the client API for Controller service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ControllerClient interface {
	Watch(ctx context.Context, opts ...grpc.CallOption) (Controller_WatchClient, error)
	PortForward(ctx context.Context, opts ...grpc.CallOption) (Controller_PortForwardClient, error)
}

type controllerClient struct {
	cc grpc.ClientConnInterface
}

func NewControllerClient(cc grpc.ClientConnInterface) ControllerClient {
	return &controllerClient{cc}
}

func (c *controllerClient) Watch(ctx context.Context, opts ...grpc.CallOption) (Controller_WatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &Controller_ServiceDesc.Streams[0], Controller_Watch_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &controllerWatchClient{stream}
	return x, nil
}

type Controller_WatchClient interface {
	Send(*WatchFromWorker) error
	Recv() (*WatchFromController, error)
	grpc.ClientStream
}

type controllerWatchClient struct {
	grpc.ClientStream
}

func (x *controllerWatchClient) Send(m *WatchFromWorker) error {
	return x.ClientStream.SendMsg(m)
}

func (x *controllerWatchClient) Recv() (*WatchFromController, error) {
	m := new(WatchFromController)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *controllerClient) PortForward(ctx context.Context, opts ...grpc.CallOption) (Controller_PortForwardClient, error) {
	stream, err := c.cc.NewStream(ctx, &Controller_ServiceDesc.Streams[1], Controller_PortForward_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &controllerPortForwardClient{stream}
	return x, nil
}

type Controller_PortForwardClient interface {
	Send(*PortForwardFromWorker) error
	Recv() (*PortForwardFromController, error)
	grpc.ClientStream
}

type controllerPortForwardClient struct {
	grpc.ClientStream
}

func (x *controllerPortForwardClient) Send(m *PortForwardFromWorker) error {
	return x.ClientStream.SendMsg(m)
}

func (x *controllerPortForwardClient) Recv() (*PortForwardFromController, error) {
	m := new(PortForwardFromController)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ControllerServer is the server API for Controller service.
// All implementations must embed UnimplementedControllerServer
// for forward compatibility
type ControllerServer interface {
	Watch(Controller_WatchServer) error
	PortForward(Controller_PortForwardServer) error
	mustEmbedUnimplementedControllerServer()
}

// UnimplementedControllerServer must be embedded to have forward compatible implementations.
type UnimplementedControllerServer struct {
}

func (UnimplementedControllerServer) Watch(Controller_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "method Watch not implemented")
}
func (UnimplementedControllerServer) PortForward(Controller_PortForwardServer) error {
	return status.Errorf(codes.Unimplemented, "method PortForward not implemented")
}
func (UnimplementedControllerServer) mustEmbedUnimplementedControllerServer() {}

// UnsafeControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ControllerServer will
// result in compilation errors.
type UnsafeControllerServer interface {
	mustEmbedUnimplementedControllerServer()
}

func RegisterControllerServer(s grpc.ServiceRegistrar, srv ControllerServer) {
	s.RegisterService(&Controller_ServiceDesc, srv)
}

func _Controller_Watch_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ControllerServer).Watch(&controllerWatchServer{stream})
}

type Controller_WatchServer interface {
	Send(*WatchFromController) error
	Recv() (*WatchFromWorker, error)
	grpc.ServerStream
}

type controllerWatchServer struct {
	grpc.ServerStream
}

func (x *controllerWatchServer) Send(m *WatchFromController) error {
	return x.ServerStream.SendMsg(m)
}

func (x *controllerWatchServer) Recv() (*WatchFromWorker, error) {
	m := new(WatchFromWorker)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Controller_PortForward_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ControllerServer).PortForward(&controllerPortForwardServer{stream})
}

type Controller_PortForwardServer interface {
	Send(*PortForwardFromController) error
	Recv() (*PortForwardFromWorker, error)
	grpc.ServerStream
}

type controllerPortForwardServer struct {
	grpc.ServerStream
}

func (x *controllerPortForwardServer) Send(m *PortForwardFromController) error {
	return x.ServerStream.SendMsg(m)
}

func (x *controllerPortForwardServer) Recv() (*PortForwardFromWorker, error) {
	m := new(PortForwardFromWorker)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Controller_ServiceDesc is the grpc.ServiceDesc for Controller service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Controller_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Controller",
	HandlerType: (*ControllerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Watch",
			Handler:       _Controller_Watch_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "PortForward",
			Handler:       _Controller_PortForward_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "orchard.proto",
}