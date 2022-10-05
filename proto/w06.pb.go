// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.7
// source: proto/w06.proto

package w06exercise

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

type Time struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp string `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Time) Reset() {
	*x = Time{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_w06_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Time) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Time) ProtoMessage() {}

func (x *Time) ProtoReflect() protoreflect.Message {
	mi := &file_proto_w06_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Time.ProtoReflect.Descriptor instead.
func (*Time) Descriptor() ([]byte, []int) {
	return file_proto_w06_proto_rawDescGZIP(), []int{0}
}

func (x *Time) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

type ClientRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientTime string `protobuf:"bytes,1,opt,name=clientTime,proto3" json:"clientTime,omitempty"`
}

func (x *ClientRequest) Reset() {
	*x = ClientRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_w06_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientRequest) ProtoMessage() {}

func (x *ClientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_w06_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientRequest.ProtoReflect.Descriptor instead.
func (*ClientRequest) Descriptor() ([]byte, []int) {
	return file_proto_w06_proto_rawDescGZIP(), []int{1}
}

func (x *ClientRequest) GetClientTime() string {
	if x != nil {
		return x.ClientTime
	}
	return ""
}

var File_proto_w06_proto protoreflect.FileDescriptor

var file_proto_w06_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x77, 0x30, 0x36, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0b, 0x77, 0x30, 0x36, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x22, 0x24,
	0x0a, 0x04, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x22, 0x2f, 0x0a, 0x0d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x32, 0x53, 0x0a, 0x13, 0x50, 0x68, 0x79, 0x73, 0x69, 0x63, 0x61,
	0x6c, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x0b,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x2e, 0x77, 0x30,
	0x36, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x77, 0x30, 0x36, 0x65, 0x78, 0x65,
	0x72, 0x63, 0x69, 0x73, 0x65, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x15, 0x5a, 0x13, 0x64, 0x73,
	0x79, 0x73, 0x77, 0x30, 0x36, 0x2f, 0x77, 0x30, 0x36, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_w06_proto_rawDescOnce sync.Once
	file_proto_w06_proto_rawDescData = file_proto_w06_proto_rawDesc
)

func file_proto_w06_proto_rawDescGZIP() []byte {
	file_proto_w06_proto_rawDescOnce.Do(func() {
		file_proto_w06_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_w06_proto_rawDescData)
	})
	return file_proto_w06_proto_rawDescData
}

var file_proto_w06_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_w06_proto_goTypes = []interface{}{
	(*Time)(nil),          // 0: w06exercise.Time
	(*ClientRequest)(nil), // 1: w06exercise.ClientRequest
}
var file_proto_w06_proto_depIdxs = []int32{
	1, // 0: w06exercise.PhysicalTimeService.currentTime:input_type -> w06exercise.ClientRequest
	0, // 1: w06exercise.PhysicalTimeService.currentTime:output_type -> w06exercise.Time
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_w06_proto_init() }
func file_proto_w06_proto_init() {
	if File_proto_w06_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_w06_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Time); i {
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
		file_proto_w06_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientRequest); i {
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
			RawDescriptor: file_proto_w06_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_w06_proto_goTypes,
		DependencyIndexes: file_proto_w06_proto_depIdxs,
		MessageInfos:      file_proto_w06_proto_msgTypes,
	}.Build()
	File_proto_w06_proto = out.File
	file_proto_w06_proto_rawDesc = nil
	file_proto_w06_proto_goTypes = nil
	file_proto_w06_proto_depIdxs = nil
}
