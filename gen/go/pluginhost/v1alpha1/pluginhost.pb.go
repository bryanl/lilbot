// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: pluginhost/v1alpha1/pluginhost.proto

package v1alpha1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Type is are the types of options.
type ActionOption_Type int32

const (
	// Default value. If this is selected, the type will default to string.
	ActionOption_TYPE_UNSPECIFIED ActionOption_Type = 0
	// String option.
	ActionOption_TYPE_STRING ActionOption_Type = 1
	// Integer option.
	ActionOption_TYPE_INTEGER ActionOption_Type = 2
	// Boolean option.
	ActionOption_TYPE_BOOLEAN ActionOption_Type = 3
)

// Enum value maps for ActionOption_Type.
var (
	ActionOption_Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "TYPE_STRING",
		2: "TYPE_INTEGER",
		3: "TYPE_BOOLEAN",
	}
	ActionOption_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED": 0,
		"TYPE_STRING":      1,
		"TYPE_INTEGER":     2,
		"TYPE_BOOLEAN":     3,
	}
)

func (x ActionOption_Type) Enum() *ActionOption_Type {
	p := new(ActionOption_Type)
	*p = x
	return p
}

func (x ActionOption_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ActionOption_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_pluginhost_v1alpha1_pluginhost_proto_enumTypes[0].Descriptor()
}

func (ActionOption_Type) Type() protoreflect.EnumType {
	return &file_pluginhost_v1alpha1_pluginhost_proto_enumTypes[0]
}

func (x ActionOption_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ActionOption_Type.Descriptor instead.
func (ActionOption_Type) EnumDescriptor() ([]byte, []int) {
	return file_pluginhost_v1alpha1_pluginhost_proto_rawDescGZIP(), []int{1, 0}
}

// Action is an action that can be called.
type Action struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name is the name of the action.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Options for the command. Required options must be listed first.
	Options []*ActionOption `protobuf:"bytes,2,rep,name=options,proto3" json:"options,omitempty"`
}

func (x *Action) Reset() {
	*x = Action{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pluginhost_v1alpha1_pluginhost_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Action) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Action) ProtoMessage() {}

func (x *Action) ProtoReflect() protoreflect.Message {
	mi := &file_pluginhost_v1alpha1_pluginhost_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Action.ProtoReflect.Descriptor instead.
func (*Action) Descriptor() ([]byte, []int) {
	return file_pluginhost_v1alpha1_pluginhost_proto_rawDescGZIP(), []int{0}
}

func (x *Action) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Action) GetOptions() []*ActionOption {
	if x != nil {
		return x.Options
	}
	return nil
}

// Option is an action option.
type ActionOption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name is the name of the option.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Description is the description of the option.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// Required sets the required status.
	Required bool `protobuf:"varint,3,opt,name=required,proto3" json:"required,omitempty"`
	// Type is the type of the option.
	Type ActionOption_Type `protobuf:"varint,4,opt,name=type,proto3,enum=pluginhost.v1alpha1.ActionOption_Type" json:"type,omitempty"` // enum
}

func (x *ActionOption) Reset() {
	*x = ActionOption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pluginhost_v1alpha1_pluginhost_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionOption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionOption) ProtoMessage() {}

func (x *ActionOption) ProtoReflect() protoreflect.Message {
	mi := &file_pluginhost_v1alpha1_pluginhost_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionOption.ProtoReflect.Descriptor instead.
func (*ActionOption) Descriptor() ([]byte, []int) {
	return file_pluginhost_v1alpha1_pluginhost_proto_rawDescGZIP(), []int{1}
}

func (x *ActionOption) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ActionOption) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ActionOption) GetRequired() bool {
	if x != nil {
		return x.Required
	}
	return false
}

func (x *ActionOption) GetType() ActionOption_Type {
	if x != nil {
		return x.Type
	}
	return ActionOption_TYPE_UNSPECIFIED
}

// Register a plugin.
type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name is the name of the plugin.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Actions are actions the plugin contains.
	Actions map[string]*Action `protobuf:"bytes,2,rep,name=actions,proto3" json:"actions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pluginhost_v1alpha1_pluginhost_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pluginhost_v1alpha1_pluginhost_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_pluginhost_v1alpha1_pluginhost_proto_rawDescGZIP(), []int{2}
}

func (x *RegisterRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RegisterRequest) GetActions() map[string]*Action {
	if x != nil {
		return x.Actions
	}
	return nil
}

type RegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RegisterResponse) Reset() {
	*x = RegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pluginhost_v1alpha1_pluginhost_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResponse) ProtoMessage() {}

func (x *RegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pluginhost_v1alpha1_pluginhost_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResponse.ProtoReflect.Descriptor instead.
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return file_pluginhost_v1alpha1_pluginhost_proto_rawDescGZIP(), []int{3}
}

var File_pluginhost_v1alpha1_pluginhost_proto protoreflect.FileDescriptor

var file_pluginhost_v1alpha1_pluginhost_proto_rawDesc = []byte{
	0x0a, 0x24, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x68, 0x6f, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x68, 0x6f, 0x73, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x68, 0x6f,
	0x73, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x59, 0x0a, 0x06, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x68, 0x6f, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xef, 0x01, 0x0a, 0x0c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08,
	0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12, 0x3a, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x68,
	0x6f, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x22, 0x51, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x52, 0x49, 0x4e,
	0x47, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x54, 0x45,
	0x47, 0x45, 0x52, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x4f,
	0x4f, 0x4c, 0x45, 0x41, 0x4e, 0x10, 0x03, 0x22, 0xcb, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x4b, 0x0a, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x31, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x68, 0x6f, 0x73, 0x74, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x57, 0x0a, 0x0c,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x31,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x68, 0x6f, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x12, 0x0a, 0x10, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x74, 0x0a, 0x17, 0x4c, 0x69, 0x6c,
	0x62, 0x6f, 0x74, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x48, 0x6f, 0x73, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x59, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x12, 0x24, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x68, 0x6f, 0x73, 0x74, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x68,
	0x6f, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x72,
	0x79, 0x61, 0x6e, 0x6c, 0x2f, 0x6c, 0x69, 0x6c, 0x62, 0x6f, 0x74, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x67, 0x6f, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pluginhost_v1alpha1_pluginhost_proto_rawDescOnce sync.Once
	file_pluginhost_v1alpha1_pluginhost_proto_rawDescData = file_pluginhost_v1alpha1_pluginhost_proto_rawDesc
)

func file_pluginhost_v1alpha1_pluginhost_proto_rawDescGZIP() []byte {
	file_pluginhost_v1alpha1_pluginhost_proto_rawDescOnce.Do(func() {
		file_pluginhost_v1alpha1_pluginhost_proto_rawDescData = protoimpl.X.CompressGZIP(file_pluginhost_v1alpha1_pluginhost_proto_rawDescData)
	})
	return file_pluginhost_v1alpha1_pluginhost_proto_rawDescData
}

var file_pluginhost_v1alpha1_pluginhost_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pluginhost_v1alpha1_pluginhost_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pluginhost_v1alpha1_pluginhost_proto_goTypes = []interface{}{
	(ActionOption_Type)(0),   // 0: pluginhost.v1alpha1.ActionOption.Type
	(*Action)(nil),           // 1: pluginhost.v1alpha1.Action
	(*ActionOption)(nil),     // 2: pluginhost.v1alpha1.ActionOption
	(*RegisterRequest)(nil),  // 3: pluginhost.v1alpha1.RegisterRequest
	(*RegisterResponse)(nil), // 4: pluginhost.v1alpha1.RegisterResponse
	nil,                      // 5: pluginhost.v1alpha1.RegisterRequest.ActionsEntry
}
var file_pluginhost_v1alpha1_pluginhost_proto_depIdxs = []int32{
	2, // 0: pluginhost.v1alpha1.Action.options:type_name -> pluginhost.v1alpha1.ActionOption
	0, // 1: pluginhost.v1alpha1.ActionOption.type:type_name -> pluginhost.v1alpha1.ActionOption.Type
	5, // 2: pluginhost.v1alpha1.RegisterRequest.actions:type_name -> pluginhost.v1alpha1.RegisterRequest.ActionsEntry
	1, // 3: pluginhost.v1alpha1.RegisterRequest.ActionsEntry.value:type_name -> pluginhost.v1alpha1.Action
	3, // 4: pluginhost.v1alpha1.LilbotPluginHostService.Register:input_type -> pluginhost.v1alpha1.RegisterRequest
	4, // 5: pluginhost.v1alpha1.LilbotPluginHostService.Register:output_type -> pluginhost.v1alpha1.RegisterResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_pluginhost_v1alpha1_pluginhost_proto_init() }
func file_pluginhost_v1alpha1_pluginhost_proto_init() {
	if File_pluginhost_v1alpha1_pluginhost_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pluginhost_v1alpha1_pluginhost_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Action); i {
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
		file_pluginhost_v1alpha1_pluginhost_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionOption); i {
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
		file_pluginhost_v1alpha1_pluginhost_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRequest); i {
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
		file_pluginhost_v1alpha1_pluginhost_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterResponse); i {
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
			RawDescriptor: file_pluginhost_v1alpha1_pluginhost_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pluginhost_v1alpha1_pluginhost_proto_goTypes,
		DependencyIndexes: file_pluginhost_v1alpha1_pluginhost_proto_depIdxs,
		EnumInfos:         file_pluginhost_v1alpha1_pluginhost_proto_enumTypes,
		MessageInfos:      file_pluginhost_v1alpha1_pluginhost_proto_msgTypes,
	}.Build()
	File_pluginhost_v1alpha1_pluginhost_proto = out.File
	file_pluginhost_v1alpha1_pluginhost_proto_rawDesc = nil
	file_pluginhost_v1alpha1_pluginhost_proto_goTypes = nil
	file_pluginhost_v1alpha1_pluginhost_proto_depIdxs = nil
}
