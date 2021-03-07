// Code generated by protoc-gen-go. DO NOT EDIT.
// source: slackbot.proto

package slack

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type SendSlackMessageRequest struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendSlackMessageRequest) Reset()         { *m = SendSlackMessageRequest{} }
func (m *SendSlackMessageRequest) String() string { return proto.CompactTextString(m) }
func (*SendSlackMessageRequest) ProtoMessage()    {}
func (*SendSlackMessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_afdd1b263620fbe7, []int{0}
}

func (m *SendSlackMessageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendSlackMessageRequest.Unmarshal(m, b)
}
func (m *SendSlackMessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendSlackMessageRequest.Marshal(b, m, deterministic)
}
func (m *SendSlackMessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendSlackMessageRequest.Merge(m, src)
}
func (m *SendSlackMessageRequest) XXX_Size() int {
	return xxx_messageInfo_SendSlackMessageRequest.Size(m)
}
func (m *SendSlackMessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendSlackMessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendSlackMessageRequest proto.InternalMessageInfo

func (m *SendSlackMessageRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type SendSlackMessageResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendSlackMessageResponse) Reset()         { *m = SendSlackMessageResponse{} }
func (m *SendSlackMessageResponse) String() string { return proto.CompactTextString(m) }
func (*SendSlackMessageResponse) ProtoMessage()    {}
func (*SendSlackMessageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_afdd1b263620fbe7, []int{1}
}

func (m *SendSlackMessageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendSlackMessageResponse.Unmarshal(m, b)
}
func (m *SendSlackMessageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendSlackMessageResponse.Marshal(b, m, deterministic)
}
func (m *SendSlackMessageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendSlackMessageResponse.Merge(m, src)
}
func (m *SendSlackMessageResponse) XXX_Size() int {
	return xxx_messageInfo_SendSlackMessageResponse.Size(m)
}
func (m *SendSlackMessageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SendSlackMessageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SendSlackMessageResponse proto.InternalMessageInfo

func (m *SendSlackMessageResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*SendSlackMessageRequest)(nil), "slack.SendSlackMessageRequest")
	proto.RegisterType((*SendSlackMessageResponse)(nil), "slack.SendSlackMessageResponse")
}

func init() {
	proto.RegisterFile("slackbot.proto", fileDescriptor_afdd1b263620fbe7)
}

var fileDescriptor_afdd1b263620fbe7 = []byte{
	// 116 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0xce, 0x49, 0x4c,
	0xce, 0x4e, 0xca, 0x2f, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0xf3, 0x95, 0x8c,
	0xb9, 0xc4, 0x83, 0x53, 0xf3, 0x52, 0x82, 0x41, 0x1c, 0xdf, 0xd4, 0xe2, 0xe2, 0xc4, 0xf4, 0xd4,
	0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x09, 0x2e, 0xf6, 0x5c, 0x88, 0x88, 0x04, 0xa3,
	0x02, 0xa3, 0x06, 0x67, 0x10, 0x8c, 0xab, 0x64, 0xc2, 0x25, 0x81, 0xa9, 0xa9, 0xb8, 0x20, 0x3f,
	0xaf, 0x38, 0x15, 0xa4, 0xab, 0xb8, 0x34, 0x39, 0x39, 0xb5, 0xb8, 0x18, 0xac, 0x8b, 0x23, 0x08,
	0xc6, 0x4d, 0x62, 0x03, 0x5b, 0x6c, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x91, 0x57, 0xa0, 0x1b,
	0x8a, 0x00, 0x00, 0x00,
}
