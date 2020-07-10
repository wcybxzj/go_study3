// Code generated by protoc-gen-go. DO NOT EDIT.
// source: person.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// 不能在同一个 包下,创建同名的 消息体.
type Teacher struct {
	Age                  int32    `protobuf:"varint,1,opt,name=age,proto3" json:"age,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Teacher) Reset()         { *m = Teacher{} }
func (m *Teacher) String() string { return proto.CompactTextString(m) }
func (*Teacher) ProtoMessage()    {}
func (*Teacher) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c9e10cf24b1156d, []int{0}
}

func (m *Teacher) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Teacher.Unmarshal(m, b)
}
func (m *Teacher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Teacher.Marshal(b, m, deterministic)
}
func (m *Teacher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Teacher.Merge(m, src)
}
func (m *Teacher) XXX_Size() int {
	return xxx_messageInfo_Teacher.Size(m)
}
func (m *Teacher) XXX_DiscardUnknown() {
	xxx_messageInfo_Teacher.DiscardUnknown(m)
}

var xxx_messageInfo_Teacher proto.InternalMessageInfo

func (m *Teacher) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *Teacher) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Teacher)(nil), "pb.Teacher")
}

func init() { proto.RegisterFile("person.proto", fileDescriptor_4c9e10cf24b1156d) }

var fileDescriptor_4c9e10cf24b1156d = []byte{
	// 125 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x48, 0x2d, 0x2a,
	0xce, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0xd2, 0xe7, 0x62,
	0x0f, 0x49, 0x4d, 0x4c, 0xce, 0x48, 0x2d, 0x12, 0x12, 0xe0, 0x62, 0x4e, 0x4c, 0x4f, 0x95, 0x60,
	0x54, 0x60, 0xd4, 0x60, 0x0d, 0x02, 0x31, 0x85, 0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73, 0x53, 0x25,
	0x98, 0x14, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x23, 0x7d, 0x2e, 0xf6, 0xe0, 0xc4, 0x4a, 0xbf,
	0xc4, 0xdc, 0x54, 0x21, 0x15, 0x2e, 0x8e, 0xe0, 0xc4, 0x4a, 0x8f, 0xd4, 0x9c, 0x9c, 0x7c, 0x21,
	0x6e, 0xbd, 0x82, 0x24, 0x3d, 0xa8, 0x49, 0x52, 0xc8, 0x9c, 0x24, 0x36, 0xb0, 0x65, 0xc6, 0x80,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x48, 0x33, 0xe6, 0x1d, 0x7c, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SayNameClient is the 2.client API for SayName service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SayNameClient interface {
	SayHello(ctx context.Context, in *Teacher, opts ...grpc.CallOption) (*Teacher, error)
}

type sayNameClient struct {
	cc *grpc.ClientConn
}

func NewSayNameClient(cc *grpc.ClientConn) SayNameClient {
	return &sayNameClient{cc}
}

func (c *sayNameClient) SayHello(ctx context.Context, in *Teacher, opts ...grpc.CallOption) (*Teacher, error) {
	out := new(Teacher)
	err := c.cc.Invoke(ctx, "/pb.SayName/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SayNameServer is the 1.server API for SayName service.
type SayNameServer interface {
	SayHello(context.Context, *Teacher) (*Teacher, error)
}

// UnimplementedSayNameServer can be embedded to have forward compatible implementations.
type UnimplementedSayNameServer struct {
}

func (*UnimplementedSayNameServer) SayHello(ctx context.Context, req *Teacher) (*Teacher, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}

func RegisterSayNameServer(s *grpc.Server, srv SayNameServer) {
	s.RegisterService(&_SayName_serviceDesc, srv)
}

func _SayName_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Teacher)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SayNameServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SayName/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SayNameServer).SayHello(ctx, req.(*Teacher))
	}
	return interceptor(ctx, in, info, handler)
}

var _SayName_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.SayName",
	HandlerType: (*SayNameServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _SayName_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "person.proto",
}