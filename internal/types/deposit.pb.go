// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        (unknown)
// source: deposit.proto

package types

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

type WithdrawalStatus int32

const (
	WithdrawalStatus_WITHDRAWAL_STATUS_UNSPECIFIED WithdrawalStatus = 0
	WithdrawalStatus_WITHDRAWAL_STATUS_PENDING     WithdrawalStatus = 1
	WithdrawalStatus_WITHDRAWAL_STATUS_PROCESSING  WithdrawalStatus = 2
	WithdrawalStatus_WITHDRAWAL_STATUS_PROCESSED   WithdrawalStatus = 3
	WithdrawalStatus_WITHDRAWAL_STATUS_FAILED      WithdrawalStatus = 4
	WithdrawalStatus_WITHDRAWAL_STATUS_INVALID     WithdrawalStatus = 5
)

// Enum value maps for WithdrawalStatus.
var (
	WithdrawalStatus_name = map[int32]string{
		0: "WITHDRAWAL_STATUS_UNSPECIFIED",
		1: "WITHDRAWAL_STATUS_PENDING",
		2: "WITHDRAWAL_STATUS_PROCESSING",
		3: "WITHDRAWAL_STATUS_PROCESSED",
		4: "WITHDRAWAL_STATUS_FAILED",
		5: "WITHDRAWAL_STATUS_INVALID",
	}
	WithdrawalStatus_value = map[string]int32{
		"WITHDRAWAL_STATUS_UNSPECIFIED": 0,
		"WITHDRAWAL_STATUS_PENDING":     1,
		"WITHDRAWAL_STATUS_PROCESSING":  2,
		"WITHDRAWAL_STATUS_PROCESSED":   3,
		"WITHDRAWAL_STATUS_FAILED":      4,
		"WITHDRAWAL_STATUS_INVALID":     5,
	}
)

func (x WithdrawalStatus) Enum() *WithdrawalStatus {
	p := new(WithdrawalStatus)
	*p = x
	return p
}

func (x WithdrawalStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WithdrawalStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_deposit_proto_enumTypes[0].Descriptor()
}

func (WithdrawalStatus) Type() protoreflect.EnumType {
	return &file_deposit_proto_enumTypes[0]
}

func (x WithdrawalStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WithdrawalStatus.Descriptor instead.
func (WithdrawalStatus) EnumDescriptor() ([]byte, []int) {
	return file_deposit_proto_rawDescGZIP(), []int{0}
}

type DepositIdentifier struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TxHash        string                 `protobuf:"bytes,1,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty"`
	TxNonce       int32                  `protobuf:"varint,2,opt,name=tx_nonce,json=txNonce,proto3" json:"tx_nonce,omitempty"`
	ChainId       string                 `protobuf:"bytes,3,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DepositIdentifier) Reset() {
	*x = DepositIdentifier{}
	mi := &file_deposit_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DepositIdentifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DepositIdentifier) ProtoMessage() {}

func (x *DepositIdentifier) ProtoReflect() protoreflect.Message {
	mi := &file_deposit_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DepositIdentifier.ProtoReflect.Descriptor instead.
func (*DepositIdentifier) Descriptor() ([]byte, []int) {
	return file_deposit_proto_rawDescGZIP(), []int{0}
}

func (x *DepositIdentifier) GetTxHash() string {
	if x != nil {
		return x.TxHash
	}
	return ""
}

func (x *DepositIdentifier) GetTxNonce() int32 {
	if x != nil {
		return x.TxNonce
	}
	return 0
}

func (x *DepositIdentifier) GetChainId() string {
	if x != nil {
		return x.ChainId
	}
	return ""
}

type WithdrawalIdentifier struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TxHash        string                 `protobuf:"bytes,1,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty"`
	ChainId       string                 `protobuf:"bytes,2,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WithdrawalIdentifier) Reset() {
	*x = WithdrawalIdentifier{}
	mi := &file_deposit_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WithdrawalIdentifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WithdrawalIdentifier) ProtoMessage() {}

func (x *WithdrawalIdentifier) ProtoReflect() protoreflect.Message {
	mi := &file_deposit_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WithdrawalIdentifier.ProtoReflect.Descriptor instead.
func (*WithdrawalIdentifier) Descriptor() ([]byte, []int) {
	return file_deposit_proto_rawDescGZIP(), []int{1}
}

func (x *WithdrawalIdentifier) GetTxHash() string {
	if x != nil {
		return x.TxHash
	}
	return ""
}

func (x *WithdrawalIdentifier) GetChainId() string {
	if x != nil {
		return x.ChainId
	}
	return ""
}

type TransferData struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	Sender           *string                `protobuf:"bytes,1,opt,name=sender,proto3,oneof" json:"sender,omitempty"`
	Receiver         string                 `protobuf:"bytes,2,opt,name=receiver,proto3" json:"receiver,omitempty"`
	DepositAmount    string                 `protobuf:"bytes,3,opt,name=deposit_amount,json=depositAmount,proto3" json:"deposit_amount,omitempty"`
	WithdrawalAmount string                 `protobuf:"bytes,4,opt,name=withdrawal_amount,json=withdrawalAmount,proto3" json:"withdrawal_amount,omitempty"`
	DepositAsset     string                 `protobuf:"bytes,5,opt,name=deposit_asset,json=depositAsset,proto3" json:"deposit_asset,omitempty"`
	WithdrawalAsset  string                 `protobuf:"bytes,6,opt,name=withdrawal_asset,json=withdrawalAsset,proto3" json:"withdrawal_asset,omitempty"`
	IsWrappedAsset   bool                   `protobuf:"varint,7,opt,name=is_wrapped_asset,json=isWrappedAsset,proto3" json:"is_wrapped_asset,omitempty"`
	DepositBlock     int64                  `protobuf:"varint,8,opt,name=deposit_block,json=depositBlock,proto3" json:"deposit_block,omitempty"`
	// used for EVM transfers
	Signature     *string `protobuf:"bytes,9,opt,name=signature,proto3,oneof" json:"signature,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TransferData) Reset() {
	*x = TransferData{}
	mi := &file_deposit_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TransferData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferData) ProtoMessage() {}

func (x *TransferData) ProtoReflect() protoreflect.Message {
	mi := &file_deposit_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferData.ProtoReflect.Descriptor instead.
func (*TransferData) Descriptor() ([]byte, []int) {
	return file_deposit_proto_rawDescGZIP(), []int{2}
}

func (x *TransferData) GetSender() string {
	if x != nil && x.Sender != nil {
		return *x.Sender
	}
	return ""
}

func (x *TransferData) GetReceiver() string {
	if x != nil {
		return x.Receiver
	}
	return ""
}

func (x *TransferData) GetDepositAmount() string {
	if x != nil {
		return x.DepositAmount
	}
	return ""
}

func (x *TransferData) GetWithdrawalAmount() string {
	if x != nil {
		return x.WithdrawalAmount
	}
	return ""
}

func (x *TransferData) GetDepositAsset() string {
	if x != nil {
		return x.DepositAsset
	}
	return ""
}

func (x *TransferData) GetWithdrawalAsset() string {
	if x != nil {
		return x.WithdrawalAsset
	}
	return ""
}

func (x *TransferData) GetIsWrappedAsset() bool {
	if x != nil {
		return x.IsWrappedAsset
	}
	return false
}

func (x *TransferData) GetDepositBlock() int64 {
	if x != nil {
		return x.DepositBlock
	}
	return 0
}

func (x *TransferData) GetSignature() string {
	if x != nil && x.Signature != nil {
		return *x.Signature
	}
	return ""
}

var File_deposit_proto protoreflect.FileDescriptor

var file_deposit_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x22, 0x62, 0x0a, 0x11, 0x44, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x17, 0x0a,
	0x07, 0x74, 0x78, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x74, 0x78, 0x48, 0x61, 0x73, 0x68, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x78, 0x5f, 0x6e, 0x6f, 0x6e,
	0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x74, 0x78, 0x4e, 0x6f, 0x6e, 0x63,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x22, 0x4a, 0x0a, 0x14,
	0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x61, 0x6c, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x78, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x78, 0x48, 0x61, 0x73, 0x68, 0x12, 0x19, 0x0a,
	0x08, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x22, 0xf6, 0x02, 0x0a, 0x0c, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x06, 0x73, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x73, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x72, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x5f, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x64, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2b, 0x0a, 0x11, 0x77, 0x69, 0x74,
	0x68, 0x64, 0x72, 0x61, 0x77, 0x61, 0x6c, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x61, 0x6c,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64,
	0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x77,
	0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x61, 0x6c, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x61,
	0x6c, 0x41, 0x73, 0x73, 0x65, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x69, 0x73, 0x5f, 0x77, 0x72, 0x61,
	0x70, 0x70, 0x65, 0x64, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0e, 0x69, 0x73, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x64, 0x41, 0x73, 0x73, 0x65, 0x74,
	0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x5f, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x21, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x2a, 0xd4, 0x01, 0x0a, 0x10, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x61, 0x6c,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x21, 0x0a, 0x1d, 0x57, 0x49, 0x54, 0x48, 0x44, 0x52,
	0x41, 0x57, 0x41, 0x4c, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1d, 0x0a, 0x19, 0x57, 0x49, 0x54,
	0x48, 0x44, 0x52, 0x41, 0x57, 0x41, 0x4c, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x50,
	0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x20, 0x0a, 0x1c, 0x57, 0x49, 0x54, 0x48,
	0x44, 0x52, 0x41, 0x57, 0x41, 0x4c, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x50, 0x52,
	0x4f, 0x43, 0x45, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x1f, 0x0a, 0x1b, 0x57, 0x49,
	0x54, 0x48, 0x44, 0x52, 0x41, 0x57, 0x41, 0x4c, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x45, 0x44, 0x10, 0x03, 0x12, 0x1c, 0x0a, 0x18, 0x57,
	0x49, 0x54, 0x48, 0x44, 0x52, 0x41, 0x57, 0x41, 0x4c, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x04, 0x12, 0x1d, 0x0a, 0x19, 0x57, 0x49, 0x54,
	0x48, 0x44, 0x52, 0x41, 0x57, 0x41, 0x4c, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x49,
	0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x05, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x79, 0x6c, 0x65, 0x2d, 0x74, 0x65, 0x61, 0x6d,
	0x2f, 0x74, 0x73, 0x73, 0x2d, 0x73, 0x76, 0x63, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_deposit_proto_rawDescOnce sync.Once
	file_deposit_proto_rawDescData = file_deposit_proto_rawDesc
)

func file_deposit_proto_rawDescGZIP() []byte {
	file_deposit_proto_rawDescOnce.Do(func() {
		file_deposit_proto_rawDescData = protoimpl.X.CompressGZIP(file_deposit_proto_rawDescData)
	})
	return file_deposit_proto_rawDescData
}

var file_deposit_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_deposit_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_deposit_proto_goTypes = []any{
	(WithdrawalStatus)(0),        // 0: deposit.WithdrawalStatus
	(*DepositIdentifier)(nil),    // 1: deposit.DepositIdentifier
	(*WithdrawalIdentifier)(nil), // 2: deposit.WithdrawalIdentifier
	(*TransferData)(nil),         // 3: deposit.TransferData
}
var file_deposit_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_deposit_proto_init() }
func file_deposit_proto_init() {
	if File_deposit_proto != nil {
		return
	}
	file_deposit_proto_msgTypes[2].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_deposit_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_deposit_proto_goTypes,
		DependencyIndexes: file_deposit_proto_depIdxs,
		EnumInfos:         file_deposit_proto_enumTypes,
		MessageInfos:      file_deposit_proto_msgTypes,
	}.Build()
	File_deposit_proto = out.File
	file_deposit_proto_rawDesc = nil
	file_deposit_proto_goTypes = nil
	file_deposit_proto_depIdxs = nil
}
