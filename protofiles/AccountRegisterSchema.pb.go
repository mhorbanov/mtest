// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.0
// source: AccountRegisterSchema.proto

package protofiles

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

type Account struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId      string      `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	RequestedBy    string      `protobuf:"bytes,2,opt,name=requested_by,json=requestedBy,proto3" json:"requested_by,omitempty"`
	Timestamp      string      `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	BeatId         string      `protobuf:"bytes,5,opt,name=beat_id,json=beatId,proto3" json:"beat_id,omitempty"`
	PAPIGroup      string      `protobuf:"bytes,7,opt,name=PAPI_group,json=PAPIGroup,proto3" json:"PAPI_group,omitempty"`
	Environment    string      `protobuf:"bytes,8,opt,name=environment,proto3" json:"environment,omitempty"`
	Resources      []*Resource `protobuf:"bytes,9,rep,name=resources,proto3" json:"resources,omitempty"`
	RequestedRoles []string    `protobuf:"bytes,11,rep,name=requested_roles,json=requestedRoles,proto3" json:"requested_roles,omitempty"`
}

func (x *Account) Reset() {
	*x = Account{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AccountRegisterSchema_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_AccountRegisterSchema_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account.ProtoReflect.Descriptor instead.
func (*Account) Descriptor() ([]byte, []int) {
	return file_AccountRegisterSchema_proto_rawDescGZIP(), []int{0}
}

func (x *Account) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *Account) GetRequestedBy() string {
	if x != nil {
		return x.RequestedBy
	}
	return ""
}

func (x *Account) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *Account) GetBeatId() string {
	if x != nil {
		return x.BeatId
	}
	return ""
}

func (x *Account) GetPAPIGroup() string {
	if x != nil {
		return x.PAPIGroup
	}
	return ""
}

func (x *Account) GetEnvironment() string {
	if x != nil {
		return x.Environment
	}
	return ""
}

func (x *Account) GetResources() []*Resource {
	if x != nil {
		return x.Resources
	}
	return nil
}

func (x *Account) GetRequestedRoles() []string {
	if x != nil {
		return x.RequestedRoles
	}
	return nil
}

type Resource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId          string `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	ResourceId         string `protobuf:"bytes,2,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	RequestedBy        string `protobuf:"bytes,3,opt,name=requested_by,json=requestedBy,proto3" json:"requested_by,omitempty"`
	RequestedTimestamp string `protobuf:"bytes,5,opt,name=requested_timestamp,json=requestedTimestamp,proto3" json:"requested_timestamp,omitempty"`
	LastUpdated        string `protobuf:"bytes,6,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	ResourceType       string `protobuf:"bytes,7,opt,name=resource_type,json=resourceType,proto3" json:"resource_type,omitempty"`
	ResourceName       string `protobuf:"bytes,8,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	ResourceArn        string `protobuf:"bytes,9,opt,name=resource_arn,json=resourceArn,proto3" json:"resource_arn,omitempty"`
	Fulfilled          bool   `protobuf:"varint,10,opt,name=fulfilled,proto3" json:"fulfilled,omitempty"`
}

func (x *Resource) Reset() {
	*x = Resource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AccountRegisterSchema_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resource) ProtoMessage() {}

func (x *Resource) ProtoReflect() protoreflect.Message {
	mi := &file_AccountRegisterSchema_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resource.ProtoReflect.Descriptor instead.
func (*Resource) Descriptor() ([]byte, []int) {
	return file_AccountRegisterSchema_proto_rawDescGZIP(), []int{1}
}

func (x *Resource) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *Resource) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *Resource) GetRequestedBy() string {
	if x != nil {
		return x.RequestedBy
	}
	return ""
}

func (x *Resource) GetRequestedTimestamp() string {
	if x != nil {
		return x.RequestedTimestamp
	}
	return ""
}

func (x *Resource) GetLastUpdated() string {
	if x != nil {
		return x.LastUpdated
	}
	return ""
}

func (x *Resource) GetResourceType() string {
	if x != nil {
		return x.ResourceType
	}
	return ""
}

func (x *Resource) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *Resource) GetResourceArn() string {
	if x != nil {
		return x.ResourceArn
	}
	return ""
}

func (x *Resource) GetFulfilled() bool {
	if x != nil {
		return x.Fulfilled
	}
	return false
}

var File_AccountRegisterSchema_proto protoreflect.FileDescriptor

var file_AccountRegisterSchema_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x95, 0x02,
	0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x17, 0x0a, 0x07, 0x62, 0x65, 0x61,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x65, 0x61, 0x74,
	0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x50, 0x41, 0x50, 0x49, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x50, 0x41, 0x50, 0x49, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x0f,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18,
	0x0b, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64,
	0x52, 0x6f, 0x6c, 0x65, 0x73, 0x22, 0xcc, 0x02, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f,
	0x62, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x2f, 0x0a, 0x13, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x12, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x61,
	0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23,
	0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x61, 0x72, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x41, 0x72, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x66, 0x69, 0x6c,
	0x6c, 0x65, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x66, 0x75, 0x6c, 0x66, 0x69,
	0x6c, 0x6c, 0x65, 0x64, 0x42, 0x03, 0x5a, 0x01, 0x2e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_AccountRegisterSchema_proto_rawDescOnce sync.Once
	file_AccountRegisterSchema_proto_rawDescData = file_AccountRegisterSchema_proto_rawDesc
)

func file_AccountRegisterSchema_proto_rawDescGZIP() []byte {
	file_AccountRegisterSchema_proto_rawDescOnce.Do(func() {
		file_AccountRegisterSchema_proto_rawDescData = protoimpl.X.CompressGZIP(file_AccountRegisterSchema_proto_rawDescData)
	})
	return file_AccountRegisterSchema_proto_rawDescData
}

var file_AccountRegisterSchema_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_AccountRegisterSchema_proto_goTypes = []interface{}{
	(*Account)(nil),  // 0: Account
	(*Resource)(nil), // 1: Resource
}
var file_AccountRegisterSchema_proto_depIdxs = []int32{
	1, // 0: Account.resources:type_name -> Resource
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_AccountRegisterSchema_proto_init() }
func file_AccountRegisterSchema_proto_init() {
	if File_AccountRegisterSchema_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_AccountRegisterSchema_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Account); i {
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
		file_AccountRegisterSchema_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resource); i {
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
			RawDescriptor: file_AccountRegisterSchema_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AccountRegisterSchema_proto_goTypes,
		DependencyIndexes: file_AccountRegisterSchema_proto_depIdxs,
		MessageInfos:      file_AccountRegisterSchema_proto_msgTypes,
	}.Build()
	File_AccountRegisterSchema_proto = out.File
	file_AccountRegisterSchema_proto_rawDesc = nil
	file_AccountRegisterSchema_proto_goTypes = nil
	file_AccountRegisterSchema_proto_depIdxs = nil
}
