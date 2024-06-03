// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.21.12
// source: tapchannelrpc/tapchannel.proto

package tapchannelrpc

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

type FundChannelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The asset amount to fund the channel with. The BTC amount is fixed and
	// cannot be customized (for now).
	AssetAmount uint64 `protobuf:"varint,1,opt,name=asset_amount,json=assetAmount,proto3" json:"asset_amount,omitempty"`
	// The asset ID to use for the channel funding.
	AssetId []byte `protobuf:"bytes,2,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty"`
	// The public key of the peer to open the channel with. Must already be
	// connected to this peer.
	PeerPubkey []byte `protobuf:"bytes,3,opt,name=peer_pubkey,json=peerPubkey,proto3" json:"peer_pubkey,omitempty"`
	// The channel funding fee rate in sat/vByte.
	FeeRateSatPerVbyte uint32 `protobuf:"varint,4,opt,name=fee_rate_sat_per_vbyte,json=feeRateSatPerVbyte,proto3" json:"fee_rate_sat_per_vbyte,omitempty"`
}

func (x *FundChannelRequest) Reset() {
	*x = FundChannelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tapchannelrpc_tapchannel_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FundChannelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FundChannelRequest) ProtoMessage() {}

func (x *FundChannelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tapchannelrpc_tapchannel_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FundChannelRequest.ProtoReflect.Descriptor instead.
func (*FundChannelRequest) Descriptor() ([]byte, []int) {
	return file_tapchannelrpc_tapchannel_proto_rawDescGZIP(), []int{0}
}

func (x *FundChannelRequest) GetAssetAmount() uint64 {
	if x != nil {
		return x.AssetAmount
	}
	return 0
}

func (x *FundChannelRequest) GetAssetId() []byte {
	if x != nil {
		return x.AssetId
	}
	return nil
}

func (x *FundChannelRequest) GetPeerPubkey() []byte {
	if x != nil {
		return x.PeerPubkey
	}
	return nil
}

func (x *FundChannelRequest) GetFeeRateSatPerVbyte() uint32 {
	if x != nil {
		return x.FeeRateSatPerVbyte
	}
	return 0
}

type FundChannelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The channel funding transaction ID.
	Txid string `protobuf:"bytes,1,opt,name=txid,proto3" json:"txid,omitempty"`
	// The index of the channel funding output in the funding transaction.
	OutputIndex int32 `protobuf:"varint,2,opt,name=output_index,json=outputIndex,proto3" json:"output_index,omitempty"`
}

func (x *FundChannelResponse) Reset() {
	*x = FundChannelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tapchannelrpc_tapchannel_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FundChannelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FundChannelResponse) ProtoMessage() {}

func (x *FundChannelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tapchannelrpc_tapchannel_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FundChannelResponse.ProtoReflect.Descriptor instead.
func (*FundChannelResponse) Descriptor() ([]byte, []int) {
	return file_tapchannelrpc_tapchannel_proto_rawDescGZIP(), []int{1}
}

func (x *FundChannelResponse) GetTxid() string {
	if x != nil {
		return x.Txid
	}
	return ""
}

func (x *FundChannelResponse) GetOutputIndex() int32 {
	if x != nil {
		return x.OutputIndex
	}
	return 0
}

type RouterSendPaymentData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The string encoded asset ID to amount mapping. Instructs the router to
	// use these assets in the given amounts for the payment. Can be empty for
	// a payment of an invoice, if the RFQ ID is set instead.
	AssetAmounts map[string]uint64 `protobuf:"bytes,1,rep,name=asset_amounts,json=assetAmounts,proto3" json:"asset_amounts,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	// The RFQ ID to use for the payment. Can be empty for a direct keysend
	// payment that doesn't involve any conversion (and thus no RFQ).
	RfqId []byte `protobuf:"bytes,2,opt,name=rfq_id,json=rfqId,proto3" json:"rfq_id,omitempty"`
}

func (x *RouterSendPaymentData) Reset() {
	*x = RouterSendPaymentData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tapchannelrpc_tapchannel_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouterSendPaymentData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouterSendPaymentData) ProtoMessage() {}

func (x *RouterSendPaymentData) ProtoReflect() protoreflect.Message {
	mi := &file_tapchannelrpc_tapchannel_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouterSendPaymentData.ProtoReflect.Descriptor instead.
func (*RouterSendPaymentData) Descriptor() ([]byte, []int) {
	return file_tapchannelrpc_tapchannel_proto_rawDescGZIP(), []int{2}
}

func (x *RouterSendPaymentData) GetAssetAmounts() map[string]uint64 {
	if x != nil {
		return x.AssetAmounts
	}
	return nil
}

func (x *RouterSendPaymentData) GetRfqId() []byte {
	if x != nil {
		return x.RfqId
	}
	return nil
}

type EncodeCustomRecordsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Input:
	//
	//	*EncodeCustomRecordsRequest_RouterSendPayment
	Input isEncodeCustomRecordsRequest_Input `protobuf_oneof:"input"`
}

func (x *EncodeCustomRecordsRequest) Reset() {
	*x = EncodeCustomRecordsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tapchannelrpc_tapchannel_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncodeCustomRecordsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncodeCustomRecordsRequest) ProtoMessage() {}

func (x *EncodeCustomRecordsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tapchannelrpc_tapchannel_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncodeCustomRecordsRequest.ProtoReflect.Descriptor instead.
func (*EncodeCustomRecordsRequest) Descriptor() ([]byte, []int) {
	return file_tapchannelrpc_tapchannel_proto_rawDescGZIP(), []int{3}
}

func (m *EncodeCustomRecordsRequest) GetInput() isEncodeCustomRecordsRequest_Input {
	if m != nil {
		return m.Input
	}
	return nil
}

func (x *EncodeCustomRecordsRequest) GetRouterSendPayment() *RouterSendPaymentData {
	if x, ok := x.GetInput().(*EncodeCustomRecordsRequest_RouterSendPayment); ok {
		return x.RouterSendPayment
	}
	return nil
}

type isEncodeCustomRecordsRequest_Input interface {
	isEncodeCustomRecordsRequest_Input()
}

type EncodeCustomRecordsRequest_RouterSendPayment struct {
	RouterSendPayment *RouterSendPaymentData `protobuf:"bytes,1,opt,name=router_send_payment,json=routerSendPayment,proto3,oneof"`
}

func (*EncodeCustomRecordsRequest_RouterSendPayment) isEncodeCustomRecordsRequest_Input() {}

type EncodeCustomRecordsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The encoded custom records in TLV format.
	CustomRecords map[uint64][]byte `protobuf:"bytes,1,rep,name=custom_records,json=customRecords,proto3" json:"custom_records,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *EncodeCustomRecordsResponse) Reset() {
	*x = EncodeCustomRecordsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tapchannelrpc_tapchannel_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncodeCustomRecordsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncodeCustomRecordsResponse) ProtoMessage() {}

func (x *EncodeCustomRecordsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tapchannelrpc_tapchannel_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncodeCustomRecordsResponse.ProtoReflect.Descriptor instead.
func (*EncodeCustomRecordsResponse) Descriptor() ([]byte, []int) {
	return file_tapchannelrpc_tapchannel_proto_rawDescGZIP(), []int{4}
}

func (x *EncodeCustomRecordsResponse) GetCustomRecords() map[uint64][]byte {
	if x != nil {
		return x.CustomRecords
	}
	return nil
}

var File_tapchannelrpc_tapchannel_proto protoreflect.FileDescriptor

var file_tapchannelrpc_tapchannel_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x74, 0x61, 0x70, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x72, 0x70, 0x63, 0x2f,
	0x74, 0x61, 0x70, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0d, 0x74, 0x61, 0x70, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x72, 0x70, 0x63, 0x22,
	0xa7, 0x01, 0x0a, 0x12, 0x46, 0x75, 0x6e, 0x64, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x61, 0x73,
	0x73, 0x65, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x70, 0x75, 0x62,
	0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x70, 0x65, 0x65, 0x72, 0x50,
	0x75, 0x62, 0x6b, 0x65, 0x79, 0x12, 0x32, 0x0a, 0x16, 0x66, 0x65, 0x65, 0x5f, 0x72, 0x61, 0x74,
	0x65, 0x5f, 0x73, 0x61, 0x74, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x76, 0x62, 0x79, 0x74, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x66, 0x65, 0x65, 0x52, 0x61, 0x74, 0x65, 0x53, 0x61,
	0x74, 0x50, 0x65, 0x72, 0x56, 0x62, 0x79, 0x74, 0x65, 0x22, 0x4c, 0x0a, 0x13, 0x46, 0x75, 0x6e,
	0x64, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x78, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x78, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0xcc, 0x01, 0x0a, 0x15, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x72, 0x53, 0x65, 0x6e, 0x64, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x5b, 0x0a, 0x0d, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x74, 0x61, 0x70, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x53,
	0x65, 0x6e, 0x64, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x41,
	0x73, 0x73, 0x65, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x0c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x15,
	0x0a, 0x06, 0x72, 0x66, 0x71, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05,
	0x72, 0x66, 0x71, 0x49, 0x64, 0x1a, 0x3f, 0x0a, 0x11, 0x41, 0x73, 0x73, 0x65, 0x74, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x7d, 0x0a, 0x1a, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x65,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x56, 0x0a, 0x13, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x5f, 0x73,
	0x65, 0x6e, 0x64, 0x5f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x24, 0x2e, 0x74, 0x61, 0x70, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x72, 0x70,
	0x63, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x53, 0x65, 0x6e, 0x64, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x11, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x72, 0x53, 0x65, 0x6e, 0x64, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x07, 0x0a, 0x05,
	0x69, 0x6e, 0x70, 0x75, 0x74, 0x22, 0xc5, 0x01, 0x0a, 0x1b, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x65,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x64, 0x0a, 0x0e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f,
	0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3d, 0x2e,
	0x74, 0x61, 0x70, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x6e,
	0x63, 0x6f, 0x64, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x1a, 0x40, 0x0a, 0x12, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0xda, 0x01,
	0x0a, 0x14, 0x54, 0x61, 0x70, 0x72, 0x6f, 0x6f, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x12, 0x54, 0x0a, 0x0b, 0x46, 0x75, 0x6e, 0x64, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x21, 0x2e, 0x74, 0x61, 0x70, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x72, 0x70, 0x63, 0x2e, 0x46, 0x75, 0x6e, 0x64, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x74, 0x61, 0x70, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x72, 0x70, 0x63, 0x2e, 0x46, 0x75, 0x6e, 0x64, 0x43, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6c, 0x0a, 0x13,
	0x45, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x73, 0x12, 0x29, 0x2e, 0x74, 0x61, 0x70, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x72, 0x70, 0x63, 0x2e, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a,
	0x2e, 0x74, 0x61, 0x70, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x72, 0x70, 0x63, 0x2e, 0x45,
	0x6e, 0x63, 0x6f, 0x64, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x6e, 0x69,
	0x6e, 0x67, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x6f, 0x74, 0x2d, 0x61,
	0x73, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x74, 0x61, 0x70, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x61, 0x70,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_tapchannelrpc_tapchannel_proto_rawDescOnce sync.Once
	file_tapchannelrpc_tapchannel_proto_rawDescData = file_tapchannelrpc_tapchannel_proto_rawDesc
)

func file_tapchannelrpc_tapchannel_proto_rawDescGZIP() []byte {
	file_tapchannelrpc_tapchannel_proto_rawDescOnce.Do(func() {
		file_tapchannelrpc_tapchannel_proto_rawDescData = protoimpl.X.CompressGZIP(file_tapchannelrpc_tapchannel_proto_rawDescData)
	})
	return file_tapchannelrpc_tapchannel_proto_rawDescData
}

var file_tapchannelrpc_tapchannel_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_tapchannelrpc_tapchannel_proto_goTypes = []interface{}{
	(*FundChannelRequest)(nil),          // 0: tapchannelrpc.FundChannelRequest
	(*FundChannelResponse)(nil),         // 1: tapchannelrpc.FundChannelResponse
	(*RouterSendPaymentData)(nil),       // 2: tapchannelrpc.RouterSendPaymentData
	(*EncodeCustomRecordsRequest)(nil),  // 3: tapchannelrpc.EncodeCustomRecordsRequest
	(*EncodeCustomRecordsResponse)(nil), // 4: tapchannelrpc.EncodeCustomRecordsResponse
	nil,                                 // 5: tapchannelrpc.RouterSendPaymentData.AssetAmountsEntry
	nil,                                 // 6: tapchannelrpc.EncodeCustomRecordsResponse.CustomRecordsEntry
}
var file_tapchannelrpc_tapchannel_proto_depIdxs = []int32{
	5, // 0: tapchannelrpc.RouterSendPaymentData.asset_amounts:type_name -> tapchannelrpc.RouterSendPaymentData.AssetAmountsEntry
	2, // 1: tapchannelrpc.EncodeCustomRecordsRequest.router_send_payment:type_name -> tapchannelrpc.RouterSendPaymentData
	6, // 2: tapchannelrpc.EncodeCustomRecordsResponse.custom_records:type_name -> tapchannelrpc.EncodeCustomRecordsResponse.CustomRecordsEntry
	0, // 3: tapchannelrpc.TaprootAssetChannels.FundChannel:input_type -> tapchannelrpc.FundChannelRequest
	3, // 4: tapchannelrpc.TaprootAssetChannels.EncodeCustomRecords:input_type -> tapchannelrpc.EncodeCustomRecordsRequest
	1, // 5: tapchannelrpc.TaprootAssetChannels.FundChannel:output_type -> tapchannelrpc.FundChannelResponse
	4, // 6: tapchannelrpc.TaprootAssetChannels.EncodeCustomRecords:output_type -> tapchannelrpc.EncodeCustomRecordsResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_tapchannelrpc_tapchannel_proto_init() }
func file_tapchannelrpc_tapchannel_proto_init() {
	if File_tapchannelrpc_tapchannel_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tapchannelrpc_tapchannel_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FundChannelRequest); i {
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
		file_tapchannelrpc_tapchannel_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FundChannelResponse); i {
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
		file_tapchannelrpc_tapchannel_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouterSendPaymentData); i {
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
		file_tapchannelrpc_tapchannel_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncodeCustomRecordsRequest); i {
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
		file_tapchannelrpc_tapchannel_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncodeCustomRecordsResponse); i {
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
	file_tapchannelrpc_tapchannel_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*EncodeCustomRecordsRequest_RouterSendPayment)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tapchannelrpc_tapchannel_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tapchannelrpc_tapchannel_proto_goTypes,
		DependencyIndexes: file_tapchannelrpc_tapchannel_proto_depIdxs,
		MessageInfos:      file_tapchannelrpc_tapchannel_proto_msgTypes,
	}.Build()
	File_tapchannelrpc_tapchannel_proto = out.File
	file_tapchannelrpc_tapchannel_proto_rawDesc = nil
	file_tapchannelrpc_tapchannel_proto_goTypes = nil
	file_tapchannelrpc_tapchannel_proto_depIdxs = nil
}
