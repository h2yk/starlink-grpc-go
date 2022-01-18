// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.6
// source: spacex/api/device/wifi.proto

package device

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

// MeshClient is the client API for Mesh service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MeshClient interface {
	MeshStream(ctx context.Context, opts ...grpc.CallOption) (Mesh_MeshStreamClient, error)
}

type meshClient struct {
	cc grpc.ClientConnInterface
}

func NewMeshClient(cc grpc.ClientConnInterface) MeshClient {
	return &meshClient{cc}
}

func (c *meshClient) MeshStream(ctx context.Context, opts ...grpc.CallOption) (Mesh_MeshStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Mesh_ServiceDesc.Streams[0], "/SpaceX.API.Device.Mesh/MeshStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &meshMeshStreamClient{stream}
	return x, nil
}

type Mesh_MeshStreamClient interface {
	Send(*ToController) error
	Recv() (*FromController, error)
	grpc.ClientStream
}

type meshMeshStreamClient struct {
	grpc.ClientStream
}

func (x *meshMeshStreamClient) Send(m *ToController) error {
	return x.ClientStream.SendMsg(m)
}

func (x *meshMeshStreamClient) Recv() (*FromController, error) {
	m := new(FromController)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MeshServer is the server API for Mesh service.
// All implementations must embed UnimplementedMeshServer
// for forward compatibility
type MeshServer interface {
	MeshStream(Mesh_MeshStreamServer) error
	mustEmbedUnimplementedMeshServer()
}

// UnimplementedMeshServer must be embedded to have forward compatible implementations.
type UnimplementedMeshServer struct {
}

func (UnimplementedMeshServer) MeshStream(Mesh_MeshStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method MeshStream not implemented")
}
func (UnimplementedMeshServer) mustEmbedUnimplementedMeshServer() {}

// UnsafeMeshServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MeshServer will
// result in compilation errors.
type UnsafeMeshServer interface {
	mustEmbedUnimplementedMeshServer()
}

func RegisterMeshServer(s grpc.ServiceRegistrar, srv MeshServer) {
	s.RegisterService(&Mesh_ServiceDesc, srv)
}

func _Mesh_MeshStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MeshServer).MeshStream(&meshMeshStreamServer{stream})
}

type Mesh_MeshStreamServer interface {
	Send(*FromController) error
	Recv() (*ToController, error)
	grpc.ServerStream
}

type meshMeshStreamServer struct {
	grpc.ServerStream
}

func (x *meshMeshStreamServer) Send(m *FromController) error {
	return x.ServerStream.SendMsg(m)
}

func (x *meshMeshStreamServer) Recv() (*ToController, error) {
	m := new(ToController)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Mesh_ServiceDesc is the grpc.ServiceDesc for Mesh service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Mesh_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "SpaceX.API.Device.Mesh",
	HandlerType: (*MeshServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "MeshStream",
			Handler:       _Mesh_MeshStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "spacex/api/device/wifi.proto",
}
