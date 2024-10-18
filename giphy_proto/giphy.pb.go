// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: giphy.proto

package giphy_proto

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

type GiphyResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Url      string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Username string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Title    string `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *GiphyResult) Reset() {
	*x = GiphyResult{}
	mi := &file_giphy_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GiphyResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GiphyResult) ProtoMessage() {}

func (x *GiphyResult) ProtoReflect() protoreflect.Message {
	mi := &file_giphy_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GiphyResult.ProtoReflect.Descriptor instead.
func (*GiphyResult) Descriptor() ([]byte, []int) {
	return file_giphy_proto_rawDescGZIP(), []int{0}
}

func (x *GiphyResult) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GiphyResult) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *GiphyResult) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *GiphyResult) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type GiphyResults struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*GiphyResult `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *GiphyResults) Reset() {
	*x = GiphyResults{}
	mi := &file_giphy_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GiphyResults) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GiphyResults) ProtoMessage() {}

func (x *GiphyResults) ProtoReflect() protoreflect.Message {
	mi := &file_giphy_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GiphyResults.ProtoReflect.Descriptor instead.
func (*GiphyResults) Descriptor() ([]byte, []int) {
	return file_giphy_proto_rawDescGZIP(), []int{1}
}

func (x *GiphyResults) GetResults() []*GiphyResult {
	if x != nil {
		return x.Results
	}
	return nil
}

var File_giphy_proto protoreflect.FileDescriptor

var file_giphy_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x67, 0x69, 0x70, 0x68, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x67,
	0x69, 0x70, 0x68, 0x79, 0x22, 0x61, 0x0a, 0x0b, 0x47, 0x69, 0x70, 0x68, 0x79, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x3c, 0x0a, 0x0c, 0x47, 0x69, 0x70, 0x68, 0x79,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x2c, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x69, 0x70, 0x68, 0x79,
	0x2e, 0x47, 0x69, 0x70, 0x68, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x07, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x42, 0x1a, 0x5a, 0x18, 0x67, 0x69, 0x70, 0x68, 0x79, 0x73, 0x63,
	0x72, 0x61, 0x70, 0x65, 0x72, 0x2f, 0x67, 0x69, 0x70, 0x68, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_giphy_proto_rawDescOnce sync.Once
	file_giphy_proto_rawDescData = file_giphy_proto_rawDesc
)

func file_giphy_proto_rawDescGZIP() []byte {
	file_giphy_proto_rawDescOnce.Do(func() {
		file_giphy_proto_rawDescData = protoimpl.X.CompressGZIP(file_giphy_proto_rawDescData)
	})
	return file_giphy_proto_rawDescData
}

var file_giphy_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_giphy_proto_goTypes = []any{
	(*GiphyResult)(nil),  // 0: giphy.GiphyResult
	(*GiphyResults)(nil), // 1: giphy.GiphyResults
}
var file_giphy_proto_depIdxs = []int32{
	0, // 0: giphy.GiphyResults.results:type_name -> giphy.GiphyResult
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_giphy_proto_init() }
func file_giphy_proto_init() {
	if File_giphy_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_giphy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_giphy_proto_goTypes,
		DependencyIndexes: file_giphy_proto_depIdxs,
		MessageInfos:      file_giphy_proto_msgTypes,
	}.Build()
	File_giphy_proto = out.File
	file_giphy_proto_rawDesc = nil
	file_giphy_proto_goTypes = nil
	file_giphy_proto_depIdxs = nil
}
