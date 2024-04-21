// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: services/otp/otp.proto

package otp

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

type GenerateOtpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId string `protobuf:"bytes,1,opt,name=accountId,proto3" json:"accountId,omitempty"`
}

func (x *GenerateOtpRequest) Reset() {
	*x = GenerateOtpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_otp_otp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateOtpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateOtpRequest) ProtoMessage() {}

func (x *GenerateOtpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_otp_otp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateOtpRequest.ProtoReflect.Descriptor instead.
func (*GenerateOtpRequest) Descriptor() ([]byte, []int) {
	return file_services_otp_otp_proto_rawDescGZIP(), []int{0}
}

func (x *GenerateOtpRequest) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

type OtpResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Password string `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *OtpResponse) Reset() {
	*x = OtpResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_otp_otp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OtpResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OtpResponse) ProtoMessage() {}

func (x *OtpResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_otp_otp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OtpResponse.ProtoReflect.Descriptor instead.
func (*OtpResponse) Descriptor() ([]byte, []int) {
	return file_services_otp_otp_proto_rawDescGZIP(), []int{1}
}

func (x *OtpResponse) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

var File_services_otp_otp_proto protoreflect.FileDescriptor

var file_services_otp_otp_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6f, 0x74, 0x70, 0x2f, 0x6f,
	0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x32, 0x0a, 0x12, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x4f, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x29, 0x0a, 0x0b,
	0x4f, 0x74, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x32, 0x39, 0x0a, 0x03, 0x4f, 0x74, 0x70, 0x12, 0x32,
	0x0a, 0x0b, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4f, 0x54, 0x50, 0x12, 0x13, 0x2e,
	0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4f, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x4f, 0x74, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x77, 0x69, 0x70, 0x64, 0x65, 0x76, 0x2d, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x6d, 0x6f, 0x6e,
	0x65, 0x79, 0x67, 0x6f, 0x70, 0x68, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2f, 0x6f, 0x74, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_otp_otp_proto_rawDescOnce sync.Once
	file_services_otp_otp_proto_rawDescData = file_services_otp_otp_proto_rawDesc
)

func file_services_otp_otp_proto_rawDescGZIP() []byte {
	file_services_otp_otp_proto_rawDescOnce.Do(func() {
		file_services_otp_otp_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_otp_otp_proto_rawDescData)
	})
	return file_services_otp_otp_proto_rawDescData
}

var file_services_otp_otp_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_services_otp_otp_proto_goTypes = []interface{}{
	(*GenerateOtpRequest)(nil), // 0: GenerateOtpRequest
	(*OtpResponse)(nil),        // 1: OtpResponse
}
var file_services_otp_otp_proto_depIdxs = []int32{
	0, // 0: Otp.GenerateOTP:input_type -> GenerateOtpRequest
	1, // 1: Otp.GenerateOTP:output_type -> OtpResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_services_otp_otp_proto_init() }
func file_services_otp_otp_proto_init() {
	if File_services_otp_otp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_otp_otp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateOtpRequest); i {
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
		file_services_otp_otp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OtpResponse); i {
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
			RawDescriptor: file_services_otp_otp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_otp_otp_proto_goTypes,
		DependencyIndexes: file_services_otp_otp_proto_depIdxs,
		MessageInfos:      file_services_otp_otp_proto_msgTypes,
	}.Build()
	File_services_otp_otp_proto = out.File
	file_services_otp_otp_proto_rawDesc = nil
	file_services_otp_otp_proto_goTypes = nil
	file_services_otp_otp_proto_depIdxs = nil
}
