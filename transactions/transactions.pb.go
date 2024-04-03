// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.1
// source: transactions/transactions.proto

package transactions

import (
	money "google.golang.org/genproto/googleapis/type/money"
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

type DepositRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrentBalance *money.Money `protobuf:"bytes,1,opt,name=currentBalance,proto3" json:"currentBalance,omitempty"`
	Amount         *money.Money `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *DepositRequest) Reset() {
	*x = DepositRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transactions_transactions_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DepositRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DepositRequest) ProtoMessage() {}

func (x *DepositRequest) ProtoReflect() protoreflect.Message {
	mi := &file_transactions_transactions_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DepositRequest.ProtoReflect.Descriptor instead.
func (*DepositRequest) Descriptor() ([]byte, []int) {
	return file_transactions_transactions_proto_rawDescGZIP(), []int{0}
}

func (x *DepositRequest) GetCurrentBalance() *money.Money {
	if x != nil {
		return x.CurrentBalance
	}
	return nil
}

func (x *DepositRequest) GetAmount() *money.Money {
	if x != nil {
		return x.Amount
	}
	return nil
}

type DepositResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NewBalance *money.Money `protobuf:"bytes,1,opt,name=newBalance,proto3" json:"newBalance,omitempty"`
}

func (x *DepositResponse) Reset() {
	*x = DepositResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transactions_transactions_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DepositResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DepositResponse) ProtoMessage() {}

func (x *DepositResponse) ProtoReflect() protoreflect.Message {
	mi := &file_transactions_transactions_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DepositResponse.ProtoReflect.Descriptor instead.
func (*DepositResponse) Descriptor() ([]byte, []int) {
	return file_transactions_transactions_proto_rawDescGZIP(), []int{1}
}

func (x *DepositResponse) GetNewBalance() *money.Money {
	if x != nil {
		return x.NewBalance
	}
	return nil
}

var File_transactions_transactions_proto protoreflect.FileDescriptor

var file_transactions_transactions_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x78, 0x0a, 0x0e, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x0e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x4d, 0x6f, 0x6e,
	0x65, 0x79, 0x52, 0x0e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x2e, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x45,
	0x0a, 0x0f, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x32, 0x0a, 0x0a, 0x6e, 0x65, 0x77, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x52, 0x0a, 0x6e, 0x65, 0x77, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x32, 0x3e, 0x0a, 0x0c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2e, 0x0a, 0x07, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x12, 0x0f, 0x2e, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x10, 0x2e, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x69, 0x70, 0x64, 0x65, 0x76, 0x2d, 0x74, 0x65, 0x63, 0x68, 0x2f,
	0x6d, 0x69, 0x63, 0x72, 0x6f, 0x62, 0x61, 0x6e, 0x6b, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transactions_transactions_proto_rawDescOnce sync.Once
	file_transactions_transactions_proto_rawDescData = file_transactions_transactions_proto_rawDesc
)

func file_transactions_transactions_proto_rawDescGZIP() []byte {
	file_transactions_transactions_proto_rawDescOnce.Do(func() {
		file_transactions_transactions_proto_rawDescData = protoimpl.X.CompressGZIP(file_transactions_transactions_proto_rawDescData)
	})
	return file_transactions_transactions_proto_rawDescData
}

var file_transactions_transactions_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_transactions_transactions_proto_goTypes = []interface{}{
	(*DepositRequest)(nil),  // 0: DepositRequest
	(*DepositResponse)(nil), // 1: DepositResponse
	(*money.Money)(nil),     // 2: google.type.Money
}
var file_transactions_transactions_proto_depIdxs = []int32{
	2, // 0: DepositRequest.currentBalance:type_name -> google.type.Money
	2, // 1: DepositRequest.amount:type_name -> google.type.Money
	2, // 2: DepositResponse.newBalance:type_name -> google.type.Money
	0, // 3: Transactions.Deposit:input_type -> DepositRequest
	1, // 4: Transactions.Deposit:output_type -> DepositResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_transactions_transactions_proto_init() }
func file_transactions_transactions_proto_init() {
	if File_transactions_transactions_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_transactions_transactions_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DepositRequest); i {
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
		file_transactions_transactions_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DepositResponse); i {
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
			RawDescriptor: file_transactions_transactions_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_transactions_transactions_proto_goTypes,
		DependencyIndexes: file_transactions_transactions_proto_depIdxs,
		MessageInfos:      file_transactions_transactions_proto_msgTypes,
	}.Build()
	File_transactions_transactions_proto = out.File
	file_transactions_transactions_proto_rawDesc = nil
	file_transactions_transactions_proto_goTypes = nil
	file_transactions_transactions_proto_depIdxs = nil
}
