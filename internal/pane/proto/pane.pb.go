// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/pane/proto/pane.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

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

type PaneControl struct {
	// Types that are valid to be assigned to Control:
	//	*PaneControl_Connect
	//	*PaneControl_WindowChange
	Control              isPaneControl_Control `protobuf_oneof:"control"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *PaneControl) Reset()         { *m = PaneControl{} }
func (m *PaneControl) String() string { return proto.CompactTextString(m) }
func (*PaneControl) ProtoMessage()    {}
func (*PaneControl) Descriptor() ([]byte, []int) {
	return fileDescriptor_65adb0e31fae1ba9, []int{0}
}

func (m *PaneControl) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaneControl.Unmarshal(m, b)
}
func (m *PaneControl) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaneControl.Marshal(b, m, deterministic)
}
func (m *PaneControl) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaneControl.Merge(m, src)
}
func (m *PaneControl) XXX_Size() int {
	return xxx_messageInfo_PaneControl.Size(m)
}
func (m *PaneControl) XXX_DiscardUnknown() {
	xxx_messageInfo_PaneControl.DiscardUnknown(m)
}

var xxx_messageInfo_PaneControl proto.InternalMessageInfo

type isPaneControl_Control interface {
	isPaneControl_Control()
}

type PaneControl_Connect struct {
	Connect *ConnectRequest `protobuf:"bytes,1,opt,name=connect,proto3,oneof"`
}

type PaneControl_WindowChange struct {
	WindowChange *WindowChange `protobuf:"bytes,2,opt,name=window_change,json=windowChange,proto3,oneof"`
}

func (*PaneControl_Connect) isPaneControl_Control() {}

func (*PaneControl_WindowChange) isPaneControl_Control() {}

func (m *PaneControl) GetControl() isPaneControl_Control {
	if m != nil {
		return m.Control
	}
	return nil
}

func (m *PaneControl) GetConnect() *ConnectRequest {
	if x, ok := m.GetControl().(*PaneControl_Connect); ok {
		return x.Connect
	}
	return nil
}

func (m *PaneControl) GetWindowChange() *WindowChange {
	if x, ok := m.GetControl().(*PaneControl_WindowChange); ok {
		return x.WindowChange
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*PaneControl) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _PaneControl_OneofMarshaler, _PaneControl_OneofUnmarshaler, _PaneControl_OneofSizer, []interface{}{
		(*PaneControl_Connect)(nil),
		(*PaneControl_WindowChange)(nil),
	}
}

func _PaneControl_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*PaneControl)
	// control
	switch x := m.Control.(type) {
	case *PaneControl_Connect:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Connect); err != nil {
			return err
		}
	case *PaneControl_WindowChange:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.WindowChange); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("PaneControl.Control has unexpected type %T", x)
	}
	return nil
}

func _PaneControl_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*PaneControl)
	switch tag {
	case 1: // control.connect
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ConnectRequest)
		err := b.DecodeMessage(msg)
		m.Control = &PaneControl_Connect{msg}
		return true, err
	case 2: // control.window_change
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(WindowChange)
		err := b.DecodeMessage(msg)
		m.Control = &PaneControl_WindowChange{msg}
		return true, err
	default:
		return false, nil
	}
}

func _PaneControl_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*PaneControl)
	// control
	switch x := m.Control.(type) {
	case *PaneControl_Connect:
		s := proto.Size(x.Connect)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *PaneControl_WindowChange:
		s := proto.Size(x.WindowChange)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type ConnectRequest struct {
	FdSocketPath         string   `protobuf:"bytes,1,opt,name=fd_socket_path,json=fdSocketPath,proto3" json:"fd_socket_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConnectRequest) Reset()         { *m = ConnectRequest{} }
func (m *ConnectRequest) String() string { return proto.CompactTextString(m) }
func (*ConnectRequest) ProtoMessage()    {}
func (*ConnectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_65adb0e31fae1ba9, []int{1}
}

func (m *ConnectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectRequest.Unmarshal(m, b)
}
func (m *ConnectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectRequest.Marshal(b, m, deterministic)
}
func (m *ConnectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectRequest.Merge(m, src)
}
func (m *ConnectRequest) XXX_Size() int {
	return xxx_messageInfo_ConnectRequest.Size(m)
}
func (m *ConnectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectRequest proto.InternalMessageInfo

func (m *ConnectRequest) GetFdSocketPath() string {
	if m != nil {
		return m.FdSocketPath
	}
	return ""
}

type WindowChange struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WindowChange) Reset()         { *m = WindowChange{} }
func (m *WindowChange) String() string { return proto.CompactTextString(m) }
func (*WindowChange) ProtoMessage()    {}
func (*WindowChange) Descriptor() ([]byte, []int) {
	return fileDescriptor_65adb0e31fae1ba9, []int{2}
}

func (m *WindowChange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WindowChange.Unmarshal(m, b)
}
func (m *WindowChange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WindowChange.Marshal(b, m, deterministic)
}
func (m *WindowChange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WindowChange.Merge(m, src)
}
func (m *WindowChange) XXX_Size() int {
	return xxx_messageInfo_WindowChange.Size(m)
}
func (m *WindowChange) XXX_DiscardUnknown() {
	xxx_messageInfo_WindowChange.DiscardUnknown(m)
}

var xxx_messageInfo_WindowChange proto.InternalMessageInfo

type DoneReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DoneReply) Reset()         { *m = DoneReply{} }
func (m *DoneReply) String() string { return proto.CompactTextString(m) }
func (*DoneReply) ProtoMessage()    {}
func (*DoneReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_65adb0e31fae1ba9, []int{3}
}

func (m *DoneReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoneReply.Unmarshal(m, b)
}
func (m *DoneReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoneReply.Marshal(b, m, deterministic)
}
func (m *DoneReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoneReply.Merge(m, src)
}
func (m *DoneReply) XXX_Size() int {
	return xxx_messageInfo_DoneReply.Size(m)
}
func (m *DoneReply) XXX_DiscardUnknown() {
	xxx_messageInfo_DoneReply.DiscardUnknown(m)
}

var xxx_messageInfo_DoneReply proto.InternalMessageInfo

func init() {
	proto.RegisterType((*PaneControl)(nil), "pane.PaneControl")
	proto.RegisterType((*ConnectRequest)(nil), "pane.ConnectRequest")
	proto.RegisterType((*WindowChange)(nil), "pane.WindowChange")
	proto.RegisterType((*DoneReply)(nil), "pane.DoneReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PaneClient is the client API for Pane service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PaneClient interface {
	ConnectPane(ctx context.Context, opts ...grpc.CallOption) (Pane_ConnectPaneClient, error)
}

type paneClient struct {
	cc *grpc.ClientConn
}

func NewPaneClient(cc *grpc.ClientConn) PaneClient {
	return &paneClient{cc}
}

func (c *paneClient) ConnectPane(ctx context.Context, opts ...grpc.CallOption) (Pane_ConnectPaneClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Pane_serviceDesc.Streams[0], "/pane.Pane/ConnectPane", opts...)
	if err != nil {
		return nil, err
	}
	x := &paneConnectPaneClient{stream}
	return x, nil
}

type Pane_ConnectPaneClient interface {
	Send(*PaneControl) error
	CloseAndRecv() (*DoneReply, error)
	grpc.ClientStream
}

type paneConnectPaneClient struct {
	grpc.ClientStream
}

func (x *paneConnectPaneClient) Send(m *PaneControl) error {
	return x.ClientStream.SendMsg(m)
}

func (x *paneConnectPaneClient) CloseAndRecv() (*DoneReply, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(DoneReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PaneServer is the server API for Pane service.
type PaneServer interface {
	ConnectPane(Pane_ConnectPaneServer) error
}

func RegisterPaneServer(s *grpc.Server, srv PaneServer) {
	s.RegisterService(&_Pane_serviceDesc, srv)
}

func _Pane_ConnectPane_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PaneServer).ConnectPane(&paneConnectPaneServer{stream})
}

type Pane_ConnectPaneServer interface {
	SendAndClose(*DoneReply) error
	Recv() (*PaneControl, error)
	grpc.ServerStream
}

type paneConnectPaneServer struct {
	grpc.ServerStream
}

func (x *paneConnectPaneServer) SendAndClose(m *DoneReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *paneConnectPaneServer) Recv() (*PaneControl, error) {
	m := new(PaneControl)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Pane_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pane.Pane",
	HandlerType: (*PaneServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ConnectPane",
			Handler:       _Pane_ConnectPane_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "internal/pane/proto/pane.proto",
}

func init() { proto.RegisterFile("internal/pane/proto/pane.proto", fileDescriptor_65adb0e31fae1ba9) }

var fileDescriptor_65adb0e31fae1ba9 = []byte{
	// 272 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x50, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x6d, 0xa4, 0x28, 0xd9, 0xc4, 0x88, 0x8b, 0x87, 0xe2, 0x41, 0x24, 0x78, 0xe8, 0x29, 0x91,
	0x4a, 0x15, 0x11, 0x2f, 0x8d, 0x87, 0x1e, 0x4b, 0x3c, 0x08, 0x5e, 0xc2, 0x36, 0x99, 0x26, 0x8b,
	0xdb, 0xd9, 0x98, 0x4e, 0x09, 0x7e, 0x80, 0xff, 0x2d, 0x99, 0xa8, 0x44, 0xf0, 0xb2, 0xfb, 0xe6,
	0xcd, 0xbe, 0x7d, 0x6f, 0x46, 0x5c, 0x68, 0x24, 0x68, 0x50, 0x99, 0xb8, 0x56, 0x08, 0x71, 0xdd,
	0x58, 0xb2, 0x0c, 0x23, 0x86, 0x72, 0xdc, 0xe1, 0xf0, 0xd3, 0x11, 0xde, 0x4a, 0x21, 0x24, 0x16,
	0xa9, 0xb1, 0x46, 0x5e, 0x8b, 0xa3, 0xdc, 0x22, 0x42, 0x4e, 0x13, 0xe7, 0xd2, 0x99, 0x7a, 0xb3,
	0xb3, 0x88, 0x35, 0x49, 0x4f, 0xa6, 0xf0, 0xbe, 0x87, 0x1d, 0x2d, 0x47, 0xe9, 0xcf, 0x33, 0x79,
	0x2f, 0x8e, 0x5b, 0x8d, 0x85, 0x6d, 0xb3, 0xbc, 0x52, 0x58, 0xc2, 0xe4, 0x80, 0x75, 0xb2, 0xd7,
	0xbd, 0x70, 0x2b, 0xe1, 0xce, 0x72, 0x94, 0xfa, 0xed, 0xa0, 0x5e, 0xb8, 0x6c, 0xd6, 0xf9, 0x86,
	0xb7, 0x22, 0xf8, 0x6b, 0x21, 0xaf, 0x44, 0xb0, 0x29, 0xb2, 0x9d, 0xcd, 0xdf, 0x80, 0xb2, 0x5a,
	0x51, 0xc5, 0x81, 0xdc, 0xd4, 0xdf, 0x14, 0xcf, 0x4c, 0xae, 0x14, 0x55, 0x61, 0x20, 0xfc, 0xa1,
	0x45, 0xe8, 0x09, 0xf7, 0xc9, 0x22, 0xa4, 0x50, 0x9b, 0x8f, 0xd9, 0xa3, 0x18, 0x77, 0xb3, 0xc9,
	0xb9, 0xf0, 0xbe, 0x3f, 0xe7, 0xf2, 0xb4, 0x8f, 0x36, 0x18, 0xfb, 0xfc, 0xa4, 0xa7, 0x7e, 0xa5,
	0xe1, 0x68, 0xea, 0x2c, 0xee, 0x5e, 0xe7, 0xa5, 0xa6, 0x6a, 0xbf, 0x8e, 0x72, 0xbb, 0x8d, 0xbb,
	0xe4, 0x5b, 0x6d, 0x0c, 0x60, 0x19, 0x93, 0x36, 0x14, 0xff, 0xb3, 0xdf, 0x07, 0x3e, 0xd7, 0x87,
	0x7c, 0xdd, 0x7c, 0x05, 0x00, 0x00, 0xff, 0xff, 0x00, 0xbc, 0x5f, 0x29, 0x83, 0x01, 0x00, 0x00,
}
