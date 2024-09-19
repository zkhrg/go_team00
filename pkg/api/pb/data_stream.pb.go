package pb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)

	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DataMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId        string  `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	Frequency        float64 `protobuf:"fixed64,2,opt,name=frequency,proto3" json:"frequency,omitempty"`
	CurrentTimestamp float64 `protobuf:"fixed64,3,opt,name=current_timestamp,json=currentTimestamp,proto3" json:"current_timestamp,omitempty"`
}

func (x *DataMessage) Reset() {
	*x = DataMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_data_stream_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataMessage) ProtoMessage() {}

func (x *DataMessage) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_data_stream_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*DataMessage) Descriptor() ([]byte, []int) {
	return file_pkg_api_data_stream_proto_rawDescGZIP(), []int{0}
}

func (x *DataMessage) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *DataMessage) GetFrequency() float64 {
	if x != nil {
		return x.Frequency
	}
	return 0
}

func (x *DataMessage) GetCurrentTimestamp() float64 {
	if x != nil {
		return x.CurrentTimestamp
	}
	return 0
}

type StreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StreamRequest) Reset() {
	*x = StreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_data_stream_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamRequest) ProtoMessage() {}

func (x *StreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_data_stream_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*StreamRequest) Descriptor() ([]byte, []int) {
	return file_pkg_api_data_stream_proto_rawDescGZIP(), []int{1}
}

var File_pkg_api_data_stream_proto protoreflect.FileDescriptor

var file_pkg_api_data_stream_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x64, 0x61, 0x74,
	0x61, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x22, 0x77, 0x0a, 0x0b, 0x44, 0x61, 0x74, 0x61, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e,
	0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x6e, 0x63, 0x79, 0x12, 0x2b, 0x0a, 0x11, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x10,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x22, 0x0f, 0x0a, 0x0d, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x32, 0x50, 0x0a, 0x0a, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12,
	0x42, 0x0a, 0x0a, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x12, 0x19, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x30, 0x01, 0x42, 0x0c, 0x5a, 0x0a, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_api_data_stream_proto_rawDescOnce sync.Once
	file_pkg_api_data_stream_proto_rawDescData = file_pkg_api_data_stream_proto_rawDesc
)

func file_pkg_api_data_stream_proto_rawDescGZIP() []byte {
	file_pkg_api_data_stream_proto_rawDescOnce.Do(func() {
		file_pkg_api_data_stream_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_api_data_stream_proto_rawDescData)
	})
	return file_pkg_api_data_stream_proto_rawDescData
}

var file_pkg_api_data_stream_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_api_data_stream_proto_goTypes = []any{
	(*DataMessage)(nil),
	(*StreamRequest)(nil),
}
var file_pkg_api_data_stream_proto_depIdxs = []int32{
	1,
	0,
	1,
	0,
	0,
	0,
	0,
}

func init() { file_pkg_api_data_stream_proto_init() }
func file_pkg_api_data_stream_proto_init() {
	if File_pkg_api_data_stream_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_api_data_stream_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*DataMessage); i {
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
		file_pkg_api_data_stream_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*StreamRequest); i {
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
			RawDescriptor: file_pkg_api_data_stream_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_api_data_stream_proto_goTypes,
		DependencyIndexes: file_pkg_api_data_stream_proto_depIdxs,
		MessageInfos:      file_pkg_api_data_stream_proto_msgTypes,
	}.Build()
	File_pkg_api_data_stream_proto = out.File
	file_pkg_api_data_stream_proto_rawDesc = nil
	file_pkg_api_data_stream_proto_goTypes = nil
	file_pkg_api_data_stream_proto_depIdxs = nil
}
