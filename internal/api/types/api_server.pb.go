// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        (unknown)
// source: api_server.proto

package types

import (
	types "github.com/hyle-team/tss-svc/internal/types"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/anypb"
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

type CheckWithdrawalResponse struct {
	state                protoimpl.MessageState      `protogen:"open.v1"`
	DepositIdentifier    *types.DepositIdentifier    `protobuf:"bytes,1,opt,name=deposit_identifier,json=depositIdentifier,proto3" json:"deposit_identifier,omitempty"`
	TransferData         *types.TransferData         `protobuf:"bytes,2,opt,name=transfer_data,json=transferData,proto3" json:"transfer_data,omitempty"`
	WithdrawalStatus     types.WithdrawalStatus      `protobuf:"varint,3,opt,name=withdrawal_status,json=withdrawalStatus,proto3,enum=deposit.WithdrawalStatus" json:"withdrawal_status,omitempty"`
	WithdrawalIdentifier *types.WithdrawalIdentifier `protobuf:"bytes,4,opt,name=withdrawal_identifier,json=withdrawalIdentifier,proto3,oneof" json:"withdrawal_identifier,omitempty"`
	unknownFields        protoimpl.UnknownFields
	sizeCache            protoimpl.SizeCache
}

func (x *CheckWithdrawalResponse) Reset() {
	*x = CheckWithdrawalResponse{}
	mi := &file_api_server_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckWithdrawalResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckWithdrawalResponse) ProtoMessage() {}

func (x *CheckWithdrawalResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_server_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckWithdrawalResponse.ProtoReflect.Descriptor instead.
func (*CheckWithdrawalResponse) Descriptor() ([]byte, []int) {
	return file_api_server_proto_rawDescGZIP(), []int{0}
}

func (x *CheckWithdrawalResponse) GetDepositIdentifier() *types.DepositIdentifier {
	if x != nil {
		return x.DepositIdentifier
	}
	return nil
}

func (x *CheckWithdrawalResponse) GetTransferData() *types.TransferData {
	if x != nil {
		return x.TransferData
	}
	return nil
}

func (x *CheckWithdrawalResponse) GetWithdrawalStatus() types.WithdrawalStatus {
	if x != nil {
		return x.WithdrawalStatus
	}
	return types.WithdrawalStatus(0)
}

func (x *CheckWithdrawalResponse) GetWithdrawalIdentifier() *types.WithdrawalIdentifier {
	if x != nil {
		return x.WithdrawalIdentifier
	}
	return nil
}

var File_api_server_proto protoreflect.FileDescriptor

var file_api_server_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x70, 0x69, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x64,
	0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdb, 0x02, 0x0a,
	0x17, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x61, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x12, 0x64, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x2e, 0x44,
	0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x52, 0x11, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x12, 0x3a, 0x0a, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x64, 0x65, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x46, 0x0a, 0x11, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x61, 0x6c, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x64, 0x65, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x2e, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x61, 0x6c, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x10, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x61,
	0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x57, 0x0a, 0x15, 0x77, 0x69, 0x74, 0x68, 0x64,
	0x72, 0x61, 0x77, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x2e, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x61, 0x6c, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x48, 0x00, 0x52, 0x14, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61,
	0x77, 0x61, 0x6c, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x88, 0x01, 0x01,
	0x42, 0x18, 0x0a, 0x16, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x61, 0x6c, 0x5f,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x32, 0xde, 0x01, 0x0a, 0x03, 0x41,
	0x50, 0x49, 0x12, 0x5a, 0x0a, 0x10, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x57, 0x69, 0x74, 0x68,
	0x64, 0x72, 0x61, 0x77, 0x61, 0x6c, 0x12, 0x1a, 0x2e, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x2e, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x0c, 0x3a, 0x01, 0x2a, 0x22, 0x07, 0x2f, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x12, 0x7b,
	0x0a, 0x0f, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x61,
	0x6c, 0x12, 0x1a, 0x2e, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x2e, 0x44, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x1a, 0x1c, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61,
	0x77, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2e, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x28, 0x12, 0x26, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2f, 0x7b, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x7b, 0x74, 0x78, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x7d,
	0x2f, 0x7b, 0x74, 0x78, 0x5f, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x7d, 0x42, 0x31, 0x5a, 0x2f, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x79, 0x6c, 0x65, 0x2d, 0x74,
	0x65, 0x61, 0x6d, 0x2f, 0x74, 0x73, 0x73, 0x2d, 0x73, 0x76, 0x63, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_server_proto_rawDescOnce sync.Once
	file_api_server_proto_rawDescData = file_api_server_proto_rawDesc
)

func file_api_server_proto_rawDescGZIP() []byte {
	file_api_server_proto_rawDescOnce.Do(func() {
		file_api_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_server_proto_rawDescData)
	})
	return file_api_server_proto_rawDescData
}

var file_api_server_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_server_proto_goTypes = []any{
	(*CheckWithdrawalResponse)(nil),    // 0: api.CheckWithdrawalResponse
	(*types.DepositIdentifier)(nil),    // 1: deposit.DepositIdentifier
	(*types.TransferData)(nil),         // 2: deposit.TransferData
	(types.WithdrawalStatus)(0),        // 3: deposit.WithdrawalStatus
	(*types.WithdrawalIdentifier)(nil), // 4: deposit.WithdrawalIdentifier
	(*emptypb.Empty)(nil),              // 5: google.protobuf.Empty
}
var file_api_server_proto_depIdxs = []int32{
	1, // 0: api.CheckWithdrawalResponse.deposit_identifier:type_name -> deposit.DepositIdentifier
	2, // 1: api.CheckWithdrawalResponse.transfer_data:type_name -> deposit.TransferData
	3, // 2: api.CheckWithdrawalResponse.withdrawal_status:type_name -> deposit.WithdrawalStatus
	4, // 3: api.CheckWithdrawalResponse.withdrawal_identifier:type_name -> deposit.WithdrawalIdentifier
	1, // 4: api.API.SubmitWithdrawal:input_type -> deposit.DepositIdentifier
	1, // 5: api.API.CheckWithdrawal:input_type -> deposit.DepositIdentifier
	5, // 6: api.API.SubmitWithdrawal:output_type -> google.protobuf.Empty
	0, // 7: api.API.CheckWithdrawal:output_type -> api.CheckWithdrawalResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_server_proto_init() }
func file_api_server_proto_init() {
	if File_api_server_proto != nil {
		return
	}
	file_api_server_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_server_proto_goTypes,
		DependencyIndexes: file_api_server_proto_depIdxs,
		MessageInfos:      file_api_server_proto_msgTypes,
	}.Build()
	File_api_server_proto = out.File
	file_api_server_proto_rawDesc = nil
	file_api_server_proto_goTypes = nil
	file_api_server_proto_depIdxs = nil
}
