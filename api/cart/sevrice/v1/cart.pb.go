// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.11.4
// source: api/cart/sevrice/v1/cart.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type DeleteCartItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ItemId int64 `protobuf:"varint,2,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
}

func (x *DeleteCartItemRequest) Reset() {
	*x = DeleteCartItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cart_sevrice_v1_cart_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCartItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCartItemRequest) ProtoMessage() {}

func (x *DeleteCartItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_cart_sevrice_v1_cart_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCartItemRequest.ProtoReflect.Descriptor instead.
func (*DeleteCartItemRequest) Descriptor() ([]byte, []int) {
	return file_api_cart_sevrice_v1_cart_proto_rawDescGZIP(), []int{0}
}

func (x *DeleteCartItemRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *DeleteCartItemRequest) GetItemId() int64 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

type DeleteCartItemReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteCartItemReply) Reset() {
	*x = DeleteCartItemReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cart_sevrice_v1_cart_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCartItemReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCartItemReply) ProtoMessage() {}

func (x *DeleteCartItemReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_cart_sevrice_v1_cart_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCartItemReply.ProtoReflect.Descriptor instead.
func (*DeleteCartItemReply) Descriptor() ([]byte, []int) {
	return file_api_cart_sevrice_v1_cart_proto_rawDescGZIP(), []int{1}
}

func (x *DeleteCartItemReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type DecrItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ItemId    int64 `protobuf:"varint,2,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	ChangeNum int64 `protobuf:"varint,3,opt,name=change_num,json=changeNum,proto3" json:"change_num,omitempty"`
}

func (x *DecrItemRequest) Reset() {
	*x = DecrItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cart_sevrice_v1_cart_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecrItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecrItemRequest) ProtoMessage() {}

func (x *DecrItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_cart_sevrice_v1_cart_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecrItemRequest.ProtoReflect.Descriptor instead.
func (*DecrItemRequest) Descriptor() ([]byte, []int) {
	return file_api_cart_sevrice_v1_cart_proto_rawDescGZIP(), []int{2}
}

func (x *DecrItemRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *DecrItemRequest) GetItemId() int64 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *DecrItemRequest) GetChangeNum() int64 {
	if x != nil {
		return x.ChangeNum
	}
	return 0
}

type DecrItemReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DecrItemReply) Reset() {
	*x = DecrItemReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cart_sevrice_v1_cart_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecrItemReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecrItemReply) ProtoMessage() {}

func (x *DecrItemReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_cart_sevrice_v1_cart_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecrItemReply.ProtoReflect.Descriptor instead.
func (*DecrItemReply) Descriptor() ([]byte, []int) {
	return file_api_cart_sevrice_v1_cart_proto_rawDescGZIP(), []int{3}
}

func (x *DecrItemReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type IncrItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ItemId    int64 `protobuf:"varint,2,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	ChangeNum int64 `protobuf:"varint,3,opt,name=change_num,json=changeNum,proto3" json:"change_num,omitempty"`
}

func (x *IncrItemRequest) Reset() {
	*x = IncrItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cart_sevrice_v1_cart_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IncrItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncrItemRequest) ProtoMessage() {}

func (x *IncrItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_cart_sevrice_v1_cart_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncrItemRequest.ProtoReflect.Descriptor instead.
func (*IncrItemRequest) Descriptor() ([]byte, []int) {
	return file_api_cart_sevrice_v1_cart_proto_rawDescGZIP(), []int{4}
}

func (x *IncrItemRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *IncrItemRequest) GetItemId() int64 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *IncrItemRequest) GetChangeNum() int64 {
	if x != nil {
		return x.ChangeNum
	}
	return 0
}

type IncrItemReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *IncrItemReply) Reset() {
	*x = IncrItemReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cart_sevrice_v1_cart_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IncrItemReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncrItemReply) ProtoMessage() {}

func (x *IncrItemReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_cart_sevrice_v1_cart_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncrItemReply.ProtoReflect.Descriptor instead.
func (*IncrItemReply) Descriptor() ([]byte, []int) {
	return file_api_cart_sevrice_v1_cart_proto_rawDescGZIP(), []int{5}
}

func (x *IncrItemReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type CreateCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId    int64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ProductId int64 `protobuf:"varint,3,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	SizeId    int64 `protobuf:"varint,4,opt,name=size_id,json=sizeId,proto3" json:"size_id,omitempty"`
	Num       int64 `protobuf:"varint,5,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *CreateCartRequest) Reset() {
	*x = CreateCartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cart_sevrice_v1_cart_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCartRequest) ProtoMessage() {}

func (x *CreateCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_cart_sevrice_v1_cart_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCartRequest.ProtoReflect.Descriptor instead.
func (*CreateCartRequest) Descriptor() ([]byte, []int) {
	return file_api_cart_sevrice_v1_cart_proto_rawDescGZIP(), []int{6}
}

func (x *CreateCartRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CreateCartRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateCartRequest) GetProductId() int64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *CreateCartRequest) GetSizeId() int64 {
	if x != nil {
		return x.SizeId
	}
	return 0
}

func (x *CreateCartRequest) GetNum() int64 {
	if x != nil {
		return x.Num
	}
	return 0
}

type CreateCartReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CreateCartReply) Reset() {
	*x = CreateCartReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cart_sevrice_v1_cart_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCartReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCartReply) ProtoMessage() {}

func (x *CreateCartReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_cart_sevrice_v1_cart_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCartReply.ProtoReflect.Descriptor instead.
func (*CreateCartReply) Descriptor() ([]byte, []int) {
	return file_api_cart_sevrice_v1_cart_proto_rawDescGZIP(), []int{7}
}

func (x *CreateCartReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type CleanCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CleanCartRequest) Reset() {
	*x = CleanCartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cart_sevrice_v1_cart_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CleanCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CleanCartRequest) ProtoMessage() {}

func (x *CleanCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_cart_sevrice_v1_cart_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CleanCartRequest.ProtoReflect.Descriptor instead.
func (*CleanCartRequest) Descriptor() ([]byte, []int) {
	return file_api_cart_sevrice_v1_cart_proto_rawDescGZIP(), []int{8}
}

func (x *CleanCartRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type CleanCartReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CleanCartReply) Reset() {
	*x = CleanCartReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cart_sevrice_v1_cart_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CleanCartReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CleanCartReply) ProtoMessage() {}

func (x *CleanCartReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_cart_sevrice_v1_cart_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CleanCartReply.ProtoReflect.Descriptor instead.
func (*CleanCartReply) Descriptor() ([]byte, []int) {
	return file_api_cart_sevrice_v1_cart_proto_rawDescGZIP(), []int{9}
}

func (x *CleanCartReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetCartRequest) Reset() {
	*x = GetCartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cart_sevrice_v1_cart_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCartRequest) ProtoMessage() {}

func (x *GetCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_cart_sevrice_v1_cart_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCartRequest.ProtoReflect.Descriptor instead.
func (*GetCartRequest) Descriptor() ([]byte, []int) {
	return file_api_cart_sevrice_v1_cart_proto_rawDescGZIP(), []int{10}
}

func (x *GetCartRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetCartReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cart []*GetCartReply_Cart `protobuf:"bytes,1,rep,name=cart,proto3" json:"cart,omitempty"`
}

func (x *GetCartReply) Reset() {
	*x = GetCartReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cart_sevrice_v1_cart_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCartReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCartReply) ProtoMessage() {}

func (x *GetCartReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_cart_sevrice_v1_cart_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCartReply.ProtoReflect.Descriptor instead.
func (*GetCartReply) Descriptor() ([]byte, []int) {
	return file_api_cart_sevrice_v1_cart_proto_rawDescGZIP(), []int{11}
}

func (x *GetCartReply) GetCart() []*GetCartReply_Cart {
	if x != nil {
		return x.Cart
	}
	return nil
}

type GetCartReply_Cart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ProductId int64 `protobuf:"varint,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	SizeId    int64 `protobuf:"varint,3,opt,name=size_id,json=sizeId,proto3" json:"size_id,omitempty"`
	Num       int64 `protobuf:"varint,4,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *GetCartReply_Cart) Reset() {
	*x = GetCartReply_Cart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cart_sevrice_v1_cart_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCartReply_Cart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCartReply_Cart) ProtoMessage() {}

func (x *GetCartReply_Cart) ProtoReflect() protoreflect.Message {
	mi := &file_api_cart_sevrice_v1_cart_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCartReply_Cart.ProtoReflect.Descriptor instead.
func (*GetCartReply_Cart) Descriptor() ([]byte, []int) {
	return file_api_cart_sevrice_v1_cart_proto_rawDescGZIP(), []int{11, 0}
}

func (x *GetCartReply_Cart) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetCartReply_Cart) GetProductId() int64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *GetCartReply_Cart) GetSizeId() int64 {
	if x != nil {
		return x.SizeId
	}
	return 0
}

func (x *GetCartReply_Cart) GetNum() int64 {
	if x != nil {
		return x.Num
	}
	return 0
}

var File_api_cart_sevrice_v1_cart_proto protoreflect.FileDescriptor

var file_api_cart_sevrice_v1_cart_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x2f, 0x73, 0x65, 0x76, 0x72, 0x69,
	0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x13, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x73, 0x65, 0x76, 0x72, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x49, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x72,
	0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x22, 0x2f,
	0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x62, 0x0a, 0x0f, 0x44, 0x65, 0x63, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x69,
	0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x69, 0x74,
	0x65, 0x6d, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x6e,
	0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x4e, 0x75, 0x6d, 0x22, 0x29, 0x0a, 0x0d, 0x44, 0x65, 0x63, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x62,
	0x0a, 0x0f, 0x49, 0x6e, 0x63, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74,
	0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x69, 0x74, 0x65,
	0x6d, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x6e, 0x75,
	0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4e,
	0x75, 0x6d, 0x22, 0x29, 0x0a, 0x0d, 0x49, 0x6e, 0x63, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x86, 0x01,
	0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x73,
	0x69, 0x7a, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x69,
	0x7a, 0x65, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x22, 0x2b, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x22, 0x0a, 0x10, 0x43, 0x6c, 0x65, 0x61, 0x6e, 0x43, 0x61, 0x72, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2a, 0x0a, 0x0e, 0x43, 0x6c, 0x65, 0x61, 0x6e,
	0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0xac, 0x01, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72,
	0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x3a, 0x0a, 0x04, 0x63, 0x61, 0x72, 0x74, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e,
	0x73, 0x65, 0x76, 0x72, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61,
	0x72, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x52, 0x04, 0x63, 0x61,
	0x72, 0x74, 0x1a, 0x60, 0x0a, 0x04, 0x43, 0x61, 0x72, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x69, 0x7a,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x69, 0x7a, 0x65,
	0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x03, 0x6e, 0x75, 0x6d, 0x32, 0xb4, 0x04, 0x0a, 0x04, 0x43, 0x61, 0x72, 0x74, 0x12, 0x6c, 0x0a,
	0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x12, 0x26, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x73, 0x65, 0x76, 0x72, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x73,
	0x65, 0x76, 0x72, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0a, 0x22, 0x05, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x54, 0x0a, 0x08, 0x49,
	0x6e, 0x63, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61,
	0x72, 0x74, 0x2e, 0x73, 0x65, 0x76, 0x72, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e,
	0x63, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x73, 0x65, 0x76, 0x72, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x63, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x54, 0x0a, 0x08, 0x44, 0x65, 0x63, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x24, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x73, 0x65, 0x76, 0x72, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x63, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x73,
	0x65, 0x76, 0x72, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x63, 0x72, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x66, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x49, 0x74, 0x65, 0x6d, 0x42, 0x79, 0x49, 0x64, 0x12, 0x2a, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x61, 0x72, 0x74, 0x2e, 0x73, 0x65, 0x76, 0x72, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x72, 0x74,
	0x2e, 0x73, 0x65, 0x76, 0x72, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x57, 0x0a, 0x09, 0x43, 0x6c, 0x65, 0x61, 0x6e, 0x43, 0x61, 0x72, 0x74, 0x12, 0x25, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x73, 0x65, 0x76, 0x72, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x6c, 0x65, 0x61, 0x6e, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x73,
	0x65, 0x76, 0x72, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x65, 0x61, 0x6e, 0x43,
	0x61, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x51, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x43,
	0x61, 0x72, 0x74, 0x12, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x73,
	0x65, 0x76, 0x72, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x61, 0x72, 0x74, 0x2e, 0x73, 0x65, 0x76, 0x72, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x34, 0x0a, 0x13, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x73, 0x65, 0x76, 0x72, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x50, 0x01, 0x5a, 0x1b, 0x6d, 0x61, 0x6c, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x61, 0x72, 0x74, 0x2f, 0x73, 0x65, 0x76, 0x72, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_cart_sevrice_v1_cart_proto_rawDescOnce sync.Once
	file_api_cart_sevrice_v1_cart_proto_rawDescData = file_api_cart_sevrice_v1_cart_proto_rawDesc
)

func file_api_cart_sevrice_v1_cart_proto_rawDescGZIP() []byte {
	file_api_cart_sevrice_v1_cart_proto_rawDescOnce.Do(func() {
		file_api_cart_sevrice_v1_cart_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_cart_sevrice_v1_cart_proto_rawDescData)
	})
	return file_api_cart_sevrice_v1_cart_proto_rawDescData
}

var file_api_cart_sevrice_v1_cart_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_api_cart_sevrice_v1_cart_proto_goTypes = []interface{}{
	(*DeleteCartItemRequest)(nil), // 0: api.cart.sevrice.v1.DeleteCartItemRequest
	(*DeleteCartItemReply)(nil),   // 1: api.cart.sevrice.v1.DeleteCartItemReply
	(*DecrItemRequest)(nil),       // 2: api.cart.sevrice.v1.DecrItemRequest
	(*DecrItemReply)(nil),         // 3: api.cart.sevrice.v1.DecrItemReply
	(*IncrItemRequest)(nil),       // 4: api.cart.sevrice.v1.IncrItemRequest
	(*IncrItemReply)(nil),         // 5: api.cart.sevrice.v1.IncrItemReply
	(*CreateCartRequest)(nil),     // 6: api.cart.sevrice.v1.CreateCartRequest
	(*CreateCartReply)(nil),       // 7: api.cart.sevrice.v1.CreateCartReply
	(*CleanCartRequest)(nil),      // 8: api.cart.sevrice.v1.CleanCartRequest
	(*CleanCartReply)(nil),        // 9: api.cart.sevrice.v1.CleanCartReply
	(*GetCartRequest)(nil),        // 10: api.cart.sevrice.v1.GetCartRequest
	(*GetCartReply)(nil),          // 11: api.cart.sevrice.v1.GetCartReply
	(*GetCartReply_Cart)(nil),     // 12: api.cart.sevrice.v1.GetCartReply.Cart
}
var file_api_cart_sevrice_v1_cart_proto_depIdxs = []int32{
	12, // 0: api.cart.sevrice.v1.GetCartReply.cart:type_name -> api.cart.sevrice.v1.GetCartReply.Cart
	6,  // 1: api.cart.sevrice.v1.Cart.CreateCart:input_type -> api.cart.sevrice.v1.CreateCartRequest
	4,  // 2: api.cart.sevrice.v1.Cart.IncrItem:input_type -> api.cart.sevrice.v1.IncrItemRequest
	2,  // 3: api.cart.sevrice.v1.Cart.DecrItem:input_type -> api.cart.sevrice.v1.DecrItemRequest
	0,  // 4: api.cart.sevrice.v1.Cart.DeleteItemById:input_type -> api.cart.sevrice.v1.DeleteCartItemRequest
	8,  // 5: api.cart.sevrice.v1.Cart.CleanCart:input_type -> api.cart.sevrice.v1.CleanCartRequest
	10, // 6: api.cart.sevrice.v1.Cart.GetCart:input_type -> api.cart.sevrice.v1.GetCartRequest
	7,  // 7: api.cart.sevrice.v1.Cart.CreateCart:output_type -> api.cart.sevrice.v1.CreateCartReply
	5,  // 8: api.cart.sevrice.v1.Cart.IncrItem:output_type -> api.cart.sevrice.v1.IncrItemReply
	3,  // 9: api.cart.sevrice.v1.Cart.DecrItem:output_type -> api.cart.sevrice.v1.DecrItemReply
	1,  // 10: api.cart.sevrice.v1.Cart.DeleteItemById:output_type -> api.cart.sevrice.v1.DeleteCartItemReply
	9,  // 11: api.cart.sevrice.v1.Cart.CleanCart:output_type -> api.cart.sevrice.v1.CleanCartReply
	11, // 12: api.cart.sevrice.v1.Cart.GetCart:output_type -> api.cart.sevrice.v1.GetCartReply
	7,  // [7:13] is the sub-list for method output_type
	1,  // [1:7] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_api_cart_sevrice_v1_cart_proto_init() }
func file_api_cart_sevrice_v1_cart_proto_init() {
	if File_api_cart_sevrice_v1_cart_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_cart_sevrice_v1_cart_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCartItemRequest); i {
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
		file_api_cart_sevrice_v1_cart_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCartItemReply); i {
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
		file_api_cart_sevrice_v1_cart_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecrItemRequest); i {
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
		file_api_cart_sevrice_v1_cart_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecrItemReply); i {
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
		file_api_cart_sevrice_v1_cart_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IncrItemRequest); i {
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
		file_api_cart_sevrice_v1_cart_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IncrItemReply); i {
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
		file_api_cart_sevrice_v1_cart_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCartRequest); i {
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
		file_api_cart_sevrice_v1_cart_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCartReply); i {
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
		file_api_cart_sevrice_v1_cart_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CleanCartRequest); i {
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
		file_api_cart_sevrice_v1_cart_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CleanCartReply); i {
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
		file_api_cart_sevrice_v1_cart_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCartRequest); i {
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
		file_api_cart_sevrice_v1_cart_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCartReply); i {
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
		file_api_cart_sevrice_v1_cart_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCartReply_Cart); i {
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
			RawDescriptor: file_api_cart_sevrice_v1_cart_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_cart_sevrice_v1_cart_proto_goTypes,
		DependencyIndexes: file_api_cart_sevrice_v1_cart_proto_depIdxs,
		MessageInfos:      file_api_cart_sevrice_v1_cart_proto_msgTypes,
	}.Build()
	File_api_cart_sevrice_v1_cart_proto = out.File
	file_api_cart_sevrice_v1_cart_proto_rawDesc = nil
	file_api_cart_sevrice_v1_cart_proto_goTypes = nil
	file_api_cart_sevrice_v1_cart_proto_depIdxs = nil
}