// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: thought/v1/thought.proto

package thoughtv1

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

type GetThoughtRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetThoughtRequest) Reset() {
	*x = GetThoughtRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thought_v1_thought_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetThoughtRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetThoughtRequest) ProtoMessage() {}

func (x *GetThoughtRequest) ProtoReflect() protoreflect.Message {
	mi := &file_thought_v1_thought_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetThoughtRequest.ProtoReflect.Descriptor instead.
func (*GetThoughtRequest) Descriptor() ([]byte, []int) {
	return file_thought_v1_thought_proto_rawDescGZIP(), []int{0}
}

type GetThoughtResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Thought string `protobuf:"bytes,1,opt,name=thought,proto3" json:"thought,omitempty"`
}

func (x *GetThoughtResponse) Reset() {
	*x = GetThoughtResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thought_v1_thought_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetThoughtResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetThoughtResponse) ProtoMessage() {}

func (x *GetThoughtResponse) ProtoReflect() protoreflect.Message {
	mi := &file_thought_v1_thought_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetThoughtResponse.ProtoReflect.Descriptor instead.
func (*GetThoughtResponse) Descriptor() ([]byte, []int) {
	return file_thought_v1_thought_proto_rawDescGZIP(), []int{1}
}

func (x *GetThoughtResponse) GetThought() string {
	if x != nil {
		return x.Thought
	}
	return ""
}

type ReceiveThoughtsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReceiveThoughtsRequest) Reset() {
	*x = ReceiveThoughtsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thought_v1_thought_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceiveThoughtsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiveThoughtsRequest) ProtoMessage() {}

func (x *ReceiveThoughtsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_thought_v1_thought_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiveThoughtsRequest.ProtoReflect.Descriptor instead.
func (*ReceiveThoughtsRequest) Descriptor() ([]byte, []int) {
	return file_thought_v1_thought_proto_rawDescGZIP(), []int{2}
}

type ReceiveThoughtsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Thought string `protobuf:"bytes,1,opt,name=thought,proto3" json:"thought,omitempty"`
}

func (x *ReceiveThoughtsResponse) Reset() {
	*x = ReceiveThoughtsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thought_v1_thought_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceiveThoughtsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiveThoughtsResponse) ProtoMessage() {}

func (x *ReceiveThoughtsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_thought_v1_thought_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiveThoughtsResponse.ProtoReflect.Descriptor instead.
func (*ReceiveThoughtsResponse) Descriptor() ([]byte, []int) {
	return file_thought_v1_thought_proto_rawDescGZIP(), []int{3}
}

func (x *ReceiveThoughtsResponse) GetThought() string {
	if x != nil {
		return x.Thought
	}
	return ""
}

type SendThoughtsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Thought string `protobuf:"bytes,1,opt,name=thought,proto3" json:"thought,omitempty"`
}

func (x *SendThoughtsRequest) Reset() {
	*x = SendThoughtsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thought_v1_thought_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendThoughtsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendThoughtsRequest) ProtoMessage() {}

func (x *SendThoughtsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_thought_v1_thought_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendThoughtsRequest.ProtoReflect.Descriptor instead.
func (*SendThoughtsRequest) Descriptor() ([]byte, []int) {
	return file_thought_v1_thought_proto_rawDescGZIP(), []int{4}
}

func (x *SendThoughtsRequest) GetThought() string {
	if x != nil {
		return x.Thought
	}
	return ""
}

type SendThoughtsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SendThoughtsResponse) Reset() {
	*x = SendThoughtsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thought_v1_thought_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendThoughtsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendThoughtsResponse) ProtoMessage() {}

func (x *SendThoughtsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_thought_v1_thought_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendThoughtsResponse.ProtoReflect.Descriptor instead.
func (*SendThoughtsResponse) Descriptor() ([]byte, []int) {
	return file_thought_v1_thought_proto_rawDescGZIP(), []int{5}
}

func (x *SendThoughtsResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ConverseThoughtsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Thought string `protobuf:"bytes,1,opt,name=thought,proto3" json:"thought,omitempty"`
}

func (x *ConverseThoughtsRequest) Reset() {
	*x = ConverseThoughtsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thought_v1_thought_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConverseThoughtsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConverseThoughtsRequest) ProtoMessage() {}

func (x *ConverseThoughtsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_thought_v1_thought_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConverseThoughtsRequest.ProtoReflect.Descriptor instead.
func (*ConverseThoughtsRequest) Descriptor() ([]byte, []int) {
	return file_thought_v1_thought_proto_rawDescGZIP(), []int{6}
}

func (x *ConverseThoughtsRequest) GetThought() string {
	if x != nil {
		return x.Thought
	}
	return ""
}

type ConverseThoughtsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Thought string `protobuf:"bytes,1,opt,name=thought,proto3" json:"thought,omitempty"`
}

func (x *ConverseThoughtsResponse) Reset() {
	*x = ConverseThoughtsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thought_v1_thought_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConverseThoughtsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConverseThoughtsResponse) ProtoMessage() {}

func (x *ConverseThoughtsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_thought_v1_thought_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConverseThoughtsResponse.ProtoReflect.Descriptor instead.
func (*ConverseThoughtsResponse) Descriptor() ([]byte, []int) {
	return file_thought_v1_thought_proto_rawDescGZIP(), []int{7}
}

func (x *ConverseThoughtsResponse) GetThought() string {
	if x != nil {
		return x.Thought
	}
	return ""
}

var File_thought_v1_thought_proto protoreflect.FileDescriptor

var file_thought_v1_thought_proto_rawDesc = []byte{
	0x0a, 0x18, 0x74, 0x68, 0x6f, 0x75, 0x67, 0x68, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x68, 0x6f,
	0x75, 0x67, 0x68, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74, 0x68, 0x6f, 0x75,
	0x67, 0x68, 0x74, 0x2e, 0x76, 0x31, 0x22, 0x13, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x54, 0x68, 0x6f,
	0x75, 0x67, 0x68, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2e, 0x0a, 0x12, 0x47,
	0x65, 0x74, 0x54, 0x68, 0x6f, 0x75, 0x67, 0x68, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x68, 0x6f, 0x75, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x74, 0x68, 0x6f, 0x75, 0x67, 0x68, 0x74, 0x22, 0x18, 0x0a, 0x16, 0x52,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x54, 0x68, 0x6f, 0x75, 0x67, 0x68, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x33, 0x0a, 0x17, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x54, 0x68, 0x6f, 0x75, 0x67, 0x68, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x74, 0x68, 0x6f, 0x75, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x74, 0x68, 0x6f, 0x75, 0x67, 0x68, 0x74, 0x22, 0x2f, 0x0a, 0x13, 0x53, 0x65,
	0x6e, 0x64, 0x54, 0x68, 0x6f, 0x75, 0x67, 0x68, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x68, 0x6f, 0x75, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x74, 0x68, 0x6f, 0x75, 0x67, 0x68, 0x74, 0x22, 0x30, 0x0a, 0x14, 0x53,
	0x65, 0x6e, 0x64, 0x54, 0x68, 0x6f, 0x75, 0x67, 0x68, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x33, 0x0a,
	0x17, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x65, 0x54, 0x68, 0x6f, 0x75, 0x67, 0x68, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x68, 0x6f, 0x75,
	0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x68, 0x6f, 0x75, 0x67,
	0x68, 0x74, 0x22, 0x34, 0x0a, 0x18, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x65, 0x54, 0x68,
	0x6f, 0x75, 0x67, 0x68, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x74, 0x68, 0x6f, 0x75, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x74, 0x68, 0x6f, 0x75, 0x67, 0x68, 0x74, 0x32, 0xf3, 0x02, 0x0a, 0x0e, 0x54, 0x68, 0x6f,
	0x75, 0x67, 0x68, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x0a, 0x47,
	0x65, 0x74, 0x54, 0x68, 0x6f, 0x75, 0x67, 0x68, 0x74, 0x12, 0x1d, 0x2e, 0x74, 0x68, 0x6f, 0x75,
	0x67, 0x68, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x68, 0x6f, 0x75, 0x67, 0x68,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x74, 0x68, 0x6f, 0x75, 0x67,
	0x68, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x68, 0x6f, 0x75, 0x67, 0x68, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5c, 0x0a, 0x0f, 0x52, 0x65, 0x63, 0x65,
	0x69, 0x76, 0x65, 0x54, 0x68, 0x6f, 0x75, 0x67, 0x68, 0x74, 0x73, 0x12, 0x22, 0x2e, 0x74, 0x68,
	0x6f, 0x75, 0x67, 0x68, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x54, 0x68, 0x6f, 0x75, 0x67, 0x68, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x23, 0x2e, 0x74, 0x68, 0x6f, 0x75, 0x67, 0x68, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x63,
	0x65, 0x69, 0x76, 0x65, 0x54, 0x68, 0x6f, 0x75, 0x67, 0x68, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x53, 0x0a, 0x0c, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x68,
	0x6f, 0x75, 0x67, 0x68, 0x74, 0x73, 0x12, 0x1f, 0x2e, 0x74, 0x68, 0x6f, 0x75, 0x67, 0x68, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x68, 0x6f, 0x75, 0x67, 0x68, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x74, 0x68, 0x6f, 0x75, 0x67, 0x68,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x68, 0x6f, 0x75, 0x67, 0x68, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x12, 0x61, 0x0a, 0x10, 0x43,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x65, 0x54, 0x68, 0x6f, 0x75, 0x67, 0x68, 0x74, 0x73, 0x12,
	0x23, 0x2e, 0x74, 0x68, 0x6f, 0x75, 0x67, 0x68, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e,
	0x76, 0x65, 0x72, 0x73, 0x65, 0x54, 0x68, 0x6f, 0x75, 0x67, 0x68, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x74, 0x68, 0x6f, 0x75, 0x67, 0x68, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x65, 0x54, 0x68, 0x6f, 0x75, 0x67, 0x68,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x97,
	0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x68, 0x6f, 0x75, 0x67, 0x68, 0x74, 0x2e, 0x76,
	0x31, 0x42, 0x0c, 0x54, 0x68, 0x6f, 0x75, 0x67, 0x68, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x2e, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x68, 0x6f,
	0x75, 0x67, 0x68, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x68, 0x6f, 0x75, 0x67, 0x68, 0x74, 0x76,
	0x31, 0xa2, 0x02, 0x03, 0x54, 0x58, 0x58, 0xaa, 0x02, 0x0a, 0x54, 0x68, 0x6f, 0x75, 0x67, 0x68,
	0x74, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0a, 0x54, 0x68, 0x6f, 0x75, 0x67, 0x68, 0x74, 0x5c, 0x56,
	0x31, 0xe2, 0x02, 0x16, 0x54, 0x68, 0x6f, 0x75, 0x67, 0x68, 0x74, 0x5c, 0x56, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0b, 0x54, 0x68, 0x6f,
	0x75, 0x67, 0x68, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_thought_v1_thought_proto_rawDescOnce sync.Once
	file_thought_v1_thought_proto_rawDescData = file_thought_v1_thought_proto_rawDesc
)

func file_thought_v1_thought_proto_rawDescGZIP() []byte {
	file_thought_v1_thought_proto_rawDescOnce.Do(func() {
		file_thought_v1_thought_proto_rawDescData = protoimpl.X.CompressGZIP(file_thought_v1_thought_proto_rawDescData)
	})
	return file_thought_v1_thought_proto_rawDescData
}

var file_thought_v1_thought_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_thought_v1_thought_proto_goTypes = []interface{}{
	(*GetThoughtRequest)(nil),        // 0: thought.v1.GetThoughtRequest
	(*GetThoughtResponse)(nil),       // 1: thought.v1.GetThoughtResponse
	(*ReceiveThoughtsRequest)(nil),   // 2: thought.v1.ReceiveThoughtsRequest
	(*ReceiveThoughtsResponse)(nil),  // 3: thought.v1.ReceiveThoughtsResponse
	(*SendThoughtsRequest)(nil),      // 4: thought.v1.SendThoughtsRequest
	(*SendThoughtsResponse)(nil),     // 5: thought.v1.SendThoughtsResponse
	(*ConverseThoughtsRequest)(nil),  // 6: thought.v1.ConverseThoughtsRequest
	(*ConverseThoughtsResponse)(nil), // 7: thought.v1.ConverseThoughtsResponse
}
var file_thought_v1_thought_proto_depIdxs = []int32{
	0, // 0: thought.v1.ThoughtService.GetThought:input_type -> thought.v1.GetThoughtRequest
	2, // 1: thought.v1.ThoughtService.ReceiveThoughts:input_type -> thought.v1.ReceiveThoughtsRequest
	4, // 2: thought.v1.ThoughtService.SendThoughts:input_type -> thought.v1.SendThoughtsRequest
	6, // 3: thought.v1.ThoughtService.ConverseThoughts:input_type -> thought.v1.ConverseThoughtsRequest
	1, // 4: thought.v1.ThoughtService.GetThought:output_type -> thought.v1.GetThoughtResponse
	3, // 5: thought.v1.ThoughtService.ReceiveThoughts:output_type -> thought.v1.ReceiveThoughtsResponse
	5, // 6: thought.v1.ThoughtService.SendThoughts:output_type -> thought.v1.SendThoughtsResponse
	7, // 7: thought.v1.ThoughtService.ConverseThoughts:output_type -> thought.v1.ConverseThoughtsResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_thought_v1_thought_proto_init() }
func file_thought_v1_thought_proto_init() {
	if File_thought_v1_thought_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_thought_v1_thought_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetThoughtRequest); i {
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
		file_thought_v1_thought_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetThoughtResponse); i {
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
		file_thought_v1_thought_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReceiveThoughtsRequest); i {
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
		file_thought_v1_thought_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReceiveThoughtsResponse); i {
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
		file_thought_v1_thought_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendThoughtsRequest); i {
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
		file_thought_v1_thought_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendThoughtsResponse); i {
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
		file_thought_v1_thought_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConverseThoughtsRequest); i {
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
		file_thought_v1_thought_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConverseThoughtsResponse); i {
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
			RawDescriptor: file_thought_v1_thought_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_thought_v1_thought_proto_goTypes,
		DependencyIndexes: file_thought_v1_thought_proto_depIdxs,
		MessageInfos:      file_thought_v1_thought_proto_msgTypes,
	}.Build()
	File_thought_v1_thought_proto = out.File
	file_thought_v1_thought_proto_rawDesc = nil
	file_thought_v1_thought_proto_goTypes = nil
	file_thought_v1_thought_proto_depIdxs = nil
}
