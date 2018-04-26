// Code generated by protoc-gen-go. DO NOT EDIT.
// source: CalCall/LZMJMessage.proto

/*
Package CalCall is a generated protocol buffer package.

It is generated from these files:
	CalCall/LZMJMessage.proto

It has these top-level messages:
	MajongCal
	CalReply
*/
package CalCall

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

// 麻将的计算调用消息
type MajongCal struct {
	Rule       string   `protobuf:"bytes,1,opt,name=Rule" json:"Rule,omitempty"`
	PlayerCard []string `protobuf:"bytes,2,rep,name=PlayerCard" json:"PlayerCard,omitempty"`
	NewCard    string   `protobuf:"bytes,3,opt,name=NewCard" json:"NewCard,omitempty"`
}

func (m *MajongCal) Reset()                    { *m = MajongCal{} }
func (m *MajongCal) String() string            { return proto.CompactTextString(m) }
func (*MajongCal) ProtoMessage()               {}
func (*MajongCal) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *MajongCal) GetRule() string {
	if m != nil {
		return m.Rule
	}
	return ""
}

func (m *MajongCal) GetPlayerCard() []string {
	if m != nil {
		return m.PlayerCard
	}
	return nil
}

func (m *MajongCal) GetNewCard() string {
	if m != nil {
		return m.NewCard
	}
	return ""
}

// 麻将计算结果反馈
type CalReply struct {
	CalResult []string `protobuf:"bytes,1,rep,name=CalResult" json:"CalResult,omitempty"`
}

func (m *CalReply) Reset()                    { *m = CalReply{} }
func (m *CalReply) String() string            { return proto.CompactTextString(m) }
func (*CalReply) ProtoMessage()               {}
func (*CalReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CalReply) GetCalResult() []string {
	if m != nil {
		return m.CalResult
	}
	return nil
}

func init() {
	proto.RegisterType((*MajongCal)(nil), "CalCall.MajongCal")
	proto.RegisterType((*CalReply)(nil), "CalCall.CalReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Greeter service

type GreeterClient interface {
	CallCal(ctx context.Context, in *MajongCal, opts ...grpc.CallOption) (*CalReply, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) CallCal(ctx context.Context, in *MajongCal, opts ...grpc.CallOption) (*CalReply, error) {
	out := new(CalReply)
	err := grpc.Invoke(ctx, "/CalCall.Greeter/CallCal", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Greeter service

type GreeterServer interface {
	CallCal(context.Context, *MajongCal) (*CalReply, error)
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_CallCal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MajongCal)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).CallCal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CalCall.Greeter/CallCal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).CallCal(ctx, req.(*MajongCal))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "CalCall.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CallCal",
			Handler:    _Greeter_CallCal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "CalCall/LZMJMessage.proto",
}

func init() { proto.RegisterFile("CalCall/LZMJMessage.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0xcd, 0x56, 0xac, 0x19, 0xbc, 0x98, 0x53, 0x14, 0x91, 0xa5, 0xa7, 0x9e, 0xa2, 0xac,
	0x67, 0x2f, 0x9b, 0x83, 0x22, 0x5b, 0x29, 0xb9, 0x88, 0xde, 0x46, 0x1d, 0x8a, 0x32, 0x9a, 0x90,
	0x74, 0xd1, 0xfd, 0x3b, 0xfe, 0x52, 0x69, 0x68, 0xab, 0xb7, 0x97, 0xf7, 0x32, 0x8f, 0x6f, 0x06,
	0x4e, 0x2c, 0xb2, 0x45, 0xe6, 0x8b, 0xcd, 0x53, 0x73, 0xd7, 0x50, 0x4a, 0xd8, 0x91, 0x09, 0xd1,
	0xf7, 0x5e, 0x95, 0x63, 0x54, 0x3d, 0x82, 0x6c, 0xf0, 0xdd, 0x7f, 0x76, 0x16, 0x59, 0x29, 0xd8,
	0x77, 0x5b, 0x26, 0x2d, 0x96, 0xa2, 0x96, 0x2e, 0x6b, 0x75, 0x0e, 0xd0, 0x32, 0xee, 0x28, 0x5a,
	0x8c, 0xaf, 0x7a, 0xb1, 0x2c, 0x6a, 0xe9, 0xfe, 0x39, 0x4a, 0x43, 0x79, 0x4f, 0x5f, 0x39, 0x2c,
	0xf2, 0xd8, 0xf4, 0xac, 0x6a, 0x38, 0xb4, 0xc8, 0x8e, 0x02, 0xef, 0xd4, 0x19, 0xc8, 0xac, 0xd3,
	0x96, 0x7b, 0x2d, 0x72, 0xc9, 0x9f, 0xb1, 0xba, 0x86, 0xf2, 0x26, 0x12, 0xf5, 0x14, 0xd5, 0x0a,
	0x06, 0x34, 0xce, 0x34, 0x66, 0x84, 0x34, 0x33, 0xe1, 0xe9, 0xf1, 0xec, 0x4d, 0xd5, 0xd5, 0xde,
	0xfa, 0x12, 0xf4, 0x9b, 0x37, 0x5d, 0x0c, 0x2f, 0x86, 0xbe, 0xf1, 0x23, 0x30, 0xa5, 0xe9, 0xdb,
	0xfa, 0x68, 0x14, 0xed, 0xb0, 0x76, 0x2b, 0x7e, 0x16, 0xc5, 0xed, 0xe6, 0xe1, 0xf9, 0x20, 0x5f,
	0xe1, 0xea, 0x37, 0x00, 0x00, 0xff, 0xff, 0x03, 0x61, 0xdb, 0xc9, 0x22, 0x01, 0x00, 0x00,
}
