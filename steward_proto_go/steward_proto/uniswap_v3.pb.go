// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.12
// source: adaptors/uniswap/uniswap_v3.proto

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

// Represents call data for the Uniswap V3 adaptor
type UniswapV3Adaptor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Function:
	//	*UniswapV3Adaptor_Swap
	//	*UniswapV3Adaptor_OracleSwap
	//	*UniswapV3Adaptor_RevokeApproval
	//	*UniswapV3Adaptor_OpenPosition_
	//	*UniswapV3Adaptor_ClosePosition_
	//	*UniswapV3Adaptor_AddToPosition_
	//	*UniswapV3Adaptor_TakeFromPosition_
	//	*UniswapV3Adaptor_CollectFees_
	Function isUniswapV3Adaptor_Function `protobuf_oneof:"function"`
}

func (x *UniswapV3Adaptor) Reset() {
	*x = UniswapV3Adaptor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adaptors_uniswap_uniswap_v3_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UniswapV3Adaptor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UniswapV3Adaptor) ProtoMessage() {}

func (x *UniswapV3Adaptor) ProtoReflect() protoreflect.Message {
	mi := &file_adaptors_uniswap_uniswap_v3_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UniswapV3Adaptor.ProtoReflect.Descriptor instead.
func (*UniswapV3Adaptor) Descriptor() ([]byte, []int) {
	return file_adaptors_uniswap_uniswap_v3_proto_rawDescGZIP(), []int{0}
}

func (m *UniswapV3Adaptor) GetFunction() isUniswapV3Adaptor_Function {
	if m != nil {
		return m.Function
	}
	return nil
}

func (x *UniswapV3Adaptor) GetSwap() *Swap {
	if x, ok := x.GetFunction().(*UniswapV3Adaptor_Swap); ok {
		return x.Swap
	}
	return nil
}

func (x *UniswapV3Adaptor) GetOracleSwap() *OracleSwap {
	if x, ok := x.GetFunction().(*UniswapV3Adaptor_OracleSwap); ok {
		return x.OracleSwap
	}
	return nil
}

func (x *UniswapV3Adaptor) GetRevokeApproval() *RevokeApproval {
	if x, ok := x.GetFunction().(*UniswapV3Adaptor_RevokeApproval); ok {
		return x.RevokeApproval
	}
	return nil
}

func (x *UniswapV3Adaptor) GetOpenPosition() *UniswapV3Adaptor_OpenPosition {
	if x, ok := x.GetFunction().(*UniswapV3Adaptor_OpenPosition_); ok {
		return x.OpenPosition
	}
	return nil
}

func (x *UniswapV3Adaptor) GetClosePosition() *UniswapV3Adaptor_ClosePosition {
	if x, ok := x.GetFunction().(*UniswapV3Adaptor_ClosePosition_); ok {
		return x.ClosePosition
	}
	return nil
}

func (x *UniswapV3Adaptor) GetAddToPosition() *UniswapV3Adaptor_AddToPosition {
	if x, ok := x.GetFunction().(*UniswapV3Adaptor_AddToPosition_); ok {
		return x.AddToPosition
	}
	return nil
}

func (x *UniswapV3Adaptor) GetTakeFromPosition() *UniswapV3Adaptor_TakeFromPosition {
	if x, ok := x.GetFunction().(*UniswapV3Adaptor_TakeFromPosition_); ok {
		return x.TakeFromPosition
	}
	return nil
}

func (x *UniswapV3Adaptor) GetCollectFees() *UniswapV3Adaptor_CollectFees {
	if x, ok := x.GetFunction().(*UniswapV3Adaptor_CollectFees_); ok {
		return x.CollectFees
	}
	return nil
}

type isUniswapV3Adaptor_Function interface {
	isUniswapV3Adaptor_Function()
}

type UniswapV3Adaptor_Swap struct {
	// Represents function `swap(ERC20 assetIn, ERC20 assetOut, uint256 amountIn, SwapRouter.Exchange exchange, bytes memory params)`
	Swap *Swap `protobuf:"bytes,1,opt,name=swap,proto3,oneof"`
}

type UniswapV3Adaptor_OracleSwap struct {
	// Represents function `oracleSwap(ERC20 assetIn, ERC20 assetOut, uint256 amountIn, SwapRouter.Exchange exchange, bytes memory params, uint64 slippage)`
	OracleSwap *OracleSwap `protobuf:"bytes,2,opt,name=oracle_swap,json=oracleSwap,proto3,oneof"`
}

type UniswapV3Adaptor_RevokeApproval struct {
	// Represents function `revokeApproval(ERC20 asset, address spender)`
	RevokeApproval *RevokeApproval `protobuf:"bytes,3,opt,name=revoke_approval,json=revokeApproval,proto3,oneof"`
}

type UniswapV3Adaptor_OpenPosition_ struct {
	// Represents function `openPosition(ERC20 token0, ERC20 token1, uint24 poolFee, uint256 amount0, uint256 amount1, uint256 min0, uint256 min1, int24 tickLower, int24 tickUpper)`
	OpenPosition *UniswapV3Adaptor_OpenPosition `protobuf:"bytes,4,opt,name=open_position,json=openPosition,proto3,oneof"`
}

type UniswapV3Adaptor_ClosePosition_ struct {
	// Represents function `closePosition(uint256 positionId, uint256 min0, uint256 min1)`
	ClosePosition *UniswapV3Adaptor_ClosePosition `protobuf:"bytes,5,opt,name=close_position,json=closePosition,proto3,oneof"`
}

type UniswapV3Adaptor_AddToPosition_ struct {
	// Represents function `addToPosition(uint256 positionId, uint256 amount0, uint256 amount1, uint256 min0, uint256 min1)`
	AddToPosition *UniswapV3Adaptor_AddToPosition `protobuf:"bytes,6,opt,name=add_to_position,json=addToPosition,proto3,oneof"`
}

type UniswapV3Adaptor_TakeFromPosition_ struct {
	// Represents function `takeFromPosition(uint256 positionId, uint128 liquidity, uint256 min0, uint256 min1, bool collectFees)`
	TakeFromPosition *UniswapV3Adaptor_TakeFromPosition `protobuf:"bytes,7,opt,name=take_from_position,json=takeFromPosition,proto3,oneof"`
}

type UniswapV3Adaptor_CollectFees_ struct {
	// Represents function `collectFees(uint256 positionId, uint128 amount0, uint128 amount1)`
	CollectFees *UniswapV3Adaptor_CollectFees `protobuf:"bytes,8,opt,name=collect_fees,json=collectFees,proto3,oneof"`
}

func (*UniswapV3Adaptor_Swap) isUniswapV3Adaptor_Function() {}

func (*UniswapV3Adaptor_OracleSwap) isUniswapV3Adaptor_Function() {}

func (*UniswapV3Adaptor_RevokeApproval) isUniswapV3Adaptor_Function() {}

func (*UniswapV3Adaptor_OpenPosition_) isUniswapV3Adaptor_Function() {}

func (*UniswapV3Adaptor_ClosePosition_) isUniswapV3Adaptor_Function() {}

func (*UniswapV3Adaptor_AddToPosition_) isUniswapV3Adaptor_Function() {}

func (*UniswapV3Adaptor_TakeFromPosition_) isUniswapV3Adaptor_Function() {}

func (*UniswapV3Adaptor_CollectFees_) isUniswapV3Adaptor_Function() {}

type UniswapV3AdaptorCalls struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Calls []*UniswapV3Adaptor `protobuf:"bytes,1,rep,name=calls,proto3" json:"calls,omitempty"`
}

func (x *UniswapV3AdaptorCalls) Reset() {
	*x = UniswapV3AdaptorCalls{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adaptors_uniswap_uniswap_v3_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UniswapV3AdaptorCalls) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UniswapV3AdaptorCalls) ProtoMessage() {}

func (x *UniswapV3AdaptorCalls) ProtoReflect() protoreflect.Message {
	mi := &file_adaptors_uniswap_uniswap_v3_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UniswapV3AdaptorCalls.ProtoReflect.Descriptor instead.
func (*UniswapV3AdaptorCalls) Descriptor() ([]byte, []int) {
	return file_adaptors_uniswap_uniswap_v3_proto_rawDescGZIP(), []int{1}
}

func (x *UniswapV3AdaptorCalls) GetCalls() []*UniswapV3Adaptor {
	if x != nil {
		return x.Calls
	}
	return nil
}

//
// Allows strategist to open up arbritray Uniswap V3 positions.
//
// Represents function openPosition(ERC20 token0, ERC20 token1, uint24 poolFee, uint256 amount0, uint256 amount1, uint256 min0, uint256 min1, int24 tickLower, int24 tickUpper)
type UniswapV3Adaptor_OpenPosition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token_0   string `protobuf:"bytes,1,opt,name=token_0,json=token0,proto3" json:"token_0,omitempty"`
	Token_1   string `protobuf:"bytes,2,opt,name=token_1,json=token1,proto3" json:"token_1,omitempty"`
	PoolFee   uint32 `protobuf:"varint,3,opt,name=pool_fee,json=poolFee,proto3" json:"pool_fee,omitempty"`
	Amount_0  string `protobuf:"bytes,4,opt,name=amount_0,json=amount0,proto3" json:"amount_0,omitempty"`
	Amount_1  string `protobuf:"bytes,5,opt,name=amount_1,json=amount1,proto3" json:"amount_1,omitempty"`
	Min_0     string `protobuf:"bytes,6,opt,name=min_0,json=min0,proto3" json:"min_0,omitempty"`
	Min_1     string `protobuf:"bytes,7,opt,name=min_1,json=min1,proto3" json:"min_1,omitempty"`
	TickLower int32  `protobuf:"varint,8,opt,name=tick_lower,json=tickLower,proto3" json:"tick_lower,omitempty"`
	TickUpper int32  `protobuf:"varint,9,opt,name=tick_upper,json=tickUpper,proto3" json:"tick_upper,omitempty"`
}

func (x *UniswapV3Adaptor_OpenPosition) Reset() {
	*x = UniswapV3Adaptor_OpenPosition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adaptors_uniswap_uniswap_v3_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UniswapV3Adaptor_OpenPosition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UniswapV3Adaptor_OpenPosition) ProtoMessage() {}

func (x *UniswapV3Adaptor_OpenPosition) ProtoReflect() protoreflect.Message {
	mi := &file_adaptors_uniswap_uniswap_v3_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UniswapV3Adaptor_OpenPosition.ProtoReflect.Descriptor instead.
func (*UniswapV3Adaptor_OpenPosition) Descriptor() ([]byte, []int) {
	return file_adaptors_uniswap_uniswap_v3_proto_rawDescGZIP(), []int{0, 0}
}

func (x *UniswapV3Adaptor_OpenPosition) GetToken_0() string {
	if x != nil {
		return x.Token_0
	}
	return ""
}

func (x *UniswapV3Adaptor_OpenPosition) GetToken_1() string {
	if x != nil {
		return x.Token_1
	}
	return ""
}

func (x *UniswapV3Adaptor_OpenPosition) GetPoolFee() uint32 {
	if x != nil {
		return x.PoolFee
	}
	return 0
}

func (x *UniswapV3Adaptor_OpenPosition) GetAmount_0() string {
	if x != nil {
		return x.Amount_0
	}
	return ""
}

func (x *UniswapV3Adaptor_OpenPosition) GetAmount_1() string {
	if x != nil {
		return x.Amount_1
	}
	return ""
}

func (x *UniswapV3Adaptor_OpenPosition) GetMin_0() string {
	if x != nil {
		return x.Min_0
	}
	return ""
}

func (x *UniswapV3Adaptor_OpenPosition) GetMin_1() string {
	if x != nil {
		return x.Min_1
	}
	return ""
}

func (x *UniswapV3Adaptor_OpenPosition) GetTickLower() int32 {
	if x != nil {
		return x.TickLower
	}
	return 0
}

func (x *UniswapV3Adaptor_OpenPosition) GetTickUpper() int32 {
	if x != nil {
		return x.TickUpper
	}
	return 0
}

//
// Allows strategist to close Uniswap V3 positions.
//
// Represents function `closePosition(uint256 positionId, uint256 min0, uint256 min1)`
type UniswapV3Adaptor_ClosePosition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PositionId string `protobuf:"bytes,1,opt,name=position_id,json=positionId,proto3" json:"position_id,omitempty"`
	Min_0      string `protobuf:"bytes,2,opt,name=min_0,json=min0,proto3" json:"min_0,omitempty"`
	Min_1      string `protobuf:"bytes,3,opt,name=min_1,json=min1,proto3" json:"min_1,omitempty"`
}

func (x *UniswapV3Adaptor_ClosePosition) Reset() {
	*x = UniswapV3Adaptor_ClosePosition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adaptors_uniswap_uniswap_v3_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UniswapV3Adaptor_ClosePosition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UniswapV3Adaptor_ClosePosition) ProtoMessage() {}

func (x *UniswapV3Adaptor_ClosePosition) ProtoReflect() protoreflect.Message {
	mi := &file_adaptors_uniswap_uniswap_v3_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UniswapV3Adaptor_ClosePosition.ProtoReflect.Descriptor instead.
func (*UniswapV3Adaptor_ClosePosition) Descriptor() ([]byte, []int) {
	return file_adaptors_uniswap_uniswap_v3_proto_rawDescGZIP(), []int{0, 1}
}

func (x *UniswapV3Adaptor_ClosePosition) GetPositionId() string {
	if x != nil {
		return x.PositionId
	}
	return ""
}

func (x *UniswapV3Adaptor_ClosePosition) GetMin_0() string {
	if x != nil {
		return x.Min_0
	}
	return ""
}

func (x *UniswapV3Adaptor_ClosePosition) GetMin_1() string {
	if x != nil {
		return x.Min_1
	}
	return ""
}

//
// Allows strategist to add to existing Uniswap V3 positions.
//
// Represents function `addToPosition(uint256 positionId, uint256 amount0, uint256 amount1, uint256 min0, uint256 min1)`
type UniswapV3Adaptor_AddToPosition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PositionId string `protobuf:"bytes,1,opt,name=position_id,json=positionId,proto3" json:"position_id,omitempty"`
	Amount_0   string `protobuf:"bytes,2,opt,name=amount_0,json=amount0,proto3" json:"amount_0,omitempty"`
	Amount_1   string `protobuf:"bytes,3,opt,name=amount_1,json=amount1,proto3" json:"amount_1,omitempty"`
	Min_0      string `protobuf:"bytes,4,opt,name=min_0,json=min0,proto3" json:"min_0,omitempty"`
	Min_1      string `protobuf:"bytes,5,opt,name=min_1,json=min1,proto3" json:"min_1,omitempty"`
}

func (x *UniswapV3Adaptor_AddToPosition) Reset() {
	*x = UniswapV3Adaptor_AddToPosition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adaptors_uniswap_uniswap_v3_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UniswapV3Adaptor_AddToPosition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UniswapV3Adaptor_AddToPosition) ProtoMessage() {}

func (x *UniswapV3Adaptor_AddToPosition) ProtoReflect() protoreflect.Message {
	mi := &file_adaptors_uniswap_uniswap_v3_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UniswapV3Adaptor_AddToPosition.ProtoReflect.Descriptor instead.
func (*UniswapV3Adaptor_AddToPosition) Descriptor() ([]byte, []int) {
	return file_adaptors_uniswap_uniswap_v3_proto_rawDescGZIP(), []int{0, 2}
}

func (x *UniswapV3Adaptor_AddToPosition) GetPositionId() string {
	if x != nil {
		return x.PositionId
	}
	return ""
}

func (x *UniswapV3Adaptor_AddToPosition) GetAmount_0() string {
	if x != nil {
		return x.Amount_0
	}
	return ""
}

func (x *UniswapV3Adaptor_AddToPosition) GetAmount_1() string {
	if x != nil {
		return x.Amount_1
	}
	return ""
}

func (x *UniswapV3Adaptor_AddToPosition) GetMin_0() string {
	if x != nil {
		return x.Min_0
	}
	return ""
}

func (x *UniswapV3Adaptor_AddToPosition) GetMin_1() string {
	if x != nil {
		return x.Min_1
	}
	return ""
}

//
// Allows strategist to take from existing Uniswap V3 positions.
//
// Represents function `takeFromPosition(uint256 positionId, uint128 liquidity, uint256 min0, uint256 min1, bool collectFees)`
type UniswapV3Adaptor_TakeFromPosition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PositionId  string `protobuf:"bytes,1,opt,name=position_id,json=positionId,proto3" json:"position_id,omitempty"`
	Liquidity   string `protobuf:"bytes,2,opt,name=liquidity,proto3" json:"liquidity,omitempty"`
	Min_0       string `protobuf:"bytes,3,opt,name=min_0,json=min0,proto3" json:"min_0,omitempty"`
	Min_1       string `protobuf:"bytes,4,opt,name=min_1,json=min1,proto3" json:"min_1,omitempty"`
	CollectFees bool   `protobuf:"varint,5,opt,name=collect_fees,json=collectFees,proto3" json:"collect_fees,omitempty"`
}

func (x *UniswapV3Adaptor_TakeFromPosition) Reset() {
	*x = UniswapV3Adaptor_TakeFromPosition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adaptors_uniswap_uniswap_v3_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UniswapV3Adaptor_TakeFromPosition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UniswapV3Adaptor_TakeFromPosition) ProtoMessage() {}

func (x *UniswapV3Adaptor_TakeFromPosition) ProtoReflect() protoreflect.Message {
	mi := &file_adaptors_uniswap_uniswap_v3_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UniswapV3Adaptor_TakeFromPosition.ProtoReflect.Descriptor instead.
func (*UniswapV3Adaptor_TakeFromPosition) Descriptor() ([]byte, []int) {
	return file_adaptors_uniswap_uniswap_v3_proto_rawDescGZIP(), []int{0, 3}
}

func (x *UniswapV3Adaptor_TakeFromPosition) GetPositionId() string {
	if x != nil {
		return x.PositionId
	}
	return ""
}

func (x *UniswapV3Adaptor_TakeFromPosition) GetLiquidity() string {
	if x != nil {
		return x.Liquidity
	}
	return ""
}

func (x *UniswapV3Adaptor_TakeFromPosition) GetMin_0() string {
	if x != nil {
		return x.Min_0
	}
	return ""
}

func (x *UniswapV3Adaptor_TakeFromPosition) GetMin_1() string {
	if x != nil {
		return x.Min_1
	}
	return ""
}

func (x *UniswapV3Adaptor_TakeFromPosition) GetCollectFees() bool {
	if x != nil {
		return x.CollectFees
	}
	return false
}

//
// Allows strategist to collect fees from existing Uniswap V3 positions.
//
// Represents function `collectFees(uint256 positionId, uint128 amount0, uint128 amount1)`
type UniswapV3Adaptor_CollectFees struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PositionId string `protobuf:"bytes,1,opt,name=position_id,json=positionId,proto3" json:"position_id,omitempty"`
	Amount_0   string `protobuf:"bytes,2,opt,name=amount_0,json=amount0,proto3" json:"amount_0,omitempty"`
	Amount_1   string `protobuf:"bytes,3,opt,name=amount_1,json=amount1,proto3" json:"amount_1,omitempty"`
}

func (x *UniswapV3Adaptor_CollectFees) Reset() {
	*x = UniswapV3Adaptor_CollectFees{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adaptors_uniswap_uniswap_v3_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UniswapV3Adaptor_CollectFees) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UniswapV3Adaptor_CollectFees) ProtoMessage() {}

func (x *UniswapV3Adaptor_CollectFees) ProtoReflect() protoreflect.Message {
	mi := &file_adaptors_uniswap_uniswap_v3_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UniswapV3Adaptor_CollectFees.ProtoReflect.Descriptor instead.
func (*UniswapV3Adaptor_CollectFees) Descriptor() ([]byte, []int) {
	return file_adaptors_uniswap_uniswap_v3_proto_rawDescGZIP(), []int{0, 4}
}

func (x *UniswapV3Adaptor_CollectFees) GetPositionId() string {
	if x != nil {
		return x.PositionId
	}
	return ""
}

func (x *UniswapV3Adaptor_CollectFees) GetAmount_0() string {
	if x != nil {
		return x.Amount_0
	}
	return ""
}

func (x *UniswapV3Adaptor_CollectFees) GetAmount_1() string {
	if x != nil {
		return x.Amount_1
	}
	return ""
}

var File_adaptors_uniswap_uniswap_v3_proto protoreflect.FileDescriptor

var file_adaptors_uniswap_uniswap_v3_proto_rawDesc = []byte{
	0x0a, 0x21, 0x61, 0x64, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x75, 0x6e, 0x69, 0x73, 0x77,
	0x61, 0x70, 0x2f, 0x75, 0x6e, 0x69, 0x73, 0x77, 0x61, 0x70, 0x5f, 0x76, 0x33, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x32, 0x1a,
	0x13, 0x61, 0x64, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe5, 0x0a, 0x0a, 0x10, 0x55, 0x6e, 0x69, 0x73, 0x77, 0x61, 0x70,
	0x56, 0x33, 0x41, 0x64, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x12, 0x26, 0x0a, 0x04, 0x73, 0x77, 0x61,
	0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x77, 0x61, 0x70, 0x48, 0x00, 0x52, 0x04, 0x73, 0x77, 0x61,
	0x70, 0x12, 0x39, 0x0a, 0x0b, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x5f, 0x73, 0x77, 0x61, 0x70,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64,
	0x2e, 0x76, 0x32, 0x2e, 0x4f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x53, 0x77, 0x61, 0x70, 0x48, 0x00,
	0x52, 0x0a, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x53, 0x77, 0x61, 0x70, 0x12, 0x45, 0x0a, 0x0f,
	0x72, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x5f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e,
	0x76, 0x32, 0x2e, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61,
	0x6c, 0x48, 0x00, 0x52, 0x0e, 0x72, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x41, 0x70, 0x70, 0x72, 0x6f,
	0x76, 0x61, 0x6c, 0x12, 0x50, 0x0a, 0x0d, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x73, 0x74, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x32, 0x2e, 0x55, 0x6e, 0x69, 0x73, 0x77, 0x61, 0x70, 0x56,
	0x33, 0x41, 0x64, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x50, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0c, 0x6f, 0x70, 0x65, 0x6e, 0x50, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x53, 0x0a, 0x0e, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x5f, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e,
	0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x32, 0x2e, 0x55, 0x6e, 0x69, 0x73, 0x77,
	0x61, 0x70, 0x56, 0x33, 0x41, 0x64, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x43, 0x6c, 0x6f, 0x73,
	0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0d, 0x63, 0x6c, 0x6f,
	0x73, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x54, 0x0a, 0x0f, 0x61, 0x64,
	0x64, 0x5f, 0x74, 0x6f, 0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x32,
	0x2e, 0x55, 0x6e, 0x69, 0x73, 0x77, 0x61, 0x70, 0x56, 0x33, 0x41, 0x64, 0x61, 0x70, 0x74, 0x6f,
	0x72, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x48,
	0x00, 0x52, 0x0d, 0x61, 0x64, 0x64, 0x54, 0x6f, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x5d, 0x0a, 0x12, 0x74, 0x61, 0x6b, 0x65, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x73,
	0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x32, 0x2e, 0x55, 0x6e, 0x69, 0x73, 0x77, 0x61,
	0x70, 0x56, 0x33, 0x41, 0x64, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x54, 0x61, 0x6b, 0x65, 0x46,
	0x72, 0x6f, 0x6d, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x10, 0x74,
	0x61, 0x6b, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x4d, 0x0a, 0x0c, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x5f, 0x66, 0x65, 0x65, 0x73, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e,
	0x76, 0x32, 0x2e, 0x55, 0x6e, 0x69, 0x73, 0x77, 0x61, 0x70, 0x56, 0x33, 0x41, 0x64, 0x61, 0x70,
	0x74, 0x6f, 0x72, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x46, 0x65, 0x65, 0x73, 0x48,
	0x00, 0x52, 0x0b, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x46, 0x65, 0x65, 0x73, 0x1a, 0xf9,
	0x01, 0x0a, 0x0c, 0x4f, 0x70, 0x65, 0x6e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x17, 0x0a, 0x07, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x30, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x30, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x5f, 0x31, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x31, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x66, 0x65, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x07, 0x70, 0x6f, 0x6f, 0x6c, 0x46, 0x65, 0x65, 0x12, 0x19, 0x0a, 0x08,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x30, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x30, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x31, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x31, 0x12, 0x13, 0x0a, 0x05, 0x6d, 0x69, 0x6e, 0x5f, 0x30, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6d, 0x69, 0x6e, 0x30, 0x12, 0x13, 0x0a, 0x05, 0x6d, 0x69, 0x6e, 0x5f, 0x31,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x69, 0x6e, 0x31, 0x12, 0x1d, 0x0a, 0x0a,
	0x74, 0x69, 0x63, 0x6b, 0x5f, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x74, 0x69, 0x63, 0x6b, 0x4c, 0x6f, 0x77, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x74,
	0x69, 0x63, 0x6b, 0x5f, 0x75, 0x70, 0x70, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x74, 0x69, 0x63, 0x6b, 0x55, 0x70, 0x70, 0x65, 0x72, 0x1a, 0x5a, 0x0a, 0x0d, 0x43, 0x6c,
	0x6f, 0x73, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x13, 0x0a, 0x05,
	0x6d, 0x69, 0x6e, 0x5f, 0x30, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x69, 0x6e,
	0x30, 0x12, 0x13, 0x0a, 0x05, 0x6d, 0x69, 0x6e, 0x5f, 0x31, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6d, 0x69, 0x6e, 0x31, 0x1a, 0x90, 0x01, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x54, 0x6f,
	0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x30, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x30, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x31,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x31, 0x12,
	0x13, 0x0a, 0x05, 0x6d, 0x69, 0x6e, 0x5f, 0x30, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6d, 0x69, 0x6e, 0x30, 0x12, 0x13, 0x0a, 0x05, 0x6d, 0x69, 0x6e, 0x5f, 0x31, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x69, 0x6e, 0x31, 0x1a, 0x9e, 0x01, 0x0a, 0x10, 0x54, 0x61,
	0x6b, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f,
	0x0a, 0x0b, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x1c, 0x0a, 0x09, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x12, 0x13, 0x0a,
	0x05, 0x6d, 0x69, 0x6e, 0x5f, 0x30, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x69,
	0x6e, 0x30, 0x12, 0x13, 0x0a, 0x05, 0x6d, 0x69, 0x6e, 0x5f, 0x31, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6d, 0x69, 0x6e, 0x31, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x5f, 0x66, 0x65, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x63,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x46, 0x65, 0x65, 0x73, 0x1a, 0x64, 0x0a, 0x0b, 0x43, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x46, 0x65, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x30, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x30, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f,
	0x31, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x31,
	0x42, 0x0a, 0x0a, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x4b, 0x0a, 0x15,
	0x55, 0x6e, 0x69, 0x73, 0x77, 0x61, 0x70, 0x56, 0x33, 0x41, 0x64, 0x61, 0x70, 0x74, 0x6f, 0x72,
	0x43, 0x61, 0x6c, 0x6c, 0x73, 0x12, 0x32, 0x0a, 0x05, 0x63, 0x61, 0x6c, 0x6c, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76,
	0x32, 0x2e, 0x55, 0x6e, 0x69, 0x73, 0x77, 0x61, 0x70, 0x56, 0x33, 0x41, 0x64, 0x61, 0x70, 0x74,
	0x6f, 0x72, 0x52, 0x05, 0x63, 0x61, 0x6c, 0x6c, 0x73, 0x42, 0x10, 0x5a, 0x0e, 0x2f, 0x73, 0x74,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_adaptors_uniswap_uniswap_v3_proto_rawDescOnce sync.Once
	file_adaptors_uniswap_uniswap_v3_proto_rawDescData = file_adaptors_uniswap_uniswap_v3_proto_rawDesc
)

func file_adaptors_uniswap_uniswap_v3_proto_rawDescGZIP() []byte {
	file_adaptors_uniswap_uniswap_v3_proto_rawDescOnce.Do(func() {
		file_adaptors_uniswap_uniswap_v3_proto_rawDescData = protoimpl.X.CompressGZIP(file_adaptors_uniswap_uniswap_v3_proto_rawDescData)
	})
	return file_adaptors_uniswap_uniswap_v3_proto_rawDescData
}

var file_adaptors_uniswap_uniswap_v3_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_adaptors_uniswap_uniswap_v3_proto_goTypes = []interface{}{
	(*UniswapV3Adaptor)(nil),                  // 0: steward.v2.UniswapV3Adaptor
	(*UniswapV3AdaptorCalls)(nil),             // 1: steward.v2.UniswapV3AdaptorCalls
	(*UniswapV3Adaptor_OpenPosition)(nil),     // 2: steward.v2.UniswapV3Adaptor.OpenPosition
	(*UniswapV3Adaptor_ClosePosition)(nil),    // 3: steward.v2.UniswapV3Adaptor.ClosePosition
	(*UniswapV3Adaptor_AddToPosition)(nil),    // 4: steward.v2.UniswapV3Adaptor.AddToPosition
	(*UniswapV3Adaptor_TakeFromPosition)(nil), // 5: steward.v2.UniswapV3Adaptor.TakeFromPosition
	(*UniswapV3Adaptor_CollectFees)(nil),      // 6: steward.v2.UniswapV3Adaptor.CollectFees
	(*Swap)(nil),                              // 7: steward.v2.Swap
	(*OracleSwap)(nil),                        // 8: steward.v2.OracleSwap
	(*RevokeApproval)(nil),                    // 9: steward.v2.RevokeApproval
}
var file_adaptors_uniswap_uniswap_v3_proto_depIdxs = []int32{
	7, // 0: steward.v2.UniswapV3Adaptor.swap:type_name -> steward.v2.Swap
	8, // 1: steward.v2.UniswapV3Adaptor.oracle_swap:type_name -> steward.v2.OracleSwap
	9, // 2: steward.v2.UniswapV3Adaptor.revoke_approval:type_name -> steward.v2.RevokeApproval
	2, // 3: steward.v2.UniswapV3Adaptor.open_position:type_name -> steward.v2.UniswapV3Adaptor.OpenPosition
	3, // 4: steward.v2.UniswapV3Adaptor.close_position:type_name -> steward.v2.UniswapV3Adaptor.ClosePosition
	4, // 5: steward.v2.UniswapV3Adaptor.add_to_position:type_name -> steward.v2.UniswapV3Adaptor.AddToPosition
	5, // 6: steward.v2.UniswapV3Adaptor.take_from_position:type_name -> steward.v2.UniswapV3Adaptor.TakeFromPosition
	6, // 7: steward.v2.UniswapV3Adaptor.collect_fees:type_name -> steward.v2.UniswapV3Adaptor.CollectFees
	0, // 8: steward.v2.UniswapV3AdaptorCalls.calls:type_name -> steward.v2.UniswapV3Adaptor
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_adaptors_uniswap_uniswap_v3_proto_init() }
func file_adaptors_uniswap_uniswap_v3_proto_init() {
	if File_adaptors_uniswap_uniswap_v3_proto != nil {
		return
	}
	file_adaptors_base_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_adaptors_uniswap_uniswap_v3_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UniswapV3Adaptor); i {
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
		file_adaptors_uniswap_uniswap_v3_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UniswapV3AdaptorCalls); i {
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
		file_adaptors_uniswap_uniswap_v3_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UniswapV3Adaptor_OpenPosition); i {
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
		file_adaptors_uniswap_uniswap_v3_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UniswapV3Adaptor_ClosePosition); i {
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
		file_adaptors_uniswap_uniswap_v3_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UniswapV3Adaptor_AddToPosition); i {
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
		file_adaptors_uniswap_uniswap_v3_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UniswapV3Adaptor_TakeFromPosition); i {
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
		file_adaptors_uniswap_uniswap_v3_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UniswapV3Adaptor_CollectFees); i {
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
	file_adaptors_uniswap_uniswap_v3_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*UniswapV3Adaptor_Swap)(nil),
		(*UniswapV3Adaptor_OracleSwap)(nil),
		(*UniswapV3Adaptor_RevokeApproval)(nil),
		(*UniswapV3Adaptor_OpenPosition_)(nil),
		(*UniswapV3Adaptor_ClosePosition_)(nil),
		(*UniswapV3Adaptor_AddToPosition_)(nil),
		(*UniswapV3Adaptor_TakeFromPosition_)(nil),
		(*UniswapV3Adaptor_CollectFees_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_adaptors_uniswap_uniswap_v3_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_adaptors_uniswap_uniswap_v3_proto_goTypes,
		DependencyIndexes: file_adaptors_uniswap_uniswap_v3_proto_depIdxs,
		MessageInfos:      file_adaptors_uniswap_uniswap_v3_proto_msgTypes,
	}.Build()
	File_adaptors_uniswap_uniswap_v3_proto = out.File
	file_adaptors_uniswap_uniswap_v3_proto_rawDesc = nil
	file_adaptors_uniswap_uniswap_v3_proto_goTypes = nil
	file_adaptors_uniswap_uniswap_v3_proto_depIdxs = nil
}
