/*
 * Copyright (c) Facebook, Inc. and its affiliates.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package grpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_go "github.com/prometheus/client_model/go"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type MetricFamilies struct {
	Families             []*_go.MetricFamily `protobuf:"bytes,1,rep,name=families,proto3" json:"families,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *MetricFamilies) Reset()         { *m = MetricFamilies{} }
func (m *MetricFamilies) String() string { return proto.CompactTextString(m) }
func (*MetricFamilies) ProtoMessage()    {}
func (*MetricFamilies) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *MetricFamilies) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricFamilies.Unmarshal(m, b)
}
func (m *MetricFamilies) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricFamilies.Marshal(b, m, deterministic)
}
func (m *MetricFamilies) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricFamilies.Merge(m, src)
}
func (m *MetricFamilies) XXX_Size() int {
	return xxx_messageInfo_MetricFamilies.Size(m)
}
func (m *MetricFamilies) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricFamilies.DiscardUnknown(m)
}

var xxx_messageInfo_MetricFamilies proto.InternalMessageInfo

func (m *MetricFamilies) GetFamilies() []*_go.MetricFamily {
	if m != nil {
		return m.Families
	}
	return nil
}

type Void struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Void) Reset()         { *m = Void{} }
func (m *Void) String() string { return proto.CompactTextString(m) }
func (*Void) ProtoMessage()    {}
func (*Void) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *Void) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Void.Unmarshal(m, b)
}
func (m *Void) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Void.Marshal(b, m, deterministic)
}
func (m *Void) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Void.Merge(m, src)
}
func (m *Void) XXX_Size() int {
	return xxx_messageInfo_Void.Size(m)
}
func (m *Void) XXX_DiscardUnknown() {
	xxx_messageInfo_Void.DiscardUnknown(m)
}

var xxx_messageInfo_Void proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MetricFamilies)(nil), "grpc.MetricFamilies")
	proto.RegisterType((*Void)(nil), "grpc.Void")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x49, 0x2f, 0x2a, 0x48, 0x96,
	0x92, 0x2c, 0xc9, 0xc8, 0x2c, 0x4a, 0xd1, 0x2d, 0x48, 0x2c, 0x2a, 0xa9, 0xd4, 0xcf, 0x4d, 0x2d,
	0x29, 0xca, 0x4c, 0x2e, 0x86, 0x28, 0x50, 0x0a, 0xe0, 0xe2, 0xf3, 0x05, 0x0b, 0xb8, 0x25, 0xe6,
	0x66, 0xe6, 0x64, 0xa6, 0x16, 0x0b, 0xd9, 0x71, 0x71, 0xa4, 0x41, 0xd9, 0x12, 0x8c, 0x0a, 0xcc,
	0x1a, 0xdc, 0x46, 0x4a, 0x7a, 0x99, 0xf9, 0x20, 0xe5, 0xb9, 0xa9, 0x25, 0x19, 0xa9, 0xa5, 0xc5,
	0x7a, 0xc9, 0x39, 0x99, 0xa9, 0x79, 0x25, 0x7a, 0x48, 0xfa, 0x2a, 0x83, 0xe0, 0x7a, 0x94, 0xd8,
	0xb8, 0x58, 0xc2, 0xf2, 0x33, 0x53, 0x8c, 0x9c, 0xb8, 0x04, 0x21, 0x2a, 0x8a, 0x9d, 0xf3, 0xf3,
	0x4a, 0x8a, 0xf2, 0x73, 0x72, 0x52, 0x8b, 0x84, 0x74, 0xb9, 0xd8, 0x9d, 0x41, 0xac, 0xe4, 0x12,
	0x21, 0x11, 0x3d, 0x90, 0xdb, 0xf4, 0x50, 0x6d, 0x97, 0xe2, 0x82, 0x88, 0x82, 0x4c, 0x50, 0x62,
	0x48, 0x62, 0x03, 0x3b, 0xd2, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x9e, 0xf7, 0x07, 0xc1, 0xd6,
	0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MetricsControllerClient is the client API for MetricsController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MetricsControllerClient interface {
	// Report a collection of metrics from a service
	Collect(ctx context.Context, in *MetricFamilies, opts ...grpc.CallOption) (*Void, error)
}

type metricsControllerClient struct {
	cc *grpc.ClientConn
}

func NewMetricsControllerClient(cc *grpc.ClientConn) MetricsControllerClient {
	return &metricsControllerClient{cc}
}

func (c *metricsControllerClient) Collect(ctx context.Context, in *MetricFamilies, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/grpc.MetricsController/Collect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetricsControllerServer is the server API for MetricsController service.
type MetricsControllerServer interface {
	// Report a collection of metrics from a service
	Collect(context.Context, *MetricFamilies) (*Void, error)
}

// UnimplementedMetricsControllerServer can be embedded to have forward compatible implementations.
type UnimplementedMetricsControllerServer struct {
}

func (*UnimplementedMetricsControllerServer) Collect(ctx context.Context, req *MetricFamilies) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Collect not implemented")
}

func RegisterMetricsControllerServer(s *grpc.Server, srv MetricsControllerServer) {
	s.RegisterService(&_MetricsController_serviceDesc, srv)
}

func _MetricsController_Collect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MetricFamilies)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsControllerServer).Collect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.MetricsController/Collect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsControllerServer).Collect(ctx, req.(*MetricFamilies))
	}
	return interceptor(ctx, in, info, handler)
}

var _MetricsController_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.MetricsController",
	HandlerType: (*MetricsControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Collect",
			Handler:    _MetricsController_Collect_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}