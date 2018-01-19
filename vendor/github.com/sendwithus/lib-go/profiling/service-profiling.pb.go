// Code generated by protoc-gen-go.
// source: service-profiling.proto
// DO NOT EDIT!

/*
Package profiling is a generated protocol buffer package.

It is generated from these files:
	service-profiling.proto

It has these top-level messages:
	Empty
	RegisterRequest
*/
package profiling

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type RegisterRequest struct {
	Name    string `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	DataUrl string `protobuf:"bytes,2,opt,name=DataUrl" json:"DataUrl,omitempty"`
	App     string `protobuf:"bytes,3,opt,name=App" json:"App,omitempty"`
}

func (m *RegisterRequest) Reset()                    { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string            { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()               {}
func (*RegisterRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RegisterRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RegisterRequest) GetDataUrl() string {
	if m != nil {
		return m.DataUrl
	}
	return ""
}

func (m *RegisterRequest) GetApp() string {
	if m != nil {
		return m.App
	}
	return ""
}

func init() {
	proto.RegisterType((*Empty)(nil), "profiling.Empty")
	proto.RegisterType((*RegisterRequest)(nil), "profiling.RegisterRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ServiceProfiling service

type ServiceProfilingClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*Empty, error)
}

type serviceProfilingClient struct {
	cc *grpc.ClientConn
}

func NewServiceProfilingClient(cc *grpc.ClientConn) ServiceProfilingClient {
	return &serviceProfilingClient{cc}
}

func (c *serviceProfilingClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/profiling.ServiceProfiling/Register", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ServiceProfiling service

type ServiceProfilingServer interface {
	Register(context.Context, *RegisterRequest) (*Empty, error)
}

func RegisterServiceProfilingServer(s *grpc.Server, srv ServiceProfilingServer) {
	s.RegisterService(&_ServiceProfiling_serviceDesc, srv)
}

func _ServiceProfiling_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceProfilingServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profiling.ServiceProfiling/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceProfilingServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ServiceProfiling_serviceDesc = grpc.ServiceDesc{
	ServiceName: "profiling.ServiceProfiling",
	HandlerType: (*ServiceProfilingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _ServiceProfiling_Register_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service-profiling.proto",
}

func init() { proto.RegisterFile("service-profiling.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 162 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2d, 0x28, 0xca, 0x4f, 0xcb, 0xcc, 0xc9, 0xcc, 0x4b, 0xd7, 0x03, 0xb2,
	0x4a, 0xf2, 0x85, 0x38, 0xe1, 0x02, 0x4a, 0xec, 0x5c, 0xac, 0xae, 0xb9, 0x05, 0x25, 0x95, 0x4a,
	0x81, 0x5c, 0xfc, 0x41, 0xa9, 0xe9, 0x99, 0xc5, 0x25, 0xa9, 0x45, 0x41, 0xa9, 0x85, 0xa5, 0xa9,
	0xc5, 0x25, 0x42, 0x42, 0x5c, 0x2c, 0x7e, 0x89, 0xb9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c,
	0x41, 0x60, 0xb6, 0x90, 0x04, 0x17, 0xbb, 0x4b, 0x62, 0x49, 0x62, 0x68, 0x51, 0x8e, 0x04, 0x13,
	0x58, 0x18, 0xc6, 0x15, 0x12, 0xe0, 0x62, 0x76, 0x2c, 0x28, 0x90, 0x60, 0x06, 0x8b, 0x82, 0x98,
	0x46, 0x7e, 0x5c, 0x02, 0xc1, 0x10, 0x17, 0x04, 0xc0, 0xec, 0x13, 0xb2, 0xe2, 0xe2, 0x80, 0x59,
	0x23, 0x24, 0xa5, 0x87, 0x70, 0x18, 0x9a, 0xdd, 0x52, 0x02, 0x48, 0x72, 0x10, 0x07, 0x32, 0x24,
	0xb1, 0x81, 0x5d, 0x6f, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xf1, 0x71, 0x07, 0xb2, 0xd8, 0x00,
	0x00, 0x00,
}
