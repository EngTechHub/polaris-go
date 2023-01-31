// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.21.9
// source: grpcapi_location.proto

package proto

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

type LocationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientIp string `protobuf:"bytes,1,opt,name=client_ip,proto3" json:"client_ip,omitempty"`
}

func (x *LocationRequest) Reset() {
	*x = LocationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcapi_location_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocationRequest) ProtoMessage() {}

func (x *LocationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpcapi_location_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocationRequest.ProtoReflect.Descriptor instead.
func (*LocationRequest) Descriptor() ([]byte, []int) {
	return file_grpcapi_location_proto_rawDescGZIP(), []int{0}
}

func (x *LocationRequest) GetClientIp() string {
	if x != nil {
		return x.ClientIp
	}
	return ""
}

type LocationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Region string `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	Zone   string `protobuf:"bytes,2,opt,name=zone,proto3" json:"zone,omitempty"`
	Campus string `protobuf:"bytes,3,opt,name=campus,proto3" json:"campus,omitempty"`
}

func (x *LocationResponse) Reset() {
	*x = LocationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcapi_location_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocationResponse) ProtoMessage() {}

func (x *LocationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpcapi_location_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocationResponse.ProtoReflect.Descriptor instead.
func (*LocationResponse) Descriptor() ([]byte, []int) {
	return file_grpcapi_location_proto_rawDescGZIP(), []int{1}
}

func (x *LocationResponse) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *LocationResponse) GetZone() string {
	if x != nil {
		return x.Zone
	}
	return ""
}

func (x *LocationResponse) GetCampus() string {
	if x != nil {
		return x.Campus
	}
	return ""
}

var File_grpcapi_location_proto protoreflect.FileDescriptor

var file_grpcapi_location_proto_rawDesc = []byte{
	0x0a, 0x16, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x2f, 0x0a, 0x0f, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x70, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x70,
	0x22, 0x56, 0x0a, 0x10, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x7a, 0x6f, 0x6e, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x32, 0x4c, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x40, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x47, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73, 0x6d, 0x65, 0x73, 0x68,
	0x2f, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x72, 0x65, 0x6d, 0x6f,
	0x74, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpcapi_location_proto_rawDescOnce sync.Once
	file_grpcapi_location_proto_rawDescData = file_grpcapi_location_proto_rawDesc
)

func file_grpcapi_location_proto_rawDescGZIP() []byte {
	file_grpcapi_location_proto_rawDescOnce.Do(func() {
		file_grpcapi_location_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpcapi_location_proto_rawDescData)
	})
	return file_grpcapi_location_proto_rawDescData
}

var file_grpcapi_location_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_grpcapi_location_proto_goTypes = []interface{}{
	(*LocationRequest)(nil),  // 0: proto.LocationRequest
	(*LocationResponse)(nil), // 1: proto.LocationResponse
}
var file_grpcapi_location_proto_depIdxs = []int32{
	0, // 0: proto.Location.GetLocation:input_type -> proto.LocationRequest
	1, // 1: proto.Location.GetLocation:output_type -> proto.LocationResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpcapi_location_proto_init() }
func file_grpcapi_location_proto_init() {
	if File_grpcapi_location_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpcapi_location_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocationRequest); i {
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
		file_grpcapi_location_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocationResponse); i {
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
			RawDescriptor: file_grpcapi_location_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpcapi_location_proto_goTypes,
		DependencyIndexes: file_grpcapi_location_proto_depIdxs,
		MessageInfos:      file_grpcapi_location_proto_msgTypes,
	}.Build()
	File_grpcapi_location_proto = out.File
	file_grpcapi_location_proto_rawDesc = nil
	file_grpcapi_location_proto_goTypes = nil
	file_grpcapi_location_proto_depIdxs = nil
}
