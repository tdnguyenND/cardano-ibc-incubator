// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: ibc/applications/transfer/v1/tx.proto

package types

import (
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	types1 "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
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

// MsgTransfer defines a msg to transfer fungible tokens (i.e Coins) between
// ICS20 enabled chains. See ICS Spec here:
// https://github.com/cosmos/ibc/tree/master/spec/app/ics-020-fungible-token-transfer#data-structures
type MsgTransfer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the port on which the packet will be sent
	SourcePort string `protobuf:"bytes,1,opt,name=source_port,json=sourcePort,proto3" json:"source_port,omitempty"`
	// the channel by which the packet will be sent
	SourceChannel string `protobuf:"bytes,2,opt,name=source_channel,json=sourceChannel,proto3" json:"source_channel,omitempty"`
	// the tokens to be transferred
	Token *types.Coin `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	// the sender address
	Sender string `protobuf:"bytes,4,opt,name=sender,proto3" json:"sender,omitempty"`
	// the recipient address on the destination chain
	Receiver string `protobuf:"bytes,5,opt,name=receiver,proto3" json:"receiver,omitempty"`
	// Timeout height relative to the current block height.
	// The timeout is disabled when set to 0.
	TimeoutHeight *types1.Height `protobuf:"bytes,6,opt,name=timeout_height,json=timeoutHeight,proto3" json:"timeout_height,omitempty"`
	// Timeout timestamp in absolute nanoseconds since unix epoch.
	// The timeout is disabled when set to 0.
	TimeoutTimestamp uint64 `protobuf:"varint,7,opt,name=timeout_timestamp,json=timeoutTimestamp,proto3" json:"timeout_timestamp,omitempty"`
	// optional memo
	Memo string `protobuf:"bytes,8,opt,name=memo,proto3" json:"memo,omitempty"`
}

func (x *MsgTransfer) Reset() {
	*x = MsgTransfer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ibc_applications_transfer_v1_tx_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgTransfer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgTransfer) ProtoMessage() {}

func (x *MsgTransfer) ProtoReflect() protoreflect.Message {
	mi := &file_ibc_applications_transfer_v1_tx_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgTransfer.ProtoReflect.Descriptor instead.
func (*MsgTransfer) Descriptor() ([]byte, []int) {
	return file_ibc_applications_transfer_v1_tx_proto_rawDescGZIP(), []int{0}
}

func (x *MsgTransfer) GetSourcePort() string {
	if x != nil {
		return x.SourcePort
	}
	return ""
}

func (x *MsgTransfer) GetSourceChannel() string {
	if x != nil {
		return x.SourceChannel
	}
	return ""
}

func (x *MsgTransfer) GetToken() *types.Coin {
	if x != nil {
		return x.Token
	}
	return nil
}

func (x *MsgTransfer) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *MsgTransfer) GetReceiver() string {
	if x != nil {
		return x.Receiver
	}
	return ""
}

func (x *MsgTransfer) GetTimeoutHeight() *types1.Height {
	if x != nil {
		return x.TimeoutHeight
	}
	return nil
}

func (x *MsgTransfer) GetTimeoutTimestamp() uint64 {
	if x != nil {
		return x.TimeoutTimestamp
	}
	return 0
}

func (x *MsgTransfer) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

// MsgTransferResponse defines the Msg/Transfer response type.
type MsgTransferResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// sequence number of the transfer packet sent
	Sequence uint64 `protobuf:"varint,1,opt,name=sequence,proto3" json:"sequence,omitempty"`
}

func (x *MsgTransferResponse) Reset() {
	*x = MsgTransferResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ibc_applications_transfer_v1_tx_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgTransferResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgTransferResponse) ProtoMessage() {}

func (x *MsgTransferResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ibc_applications_transfer_v1_tx_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgTransferResponse.ProtoReflect.Descriptor instead.
func (*MsgTransferResponse) Descriptor() ([]byte, []int) {
	return file_ibc_applications_transfer_v1_tx_proto_rawDescGZIP(), []int{1}
}

func (x *MsgTransferResponse) GetSequence() uint64 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

var File_ibc_applications_transfer_v1_tx_proto protoreflect.FileDescriptor

var file_ibc_applications_transfer_v1_tx_proto_rawDesc = []byte{
	0x0a, 0x25, 0x69, 0x62, 0x63, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x74,
	0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x69, 0x62, 0x63, 0x2e, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x69, 0x62, 0x63,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbe, 0x03, 0x0a,
	0x0b, 0x4d, 0x73, 0x67, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x37, 0x0a, 0x0b,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x16, 0xf2, 0xde, 0x1f, 0x12, 0x79, 0x61, 0x6d, 0x6c, 0x3a, 0x22, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x40, 0x0a, 0x0e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x19, 0xf2,
	0xde, 0x1f, 0x15, 0x79, 0x61, 0x6d, 0x6c, 0x3a, 0x22, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x22, 0x52, 0x0d, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x35, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x69,
	0x6e, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x72, 0x12, 0x60, 0x0a, 0x0e, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x5f, 0x68, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x69, 0x62, 0x63,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x42, 0x1d, 0xc8, 0xde, 0x1f, 0x00, 0xf2, 0xde, 0x1f, 0x15,
	0x79, 0x61, 0x6d, 0x6c, 0x3a, 0x22, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x5f, 0x68, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x22, 0x52, 0x0d, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x48, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x12, 0x49, 0x0a, 0x11, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x42,
	0x1c, 0xf2, 0xde, 0x1f, 0x18, 0x79, 0x61, 0x6d, 0x6c, 0x3a, 0x22, 0x74, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x52, 0x10, 0x74,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12,
	0x12, 0x0a, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d,
	0x65, 0x6d, 0x6f, 0x3a, 0x08, 0x88, 0xa0, 0x1f, 0x00, 0xe8, 0xa0, 0x1f, 0x00, 0x22, 0x31, 0x0a,
	0x13, 0x4d, 0x73, 0x67, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65,
	0x32, 0x6f, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x12, 0x68, 0x0a, 0x08, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x12, 0x29, 0x2e, 0x69, 0x62, 0x63, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x1a, 0x31,
	0x2e, 0x69, 0x62, 0x63, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x73,
	0x67, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x69, 0x62, 0x63, 0x2d, 0x67, 0x6f, 0x2f, 0x76, 0x37,
	0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ibc_applications_transfer_v1_tx_proto_rawDescOnce sync.Once
	file_ibc_applications_transfer_v1_tx_proto_rawDescData = file_ibc_applications_transfer_v1_tx_proto_rawDesc
)

func file_ibc_applications_transfer_v1_tx_proto_rawDescGZIP() []byte {
	file_ibc_applications_transfer_v1_tx_proto_rawDescOnce.Do(func() {
		file_ibc_applications_transfer_v1_tx_proto_rawDescData = protoimpl.X.CompressGZIP(file_ibc_applications_transfer_v1_tx_proto_rawDescData)
	})
	return file_ibc_applications_transfer_v1_tx_proto_rawDescData
}

var file_ibc_applications_transfer_v1_tx_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ibc_applications_transfer_v1_tx_proto_goTypes = []interface{}{
	(*MsgTransfer)(nil),         // 0: ibc.applications.transfer.v1.MsgTransfer
	(*MsgTransferResponse)(nil), // 1: ibc.applications.transfer.v1.MsgTransferResponse
	(*types.Coin)(nil),          // 2: cosmos.base.v1beta1.Coin
	(*types1.Height)(nil),       // 3: ibc.core.client.v1.Height
}
var file_ibc_applications_transfer_v1_tx_proto_depIdxs = []int32{
	2, // 0: ibc.applications.transfer.v1.MsgTransfer.token:type_name -> cosmos.base.v1beta1.Coin
	3, // 1: ibc.applications.transfer.v1.MsgTransfer.timeout_height:type_name -> ibc.core.client.v1.Height
	0, // 2: ibc.applications.transfer.v1.Msg.Transfer:input_type -> ibc.applications.transfer.v1.MsgTransfer
	1, // 3: ibc.applications.transfer.v1.Msg.Transfer:output_type -> ibc.applications.transfer.v1.MsgTransferResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ibc_applications_transfer_v1_tx_proto_init() }
func file_ibc_applications_transfer_v1_tx_proto_init() {
	if File_ibc_applications_transfer_v1_tx_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ibc_applications_transfer_v1_tx_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgTransfer); i {
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
		file_ibc_applications_transfer_v1_tx_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgTransferResponse); i {
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
			RawDescriptor: file_ibc_applications_transfer_v1_tx_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ibc_applications_transfer_v1_tx_proto_goTypes,
		DependencyIndexes: file_ibc_applications_transfer_v1_tx_proto_depIdxs,
		MessageInfos:      file_ibc_applications_transfer_v1_tx_proto_msgTypes,
	}.Build()
	File_ibc_applications_transfer_v1_tx_proto = out.File
	file_ibc_applications_transfer_v1_tx_proto_rawDesc = nil
	file_ibc_applications_transfer_v1_tx_proto_goTypes = nil
	file_ibc_applications_transfer_v1_tx_proto_depIdxs = nil
}
