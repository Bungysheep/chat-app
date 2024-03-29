// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chat-protobuf/chat.proto

package chat

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Message struct {
	SendId               string               `protobuf:"bytes,1,opt,name=send_id,json=sendId,proto3" json:"send_id,omitempty"`
	Text                 string               `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	SentAt               *timestamp.Timestamp `protobuf:"bytes,3,opt,name=sent_at,json=sentAt,proto3" json:"sent_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_f5b84c3e519282d7, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetSendId() string {
	if m != nil {
		return m.SendId
	}
	return ""
}

func (m *Message) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Message) GetSentAt() *timestamp.Timestamp {
	if m != nil {
		return m.SentAt
	}
	return nil
}

type SendRequest struct {
	SubId                string   `protobuf:"bytes,1,opt,name=sub_id,json=subId,proto3" json:"sub_id,omitempty"`
	Message              *Message `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendRequest) Reset()         { *m = SendRequest{} }
func (m *SendRequest) String() string { return proto.CompactTextString(m) }
func (*SendRequest) ProtoMessage()    {}
func (*SendRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f5b84c3e519282d7, []int{1}
}

func (m *SendRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendRequest.Unmarshal(m, b)
}
func (m *SendRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendRequest.Marshal(b, m, deterministic)
}
func (m *SendRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendRequest.Merge(m, src)
}
func (m *SendRequest) XXX_Size() int {
	return xxx_messageInfo_SendRequest.Size(m)
}
func (m *SendRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendRequest proto.InternalMessageInfo

func (m *SendRequest) GetSubId() string {
	if m != nil {
		return m.SubId
	}
	return ""
}

func (m *SendRequest) GetMessage() *Message {
	if m != nil {
		return m.Message
	}
	return nil
}

type SendResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendResponse) Reset()         { *m = SendResponse{} }
func (m *SendResponse) String() string { return proto.CompactTextString(m) }
func (*SendResponse) ProtoMessage()    {}
func (*SendResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f5b84c3e519282d7, []int{2}
}

func (m *SendResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendResponse.Unmarshal(m, b)
}
func (m *SendResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendResponse.Marshal(b, m, deterministic)
}
func (m *SendResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendResponse.Merge(m, src)
}
func (m *SendResponse) XXX_Size() int {
	return xxx_messageInfo_SendResponse.Size(m)
}
func (m *SendResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SendResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SendResponse proto.InternalMessageInfo

type SubscribeRequest struct {
	SubId                string   `protobuf:"bytes,1,opt,name=sub_id,json=subId,proto3" json:"sub_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubscribeRequest) Reset()         { *m = SubscribeRequest{} }
func (m *SubscribeRequest) String() string { return proto.CompactTextString(m) }
func (*SubscribeRequest) ProtoMessage()    {}
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f5b84c3e519282d7, []int{3}
}

func (m *SubscribeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeRequest.Unmarshal(m, b)
}
func (m *SubscribeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeRequest.Marshal(b, m, deterministic)
}
func (m *SubscribeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeRequest.Merge(m, src)
}
func (m *SubscribeRequest) XXX_Size() int {
	return xxx_messageInfo_SubscribeRequest.Size(m)
}
func (m *SubscribeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeRequest proto.InternalMessageInfo

func (m *SubscribeRequest) GetSubId() string {
	if m != nil {
		return m.SubId
	}
	return ""
}

type SubscribeResponse struct {
	Message              *Message `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubscribeResponse) Reset()         { *m = SubscribeResponse{} }
func (m *SubscribeResponse) String() string { return proto.CompactTextString(m) }
func (*SubscribeResponse) ProtoMessage()    {}
func (*SubscribeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f5b84c3e519282d7, []int{4}
}

func (m *SubscribeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeResponse.Unmarshal(m, b)
}
func (m *SubscribeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeResponse.Marshal(b, m, deterministic)
}
func (m *SubscribeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeResponse.Merge(m, src)
}
func (m *SubscribeResponse) XXX_Size() int {
	return xxx_messageInfo_SubscribeResponse.Size(m)
}
func (m *SubscribeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeResponse proto.InternalMessageInfo

func (m *SubscribeResponse) GetMessage() *Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "chat_protobuf.Message")
	proto.RegisterType((*SendRequest)(nil), "chat_protobuf.SendRequest")
	proto.RegisterType((*SendResponse)(nil), "chat_protobuf.SendResponse")
	proto.RegisterType((*SubscribeRequest)(nil), "chat_protobuf.SubscribeRequest")
	proto.RegisterType((*SubscribeResponse)(nil), "chat_protobuf.SubscribeResponse")
}

func init() { proto.RegisterFile("chat-protobuf/chat.proto", fileDescriptor_f5b84c3e519282d7) }

var fileDescriptor_f5b84c3e519282d7 = []byte{
	// 303 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x8f, 0xcd, 0x4e, 0x83, 0x40,
	0x10, 0xc7, 0x5d, 0xad, 0x34, 0x1d, 0xd4, 0xe8, 0x26, 0x2a, 0xc1, 0x43, 0x1b, 0x4e, 0xf5, 0xe0,
	0xb6, 0x69, 0x9f, 0xa0, 0x1a, 0x0f, 0x3d, 0x78, 0xa1, 0xc6, 0x83, 0x17, 0xc2, 0xc2, 0x48, 0x89,
	0xf2, 0x21, 0x3b, 0x18, 0x1f, 0xc8, 0x07, 0x35, 0x2c, 0xa5, 0x22, 0x51, 0xe3, 0x6d, 0x67, 0xe7,
	0x37, 0xff, 0x0f, 0xb0, 0x82, 0xb5, 0x4f, 0x57, 0x79, 0x91, 0x51, 0x26, 0xcb, 0xa7, 0x49, 0x35,
	0x09, 0x3d, 0xf1, 0xc3, 0xea, 0xed, 0x35, 0x1b, 0x7b, 0x18, 0x65, 0x59, 0xf4, 0x82, 0x93, 0x2d,
	0x4a, 0x71, 0x82, 0x8a, 0xfc, 0x24, 0xaf, 0x79, 0xe7, 0x19, 0xfa, 0x77, 0xa8, 0x94, 0x1f, 0x21,
	0x3f, 0x87, 0xbe, 0xc2, 0x34, 0xf4, 0xe2, 0xd0, 0x62, 0x23, 0x36, 0x1e, 0xb8, 0x46, 0x35, 0x2e,
	0x43, 0xce, 0xa1, 0x47, 0xf8, 0x4e, 0xd6, 0xae, 0xfe, 0xd5, 0x6f, 0x3e, 0xd7, 0x30, 0x79, 0x3e,
	0x59, 0x7b, 0x23, 0x36, 0x36, 0x67, 0xb6, 0xa8, 0xad, 0x44, 0x63, 0x25, 0xee, 0x1b, 0x2b, 0x2d,
	0x44, 0x0b, 0x72, 0x1e, 0xc0, 0x5c, 0x61, 0x1a, 0xba, 0xf8, 0x5a, 0xa2, 0x22, 0x7e, 0x0a, 0x86,
	0x2a, 0xe5, 0x97, 0xdf, 0xbe, 0x2a, 0xe5, 0x32, 0xe4, 0x53, 0xe8, 0x27, 0x75, 0x24, 0xed, 0x68,
	0xce, 0xce, 0xc4, 0xb7, 0x52, 0x62, 0x13, 0xd8, 0x6d, 0x30, 0xe7, 0x08, 0x0e, 0x6a, 0x5d, 0x95,
	0x67, 0xa9, 0x42, 0xe7, 0x12, 0x8e, 0x57, 0xa5, 0x54, 0x41, 0x11, 0x4b, 0xfc, 0xdb, 0xcc, 0xb9,
	0x85, 0x93, 0x16, 0x5a, 0xdf, 0xb7, 0x13, 0xb0, 0x7f, 0x25, 0x98, 0x7d, 0x30, 0x30, 0x6f, 0xd6,
	0x3e, 0xad, 0xb0, 0x78, 0x8b, 0x03, 0xe4, 0x0b, 0xe8, 0x55, 0x89, 0xb8, 0xdd, 0x39, 0x6c, 0xd5,
	0xb7, 0x2f, 0x7e, 0xdc, 0x6d, 0x2a, 0xec, 0x70, 0x17, 0x06, 0xdb, 0x64, 0x7c, 0xd8, 0x65, 0x3b,
	0xf5, 0xec, 0xd1, 0xef, 0x40, 0xa3, 0x38, 0x65, 0xd7, 0xc6, 0x63, 0xaf, 0xc2, 0xa4, 0xa1, 0xb9,
	0xf9, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0xcc, 0x12, 0x26, 0xfd, 0x48, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatServiceClient interface {
	Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendResponse, error)
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (ChatService_SubscribeClient, error)
}

type chatServiceClient struct {
	cc *grpc.ClientConn
}

func NewChatServiceClient(cc *grpc.ClientConn) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendResponse, error) {
	out := new(SendResponse)
	err := c.cc.Invoke(ctx, "/chat_protobuf.ChatService/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (ChatService_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ChatService_serviceDesc.Streams[0], "/chat_protobuf.ChatService/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ChatService_SubscribeClient interface {
	Recv() (*SubscribeResponse, error)
	grpc.ClientStream
}

type chatServiceSubscribeClient struct {
	grpc.ClientStream
}

func (x *chatServiceSubscribeClient) Recv() (*SubscribeResponse, error) {
	m := new(SubscribeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatServiceServer is the server API for ChatService service.
type ChatServiceServer interface {
	Send(context.Context, *SendRequest) (*SendResponse, error)
	Subscribe(*SubscribeRequest, ChatService_SubscribeServer) error
}

// UnimplementedChatServiceServer can be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (*UnimplementedChatServiceServer) Send(ctx context.Context, req *SendRequest) (*SendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (*UnimplementedChatServiceServer) Subscribe(req *SubscribeRequest, srv ChatService_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}

func RegisterChatServiceServer(s *grpc.Server, srv ChatServiceServer) {
	s.RegisterService(&_ChatService_serviceDesc, srv)
}

func _ChatService_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat_protobuf.ChatService/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).Send(ctx, req.(*SendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatServiceServer).Subscribe(m, &chatServiceSubscribeServer{stream})
}

type ChatService_SubscribeServer interface {
	Send(*SubscribeResponse) error
	grpc.ServerStream
}

type chatServiceSubscribeServer struct {
	grpc.ServerStream
}

func (x *chatServiceSubscribeServer) Send(m *SubscribeResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _ChatService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chat_protobuf.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _ChatService_Send_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _ChatService_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "chat-protobuf/chat.proto",
}
