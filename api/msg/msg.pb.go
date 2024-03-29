// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: api/msg/msg.proto

package msg

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SubscribeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 消息id
	MsgId uint64 `protobuf:"varint,1,opt,name=msgId,proto3" json:"msgId,omitempty"`
	// 主题列表
	Topics []string `protobuf:"bytes,2,rep,name=topics,proto3" json:"topics,omitempty"`
}

func (x *SubscribeReq) Reset() {
	*x = SubscribeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_msg_msg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeReq) ProtoMessage() {}

func (x *SubscribeReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_msg_msg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeReq.ProtoReflect.Descriptor instead.
func (*SubscribeReq) Descriptor() ([]byte, []int) {
	return file_api_msg_msg_proto_rawDescGZIP(), []int{0}
}

func (x *SubscribeReq) GetMsgId() uint64 {
	if x != nil {
		return x.MsgId
	}
	return 0
}

func (x *SubscribeReq) GetTopics() []string {
	if x != nil {
		return x.Topics
	}
	return nil
}

type SubscribeAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 消息id
	MsgId uint64 `protobuf:"varint,1,opt,name=msgId,proto3" json:"msgId,omitempty"`
}

func (x *SubscribeAck) Reset() {
	*x = SubscribeAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_msg_msg_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeAck) ProtoMessage() {}

func (x *SubscribeAck) ProtoReflect() protoreflect.Message {
	mi := &file_api_msg_msg_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeAck.ProtoReflect.Descriptor instead.
func (*SubscribeAck) Descriptor() ([]byte, []int) {
	return file_api_msg_msg_proto_rawDescGZIP(), []int{1}
}

func (x *SubscribeAck) GetMsgId() uint64 {
	if x != nil {
		return x.MsgId
	}
	return 0
}

type UnsubscribeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 消息id
	MsgId string `protobuf:"bytes,1,opt,name=msgId,proto3" json:"msgId,omitempty"`
	// 主题列表
	Topics []string `protobuf:"bytes,2,rep,name=topics,proto3" json:"topics,omitempty"`
}

func (x *UnsubscribeReq) Reset() {
	*x = UnsubscribeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_msg_msg_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnsubscribeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnsubscribeReq) ProtoMessage() {}

func (x *UnsubscribeReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_msg_msg_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnsubscribeReq.ProtoReflect.Descriptor instead.
func (*UnsubscribeReq) Descriptor() ([]byte, []int) {
	return file_api_msg_msg_proto_rawDescGZIP(), []int{2}
}

func (x *UnsubscribeReq) GetMsgId() string {
	if x != nil {
		return x.MsgId
	}
	return ""
}

func (x *UnsubscribeReq) GetTopics() []string {
	if x != nil {
		return x.Topics
	}
	return nil
}

type UnsubscribeAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 消息id
	MsgId uint64 `protobuf:"varint,1,opt,name=msgId,proto3" json:"msgId,omitempty"`
}

func (x *UnsubscribeAck) Reset() {
	*x = UnsubscribeAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_msg_msg_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnsubscribeAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnsubscribeAck) ProtoMessage() {}

func (x *UnsubscribeAck) ProtoReflect() protoreflect.Message {
	mi := &file_api_msg_msg_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnsubscribeAck.ProtoReflect.Descriptor instead.
func (*UnsubscribeAck) Descriptor() ([]byte, []int) {
	return file_api_msg_msg_proto_rawDescGZIP(), []int{3}
}

func (x *UnsubscribeAck) GetMsgId() uint64 {
	if x != nil {
		return x.MsgId
	}
	return 0
}

type PublishReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 主题通道
	Topic string `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	// 消息id
	MsgId uint64 `protobuf:"varint,2,opt,name=msgId,proto3" json:"msgId,omitempty"`
	// 消息内容
	Payload []byte `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *PublishReq) Reset() {
	*x = PublishReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_msg_msg_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishReq) ProtoMessage() {}

func (x *PublishReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_msg_msg_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishReq.ProtoReflect.Descriptor instead.
func (*PublishReq) Descriptor() ([]byte, []int) {
	return file_api_msg_msg_proto_rawDescGZIP(), []int{4}
}

func (x *PublishReq) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *PublishReq) GetMsgId() uint64 {
	if x != nil {
		return x.MsgId
	}
	return 0
}

func (x *PublishReq) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

type PublishAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 消息id
	MsgId uint64 `protobuf:"varint,1,opt,name=msgId,proto3" json:"msgId,omitempty"`
}

func (x *PublishAck) Reset() {
	*x = PublishAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_msg_msg_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishAck) ProtoMessage() {}

func (x *PublishAck) ProtoReflect() protoreflect.Message {
	mi := &file_api_msg_msg_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishAck.ProtoReflect.Descriptor instead.
func (*PublishAck) Descriptor() ([]byte, []int) {
	return file_api_msg_msg_proto_rawDescGZIP(), []int{5}
}

func (x *PublishAck) GetMsgId() uint64 {
	if x != nil {
		return x.MsgId
	}
	return 0
}

type PingReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PingReq) Reset() {
	*x = PingReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_msg_msg_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingReq) ProtoMessage() {}

func (x *PingReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_msg_msg_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingReq.ProtoReflect.Descriptor instead.
func (*PingReq) Descriptor() ([]byte, []int) {
	return file_api_msg_msg_proto_rawDescGZIP(), []int{6}
}

type PingAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PingAck) Reset() {
	*x = PingAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_msg_msg_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingAck) ProtoMessage() {}

func (x *PingAck) ProtoReflect() protoreflect.Message {
	mi := &file_api_msg_msg_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingAck.ProtoReflect.Descriptor instead.
func (*PingAck) Descriptor() ([]byte, []int) {
	return file_api_msg_msg_proto_rawDescGZIP(), []int{7}
}

type ConnectReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ConnectReq) Reset() {
	*x = ConnectReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_msg_msg_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectReq) ProtoMessage() {}

func (x *ConnectReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_msg_msg_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectReq.ProtoReflect.Descriptor instead.
func (*ConnectReq) Descriptor() ([]byte, []int) {
	return file_api_msg_msg_proto_rawDescGZIP(), []int{8}
}

type ConnectAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ConnectAck) Reset() {
	*x = ConnectAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_msg_msg_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectAck) ProtoMessage() {}

func (x *ConnectAck) ProtoReflect() protoreflect.Message {
	mi := &file_api_msg_msg_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectAck.ProtoReflect.Descriptor instead.
func (*ConnectAck) Descriptor() ([]byte, []int) {
	return file_api_msg_msg_proto_rawDescGZIP(), []int{9}
}

type DisconnectReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DisconnectReq) Reset() {
	*x = DisconnectReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_msg_msg_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisconnectReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisconnectReq) ProtoMessage() {}

func (x *DisconnectReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_msg_msg_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisconnectReq.ProtoReflect.Descriptor instead.
func (*DisconnectReq) Descriptor() ([]byte, []int) {
	return file_api_msg_msg_proto_rawDescGZIP(), []int{10}
}

type DisconnectAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DisconnectAck) Reset() {
	*x = DisconnectAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_msg_msg_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisconnectAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisconnectAck) ProtoMessage() {}

func (x *DisconnectAck) ProtoReflect() protoreflect.Message {
	mi := &file_api_msg_msg_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisconnectAck.ProtoReflect.Descriptor instead.
func (*DisconnectAck) Descriptor() ([]byte, []int) {
	return file_api_msg_msg_proto_rawDescGZIP(), []int{11}
}

type PushMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户id
	Topic string `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	// 消息id
	MsgId uint64 `protobuf:"varint,2,opt,name=msgId,proto3" json:"msgId,omitempty"`
	// 消息内容
	Payload []byte `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *PushMsg) Reset() {
	*x = PushMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_msg_msg_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushMsg) ProtoMessage() {}

func (x *PushMsg) ProtoReflect() protoreflect.Message {
	mi := &file_api_msg_msg_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushMsg.ProtoReflect.Descriptor instead.
func (*PushMsg) Descriptor() ([]byte, []int) {
	return file_api_msg_msg_proto_rawDescGZIP(), []int{12}
}

func (x *PushMsg) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *PushMsg) GetMsgId() uint64 {
	if x != nil {
		return x.MsgId
	}
	return 0
}

func (x *PushMsg) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

type PushMsgResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PushMsgResp) Reset() {
	*x = PushMsgResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_msg_msg_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushMsgResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushMsgResp) ProtoMessage() {}

func (x *PushMsgResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_msg_msg_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushMsgResp.ProtoReflect.Descriptor instead.
func (*PushMsgResp) Descriptor() ([]byte, []int) {
	return file_api_msg_msg_proto_rawDescGZIP(), []int{13}
}

type BroadcastResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BroadcastResp) Reset() {
	*x = BroadcastResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_msg_msg_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BroadcastResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastResp) ProtoMessage() {}

func (x *BroadcastResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_msg_msg_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastResp.ProtoReflect.Descriptor instead.
func (*BroadcastResp) Descriptor() ([]byte, []int) {
	return file_api_msg_msg_proto_rawDescGZIP(), []int{14}
}

var File_api_msg_msg_proto protoreflect.FileDescriptor

var file_api_msg_msg_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x73, 0x67, 0x2f, 0x6d, 0x73, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x72, 0x65, 0x61, 0x6c, 0x74, 0x69, 0x6d, 0x65, 0x78, 0x2e, 0x6d,
	0x73, 0x67, 0x22, 0x3c, 0x0a, 0x0c, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52,
	0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x73, 0x67, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x05, 0x6d, 0x73, 0x67, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x6f, 0x70, 0x69,
	0x63, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73,
	0x22, 0x24, 0x0a, 0x0c, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x41, 0x63, 0x6b,
	0x12, 0x14, 0x0a, 0x05, 0x6d, 0x73, 0x67, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x05, 0x6d, 0x73, 0x67, 0x49, 0x64, 0x22, 0x3e, 0x0a, 0x0e, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x73, 0x67, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x73, 0x67, 0x49, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06,
	0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x22, 0x26, 0x0a, 0x0e, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x41, 0x63, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x73, 0x67, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6d, 0x73, 0x67, 0x49, 0x64, 0x22, 0x52,
	0x0a, 0x0a, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70,
	0x69, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x73, 0x67, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x05, 0x6d, 0x73, 0x67, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x22, 0x22, 0x0a, 0x0a, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x41, 0x63, 0x6b,
	0x12, 0x14, 0x0a, 0x05, 0x6d, 0x73, 0x67, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x05, 0x6d, 0x73, 0x67, 0x49, 0x64, 0x22, 0x09, 0x0a, 0x07, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x22, 0x09, 0x0a, 0x07, 0x50, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x6b, 0x22, 0x0c, 0x0a, 0x0a,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x22, 0x0c, 0x0a, 0x0a, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x41, 0x63, 0x6b, 0x22, 0x0f, 0x0a, 0x0d, 0x44, 0x69, 0x73, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x22, 0x0f, 0x0a, 0x0d, 0x44, 0x69, 0x73,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x41, 0x63, 0x6b, 0x22, 0x4f, 0x0a, 0x07, 0x50, 0x75,
	0x73, 0x68, 0x4d, 0x73, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x6d,
	0x73, 0x67, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6d, 0x73, 0x67, 0x49,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x0d, 0x0a, 0x0b, 0x50,
	0x75, 0x73, 0x68, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x73, 0x70, 0x22, 0x0f, 0x0a, 0x0d, 0x42, 0x72,
	0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x32, 0xa4, 0x03, 0x0a, 0x0a,
	0x4d, 0x73, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x07, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x19, 0x2e, 0x72, 0x65, 0x61, 0x6c, 0x74, 0x69, 0x6d, 0x65,
	0x78, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71,
	0x1a, 0x19, 0x2e, 0x72, 0x65, 0x61, 0x6c, 0x74, 0x69, 0x6d, 0x65, 0x78, 0x2e, 0x6d, 0x73, 0x67,
	0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x41, 0x63, 0x6b, 0x12, 0x45, 0x0a, 0x09, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x1b, 0x2e, 0x72, 0x65, 0x61, 0x6c, 0x74,
	0x69, 0x6d, 0x65, 0x78, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x72, 0x65, 0x61, 0x6c, 0x74, 0x69, 0x6d, 0x65,
	0x78, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x41,
	0x63, 0x6b, 0x12, 0x4b, 0x0a, 0x0b, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x12, 0x1d, 0x2e, 0x72, 0x65, 0x61, 0x6c, 0x74, 0x69, 0x6d, 0x65, 0x78, 0x2e, 0x6d, 0x73,
	0x67, 0x2e, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71,
	0x1a, 0x1d, 0x2e, 0x72, 0x65, 0x61, 0x6c, 0x74, 0x69, 0x6d, 0x65, 0x78, 0x2e, 0x6d, 0x73, 0x67,
	0x2e, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x41, 0x63, 0x6b, 0x12,
	0x3f, 0x0a, 0x07, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x12, 0x19, 0x2e, 0x72, 0x65, 0x61,
	0x6c, 0x74, 0x69, 0x6d, 0x65, 0x78, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x72, 0x65, 0x61, 0x6c, 0x74, 0x69, 0x6d, 0x65,
	0x78, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x41, 0x63, 0x6b,
	0x12, 0x36, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x2e, 0x72, 0x65, 0x61, 0x6c, 0x74,
	0x69, 0x6d, 0x65, 0x78, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71,
	0x1a, 0x16, 0x2e, 0x72, 0x65, 0x61, 0x6c, 0x74, 0x69, 0x6d, 0x65, 0x78, 0x2e, 0x6d, 0x73, 0x67,
	0x2e, 0x50, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x6b, 0x12, 0x48, 0x0a, 0x0a, 0x44, 0x69, 0x73, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x1c, 0x2e, 0x72, 0x65, 0x61, 0x6c, 0x74, 0x69, 0x6d,
	0x65, 0x78, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x72, 0x65, 0x61, 0x6c, 0x74, 0x69, 0x6d, 0x65, 0x78,
	0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x41,
	0x63, 0x6b, 0x32, 0x8c, 0x01, 0x0a, 0x0b, 0x50, 0x75, 0x73, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x41, 0x0a, 0x09, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x12,
	0x16, 0x2e, 0x72, 0x65, 0x61, 0x6c, 0x74, 0x69, 0x6d, 0x65, 0x78, 0x2e, 0x6d, 0x73, 0x67, 0x2e,
	0x50, 0x75, 0x73, 0x68, 0x4d, 0x73, 0x67, 0x1a, 0x1c, 0x2e, 0x72, 0x65, 0x61, 0x6c, 0x74, 0x69,
	0x6d, 0x65, 0x78, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3a, 0x0a, 0x04, 0x50, 0x75, 0x73, 0x68, 0x12, 0x16, 0x2e,
	0x72, 0x65, 0x61, 0x6c, 0x74, 0x69, 0x6d, 0x65, 0x78, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x50, 0x75,
	0x73, 0x68, 0x4d, 0x73, 0x67, 0x1a, 0x1a, 0x2e, 0x72, 0x65, 0x61, 0x6c, 0x74, 0x69, 0x6d, 0x65,
	0x78, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6c, 0x79, 0x6f, 0x75, 0x6e, 0x74, 0x68, 0x7a, 0x7a, 0x7a, 0x2f, 0x72, 0x65, 0x61, 0x6c, 0x74,
	0x69, 0x6d, 0x65, 0x78, 0x2f, 0x6d, 0x73, 0x67, 0x3b, 0x6d, 0x73, 0x67, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_msg_msg_proto_rawDescOnce sync.Once
	file_api_msg_msg_proto_rawDescData = file_api_msg_msg_proto_rawDesc
)

func file_api_msg_msg_proto_rawDescGZIP() []byte {
	file_api_msg_msg_proto_rawDescOnce.Do(func() {
		file_api_msg_msg_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_msg_msg_proto_rawDescData)
	})
	return file_api_msg_msg_proto_rawDescData
}

var file_api_msg_msg_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_api_msg_msg_proto_goTypes = []interface{}{
	(*SubscribeReq)(nil),   // 0: realtimex.msg.SubscribeReq
	(*SubscribeAck)(nil),   // 1: realtimex.msg.SubscribeAck
	(*UnsubscribeReq)(nil), // 2: realtimex.msg.UnsubscribeReq
	(*UnsubscribeAck)(nil), // 3: realtimex.msg.UnsubscribeAck
	(*PublishReq)(nil),     // 4: realtimex.msg.PublishReq
	(*PublishAck)(nil),     // 5: realtimex.msg.PublishAck
	(*PingReq)(nil),        // 6: realtimex.msg.PingReq
	(*PingAck)(nil),        // 7: realtimex.msg.PingAck
	(*ConnectReq)(nil),     // 8: realtimex.msg.ConnectReq
	(*ConnectAck)(nil),     // 9: realtimex.msg.ConnectAck
	(*DisconnectReq)(nil),  // 10: realtimex.msg.DisconnectReq
	(*DisconnectAck)(nil),  // 11: realtimex.msg.DisconnectAck
	(*PushMsg)(nil),        // 12: realtimex.msg.PushMsg
	(*PushMsgResp)(nil),    // 13: realtimex.msg.PushMsgResp
	(*BroadcastResp)(nil),  // 14: realtimex.msg.BroadcastResp
}
var file_api_msg_msg_proto_depIdxs = []int32{
	8,  // 0: realtimex.msg.MsgService.Connect:input_type -> realtimex.msg.ConnectReq
	0,  // 1: realtimex.msg.MsgService.Subscribe:input_type -> realtimex.msg.SubscribeReq
	2,  // 2: realtimex.msg.MsgService.Unsubscribe:input_type -> realtimex.msg.UnsubscribeReq
	4,  // 3: realtimex.msg.MsgService.Publish:input_type -> realtimex.msg.PublishReq
	6,  // 4: realtimex.msg.MsgService.Ping:input_type -> realtimex.msg.PingReq
	10, // 5: realtimex.msg.MsgService.Disconnect:input_type -> realtimex.msg.DisconnectReq
	12, // 6: realtimex.msg.PushService.Broadcast:input_type -> realtimex.msg.PushMsg
	12, // 7: realtimex.msg.PushService.Push:input_type -> realtimex.msg.PushMsg
	9,  // 8: realtimex.msg.MsgService.Connect:output_type -> realtimex.msg.ConnectAck
	1,  // 9: realtimex.msg.MsgService.Subscribe:output_type -> realtimex.msg.SubscribeAck
	3,  // 10: realtimex.msg.MsgService.Unsubscribe:output_type -> realtimex.msg.UnsubscribeAck
	5,  // 11: realtimex.msg.MsgService.Publish:output_type -> realtimex.msg.PublishAck
	7,  // 12: realtimex.msg.MsgService.Ping:output_type -> realtimex.msg.PingAck
	11, // 13: realtimex.msg.MsgService.Disconnect:output_type -> realtimex.msg.DisconnectAck
	14, // 14: realtimex.msg.PushService.Broadcast:output_type -> realtimex.msg.BroadcastResp
	13, // 15: realtimex.msg.PushService.Push:output_type -> realtimex.msg.PushMsgResp
	8,  // [8:16] is the sub-list for method output_type
	0,  // [0:8] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_api_msg_msg_proto_init() }
func file_api_msg_msg_proto_init() {
	if File_api_msg_msg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_msg_msg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_msg_msg_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeAck); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_msg_msg_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnsubscribeReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_msg_msg_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnsubscribeAck); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_msg_msg_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_msg_msg_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishAck); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_msg_msg_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_msg_msg_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingAck); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_msg_msg_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_msg_msg_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectAck); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_msg_msg_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisconnectReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_msg_msg_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisconnectAck); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_msg_msg_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushMsg); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_msg_msg_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushMsgResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_msg_msg_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BroadcastResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_msg_msg_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_api_msg_msg_proto_goTypes,
		DependencyIndexes: file_api_msg_msg_proto_depIdxs,
		MessageInfos:      file_api_msg_msg_proto_msgTypes,
	}.Build()
	File_api_msg_msg_proto = out.File
	file_api_msg_msg_proto_rawDesc = nil
	file_api_msg_msg_proto_goTypes = nil
	file_api_msg_msg_proto_depIdxs = nil
}
