// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0-devel
// 	protoc        v3.19.4
// source: pkg/wanikani/pb/wanikani.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Card struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	WanikaniId int32    `protobuf:"varint,2,opt,name=wanikaniId,proto3" json:"wanikaniId,omitempty"`
	Type       string   `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Characters string   `protobuf:"bytes,4,opt,name=characters,proto3" json:"characters,omitempty"`
	Meanings   []string `protobuf:"bytes,5,rep,name=meanings,proto3" json:"meanings,omitempty"`
	Readings   []string `protobuf:"bytes,6,rep,name=readings,proto3" json:"readings,omitempty"`
	Audio      string   `protobuf:"bytes,7,opt,name=audio,proto3" json:"audio,omitempty"`
}

func (x *Card) Reset() {
	*x = Card{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_wanikani_pb_wanikani_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Card) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Card) ProtoMessage() {}

func (x *Card) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_wanikani_pb_wanikani_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Card.ProtoReflect.Descriptor instead.
func (*Card) Descriptor() ([]byte, []int) {
	return file_pkg_wanikani_pb_wanikani_proto_rawDescGZIP(), []int{0}
}

func (x *Card) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Card) GetWanikaniId() int32 {
	if x != nil {
		return x.WanikaniId
	}
	return 0
}

func (x *Card) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Card) GetCharacters() string {
	if x != nil {
		return x.Characters
	}
	return ""
}

func (x *Card) GetMeanings() []string {
	if x != nil {
		return x.Meanings
	}
	return nil
}

func (x *Card) GetReadings() []string {
	if x != nil {
		return x.Readings
	}
	return nil
}

func (x *Card) GetAudio() string {
	if x != nil {
		return x.Audio
	}
	return ""
}

type FindRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Level int32  `protobuf:"varint,1,opt,name=level,proto3" json:"level,omitempty"`
	Type  string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *FindRequest) Reset() {
	*x = FindRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_wanikani_pb_wanikani_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindRequest) ProtoMessage() {}

func (x *FindRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_wanikani_pb_wanikani_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindRequest.ProtoReflect.Descriptor instead.
func (*FindRequest) Descriptor() ([]byte, []int) {
	return file_pkg_wanikani_pb_wanikani_proto_rawDescGZIP(), []int{1}
}

func (x *FindRequest) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *FindRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type FindResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int32   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string  `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Data   []*Card `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *FindResponse) Reset() {
	*x = FindResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_wanikani_pb_wanikani_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindResponse) ProtoMessage() {}

func (x *FindResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_wanikani_pb_wanikani_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindResponse.ProtoReflect.Descriptor instead.
func (*FindResponse) Descriptor() ([]byte, []int) {
	return file_pkg_wanikani_pb_wanikani_proto_rawDescGZIP(), []int{2}
}

func (x *FindResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *FindResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *FindResponse) GetData() []*Card {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_pkg_wanikani_pb_wanikani_proto protoreflect.FileDescriptor

var file_pkg_wanikani_pb_wanikani_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x6b, 0x67, 0x2f, 0x77, 0x61, 0x6e, 0x69, 0x6b, 0x61, 0x6e, 0x69, 0x2f, 0x70,
	0x62, 0x2f, 0x77, 0x61, 0x6e, 0x69, 0x6b, 0x61, 0x6e, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x77, 0x61, 0x6e, 0x69, 0x6b, 0x61, 0x6e, 0x69, 0x22, 0xb8, 0x01, 0x0a, 0x04, 0x43,
	0x61, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x77, 0x61, 0x6e, 0x69, 0x6b, 0x61, 0x6e, 0x69, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x77, 0x61, 0x6e, 0x69, 0x6b, 0x61, 0x6e,
	0x69, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x72, 0x61,
	0x63, 0x74, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x68, 0x61,
	0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x61, 0x6e, 0x69,
	0x6e, 0x67, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x61, 0x6e, 0x69,
	0x6e, 0x67, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x61, 0x75, 0x64, 0x69, 0x6f, 0x22, 0x37, 0x0a, 0x0b, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x60,
	0x0a, 0x0c, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x22, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x77, 0x61, 0x6e,
	0x69, 0x6b, 0x61, 0x6e, 0x69, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x32, 0x4a, 0x0a, 0x0f, 0x57, 0x61, 0x6e, 0x69, 0x6b, 0x61, 0x6e, 0x69, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x04, 0x46, 0x69, 0x6e, 0x64, 0x12, 0x15, 0x2e, 0x77, 0x61,
	0x6e, 0x69, 0x6b, 0x61, 0x6e, 0x69, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x77, 0x61, 0x6e, 0x69, 0x6b, 0x61, 0x6e, 0x69, 0x2e, 0x46, 0x69,
	0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x13, 0x5a, 0x11,
	0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x77, 0x61, 0x6e, 0x69, 0x6b, 0x61, 0x6e, 0x69, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_wanikani_pb_wanikani_proto_rawDescOnce sync.Once
	file_pkg_wanikani_pb_wanikani_proto_rawDescData = file_pkg_wanikani_pb_wanikani_proto_rawDesc
)

func file_pkg_wanikani_pb_wanikani_proto_rawDescGZIP() []byte {
	file_pkg_wanikani_pb_wanikani_proto_rawDescOnce.Do(func() {
		file_pkg_wanikani_pb_wanikani_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_wanikani_pb_wanikani_proto_rawDescData)
	})
	return file_pkg_wanikani_pb_wanikani_proto_rawDescData
}

var file_pkg_wanikani_pb_wanikani_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pkg_wanikani_pb_wanikani_proto_goTypes = []interface{}{
	(*Card)(nil),         // 0: wanikani.Card
	(*FindRequest)(nil),  // 1: wanikani.FindRequest
	(*FindResponse)(nil), // 2: wanikani.FindResponse
}
var file_pkg_wanikani_pb_wanikani_proto_depIdxs = []int32{
	0, // 0: wanikani.FindResponse.data:type_name -> wanikani.Card
	1, // 1: wanikani.WanikaniService.Find:input_type -> wanikani.FindRequest
	2, // 2: wanikani.WanikaniService.Find:output_type -> wanikani.FindResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_wanikani_pb_wanikani_proto_init() }
func file_pkg_wanikani_pb_wanikani_proto_init() {
	if File_pkg_wanikani_pb_wanikani_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_wanikani_pb_wanikani_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Card); i {
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
		file_pkg_wanikani_pb_wanikani_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindRequest); i {
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
		file_pkg_wanikani_pb_wanikani_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindResponse); i {
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
			RawDescriptor: file_pkg_wanikani_pb_wanikani_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_wanikani_pb_wanikani_proto_goTypes,
		DependencyIndexes: file_pkg_wanikani_pb_wanikani_proto_depIdxs,
		MessageInfos:      file_pkg_wanikani_pb_wanikani_proto_msgTypes,
	}.Build()
	File_pkg_wanikani_pb_wanikani_proto = out.File
	file_pkg_wanikani_pb_wanikani_proto_rawDesc = nil
	file_pkg_wanikani_pb_wanikani_proto_goTypes = nil
	file_pkg_wanikani_pb_wanikani_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// WanikaniServiceClient is the client API for WanikaniService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WanikaniServiceClient interface {
	Find(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*FindResponse, error)
}

type wanikaniServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWanikaniServiceClient(cc grpc.ClientConnInterface) WanikaniServiceClient {
	return &wanikaniServiceClient{cc}
}

func (c *wanikaniServiceClient) Find(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*FindResponse, error) {
	out := new(FindResponse)
	err := c.cc.Invoke(ctx, "/wanikani.WanikaniService/Find", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WanikaniServiceServer is the server API for WanikaniService service.
type WanikaniServiceServer interface {
	Find(context.Context, *FindRequest) (*FindResponse, error)
}

// UnimplementedWanikaniServiceServer can be embedded to have forward compatible implementations.
type UnimplementedWanikaniServiceServer struct {
}

func (*UnimplementedWanikaniServiceServer) Find(context.Context, *FindRequest) (*FindResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Find not implemented")
}

func RegisterWanikaniServiceServer(s *grpc.Server, srv WanikaniServiceServer) {
	s.RegisterService(&_WanikaniService_serviceDesc, srv)
}

func _WanikaniService_Find_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WanikaniServiceServer).Find(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wanikani.WanikaniService/Find",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WanikaniServiceServer).Find(ctx, req.(*FindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _WanikaniService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "wanikani.WanikaniService",
	HandlerType: (*WanikaniServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Find",
			Handler:    _WanikaniService_Find_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/wanikani/pb/wanikani.proto",
}
