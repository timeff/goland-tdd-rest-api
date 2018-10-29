// Code generated by protoc-gen-go. DO NOT EDIT.
// source: transport/gRPC/todo.proto

package gRPC

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type TodosRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TodosRequest) Reset()         { *m = TodosRequest{} }
func (m *TodosRequest) String() string { return proto.CompactTextString(m) }
func (*TodosRequest) ProtoMessage()    {}
func (*TodosRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9dd00506811c099f, []int{0}
}

func (m *TodosRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TodosRequest.Unmarshal(m, b)
}
func (m *TodosRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TodosRequest.Marshal(b, m, deterministic)
}
func (m *TodosRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TodosRequest.Merge(m, src)
}
func (m *TodosRequest) XXX_Size() int {
	return xxx_messageInfo_TodosRequest.Size(m)
}
func (m *TodosRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TodosRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TodosRequest proto.InternalMessageInfo

type TodosReply struct {
	Todo                 []*TodoReply `protobuf:"bytes,1,rep,name=Todo,proto3" json:"Todo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *TodosReply) Reset()         { *m = TodosReply{} }
func (m *TodosReply) String() string { return proto.CompactTextString(m) }
func (*TodosReply) ProtoMessage()    {}
func (*TodosReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_9dd00506811c099f, []int{1}
}

func (m *TodosReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TodosReply.Unmarshal(m, b)
}
func (m *TodosReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TodosReply.Marshal(b, m, deterministic)
}
func (m *TodosReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TodosReply.Merge(m, src)
}
func (m *TodosReply) XXX_Size() int {
	return xxx_messageInfo_TodosReply.Size(m)
}
func (m *TodosReply) XXX_DiscardUnknown() {
	xxx_messageInfo_TodosReply.DiscardUnknown(m)
}

var xxx_messageInfo_TodosReply proto.InternalMessageInfo

func (m *TodosReply) GetTodo() []*TodoReply {
	if m != nil {
		return m.Todo
	}
	return nil
}

type TodoReply struct {
	ID                   int64                `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Content              string               `protobuf:"bytes,2,opt,name=Content,proto3" json:"Content,omitempty"`
	Done                 bool                 `protobuf:"varint,3,opt,name=Done,proto3" json:"Done,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	DeletedAt            *timestamp.Timestamp `protobuf:"bytes,6,opt,name=DeletedAt,proto3" json:"DeletedAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TodoReply) Reset()         { *m = TodoReply{} }
func (m *TodoReply) String() string { return proto.CompactTextString(m) }
func (*TodoReply) ProtoMessage()    {}
func (*TodoReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_9dd00506811c099f, []int{2}
}

func (m *TodoReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TodoReply.Unmarshal(m, b)
}
func (m *TodoReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TodoReply.Marshal(b, m, deterministic)
}
func (m *TodoReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TodoReply.Merge(m, src)
}
func (m *TodoReply) XXX_Size() int {
	return xxx_messageInfo_TodoReply.Size(m)
}
func (m *TodoReply) XXX_DiscardUnknown() {
	xxx_messageInfo_TodoReply.DiscardUnknown(m)
}

var xxx_messageInfo_TodoReply proto.InternalMessageInfo

func (m *TodoReply) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *TodoReply) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *TodoReply) GetDone() bool {
	if m != nil {
		return m.Done
	}
	return false
}

func (m *TodoReply) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *TodoReply) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *TodoReply) GetDeletedAt() *timestamp.Timestamp {
	if m != nil {
		return m.DeletedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*TodosRequest)(nil), "gRPC.TodosRequest")
	proto.RegisterType((*TodosReply)(nil), "gRPC.TodosReply")
	proto.RegisterType((*TodoReply)(nil), "gRPC.TodoReply")
}

func init() { proto.RegisterFile("transport/gRPC/todo.proto", fileDescriptor_9dd00506811c099f) }

var fileDescriptor_9dd00506811c099f = []byte{
	// 271 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x41, 0x4f, 0xbb, 0x40,
	0x10, 0xc5, 0xff, 0x5b, 0xf8, 0x57, 0x3b, 0x35, 0xd5, 0xec, 0x69, 0xe5, 0x22, 0xc1, 0x0b, 0x17,
	0x97, 0x88, 0x31, 0xf1, 0x6a, 0x20, 0x31, 0xbd, 0x99, 0x4d, 0xfd, 0x00, 0x34, 0x8c, 0xc4, 0x84,
	0x32, 0x2b, 0x4c, 0x0f, 0xfd, 0xe4, 0x5e, 0xcd, 0x42, 0xa1, 0xf1, 0xd4, 0xdb, 0xbe, 0xb7, 0xef,
	0x37, 0xc9, 0x7b, 0x70, 0xcb, 0x6d, 0xd1, 0x74, 0x96, 0x5a, 0x4e, 0x2a, 0xf3, 0x9e, 0x25, 0x4c,
	0x25, 0x69, 0xdb, 0x12, 0x93, 0xf4, 0x9d, 0x11, 0xdc, 0x55, 0x44, 0x55, 0x8d, 0x49, 0xef, 0x6d,
	0xf7, 0x9f, 0x09, 0x7f, 0xed, 0xb0, 0xe3, 0x62, 0x67, 0x87, 0x58, 0xb4, 0x82, 0xab, 0x0d, 0x95,
	0xd4, 0x19, 0xfc, 0xde, 0x63, 0xc7, 0xd1, 0x23, 0xc0, 0x51, 0xdb, 0xfa, 0x20, 0xef, 0xc1, 0x77,
	0x4a, 0x89, 0xd0, 0x8b, 0x97, 0xe9, 0xb5, 0x76, 0x37, 0xb5, 0x73, 0xfa, 0x6f, 0xd3, 0x7f, 0x46,
	0x3f, 0x02, 0x16, 0x93, 0x27, 0x57, 0x30, 0x5b, 0xe7, 0x4a, 0x84, 0x22, 0xf6, 0xcc, 0x6c, 0x9d,
	0x4b, 0x05, 0x17, 0x19, 0x35, 0x8c, 0x0d, 0xab, 0x59, 0x28, 0xe2, 0x85, 0x19, 0xa5, 0x94, 0xe0,
	0xe7, 0xd4, 0xa0, 0xf2, 0x42, 0x11, 0x5f, 0x9a, 0xfe, 0x2d, 0x5f, 0x60, 0x91, 0xb5, 0x58, 0x30,
	0x96, 0xaf, 0xac, 0xfc, 0x50, 0xc4, 0xcb, 0x34, 0xd0, 0x43, 0x07, 0x3d, 0x76, 0xd0, 0x9b, 0xb1,
	0x83, 0x39, 0x85, 0x1d, 0xf9, 0x61, 0xcb, 0x23, 0xf9, 0xff, 0x3c, 0x39, 0x85, 0x1d, 0x99, 0x63,
	0x8d, 0x03, 0x39, 0x3f, 0x4f, 0x4e, 0xe1, 0xf4, 0x79, 0x98, 0x47, 0x3e, 0x80, 0xf7, 0x86, 0x2c,
	0xe5, 0x69, 0x9f, 0x71, 0xcf, 0xe0, 0xe6, 0x8f, 0x67, 0xeb, 0x43, 0xf4, 0x6f, 0x3b, 0xef, 0xaf,
	0x3e, 0xfd, 0x06, 0x00, 0x00, 0xff, 0xff, 0x01, 0x8b, 0x5b, 0x14, 0xbe, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TodoClient is the client API for Todo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TodoClient interface {
	Get(ctx context.Context, in *TodosRequest, opts ...grpc.CallOption) (*TodosReply, error)
}

type todoClient struct {
	cc *grpc.ClientConn
}

func NewTodoClient(cc *grpc.ClientConn) TodoClient {
	return &todoClient{cc}
}

func (c *todoClient) Get(ctx context.Context, in *TodosRequest, opts ...grpc.CallOption) (*TodosReply, error) {
	out := new(TodosReply)
	err := c.cc.Invoke(ctx, "/gRPC.Todo/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoServer is the server API for Todo service.
type TodoServer interface {
	Get(context.Context, *TodosRequest) (*TodosReply, error)
}

func RegisterTodoServer(s *grpc.Server, srv TodoServer) {
	s.RegisterService(&_Todo_serviceDesc, srv)
}

func _Todo_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TodosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gRPC.Todo/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServer).Get(ctx, req.(*TodosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Todo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gRPC.Todo",
	HandlerType: (*TodoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Todo_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transport/gRPC/todo.proto",
}
