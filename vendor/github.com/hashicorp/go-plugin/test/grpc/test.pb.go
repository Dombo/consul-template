// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

package grpctest

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TestRequest struct {
	Input                int32    `protobuf:"varint,1,opt,name=Input,proto3" json:"Input,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestRequest) Reset()         { *m = TestRequest{} }
func (m *TestRequest) String() string { return proto.CompactTextString(m) }
func (*TestRequest) ProtoMessage()    {}
func (*TestRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{0}
}

func (m *TestRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestRequest.Unmarshal(m, b)
}
func (m *TestRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestRequest.Marshal(b, m, deterministic)
}
func (m *TestRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestRequest.Merge(m, src)
}
func (m *TestRequest) XXX_Size() int {
	return xxx_messageInfo_TestRequest.Size(m)
}
func (m *TestRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TestRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TestRequest proto.InternalMessageInfo

func (m *TestRequest) GetInput() int32 {
	if m != nil {
		return m.Input
	}
	return 0
}

type TestResponse struct {
	Output               int32    `protobuf:"varint,2,opt,name=Output,proto3" json:"Output,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestResponse) Reset()         { *m = TestResponse{} }
func (m *TestResponse) String() string { return proto.CompactTextString(m) }
func (*TestResponse) ProtoMessage()    {}
func (*TestResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{1}
}

func (m *TestResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestResponse.Unmarshal(m, b)
}
func (m *TestResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestResponse.Marshal(b, m, deterministic)
}
func (m *TestResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestResponse.Merge(m, src)
}
func (m *TestResponse) XXX_Size() int {
	return xxx_messageInfo_TestResponse.Size(m)
}
func (m *TestResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TestResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TestResponse proto.InternalMessageInfo

func (m *TestResponse) GetOutput() int32 {
	if m != nil {
		return m.Output
	}
	return 0
}

type PrintKVRequest struct {
	Key string `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
	// Types that are valid to be assigned to Value:
	//	*PrintKVRequest_ValueString
	//	*PrintKVRequest_ValueInt
	Value                isPrintKVRequest_Value `protobuf_oneof:"Value"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *PrintKVRequest) Reset()         { *m = PrintKVRequest{} }
func (m *PrintKVRequest) String() string { return proto.CompactTextString(m) }
func (*PrintKVRequest) ProtoMessage()    {}
func (*PrintKVRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{2}
}

func (m *PrintKVRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrintKVRequest.Unmarshal(m, b)
}
func (m *PrintKVRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrintKVRequest.Marshal(b, m, deterministic)
}
func (m *PrintKVRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrintKVRequest.Merge(m, src)
}
func (m *PrintKVRequest) XXX_Size() int {
	return xxx_messageInfo_PrintKVRequest.Size(m)
}
func (m *PrintKVRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PrintKVRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PrintKVRequest proto.InternalMessageInfo

func (m *PrintKVRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type isPrintKVRequest_Value interface {
	isPrintKVRequest_Value()
}

type PrintKVRequest_ValueString struct {
	ValueString string `protobuf:"bytes,2,opt,name=ValueString,proto3,oneof"`
}

type PrintKVRequest_ValueInt struct {
	ValueInt int32 `protobuf:"varint,3,opt,name=ValueInt,proto3,oneof"`
}

func (*PrintKVRequest_ValueString) isPrintKVRequest_Value() {}

func (*PrintKVRequest_ValueInt) isPrintKVRequest_Value() {}

func (m *PrintKVRequest) GetValue() isPrintKVRequest_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *PrintKVRequest) GetValueString() string {
	if x, ok := m.GetValue().(*PrintKVRequest_ValueString); ok {
		return x.ValueString
	}
	return ""
}

func (m *PrintKVRequest) GetValueInt() int32 {
	if x, ok := m.GetValue().(*PrintKVRequest_ValueInt); ok {
		return x.ValueInt
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*PrintKVRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _PrintKVRequest_OneofMarshaler, _PrintKVRequest_OneofUnmarshaler, _PrintKVRequest_OneofSizer, []interface{}{
		(*PrintKVRequest_ValueString)(nil),
		(*PrintKVRequest_ValueInt)(nil),
	}
}

func _PrintKVRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*PrintKVRequest)
	// Value
	switch x := m.Value.(type) {
	case *PrintKVRequest_ValueString:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.ValueString)
	case *PrintKVRequest_ValueInt:
		b.EncodeVarint(3<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.ValueInt))
	case nil:
	default:
		return fmt.Errorf("PrintKVRequest.Value has unexpected type %T", x)
	}
	return nil
}

func _PrintKVRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*PrintKVRequest)
	switch tag {
	case 2: // Value.ValueString
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Value = &PrintKVRequest_ValueString{x}
		return true, err
	case 3: // Value.ValueInt
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &PrintKVRequest_ValueInt{int32(x)}
		return true, err
	default:
		return false, nil
	}
}

func _PrintKVRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*PrintKVRequest)
	// Value
	switch x := m.Value.(type) {
	case *PrintKVRequest_ValueString:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.ValueString)))
		n += len(x.ValueString)
	case *PrintKVRequest_ValueInt:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.ValueInt))
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type PrintKVResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrintKVResponse) Reset()         { *m = PrintKVResponse{} }
func (m *PrintKVResponse) String() string { return proto.CompactTextString(m) }
func (*PrintKVResponse) ProtoMessage()    {}
func (*PrintKVResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{3}
}

func (m *PrintKVResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrintKVResponse.Unmarshal(m, b)
}
func (m *PrintKVResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrintKVResponse.Marshal(b, m, deterministic)
}
func (m *PrintKVResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrintKVResponse.Merge(m, src)
}
func (m *PrintKVResponse) XXX_Size() int {
	return xxx_messageInfo_PrintKVResponse.Size(m)
}
func (m *PrintKVResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PrintKVResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PrintKVResponse proto.InternalMessageInfo

type BidirectionalRequest struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BidirectionalRequest) Reset()         { *m = BidirectionalRequest{} }
func (m *BidirectionalRequest) String() string { return proto.CompactTextString(m) }
func (*BidirectionalRequest) ProtoMessage()    {}
func (*BidirectionalRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{4}
}

func (m *BidirectionalRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BidirectionalRequest.Unmarshal(m, b)
}
func (m *BidirectionalRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BidirectionalRequest.Marshal(b, m, deterministic)
}
func (m *BidirectionalRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BidirectionalRequest.Merge(m, src)
}
func (m *BidirectionalRequest) XXX_Size() int {
	return xxx_messageInfo_BidirectionalRequest.Size(m)
}
func (m *BidirectionalRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BidirectionalRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BidirectionalRequest proto.InternalMessageInfo

func (m *BidirectionalRequest) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type BidirectionalResponse struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BidirectionalResponse) Reset()         { *m = BidirectionalResponse{} }
func (m *BidirectionalResponse) String() string { return proto.CompactTextString(m) }
func (*BidirectionalResponse) ProtoMessage()    {}
func (*BidirectionalResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{5}
}

func (m *BidirectionalResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BidirectionalResponse.Unmarshal(m, b)
}
func (m *BidirectionalResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BidirectionalResponse.Marshal(b, m, deterministic)
}
func (m *BidirectionalResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BidirectionalResponse.Merge(m, src)
}
func (m *BidirectionalResponse) XXX_Size() int {
	return xxx_messageInfo_BidirectionalResponse.Size(m)
}
func (m *BidirectionalResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BidirectionalResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BidirectionalResponse proto.InternalMessageInfo

func (m *BidirectionalResponse) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type PingRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingRequest) Reset()         { *m = PingRequest{} }
func (m *PingRequest) String() string { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()    {}
func (*PingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{6}
}

func (m *PingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingRequest.Unmarshal(m, b)
}
func (m *PingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingRequest.Marshal(b, m, deterministic)
}
func (m *PingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingRequest.Merge(m, src)
}
func (m *PingRequest) XXX_Size() int {
	return xxx_messageInfo_PingRequest.Size(m)
}
func (m *PingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PingRequest proto.InternalMessageInfo

type PongResponse struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PongResponse) Reset()         { *m = PongResponse{} }
func (m *PongResponse) String() string { return proto.CompactTextString(m) }
func (*PongResponse) ProtoMessage()    {}
func (*PongResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{7}
}

func (m *PongResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PongResponse.Unmarshal(m, b)
}
func (m *PongResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PongResponse.Marshal(b, m, deterministic)
}
func (m *PongResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PongResponse.Merge(m, src)
}
func (m *PongResponse) XXX_Size() int {
	return xxx_messageInfo_PongResponse.Size(m)
}
func (m *PongResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PongResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PongResponse proto.InternalMessageInfo

func (m *PongResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*TestRequest)(nil), "grpctest.TestRequest")
	proto.RegisterType((*TestResponse)(nil), "grpctest.TestResponse")
	proto.RegisterType((*PrintKVRequest)(nil), "grpctest.PrintKVRequest")
	proto.RegisterType((*PrintKVResponse)(nil), "grpctest.PrintKVResponse")
	proto.RegisterType((*BidirectionalRequest)(nil), "grpctest.BidirectionalRequest")
	proto.RegisterType((*BidirectionalResponse)(nil), "grpctest.BidirectionalResponse")
	proto.RegisterType((*PingRequest)(nil), "grpctest.PingRequest")
	proto.RegisterType((*PongResponse)(nil), "grpctest.PongResponse")
}

func init() { proto.RegisterFile("test.proto", fileDescriptor_c161fcfdc0c3ff1e) }

var fileDescriptor_c161fcfdc0c3ff1e = []byte{
	// 355 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xcd, 0x4e, 0xc2, 0x40,
	0x14, 0x85, 0xdb, 0x02, 0x05, 0x2e, 0x3f, 0xe2, 0x04, 0x08, 0x12, 0xa3, 0x64, 0x4c, 0x90, 0x15,
	0x31, 0xb8, 0x30, 0x2e, 0x4c, 0x0c, 0xba, 0x80, 0xb0, 0x90, 0x0c, 0x86, 0x3d, 0x3f, 0x93, 0x66,
	0x12, 0x3a, 0xad, 0x9d, 0xe9, 0xc2, 0x17, 0xf1, 0x79, 0xcd, 0x0c, 0x6d, 0x19, 0x08, 0x2e, 0xdc,
	0xdd, 0x73, 0x7b, 0x72, 0xe6, 0x9e, 0x2f, 0x05, 0x90, 0x54, 0xc8, 0x61, 0x18, 0x05, 0x32, 0x40,
	0x25, 0x2f, 0x0a, 0x37, 0x4a, 0xe3, 0x3b, 0xa8, 0x7c, 0x52, 0x21, 0x09, 0xfd, 0x8a, 0xa9, 0x90,
	0xa8, 0x09, 0x85, 0x29, 0x0f, 0x63, 0xd9, 0xb1, 0x7b, 0xf6, 0xa0, 0x40, 0xf6, 0x02, 0xf7, 0xa1,
	0xba, 0x37, 0x89, 0x30, 0xe0, 0x82, 0xa2, 0x36, 0xb8, 0x1f, 0xb1, 0x54, 0x36, 0x47, 0xdb, 0x12,
	0x85, 0x7d, 0xa8, 0xcf, 0x23, 0xc6, 0xe5, 0x6c, 0x99, 0xe6, 0x35, 0x20, 0x37, 0xa3, 0xdf, 0x3a,
	0xad, 0x4c, 0xd4, 0x88, 0x30, 0x54, 0x96, 0xab, 0x5d, 0x4c, 0x17, 0x32, 0x62, 0xdc, 0xd3, 0x01,
	0xe5, 0x89, 0x45, 0xcc, 0x25, 0xba, 0x86, 0x92, 0x96, 0x53, 0x2e, 0x3b, 0x39, 0xf5, 0xc2, 0xc4,
	0x22, 0xd9, 0x66, 0x5c, 0x84, 0x82, 0x9e, 0xf1, 0x25, 0x5c, 0x64, 0xcf, 0xed, 0x2f, 0xc3, 0x7d,
	0x68, 0x8e, 0xd9, 0x96, 0x45, 0x74, 0x23, 0x59, 0xc0, 0x57, 0xbb, 0xf4, 0x8e, 0x3a, 0x38, 0x6c,
	0xab, 0xcf, 0xa8, 0x11, 0x87, 0x6d, 0xf1, 0x3d, 0xb4, 0x4e, 0x7c, 0x49, 0xb5, 0x53, 0x63, 0x0d,
	0x2a, 0x73, 0xc6, 0xbd, 0x24, 0x07, 0xf7, 0xa0, 0x3a, 0x0f, 0x94, 0x4c, 0xec, 0x0d, 0xc8, 0xf9,
	0xc2, 0x4b, 0xfb, 0xf9, 0xc2, 0x1b, 0xfd, 0x38, 0x90, 0x57, 0xb0, 0xd0, 0x33, 0xb8, 0xef, 0x41,
	0xbc, 0xde, 0x51, 0xd4, 0x1a, 0xa6, 0xb8, 0x87, 0x06, 0xeb, 0x6e, 0xfb, 0x74, 0x9d, 0x74, 0xb0,
	0xd0, 0x2b, 0x14, 0x93, 0x62, 0xa8, 0x73, 0x30, 0x1d, 0xa3, 0xed, 0x5e, 0x9d, 0xf9, 0x92, 0x25,
	0x10, 0xa8, 0x1d, 0xf5, 0x43, 0x37, 0x07, 0xf7, 0x39, 0x40, 0xdd, 0xdb, 0x3f, 0xbf, 0x67, 0x99,
	0x2f, 0xe0, 0x2e, 0x64, 0x44, 0x57, 0xfe, 0xbf, 0x0b, 0x0d, 0xec, 0x07, 0x7b, 0xf4, 0x06, 0x25,
	0x45, 0x52, 0xe1, 0x43, 0x4f, 0x90, 0x57, 0xb3, 0x19, 0x64, 0x50, 0x36, 0x83, 0x4c, 0xda, 0xd8,
	0x5a, 0xbb, 0xfa, 0xff, 0x7d, 0xfc, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xf0, 0x59, 0x20, 0xc7, 0xcd,
	0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TestClient is the client API for Test service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TestClient interface {
	Double(ctx context.Context, in *TestRequest, opts ...grpc.CallOption) (*TestResponse, error)
	PrintKV(ctx context.Context, in *PrintKVRequest, opts ...grpc.CallOption) (*PrintKVResponse, error)
	Bidirectional(ctx context.Context, in *BidirectionalRequest, opts ...grpc.CallOption) (*BidirectionalResponse, error)
	Stream(ctx context.Context, opts ...grpc.CallOption) (Test_StreamClient, error)
}

type testClient struct {
	cc *grpc.ClientConn
}

func NewTestClient(cc *grpc.ClientConn) TestClient {
	return &testClient{cc}
}

func (c *testClient) Double(ctx context.Context, in *TestRequest, opts ...grpc.CallOption) (*TestResponse, error) {
	out := new(TestResponse)
	err := c.cc.Invoke(ctx, "/grpctest.Test/Double", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testClient) PrintKV(ctx context.Context, in *PrintKVRequest, opts ...grpc.CallOption) (*PrintKVResponse, error) {
	out := new(PrintKVResponse)
	err := c.cc.Invoke(ctx, "/grpctest.Test/PrintKV", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testClient) Bidirectional(ctx context.Context, in *BidirectionalRequest, opts ...grpc.CallOption) (*BidirectionalResponse, error) {
	out := new(BidirectionalResponse)
	err := c.cc.Invoke(ctx, "/grpctest.Test/Bidirectional", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testClient) Stream(ctx context.Context, opts ...grpc.CallOption) (Test_StreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Test_serviceDesc.Streams[0], "/grpctest.Test/Stream", opts...)
	if err != nil {
		return nil, err
	}
	x := &testStreamClient{stream}
	return x, nil
}

type Test_StreamClient interface {
	Send(*TestRequest) error
	Recv() (*TestResponse, error)
	grpc.ClientStream
}

type testStreamClient struct {
	grpc.ClientStream
}

func (x *testStreamClient) Send(m *TestRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testStreamClient) Recv() (*TestResponse, error) {
	m := new(TestResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TestServer is the server API for Test service.
type TestServer interface {
	Double(context.Context, *TestRequest) (*TestResponse, error)
	PrintKV(context.Context, *PrintKVRequest) (*PrintKVResponse, error)
	Bidirectional(context.Context, *BidirectionalRequest) (*BidirectionalResponse, error)
	Stream(Test_StreamServer) error
}

func RegisterTestServer(s *grpc.Server, srv TestServer) {
	s.RegisterService(&_Test_serviceDesc, srv)
}

func _Test_Double_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServer).Double(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpctest.Test/Double",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServer).Double(ctx, req.(*TestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Test_PrintKV_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrintKVRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServer).PrintKV(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpctest.Test/PrintKV",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServer).PrintKV(ctx, req.(*PrintKVRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Test_Bidirectional_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BidirectionalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServer).Bidirectional(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpctest.Test/Bidirectional",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServer).Bidirectional(ctx, req.(*BidirectionalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Test_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestServer).Stream(&testStreamServer{stream})
}

type Test_StreamServer interface {
	Send(*TestResponse) error
	Recv() (*TestRequest, error)
	grpc.ServerStream
}

type testStreamServer struct {
	grpc.ServerStream
}

func (x *testStreamServer) Send(m *TestResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testStreamServer) Recv() (*TestRequest, error) {
	m := new(TestRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Test_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpctest.Test",
	HandlerType: (*TestServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Double",
			Handler:    _Test_Double_Handler,
		},
		{
			MethodName: "PrintKV",
			Handler:    _Test_PrintKV_Handler,
		},
		{
			MethodName: "Bidirectional",
			Handler:    _Test_Bidirectional_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _Test_Stream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "test.proto",
}

// PingPongClient is the client API for PingPong service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PingPongClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PongResponse, error)
}

type pingPongClient struct {
	cc *grpc.ClientConn
}

func NewPingPongClient(cc *grpc.ClientConn) PingPongClient {
	return &pingPongClient{cc}
}

func (c *pingPongClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PongResponse, error) {
	out := new(PongResponse)
	err := c.cc.Invoke(ctx, "/grpctest.PingPong/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PingPongServer is the server API for PingPong service.
type PingPongServer interface {
	Ping(context.Context, *PingRequest) (*PongResponse, error)
}

func RegisterPingPongServer(s *grpc.Server, srv PingPongServer) {
	s.RegisterService(&_PingPong_serviceDesc, srv)
}

func _PingPong_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PingPongServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpctest.PingPong/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PingPongServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PingPong_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpctest.PingPong",
	HandlerType: (*PingPongServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _PingPong_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test.proto",
}
