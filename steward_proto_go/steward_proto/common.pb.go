//
// Protos for common message types.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        v3.21.12
// source: common.proto

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

// Exchange selector
type Exchange int32

const (
	Exchange_EXCHANGE_UNSPECIFIED Exchange = 0
	// Represents Uniswap V2
	Exchange_EXCHANGE_UNIV2 Exchange = 1
	// Represents Uniswap V3
	Exchange_EXCHANGE_UNIV3 Exchange = 2
)

// Enum value maps for Exchange.
var (
	Exchange_name = map[int32]string{
		0: "EXCHANGE_UNSPECIFIED",
		1: "EXCHANGE_UNIV2",
		2: "EXCHANGE_UNIV3",
	}
	Exchange_value = map[string]int32{
		"EXCHANGE_UNSPECIFIED": 0,
		"EXCHANGE_UNIV2":       1,
		"EXCHANGE_UNIV3":       2,
	}
)

func (x Exchange) Enum() *Exchange {
	p := new(Exchange)
	*p = x
	return p
}

func (x Exchange) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Exchange) Descriptor() protoreflect.EnumDescriptor {
	return file_common_proto_enumTypes[0].Descriptor()
}

func (Exchange) Type() protoreflect.EnumType {
	return &file_common_proto_enumTypes[0]
}

func (x Exchange) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Exchange.Descriptor instead.
func (Exchange) EnumDescriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

// Represents swap parameters for UniswapV2
type UniV2SwapParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Array of addresses dictating what swap path to follow
	Path []string `protobuf:"bytes,1,rep,name=path,proto3" json:"path,omitempty"`
	// Amount of the first asset in the path to swap
	Amount string `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
	// The minimum amount of the last asset in the path to receive
	AmountOutMin string `protobuf:"bytes,3,opt,name=amount_out_min,json=amountOutMin,proto3" json:"amount_out_min,omitempty"`
}

func (x *UniV2SwapParams) Reset() {
	*x = UniV2SwapParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UniV2SwapParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UniV2SwapParams) ProtoMessage() {}

func (x *UniV2SwapParams) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UniV2SwapParams.ProtoReflect.Descriptor instead.
func (*UniV2SwapParams) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

func (x *UniV2SwapParams) GetPath() []string {
	if x != nil {
		return x.Path
	}
	return nil
}

func (x *UniV2SwapParams) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *UniV2SwapParams) GetAmountOutMin() string {
	if x != nil {
		return x.AmountOutMin
	}
	return ""
}

// Represents swap parameters for UniswapV3
type UniV3SwapParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Array of addresses dictating what swap path to follow
	Path []string `protobuf:"bytes,1,rep,name=path,proto3" json:"path,omitempty"`
	// Array of pool fees dictating what swap pools to use
	PoolFees []uint32 `protobuf:"varint,2,rep,packed,name=pool_fees,json=poolFees,proto3" json:"pool_fees,omitempty"`
	// Amount of the first asset in the path to swap
	Amount string `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
	// The minimum amount of the last asset in the path to receive
	AmountOutMin string `protobuf:"bytes,4,opt,name=amount_out_min,json=amountOutMin,proto3" json:"amount_out_min,omitempty"`
}

func (x *UniV3SwapParams) Reset() {
	*x = UniV3SwapParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UniV3SwapParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UniV3SwapParams) ProtoMessage() {}

func (x *UniV3SwapParams) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UniV3SwapParams.ProtoReflect.Descriptor instead.
func (*UniV3SwapParams) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{1}
}

func (x *UniV3SwapParams) GetPath() []string {
	if x != nil {
		return x.Path
	}
	return nil
}

func (x *UniV3SwapParams) GetPoolFees() []uint32 {
	if x != nil {
		return x.PoolFees
	}
	return nil
}

func (x *UniV3SwapParams) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *UniV3SwapParams) GetAmountOutMin() string {
	if x != nil {
		return x.AmountOutMin
	}
	return ""
}

// Represents swap parameters for an exchange
type SwapParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Params:
	//
	//	*SwapParams_Univ2Params
	//	*SwapParams_Univ3Params
	Params isSwapParams_Params `protobuf_oneof:"params"`
}

func (x *SwapParams) Reset() {
	*x = SwapParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SwapParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SwapParams) ProtoMessage() {}

func (x *SwapParams) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SwapParams.ProtoReflect.Descriptor instead.
func (*SwapParams) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{2}
}

func (m *SwapParams) GetParams() isSwapParams_Params {
	if m != nil {
		return m.Params
	}
	return nil
}

func (x *SwapParams) GetUniv2Params() *UniV2SwapParams {
	if x, ok := x.GetParams().(*SwapParams_Univ2Params); ok {
		return x.Univ2Params
	}
	return nil
}

func (x *SwapParams) GetUniv3Params() *UniV3SwapParams {
	if x, ok := x.GetParams().(*SwapParams_Univ3Params); ok {
		return x.Univ3Params
	}
	return nil
}

type isSwapParams_Params interface {
	isSwapParams_Params()
}

type SwapParams_Univ2Params struct {
	// Params for a Uniswap V2 swap
	Univ2Params *UniV2SwapParams `protobuf:"bytes,1,opt,name=univ2_params,json=univ2Params,proto3,oneof"`
}

type SwapParams_Univ3Params struct {
	// Params for a Uniswap V3 swap
	Univ3Params *UniV3SwapParams `protobuf:"bytes,2,opt,name=univ3_params,json=univ3Params,proto3,oneof"`
}

func (*SwapParams_Univ2Params) isSwapParams_Params() {}

func (*SwapParams_Univ3Params) isSwapParams_Params() {}

// Represents oracle swap parameters for UniswapV2
type UniV2OracleSwapParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Array of addresses dictating what swap path to follow
	Path []string `protobuf:"bytes,1,rep,name=path,proto3" json:"path,omitempty"`
}

func (x *UniV2OracleSwapParams) Reset() {
	*x = UniV2OracleSwapParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UniV2OracleSwapParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UniV2OracleSwapParams) ProtoMessage() {}

func (x *UniV2OracleSwapParams) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UniV2OracleSwapParams.ProtoReflect.Descriptor instead.
func (*UniV2OracleSwapParams) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{3}
}

func (x *UniV2OracleSwapParams) GetPath() []string {
	if x != nil {
		return x.Path
	}
	return nil
}

// Represents oracle swap parameters for UniswapV3
type UniV3OracleSwapParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Array of addresses dictating what swap path to follow
	Path []string `protobuf:"bytes,1,rep,name=path,proto3" json:"path,omitempty"`
	// Array of pool fees dictating what swap pools to use
	PoolFees []uint32 `protobuf:"varint,2,rep,packed,name=pool_fees,json=poolFees,proto3" json:"pool_fees,omitempty"`
}

func (x *UniV3OracleSwapParams) Reset() {
	*x = UniV3OracleSwapParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UniV3OracleSwapParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UniV3OracleSwapParams) ProtoMessage() {}

func (x *UniV3OracleSwapParams) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UniV3OracleSwapParams.ProtoReflect.Descriptor instead.
func (*UniV3OracleSwapParams) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{4}
}

func (x *UniV3OracleSwapParams) GetPath() []string {
	if x != nil {
		return x.Path
	}
	return nil
}

func (x *UniV3OracleSwapParams) GetPoolFees() []uint32 {
	if x != nil {
		return x.PoolFees
	}
	return nil
}

// Represents swap params for BaseAdaptor.oracleSwap()
type OracleSwapParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Params:
	//
	//	*OracleSwapParams_Univ2Params
	//	*OracleSwapParams_Univ3Params
	Params isOracleSwapParams_Params `protobuf_oneof:"params"`
}

func (x *OracleSwapParams) Reset() {
	*x = OracleSwapParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OracleSwapParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OracleSwapParams) ProtoMessage() {}

func (x *OracleSwapParams) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OracleSwapParams.ProtoReflect.Descriptor instead.
func (*OracleSwapParams) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{5}
}

func (m *OracleSwapParams) GetParams() isOracleSwapParams_Params {
	if m != nil {
		return m.Params
	}
	return nil
}

func (x *OracleSwapParams) GetUniv2Params() *UniV2OracleSwapParams {
	if x, ok := x.GetParams().(*OracleSwapParams_Univ2Params); ok {
		return x.Univ2Params
	}
	return nil
}

func (x *OracleSwapParams) GetUniv3Params() *UniV3OracleSwapParams {
	if x, ok := x.GetParams().(*OracleSwapParams_Univ3Params); ok {
		return x.Univ3Params
	}
	return nil
}

type isOracleSwapParams_Params interface {
	isOracleSwapParams_Params()
}

type OracleSwapParams_Univ2Params struct {
	Univ2Params *UniV2OracleSwapParams `protobuf:"bytes,1,opt,name=univ2_params,json=univ2Params,proto3,oneof"`
}

type OracleSwapParams_Univ3Params struct {
	Univ3Params *UniV3OracleSwapParams `protobuf:"bytes,2,opt,name=univ3_params,json=univ3Params,proto3,oneof"`
}

func (*OracleSwapParams_Univ2Params) isOracleSwapParams_Params() {}

func (*OracleSwapParams_Univ3Params) isOracleSwapParams_Params() {}

var File_common_proto protoreflect.FileDescriptor

var file_common_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a,
	0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x32, 0x22, 0x63, 0x0a, 0x0f, 0x55, 0x6e,
	0x69, 0x56, 0x32, 0x53, 0x77, 0x61, 0x70, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x6f, 0x75, 0x74, 0x5f, 0x6d, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x4f, 0x75, 0x74, 0x4d, 0x69, 0x6e, 0x22,
	0x80, 0x01, 0x0a, 0x0f, 0x55, 0x6e, 0x69, 0x56, 0x33, 0x53, 0x77, 0x61, 0x70, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6f, 0x6f, 0x6c, 0x5f,
	0x66, 0x65, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x6f, 0x6f, 0x6c,
	0x46, 0x65, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x0e,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6f, 0x75, 0x74, 0x5f, 0x6d, 0x69, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x4f, 0x75, 0x74, 0x4d,
	0x69, 0x6e, 0x22, 0x9a, 0x01, 0x0a, 0x0a, 0x53, 0x77, 0x61, 0x70, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x12, 0x40, 0x0a, 0x0c, 0x75, 0x6e, 0x69, 0x76, 0x32, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x2e, 0x76, 0x32, 0x2e, 0x55, 0x6e, 0x69, 0x56, 0x32, 0x53, 0x77, 0x61, 0x70, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x48, 0x00, 0x52, 0x0b, 0x75, 0x6e, 0x69, 0x76, 0x32, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x12, 0x40, 0x0a, 0x0c, 0x75, 0x6e, 0x69, 0x76, 0x33, 0x5f, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x74, 0x65, 0x77,
	0x61, 0x72, 0x64, 0x2e, 0x76, 0x32, 0x2e, 0x55, 0x6e, 0x69, 0x56, 0x33, 0x53, 0x77, 0x61, 0x70,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x48, 0x00, 0x52, 0x0b, 0x75, 0x6e, 0x69, 0x76, 0x33, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x42, 0x08, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22,
	0x2b, 0x0a, 0x15, 0x55, 0x6e, 0x69, 0x56, 0x32, 0x4f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x53, 0x77,
	0x61, 0x70, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x48, 0x0a, 0x15,
	0x55, 0x6e, 0x69, 0x56, 0x33, 0x4f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x53, 0x77, 0x61, 0x70, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6f, 0x6f,
	0x6c, 0x5f, 0x66, 0x65, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x6f,
	0x6f, 0x6c, 0x46, 0x65, 0x65, 0x73, 0x22, 0xac, 0x01, 0x0a, 0x10, 0x4f, 0x72, 0x61, 0x63, 0x6c,
	0x65, 0x53, 0x77, 0x61, 0x70, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x46, 0x0a, 0x0c, 0x75,
	0x6e, 0x69, 0x76, 0x32, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x21, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x32, 0x2e, 0x55,
	0x6e, 0x69, 0x56, 0x32, 0x4f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x53, 0x77, 0x61, 0x70, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x48, 0x00, 0x52, 0x0b, 0x75, 0x6e, 0x69, 0x76, 0x32, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x12, 0x46, 0x0a, 0x0c, 0x75, 0x6e, 0x69, 0x76, 0x33, 0x5f, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x73, 0x74, 0x65, 0x77,
	0x61, 0x72, 0x64, 0x2e, 0x76, 0x32, 0x2e, 0x55, 0x6e, 0x69, 0x56, 0x33, 0x4f, 0x72, 0x61, 0x63,
	0x6c, 0x65, 0x53, 0x77, 0x61, 0x70, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x48, 0x00, 0x52, 0x0b,
	0x75, 0x6e, 0x69, 0x76, 0x33, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x42, 0x08, 0x0a, 0x06, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x2a, 0x4c, 0x0a, 0x08, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x12, 0x18, 0x0a, 0x14, 0x45, 0x58, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x45,
	0x58, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x55, 0x4e, 0x49, 0x56, 0x32, 0x10, 0x01, 0x12,
	0x12, 0x0a, 0x0e, 0x45, 0x58, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x55, 0x4e, 0x49, 0x56,
	0x33, 0x10, 0x02, 0x42, 0x10, 0x5a, 0x0e, 0x2f, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_proto_rawDescOnce sync.Once
	file_common_proto_rawDescData = file_common_proto_rawDesc
)

func file_common_proto_rawDescGZIP() []byte {
	file_common_proto_rawDescOnce.Do(func() {
		file_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_proto_rawDescData)
	})
	return file_common_proto_rawDescData
}

var file_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_common_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_common_proto_goTypes = []interface{}{
	(Exchange)(0),                 // 0: steward.v2.Exchange
	(*UniV2SwapParams)(nil),       // 1: steward.v2.UniV2SwapParams
	(*UniV3SwapParams)(nil),       // 2: steward.v2.UniV3SwapParams
	(*SwapParams)(nil),            // 3: steward.v2.SwapParams
	(*UniV2OracleSwapParams)(nil), // 4: steward.v2.UniV2OracleSwapParams
	(*UniV3OracleSwapParams)(nil), // 5: steward.v2.UniV3OracleSwapParams
	(*OracleSwapParams)(nil),      // 6: steward.v2.OracleSwapParams
}
var file_common_proto_depIdxs = []int32{
	1, // 0: steward.v2.SwapParams.univ2_params:type_name -> steward.v2.UniV2SwapParams
	2, // 1: steward.v2.SwapParams.univ3_params:type_name -> steward.v2.UniV3SwapParams
	4, // 2: steward.v2.OracleSwapParams.univ2_params:type_name -> steward.v2.UniV2OracleSwapParams
	5, // 3: steward.v2.OracleSwapParams.univ3_params:type_name -> steward.v2.UniV3OracleSwapParams
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_common_proto_init() }
func file_common_proto_init() {
	if File_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UniV2SwapParams); i {
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
		file_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UniV3SwapParams); i {
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
		file_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SwapParams); i {
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
		file_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UniV2OracleSwapParams); i {
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
		file_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UniV3OracleSwapParams); i {
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
		file_common_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OracleSwapParams); i {
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
	file_common_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*SwapParams_Univ2Params)(nil),
		(*SwapParams_Univ3Params)(nil),
	}
	file_common_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*OracleSwapParams_Univ2Params)(nil),
		(*OracleSwapParams_Univ3Params)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_proto_goTypes,
		DependencyIndexes: file_common_proto_depIdxs,
		EnumInfos:         file_common_proto_enumTypes,
		MessageInfos:      file_common_proto_msgTypes,
	}.Build()
	File_common_proto = out.File
	file_common_proto_rawDesc = nil
	file_common_proto_goTypes = nil
	file_common_proto_depIdxs = nil
}
