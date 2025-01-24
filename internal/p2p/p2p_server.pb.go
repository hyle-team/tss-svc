// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        (unknown)
// source: p2p_server.proto

package p2p

import (
	_ "github.com/hyle-team/tss-svc/internal/types"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PartyStatus int32

const (
	PartyStatus_PS_UNKNOWN PartyStatus = 0
	PartyStatus_PS_KEYGEN  PartyStatus = 1
	PartyStatus_PS_SIGN    PartyStatus = 2
)

// Enum value maps for PartyStatus.
var (
	PartyStatus_name = map[int32]string{
		0: "PS_UNKNOWN",
		1: "PS_KEYGEN",
		2: "PS_SIGN",
	}
	PartyStatus_value = map[string]int32{
		"PS_UNKNOWN": 0,
		"PS_KEYGEN":  1,
		"PS_SIGN":    2,
	}
)

func (x PartyStatus) Enum() *PartyStatus {
	p := new(PartyStatus)
	*p = x
	return p
}

func (x PartyStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PartyStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_p2p_server_proto_enumTypes[0].Descriptor()
}

func (PartyStatus) Type() protoreflect.EnumType {
	return &file_p2p_server_proto_enumTypes[0]
}

func (x PartyStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PartyStatus.Descriptor instead.
func (PartyStatus) EnumDescriptor() ([]byte, []int) {
	return file_p2p_server_proto_rawDescGZIP(), []int{0}
}

type RequestType int32

const (
	RequestType_RT_KEYGEN     RequestType = 0
	RequestType_RT_SIGN       RequestType = 1
	RequestType_RT_PROPOSAL   RequestType = 2
	RequestType_RT_ACCEPTANCE RequestType = 3
	RequestType_RT_SIGN_START RequestType = 4
)

// Enum value maps for RequestType.
var (
	RequestType_name = map[int32]string{
		0: "RT_KEYGEN",
		1: "RT_SIGN",
		2: "RT_PROPOSAL",
		3: "RT_ACCEPTANCE",
		4: "RT_SIGN_START",
	}
	RequestType_value = map[string]int32{
		"RT_KEYGEN":     0,
		"RT_SIGN":       1,
		"RT_PROPOSAL":   2,
		"RT_ACCEPTANCE": 3,
		"RT_SIGN_START": 4,
	}
)

func (x RequestType) Enum() *RequestType {
	p := new(RequestType)
	*p = x
	return p
}

func (x RequestType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RequestType) Descriptor() protoreflect.EnumDescriptor {
	return file_p2p_server_proto_enumTypes[1].Descriptor()
}

func (RequestType) Type() protoreflect.EnumType {
	return &file_p2p_server_proto_enumTypes[1]
}

func (x RequestType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RequestType.Descriptor instead.
func (RequestType) EnumDescriptor() ([]byte, []int) {
	return file_p2p_server_proto_rawDescGZIP(), []int{1}
}

type StatusResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        PartyStatus            `protobuf:"varint,1,opt,name=status,proto3,enum=p2p.PartyStatus" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StatusResponse) Reset() {
	*x = StatusResponse{}
	mi := &file_p2p_server_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusResponse) ProtoMessage() {}

func (x *StatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_server_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusResponse.ProtoReflect.Descriptor instead.
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return file_p2p_server_proto_rawDescGZIP(), []int{0}
}

func (x *StatusResponse) GetStatus() PartyStatus {
	if x != nil {
		return x.Status
	}
	return PartyStatus_PS_UNKNOWN
}

type SubmitRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Sender        string                 `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	SessionId     string                 `protobuf:"bytes,2,opt,name=sessionId,proto3" json:"sessionId,omitempty"`
	Type          RequestType            `protobuf:"varint,3,opt,name=type,proto3,enum=p2p.RequestType" json:"type,omitempty"`
	Data          *anypb.Any             `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SubmitRequest) Reset() {
	*x = SubmitRequest{}
	mi := &file_p2p_server_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubmitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitRequest) ProtoMessage() {}

func (x *SubmitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_server_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitRequest.ProtoReflect.Descriptor instead.
func (*SubmitRequest) Descriptor() ([]byte, []int) {
	return file_p2p_server_proto_rawDescGZIP(), []int{1}
}

func (x *SubmitRequest) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *SubmitRequest) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *SubmitRequest) GetType() RequestType {
	if x != nil {
		return x.Type
	}
	return RequestType_RT_KEYGEN
}

func (x *SubmitRequest) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

type TssData struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Data          []byte                 `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	IsBroadcast   bool                   `protobuf:"varint,2,opt,name=isBroadcast,proto3" json:"isBroadcast,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TssData) Reset() {
	*x = TssData{}
	mi := &file_p2p_server_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TssData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TssData) ProtoMessage() {}

func (x *TssData) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_server_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TssData.ProtoReflect.Descriptor instead.
func (*TssData) Descriptor() ([]byte, []int) {
	return file_p2p_server_proto_rawDescGZIP(), []int{2}
}

func (x *TssData) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *TssData) GetIsBroadcast() bool {
	if x != nil {
		return x.IsBroadcast
	}
	return false
}

type SignStartData struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Parties       []string               `protobuf:"bytes,1,rep,name=parties,proto3" json:"parties,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SignStartData) Reset() {
	*x = SignStartData{}
	mi := &file_p2p_server_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SignStartData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignStartData) ProtoMessage() {}

func (x *SignStartData) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_server_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignStartData.ProtoReflect.Descriptor instead.
func (*SignStartData) Descriptor() ([]byte, []int) {
	return file_p2p_server_proto_rawDescGZIP(), []int{3}
}

func (x *SignStartData) GetParties() []string {
	if x != nil {
		return x.Parties
	}
	return nil
}

type AcceptanceData struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Accepted      bool                   `protobuf:"varint,1,opt,name=accepted,proto3" json:"accepted,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AcceptanceData) Reset() {
	*x = AcceptanceData{}
	mi := &file_p2p_server_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AcceptanceData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcceptanceData) ProtoMessage() {}

func (x *AcceptanceData) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_server_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcceptanceData.ProtoReflect.Descriptor instead.
func (*AcceptanceData) Descriptor() ([]byte, []int) {
	return file_p2p_server_proto_rawDescGZIP(), []int{4}
}

func (x *AcceptanceData) GetAccepted() bool {
	if x != nil {
		return x.Accepted
	}
	return false
}

var File_p2p_server_proto protoreflect.FileDescriptor

var file_p2p_server_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x32, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x03, 0x70, 0x32, 0x70, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0d, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3a,
	0x0a, 0x0e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x28, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x10, 0x2e, 0x70, 0x32, 0x70, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x79, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x95, 0x01, 0x0a, 0x0d, 0x53,
	0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x24, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x10, 0x2e, 0x70, 0x32, 0x70, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x3f, 0x0a, 0x07, 0x54, 0x73, 0x73, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x73, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63,
	0x61, 0x73, 0x74, 0x22, 0x29, 0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x72, 0x74, 0x69, 0x65, 0x73, 0x22, 0x2c,
	0x0a, 0x0e, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x1a, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x2a, 0x39, 0x0a, 0x0b,
	0x50, 0x61, 0x72, 0x74, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0e, 0x0a, 0x0a, 0x50,
	0x53, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x50,
	0x53, 0x5f, 0x4b, 0x45, 0x59, 0x47, 0x45, 0x4e, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x53,
	0x5f, 0x53, 0x49, 0x47, 0x4e, 0x10, 0x02, 0x2a, 0x60, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x52, 0x54, 0x5f, 0x4b, 0x45, 0x59,
	0x47, 0x45, 0x4e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x54, 0x5f, 0x53, 0x49, 0x47, 0x4e,
	0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x52, 0x54, 0x5f, 0x50, 0x52, 0x4f, 0x50, 0x4f, 0x53, 0x41,
	0x4c, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x52, 0x54, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x50, 0x54,
	0x41, 0x4e, 0x43, 0x45, 0x10, 0x03, 0x12, 0x11, 0x0a, 0x0d, 0x52, 0x54, 0x5f, 0x53, 0x49, 0x47,
	0x4e, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x10, 0x04, 0x32, 0x76, 0x0a, 0x03, 0x50, 0x32, 0x50,
	0x12, 0x37, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x13, 0x2e, 0x70, 0x32, 0x70, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x06, 0x53, 0x75, 0x62,
	0x6d, 0x69, 0x74, 0x12, 0x12, 0x2e, 0x70, 0x32, 0x70, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x68, 0x79, 0x6c, 0x65, 0x2d, 0x74, 0x65, 0x61, 0x6d, 0x2f, 0x74, 0x73, 0x73, 0x2d, 0x73, 0x76,
	0x63, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x32, 0x70, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_p2p_server_proto_rawDescOnce sync.Once
	file_p2p_server_proto_rawDescData = file_p2p_server_proto_rawDesc
)

func file_p2p_server_proto_rawDescGZIP() []byte {
	file_p2p_server_proto_rawDescOnce.Do(func() {
		file_p2p_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_p2p_server_proto_rawDescData)
	})
	return file_p2p_server_proto_rawDescData
}

var file_p2p_server_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_p2p_server_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_p2p_server_proto_goTypes = []any{
	(PartyStatus)(0),       // 0: p2p.PartyStatus
	(RequestType)(0),       // 1: p2p.RequestType
	(*StatusResponse)(nil), // 2: p2p.StatusResponse
	(*SubmitRequest)(nil),  // 3: p2p.SubmitRequest
	(*TssData)(nil),        // 4: p2p.TssData
	(*SignStartData)(nil),  // 5: p2p.SignStartData
	(*AcceptanceData)(nil), // 6: p2p.AcceptanceData
	(*anypb.Any)(nil),      // 7: google.protobuf.Any
	(*emptypb.Empty)(nil),  // 8: google.protobuf.Empty
}
var file_p2p_server_proto_depIdxs = []int32{
	0, // 0: p2p.StatusResponse.status:type_name -> p2p.PartyStatus
	1, // 1: p2p.SubmitRequest.type:type_name -> p2p.RequestType
	7, // 2: p2p.SubmitRequest.data:type_name -> google.protobuf.Any
	8, // 3: p2p.P2P.Status:input_type -> google.protobuf.Empty
	3, // 4: p2p.P2P.Submit:input_type -> p2p.SubmitRequest
	2, // 5: p2p.P2P.Status:output_type -> p2p.StatusResponse
	8, // 6: p2p.P2P.Submit:output_type -> google.protobuf.Empty
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_p2p_server_proto_init() }
func file_p2p_server_proto_init() {
	if File_p2p_server_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_p2p_server_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_p2p_server_proto_goTypes,
		DependencyIndexes: file_p2p_server_proto_depIdxs,
		EnumInfos:         file_p2p_server_proto_enumTypes,
		MessageInfos:      file_p2p_server_proto_msgTypes,
	}.Build()
	File_p2p_server_proto = out.File
	file_p2p_server_proto_rawDesc = nil
	file_p2p_server_proto_goTypes = nil
	file_p2p_server_proto_depIdxs = nil
}
