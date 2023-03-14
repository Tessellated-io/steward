//
// Protos for function calls in the BaseAdaptor interface.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        v3.21.12
// source: adaptors/base.proto

package steward_proto

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

// Helper function that allows swaps using the Swap Router
//
// Represents function `swap(ERC20 assetIn, ERC20 assetOut, uint256 amountIn, SwapRouter.Exchange exchange, bytes memory params)`
type Swap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Asset to swap from
	AssetIn string `protobuf:"bytes,1,opt,name=asset_in,json=assetIn,proto3" json:"asset_in,omitempty"`
	// Asset to swap to
	AssetOut string `protobuf:"bytes,2,opt,name=asset_out,json=assetOut,proto3" json:"asset_out,omitempty"`
	// Amount to swap
	AmountIn string `protobuf:"bytes,3,opt,name=amount_in,json=amountIn,proto3" json:"amount_in,omitempty"`
	// The exchange to make the swap on
	Exchange Exchange `protobuf:"varint,4,opt,name=exchange,proto3,enum=steward.v2.Exchange" json:"exchange,omitempty"`
	// The parameters for the swap
	Params *SwapParams `protobuf:"bytes,5,opt,name=params,proto3" json:"params,omitempty"`
}

func (x *Swap) Reset() {
	*x = Swap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adaptors_base_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Swap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Swap) ProtoMessage() {}

func (x *Swap) ProtoReflect() protoreflect.Message {
	mi := &file_adaptors_base_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Swap.ProtoReflect.Descriptor instead.
func (*Swap) Descriptor() ([]byte, []int) {
	return file_adaptors_base_proto_rawDescGZIP(), []int{0}
}

func (x *Swap) GetAssetIn() string {
	if x != nil {
		return x.AssetIn
	}
	return ""
}

func (x *Swap) GetAssetOut() string {
	if x != nil {
		return x.AssetOut
	}
	return ""
}

func (x *Swap) GetAmountIn() string {
	if x != nil {
		return x.AmountIn
	}
	return ""
}

func (x *Swap) GetExchange() Exchange {
	if x != nil {
		return x.Exchange
	}
	return Exchange_EXCHANGE_UNSPECIFIED
}

func (x *Swap) GetParams() *SwapParams {
	if x != nil {
		return x.Params
	}
	return nil
}

// Helper function to make safe "blind" Uniswap Swaps by comparing value in vs value out of the swap.
//
// Represents function `oracleSwap(ERC20 assetIn, ERC20 assetOut, uint256 amountIn, SwapRouter.Exchange exchange, bytes memory params, uint64 slippage)`
type OracleSwap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Asset to swap from
	AssetIn string `protobuf:"bytes,1,opt,name=asset_in,json=assetIn,proto3" json:"asset_in,omitempty"`
	// Asset to swap to
	AssetOut string `protobuf:"bytes,2,opt,name=asset_out,json=assetOut,proto3" json:"asset_out,omitempty"`
	// Amount to swap
	AmountIn string `protobuf:"bytes,3,opt,name=amount_in,json=amountIn,proto3" json:"amount_in,omitempty"`
	// The exchange to make the swap on
	Exchange Exchange `protobuf:"varint,4,opt,name=exchange,proto3,enum=steward.v2.Exchange" json:"exchange,omitempty"`
	// The parameters for the swap
	Params *OracleSwapParams `protobuf:"bytes,5,opt,name=params,proto3" json:"params,omitempty"`
	// The slippage allowed for the swap
	Slippage uint64 `protobuf:"varint,6,opt,name=slippage,proto3" json:"slippage,omitempty"`
}

func (x *OracleSwap) Reset() {
	*x = OracleSwap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adaptors_base_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OracleSwap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OracleSwap) ProtoMessage() {}

func (x *OracleSwap) ProtoReflect() protoreflect.Message {
	mi := &file_adaptors_base_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OracleSwap.ProtoReflect.Descriptor instead.
func (*OracleSwap) Descriptor() ([]byte, []int) {
	return file_adaptors_base_proto_rawDescGZIP(), []int{1}
}

func (x *OracleSwap) GetAssetIn() string {
	if x != nil {
		return x.AssetIn
	}
	return ""
}

func (x *OracleSwap) GetAssetOut() string {
	if x != nil {
		return x.AssetOut
	}
	return ""
}

func (x *OracleSwap) GetAmountIn() string {
	if x != nil {
		return x.AmountIn
	}
	return ""
}

func (x *OracleSwap) GetExchange() Exchange {
	if x != nil {
		return x.Exchange
	}
	return Exchange_EXCHANGE_UNSPECIFIED
}

func (x *OracleSwap) GetParams() *OracleSwapParams {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *OracleSwap) GetSlippage() uint64 {
	if x != nil {
		return x.Slippage
	}
	return 0
}

// Allows strategists to zero out an approval for a given `asset`.
//
// Represents function `revokeApproval(ERC20 asset, address spender)`
type RevokeApproval struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ERC20 Asset to revoke spender's approval for
	Asset string `protobuf:"bytes,1,opt,name=asset,proto3" json:"asset,omitempty"`
	// The spender to revoke approval of asset for
	Spender string `protobuf:"bytes,2,opt,name=spender,proto3" json:"spender,omitempty"`
}

func (x *RevokeApproval) Reset() {
	*x = RevokeApproval{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adaptors_base_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RevokeApproval) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RevokeApproval) ProtoMessage() {}

func (x *RevokeApproval) ProtoReflect() protoreflect.Message {
	mi := &file_adaptors_base_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RevokeApproval.ProtoReflect.Descriptor instead.
func (*RevokeApproval) Descriptor() ([]byte, []int) {
	return file_adaptors_base_proto_rawDescGZIP(), []int{2}
}

func (x *RevokeApproval) GetAsset() string {
	if x != nil {
		return x.Asset
	}
	return ""
}

func (x *RevokeApproval) GetSpender() string {
	if x != nil {
		return x.Spender
	}
	return ""
}

var File_adaptors_base_proto protoreflect.FileDescriptor

var file_adaptors_base_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x64, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76,
	0x32, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xbd, 0x01, 0x0a, 0x04, 0x53, 0x77, 0x61, 0x70, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x5f, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x49, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x6f, 0x75, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x73, 0x73, 0x65, 0x74, 0x4f, 0x75, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x12, 0x30, 0x0a,
	0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x14, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x32, 0x2e, 0x45, 0x78, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12,
	0x2e, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x77, 0x61,
	0x70, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22,
	0xe5, 0x01, 0x0a, 0x0a, 0x4f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x53, 0x77, 0x61, 0x70, 0x12, 0x19,
	0x0a, 0x08, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x73, 0x73, 0x65, 0x74, 0x49, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x5f, 0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x73,
	0x73, 0x65, 0x74, 0x4f, 0x75, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x6e, 0x12, 0x30, 0x0a, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e,
	0x76, 0x32, 0x2e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x08, 0x65, 0x78, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x34, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e,
	0x76, 0x32, 0x2e, 0x4f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x53, 0x77, 0x61, 0x70, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x73,
	0x6c, 0x69, 0x70, 0x70, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x73,
	0x6c, 0x69, 0x70, 0x70, 0x61, 0x67, 0x65, 0x22, 0x40, 0x0a, 0x0e, 0x52, 0x65, 0x76, 0x6f, 0x6b,
	0x65, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x73, 0x73, 0x65, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x42, 0x10, 0x5a, 0x0e, 0x2f, 0x73, 0x74,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_adaptors_base_proto_rawDescOnce sync.Once
	file_adaptors_base_proto_rawDescData = file_adaptors_base_proto_rawDesc
)

func file_adaptors_base_proto_rawDescGZIP() []byte {
	file_adaptors_base_proto_rawDescOnce.Do(func() {
		file_adaptors_base_proto_rawDescData = protoimpl.X.CompressGZIP(file_adaptors_base_proto_rawDescData)
	})
	return file_adaptors_base_proto_rawDescData
}

var file_adaptors_base_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_adaptors_base_proto_goTypes = []interface{}{
	(*Swap)(nil),             // 0: steward.v2.Swap
	(*OracleSwap)(nil),       // 1: steward.v2.OracleSwap
	(*RevokeApproval)(nil),   // 2: steward.v2.RevokeApproval
	(Exchange)(0),            // 3: steward.v2.Exchange
	(*SwapParams)(nil),       // 4: steward.v2.SwapParams
	(*OracleSwapParams)(nil), // 5: steward.v2.OracleSwapParams
}
var file_adaptors_base_proto_depIdxs = []int32{
	3, // 0: steward.v2.Swap.exchange:type_name -> steward.v2.Exchange
	4, // 1: steward.v2.Swap.params:type_name -> steward.v2.SwapParams
	3, // 2: steward.v2.OracleSwap.exchange:type_name -> steward.v2.Exchange
	5, // 3: steward.v2.OracleSwap.params:type_name -> steward.v2.OracleSwapParams
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_adaptors_base_proto_init() }
func file_adaptors_base_proto_init() {
	if File_adaptors_base_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_adaptors_base_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Swap); i {
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
		file_adaptors_base_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OracleSwap); i {
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
		file_adaptors_base_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RevokeApproval); i {
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
			RawDescriptor: file_adaptors_base_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_adaptors_base_proto_goTypes,
		DependencyIndexes: file_adaptors_base_proto_depIdxs,
		MessageInfos:      file_adaptors_base_proto_msgTypes,
	}.Build()
	File_adaptors_base_proto = out.File
	file_adaptors_base_proto_rawDesc = nil
	file_adaptors_base_proto_goTypes = nil
	file_adaptors_base_proto_depIdxs = nil
}
