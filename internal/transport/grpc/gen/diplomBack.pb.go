// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: diplomBack.proto

package mainBack

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

type EmptyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyRequest) Reset() {
	*x = EmptyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_diplomBack_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyRequest) ProtoMessage() {}

func (x *EmptyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_diplomBack_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyRequest.ProtoReflect.Descriptor instead.
func (*EmptyRequest) Descriptor() ([]byte, []int) {
	return file_diplomBack_proto_rawDescGZIP(), []int{0}
}

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login    string `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_diplomBack_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_diplomBack_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_diplomBack_proto_rawDescGZIP(), []int{1}
}

func (x *LoginRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Access  string `protobuf:"bytes,1,opt,name=access,proto3" json:"access,omitempty"`
	Refresh string `protobuf:"bytes,2,opt,name=refresh,proto3" json:"refresh,omitempty"`
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_diplomBack_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_diplomBack_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_diplomBack_proto_rawDescGZIP(), []int{2}
}

func (x *LoginResponse) GetAccess() string {
	if x != nil {
		return x.Access
	}
	return ""
}

func (x *LoginResponse) GetRefresh() string {
	if x != nil {
		return x.Refresh
	}
	return ""
}

type RequestsCount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int32 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *RequestsCount) Reset() {
	*x = RequestsCount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_diplomBack_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestsCount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestsCount) ProtoMessage() {}

func (x *RequestsCount) ProtoReflect() protoreflect.Message {
	mi := &file_diplomBack_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestsCount.ProtoReflect.Descriptor instead.
func (*RequestsCount) Descriptor() ([]byte, []int) {
	return file_diplomBack_proto_rawDescGZIP(), []int{3}
}

func (x *RequestsCount) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_diplomBack_proto protoreflect.FileDescriptor

var file_diplomBack_proto_rawDesc = []byte{
	0x0a, 0x10, 0x64, 0x69, 0x70, 0x6c, 0x6f, 0x6d, 0x42, 0x61, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x08, 0x6d, 0x61, 0x69, 0x6e, 0x42, 0x61, 0x63, 0x6b, 0x22, 0x0e, 0x0a, 0x0c,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x40, 0x0a, 0x0c,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67,
	0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x41,
	0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x22, 0x25, 0x0a, 0x0d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0x81, 0x01, 0x0a, 0x08, 0x6d, 0x61, 0x69,
	0x6e, 0x42, 0x61, 0x63, 0x6b, 0x12, 0x38, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x16,
	0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x42, 0x61, 0x63, 0x6b, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x42, 0x61, 0x63,
	0x6b, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3b, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x73, 0x12, 0x16, 0x2e, 0x6d, 0x61,
	0x69, 0x6e, 0x42, 0x61, 0x63, 0x6b, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x42, 0x61, 0x63, 0x6b, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x22, 0x5a, 0x20,
	0x64, 0x69, 0x70, 0x6c, 0x6f, 0x6d, 0x50, 0x6c, 0x75, 0x67, 0x42, 0x61, 0x63, 0x6b, 0x2e, 0x6d,
	0x61, 0x69, 0x6e, 0x42, 0x61, 0x63, 0x6b, 0x3b, 0x6d, 0x61, 0x69, 0x6e, 0x42, 0x61, 0x63, 0x6b,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_diplomBack_proto_rawDescOnce sync.Once
	file_diplomBack_proto_rawDescData = file_diplomBack_proto_rawDesc
)

func file_diplomBack_proto_rawDescGZIP() []byte {
	file_diplomBack_proto_rawDescOnce.Do(func() {
		file_diplomBack_proto_rawDescData = protoimpl.X.CompressGZIP(file_diplomBack_proto_rawDescData)
	})
	return file_diplomBack_proto_rawDescData
}

var file_diplomBack_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_diplomBack_proto_goTypes = []interface{}{
	(*EmptyRequest)(nil),  // 0: mainBack.EmptyRequest
	(*LoginRequest)(nil),  // 1: mainBack.LoginRequest
	(*LoginResponse)(nil), // 2: mainBack.LoginResponse
	(*RequestsCount)(nil), // 3: mainBack.RequestsCount
}
var file_diplomBack_proto_depIdxs = []int32{
	1, // 0: mainBack.mainBack.Login:input_type -> mainBack.LoginRequest
	0, // 1: mainBack.mainBack.GetRouts:input_type -> mainBack.EmptyRequest
	2, // 2: mainBack.mainBack.Login:output_type -> mainBack.LoginResponse
	3, // 3: mainBack.mainBack.GetRouts:output_type -> mainBack.RequestsCount
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_diplomBack_proto_init() }
func file_diplomBack_proto_init() {
	if File_diplomBack_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_diplomBack_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyRequest); i {
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
		file_diplomBack_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_diplomBack_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResponse); i {
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
		file_diplomBack_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestsCount); i {
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
			RawDescriptor: file_diplomBack_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_diplomBack_proto_goTypes,
		DependencyIndexes: file_diplomBack_proto_depIdxs,
		MessageInfos:      file_diplomBack_proto_msgTypes,
	}.Build()
	File_diplomBack_proto = out.File
	file_diplomBack_proto_rawDesc = nil
	file_diplomBack_proto_goTypes = nil
	file_diplomBack_proto_depIdxs = nil
}
