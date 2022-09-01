// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: proto/stackcoder.proto

package stackcoder

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

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Say string `protobuf:"bytes,1,opt,name=say,proto3" json:"say,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_stackcoder_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_proto_stackcoder_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_proto_stackcoder_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetSay() string {
	if x != nil {
		return x.Say
	}
	return ""
}

type UsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UsersRequest) Reset() {
	*x = UsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_stackcoder_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersRequest) ProtoMessage() {}

func (x *UsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_stackcoder_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersRequest.ProtoReflect.Descriptor instead.
func (*UsersRequest) Descriptor() ([]byte, []int) {
	return file_proto_stackcoder_proto_rawDescGZIP(), []int{1}
}

func (x *UsersRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type HintsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *HintsRequest) Reset() {
	*x = HintsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_stackcoder_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HintsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HintsRequest) ProtoMessage() {}

func (x *HintsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_stackcoder_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HintsRequest.ProtoReflect.Descriptor instead.
func (*HintsRequest) Descriptor() ([]byte, []int) {
	return file_proto_stackcoder_proto_rawDescGZIP(), []int{2}
}

func (x *HintsRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type TutorialsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *TutorialsRequest) Reset() {
	*x = TutorialsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_stackcoder_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TutorialsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TutorialsRequest) ProtoMessage() {}

func (x *TutorialsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_stackcoder_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TutorialsRequest.ProtoReflect.Descriptor instead.
func (*TutorialsRequest) Descriptor() ([]byte, []int) {
	return file_proto_stackcoder_proto_rawDescGZIP(), []int{3}
}

func (x *TutorialsRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_stackcoder_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_proto_stackcoder_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_proto_stackcoder_proto_rawDescGZIP(), []int{4}
}

func (x *Request) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_stackcoder_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_proto_stackcoder_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_proto_stackcoder_proto_rawDescGZIP(), []int{5}
}

func (x *Response) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_proto_stackcoder_proto protoreflect.FileDescriptor

var file_proto_stackcoder_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x63, 0x6f, 0x64,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x63,
	0x6f, 0x64, 0x65, 0x72, 0x22, 0x1b, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x73, 0x61, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x61,
	0x79, 0x22, 0x1e, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x1e, 0x0a, 0x0c, 0x48, 0x69, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x22, 0x0a, 0x10, 0x54, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1d, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x1c, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d,
	0x73, 0x67, 0x32, 0xfa, 0x01, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x63, 0x6f, 0x64, 0x65,
	0x72, 0x12, 0x33, 0x0a, 0x04, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x13, 0x2e, 0x73, 0x74, 0x61, 0x63,
	0x6b, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14,
	0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x05, 0x48, 0x69, 0x6e, 0x74, 0x73, 0x12,
	0x18, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x2e, 0x48, 0x69, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x73, 0x74, 0x61, 0x63,
	0x6b, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x41, 0x0a, 0x09, 0x54, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x1c,
	0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x2e, 0x54, 0x75, 0x74, 0x6f,
	0x72, 0x69, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x73,
	0x74, 0x61, 0x63, 0x6b, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x18, 0x2e,
	0x73, 0x74, 0x61, 0x63, 0x6b, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x63,
	0x6f, 0x64, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x14, 0x5a, 0x12, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x73, 0x74, 0x61, 0x63, 0x6b,
	0x63, 0x6f, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_stackcoder_proto_rawDescOnce sync.Once
	file_proto_stackcoder_proto_rawDescData = file_proto_stackcoder_proto_rawDesc
)

func file_proto_stackcoder_proto_rawDescGZIP() []byte {
	file_proto_stackcoder_proto_rawDescOnce.Do(func() {
		file_proto_stackcoder_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_stackcoder_proto_rawDescData)
	})
	return file_proto_stackcoder_proto_rawDescData
}

var file_proto_stackcoder_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_stackcoder_proto_goTypes = []interface{}{
	(*Message)(nil),          // 0: stackcoder.Message
	(*UsersRequest)(nil),     // 1: stackcoder.UsersRequest
	(*HintsRequest)(nil),     // 2: stackcoder.HintsRequest
	(*TutorialsRequest)(nil), // 3: stackcoder.TutorialsRequest
	(*Request)(nil),          // 4: stackcoder.Request
	(*Response)(nil),         // 5: stackcoder.Response
}
var file_proto_stackcoder_proto_depIdxs = []int32{
	4, // 0: stackcoder.Stackcoder.Call:input_type -> stackcoder.Request
	2, // 1: stackcoder.Stackcoder.Hints:input_type -> stackcoder.HintsRequest
	3, // 2: stackcoder.Stackcoder.Tutorials:input_type -> stackcoder.TutorialsRequest
	1, // 3: stackcoder.Stackcoder.Users:input_type -> stackcoder.UsersRequest
	5, // 4: stackcoder.Stackcoder.Call:output_type -> stackcoder.Response
	5, // 5: stackcoder.Stackcoder.Hints:output_type -> stackcoder.Response
	5, // 6: stackcoder.Stackcoder.Tutorials:output_type -> stackcoder.Response
	5, // 7: stackcoder.Stackcoder.Users:output_type -> stackcoder.Response
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_stackcoder_proto_init() }
func file_proto_stackcoder_proto_init() {
	if File_proto_stackcoder_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_stackcoder_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_proto_stackcoder_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsersRequest); i {
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
		file_proto_stackcoder_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HintsRequest); i {
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
		file_proto_stackcoder_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TutorialsRequest); i {
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
		file_proto_stackcoder_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_proto_stackcoder_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_proto_stackcoder_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_stackcoder_proto_goTypes,
		DependencyIndexes: file_proto_stackcoder_proto_depIdxs,
		MessageInfos:      file_proto_stackcoder_proto_msgTypes,
	}.Build()
	File_proto_stackcoder_proto = out.File
	file_proto_stackcoder_proto_rawDesc = nil
	file_proto_stackcoder_proto_goTypes = nil
	file_proto_stackcoder_proto_depIdxs = nil
}