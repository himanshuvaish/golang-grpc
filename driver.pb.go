// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: driver.proto

package golang_grpc

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

type Driver struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Wins  int32  `protobuf:"varint,3,opt,name=wins,proto3" json:"wins,omitempty"`
	Poles int32  `protobuf:"varint,4,opt,name=poles,proto3" json:"poles,omitempty"`
}

func (x *Driver) Reset() {
	*x = Driver{}
	if protoimpl.UnsafeEnabled {
		mi := &file_driver_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Driver) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Driver) ProtoMessage() {}

func (x *Driver) ProtoReflect() protoreflect.Message {
	mi := &file_driver_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Driver.ProtoReflect.Descriptor instead.
func (*Driver) Descriptor() ([]byte, []int) {
	return file_driver_proto_rawDescGZIP(), []int{0}
}

func (x *Driver) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Driver) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Driver) GetWins() int32 {
	if x != nil {
		return x.Wins
	}
	return 0
}

func (x *Driver) GetPoles() int32 {
	if x != nil {
		return x.Poles
	}
	return 0
}

type AddDriverRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Driver *Driver `protobuf:"bytes,1,opt,name=driver,proto3" json:"driver,omitempty"`
}

func (x *AddDriverRequest) Reset() {
	*x = AddDriverRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_driver_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddDriverRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddDriverRequest) ProtoMessage() {}

func (x *AddDriverRequest) ProtoReflect() protoreflect.Message {
	mi := &file_driver_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddDriverRequest.ProtoReflect.Descriptor instead.
func (*AddDriverRequest) Descriptor() ([]byte, []int) {
	return file_driver_proto_rawDescGZIP(), []int{1}
}

func (x *AddDriverRequest) GetDriver() *Driver {
	if x != nil {
		return x.Driver
	}
	return nil
}

type AddDriverResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Driver *Driver `protobuf:"bytes,1,opt,name=driver,proto3" json:"driver,omitempty"`
}

func (x *AddDriverResponse) Reset() {
	*x = AddDriverResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_driver_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddDriverResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddDriverResponse) ProtoMessage() {}

func (x *AddDriverResponse) ProtoReflect() protoreflect.Message {
	mi := &file_driver_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddDriverResponse.ProtoReflect.Descriptor instead.
func (*AddDriverResponse) Descriptor() ([]byte, []int) {
	return file_driver_proto_rawDescGZIP(), []int{2}
}

func (x *AddDriverResponse) GetDriver() *Driver {
	if x != nil {
		return x.Driver
	}
	return nil
}

type GetDriverRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetDriverRequest) Reset() {
	*x = GetDriverRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_driver_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDriverRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDriverRequest) ProtoMessage() {}

func (x *GetDriverRequest) ProtoReflect() protoreflect.Message {
	mi := &file_driver_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDriverRequest.ProtoReflect.Descriptor instead.
func (*GetDriverRequest) Descriptor() ([]byte, []int) {
	return file_driver_proto_rawDescGZIP(), []int{3}
}

func (x *GetDriverRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetDriverResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Driver *Driver `protobuf:"bytes,1,opt,name=driver,proto3" json:"driver,omitempty"`
}

func (x *GetDriverResponse) Reset() {
	*x = GetDriverResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_driver_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDriverResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDriverResponse) ProtoMessage() {}

func (x *GetDriverResponse) ProtoReflect() protoreflect.Message {
	mi := &file_driver_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDriverResponse.ProtoReflect.Descriptor instead.
func (*GetDriverResponse) Descriptor() ([]byte, []int) {
	return file_driver_proto_rawDescGZIP(), []int{4}
}

func (x *GetDriverResponse) GetDriver() *Driver {
	if x != nil {
		return x.Driver
	}
	return nil
}

type UpdateDriverRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Wins  int32 `protobuf:"varint,2,opt,name=wins,proto3" json:"wins,omitempty"`
	Poles int32 `protobuf:"varint,3,opt,name=poles,proto3" json:"poles,omitempty"`
}

func (x *UpdateDriverRequest) Reset() {
	*x = UpdateDriverRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_driver_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDriverRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDriverRequest) ProtoMessage() {}

func (x *UpdateDriverRequest) ProtoReflect() protoreflect.Message {
	mi := &file_driver_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDriverRequest.ProtoReflect.Descriptor instead.
func (*UpdateDriverRequest) Descriptor() ([]byte, []int) {
	return file_driver_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateDriverRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateDriverRequest) GetWins() int32 {
	if x != nil {
		return x.Wins
	}
	return 0
}

func (x *UpdateDriverRequest) GetPoles() int32 {
	if x != nil {
		return x.Poles
	}
	return 0
}

type UpdateDriverResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Driver *Driver `protobuf:"bytes,1,opt,name=driver,proto3" json:"driver,omitempty"`
}

func (x *UpdateDriverResponse) Reset() {
	*x = UpdateDriverResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_driver_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDriverResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDriverResponse) ProtoMessage() {}

func (x *UpdateDriverResponse) ProtoReflect() protoreflect.Message {
	mi := &file_driver_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDriverResponse.ProtoReflect.Descriptor instead.
func (*UpdateDriverResponse) Descriptor() ([]byte, []int) {
	return file_driver_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateDriverResponse) GetDriver() *Driver {
	if x != nil {
		return x.Driver
	}
	return nil
}

type DeleteDriverRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteDriverRequest) Reset() {
	*x = DeleteDriverRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_driver_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDriverRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDriverRequest) ProtoMessage() {}

func (x *DeleteDriverRequest) ProtoReflect() protoreflect.Message {
	mi := &file_driver_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDriverRequest.ProtoReflect.Descriptor instead.
func (*DeleteDriverRequest) Descriptor() ([]byte, []int) {
	return file_driver_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteDriverRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteDriverResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteDriverResponse) Reset() {
	*x = DeleteDriverResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_driver_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDriverResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDriverResponse) ProtoMessage() {}

func (x *DeleteDriverResponse) ProtoReflect() protoreflect.Message {
	mi := &file_driver_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDriverResponse.ProtoReflect.Descriptor instead.
func (*DeleteDriverResponse) Descriptor() ([]byte, []int) {
	return file_driver_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteDriverResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type ListDriversRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListDriversRequest) Reset() {
	*x = ListDriversRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_driver_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDriversRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDriversRequest) ProtoMessage() {}

func (x *ListDriversRequest) ProtoReflect() protoreflect.Message {
	mi := &file_driver_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDriversRequest.ProtoReflect.Descriptor instead.
func (*ListDriversRequest) Descriptor() ([]byte, []int) {
	return file_driver_proto_rawDescGZIP(), []int{9}
}

type ListDriversResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Drivers []*Driver `protobuf:"bytes,1,rep,name=drivers,proto3" json:"drivers,omitempty"`
}

func (x *ListDriversResponse) Reset() {
	*x = ListDriversResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_driver_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDriversResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDriversResponse) ProtoMessage() {}

func (x *ListDriversResponse) ProtoReflect() protoreflect.Message {
	mi := &file_driver_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDriversResponse.ProtoReflect.Descriptor instead.
func (*ListDriversResponse) Descriptor() ([]byte, []int) {
	return file_driver_proto_rawDescGZIP(), []int{10}
}

func (x *ListDriversResponse) GetDrivers() []*Driver {
	if x != nil {
		return x.Drivers
	}
	return nil
}

var File_driver_proto protoreflect.FileDescriptor

var file_driver_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x22, 0x56, 0x0a, 0x06, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x77, 0x69, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x77, 0x69, 0x6e, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6f, 0x6c, 0x65,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x6f, 0x6c, 0x65, 0x73, 0x22, 0x3a,
	0x0a, 0x10, 0x41, 0x64, 0x64, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x26, 0x0a, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x44, 0x72, 0x69, 0x76,
	0x65, 0x72, 0x52, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x22, 0x3b, 0x0a, 0x11, 0x41, 0x64,
	0x64, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x26, 0x0a, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x52,
	0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x22, 0x22, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x44, 0x72,
	0x69, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3b, 0x0a, 0x11, 0x47,
	0x65, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x26, 0x0a, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72,
	0x52, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x22, 0x4f, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x77, 0x69, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x77,
	0x69, 0x6e, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x70, 0x6f, 0x6c, 0x65, 0x73, 0x22, 0x3e, 0x0a, 0x14, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x26, 0x0a, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x44, 0x72, 0x69, 0x76, 0x65,
	0x72, 0x52, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x22, 0x25, 0x0a, 0x13, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x30, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x22, 0x14, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3f, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74,
	0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x28, 0x0a, 0x07, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72,
	0x52, 0x07, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x32, 0xfb, 0x02, 0x0a, 0x0d, 0x44, 0x72,
	0x69, 0x76, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x09, 0x41,
	0x64, 0x64, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x18, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65,
	0x72, 0x2e, 0x41, 0x64, 0x64, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x44,
	0x72, 0x69, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x42, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x18, 0x2e, 0x64,
	0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e,
	0x47, 0x65, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x72, 0x69,
	0x76, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x4b, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72,
	0x12, 0x1b, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e,
	0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x72, 0x69,
	0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x48, 0x0a,
	0x0b, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x12, 0x1a, 0x2e, 0x64,
	0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65,
	0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x3a, 0x5a, 0x38, 0x68, 0x74, 0x74, 0x70, 0x73,
	0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x69,
	0x6d, 0x61, 0x6e, 0x73, 0x68, 0x75, 0x76, 0x61, 0x69, 0x73, 0x68, 0x2f, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x3b, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x5f, 0x67,
	0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_driver_proto_rawDescOnce sync.Once
	file_driver_proto_rawDescData = file_driver_proto_rawDesc
)

func file_driver_proto_rawDescGZIP() []byte {
	file_driver_proto_rawDescOnce.Do(func() {
		file_driver_proto_rawDescData = protoimpl.X.CompressGZIP(file_driver_proto_rawDescData)
	})
	return file_driver_proto_rawDescData
}

var file_driver_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_driver_proto_goTypes = []interface{}{
	(*Driver)(nil),               // 0: driver.Driver
	(*AddDriverRequest)(nil),     // 1: driver.AddDriverRequest
	(*AddDriverResponse)(nil),    // 2: driver.AddDriverResponse
	(*GetDriverRequest)(nil),     // 3: driver.GetDriverRequest
	(*GetDriverResponse)(nil),    // 4: driver.GetDriverResponse
	(*UpdateDriverRequest)(nil),  // 5: driver.UpdateDriverRequest
	(*UpdateDriverResponse)(nil), // 6: driver.UpdateDriverResponse
	(*DeleteDriverRequest)(nil),  // 7: driver.DeleteDriverRequest
	(*DeleteDriverResponse)(nil), // 8: driver.DeleteDriverResponse
	(*ListDriversRequest)(nil),   // 9: driver.ListDriversRequest
	(*ListDriversResponse)(nil),  // 10: driver.ListDriversResponse
}
var file_driver_proto_depIdxs = []int32{
	0,  // 0: driver.AddDriverRequest.driver:type_name -> driver.Driver
	0,  // 1: driver.AddDriverResponse.driver:type_name -> driver.Driver
	0,  // 2: driver.GetDriverResponse.driver:type_name -> driver.Driver
	0,  // 3: driver.UpdateDriverResponse.driver:type_name -> driver.Driver
	0,  // 4: driver.ListDriversResponse.drivers:type_name -> driver.Driver
	1,  // 5: driver.DriverService.AddDriver:input_type -> driver.AddDriverRequest
	3,  // 6: driver.DriverService.GetDriver:input_type -> driver.GetDriverRequest
	5,  // 7: driver.DriverService.UpdateDriver:input_type -> driver.UpdateDriverRequest
	7,  // 8: driver.DriverService.DeleteDriver:input_type -> driver.DeleteDriverRequest
	9,  // 9: driver.DriverService.ListDrivers:input_type -> driver.ListDriversRequest
	2,  // 10: driver.DriverService.AddDriver:output_type -> driver.AddDriverResponse
	4,  // 11: driver.DriverService.GetDriver:output_type -> driver.GetDriverResponse
	6,  // 12: driver.DriverService.UpdateDriver:output_type -> driver.UpdateDriverResponse
	8,  // 13: driver.DriverService.DeleteDriver:output_type -> driver.DeleteDriverResponse
	10, // 14: driver.DriverService.ListDrivers:output_type -> driver.ListDriversResponse
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_driver_proto_init() }
func file_driver_proto_init() {
	if File_driver_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_driver_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Driver); i {
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
		file_driver_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddDriverRequest); i {
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
		file_driver_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddDriverResponse); i {
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
		file_driver_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDriverRequest); i {
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
		file_driver_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDriverResponse); i {
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
		file_driver_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDriverRequest); i {
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
		file_driver_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDriverResponse); i {
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
		file_driver_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDriverRequest); i {
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
		file_driver_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDriverResponse); i {
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
		file_driver_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDriversRequest); i {
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
		file_driver_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDriversResponse); i {
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
			RawDescriptor: file_driver_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_driver_proto_goTypes,
		DependencyIndexes: file_driver_proto_depIdxs,
		MessageInfos:      file_driver_proto_msgTypes,
	}.Build()
	File_driver_proto = out.File
	file_driver_proto_rawDesc = nil
	file_driver_proto_goTypes = nil
	file_driver_proto_depIdxs = nil
}
