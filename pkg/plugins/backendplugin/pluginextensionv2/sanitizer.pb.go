// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: sanitizer.proto

package pluginextensionv2

import (
	context "context"
	reflect "reflect"
	sync "sync"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SanitizeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename   string `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	Content    []byte `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	ConfigType string `protobuf:"bytes,3,opt,name=configType,proto3" json:"configType,omitempty"` // DOMPurify, ...
	Config     []byte `protobuf:"bytes,4,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *SanitizeRequest) Reset() {
	*x = SanitizeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sanitizer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SanitizeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SanitizeRequest) ProtoMessage() {}

func (x *SanitizeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sanitizer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SanitizeRequest.ProtoReflect.Descriptor instead.
func (*SanitizeRequest) Descriptor() ([]byte, []int) {
	return file_sanitizer_proto_rawDescGZIP(), []int{0}
}

func (x *SanitizeRequest) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *SanitizeRequest) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *SanitizeRequest) GetConfigType() string {
	if x != nil {
		return x.ConfigType
	}
	return ""
}

func (x *SanitizeRequest) GetConfig() []byte {
	if x != nil {
		return x.Config
	}
	return nil
}

type SanitizeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error     string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Sanitized []byte `protobuf:"bytes,2,opt,name=sanitized,proto3" json:"sanitized,omitempty"`
}

func (x *SanitizeResponse) Reset() {
	*x = SanitizeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sanitizer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SanitizeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SanitizeResponse) ProtoMessage() {}

func (x *SanitizeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sanitizer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SanitizeResponse.ProtoReflect.Descriptor instead.
func (*SanitizeResponse) Descriptor() ([]byte, []int) {
	return file_sanitizer_proto_rawDescGZIP(), []int{1}
}

func (x *SanitizeResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *SanitizeResponse) GetSanitized() []byte {
	if x != nil {
		return x.Sanitized
	}
	return nil
}

var File_sanitizer_proto protoreflect.FileDescriptor

var file_sanitizer_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x73, 0x61, 0x6e, 0x69, 0x74, 0x69, 0x7a, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x11, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x76, 0x32, 0x22, 0x7f, 0x0a, 0x0f, 0x53, 0x61, 0x6e, 0x69, 0x74, 0x69, 0x7a, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x0a,
	0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x46, 0x0a, 0x10, 0x53, 0x61, 0x6e, 0x69, 0x74, 0x69, 0x7a,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12,
	0x1c, 0x0a, 0x09, 0x73, 0x61, 0x6e, 0x69, 0x74, 0x69, 0x7a, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x09, 0x73, 0x61, 0x6e, 0x69, 0x74, 0x69, 0x7a, 0x65, 0x64, 0x32, 0x60, 0x0a,
	0x09, 0x53, 0x61, 0x6e, 0x69, 0x74, 0x69, 0x7a, 0x65, 0x72, 0x12, 0x53, 0x0a, 0x08, 0x53, 0x61,
	0x6e, 0x69, 0x74, 0x69, 0x7a, 0x65, 0x12, 0x22, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x76, 0x32, 0x2e, 0x53, 0x61, 0x6e, 0x69, 0x74,
	0x69, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x76, 0x32, 0x2e, 0x53,
	0x61, 0x6e, 0x69, 0x74, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x15, 0x5a, 0x13, 0x2e, 0x3b, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x76, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sanitizer_proto_rawDescOnce sync.Once
	file_sanitizer_proto_rawDescData = file_sanitizer_proto_rawDesc
)

func file_sanitizer_proto_rawDescGZIP() []byte {
	file_sanitizer_proto_rawDescOnce.Do(func() {
		file_sanitizer_proto_rawDescData = protoimpl.X.CompressGZIP(file_sanitizer_proto_rawDescData)
	})
	return file_sanitizer_proto_rawDescData
}

var file_sanitizer_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_sanitizer_proto_goTypes = []any{
	(*SanitizeRequest)(nil),  // 0: pluginextensionv2.SanitizeRequest
	(*SanitizeResponse)(nil), // 1: pluginextensionv2.SanitizeResponse
}
var file_sanitizer_proto_depIdxs = []int32{
	0, // 0: pluginextensionv2.Sanitizer.Sanitize:input_type -> pluginextensionv2.SanitizeRequest
	1, // 1: pluginextensionv2.Sanitizer.Sanitize:output_type -> pluginextensionv2.SanitizeResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sanitizer_proto_init() }
func file_sanitizer_proto_init() {
	if File_sanitizer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sanitizer_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*SanitizeRequest); i {
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
		file_sanitizer_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*SanitizeResponse); i {
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
			RawDescriptor: file_sanitizer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sanitizer_proto_goTypes,
		DependencyIndexes: file_sanitizer_proto_depIdxs,
		MessageInfos:      file_sanitizer_proto_msgTypes,
	}.Build()
	File_sanitizer_proto = out.File
	file_sanitizer_proto_rawDesc = nil
	file_sanitizer_proto_goTypes = nil
	file_sanitizer_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SanitizerClient is the client API for Sanitizer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SanitizerClient interface {
	Sanitize(ctx context.Context, in *SanitizeRequest, opts ...grpc.CallOption) (*SanitizeResponse, error)
}

type sanitizerClient struct {
	cc grpc.ClientConnInterface
}

func NewSanitizerClient(cc grpc.ClientConnInterface) SanitizerClient {
	return &sanitizerClient{cc}
}

func (c *sanitizerClient) Sanitize(ctx context.Context, in *SanitizeRequest, opts ...grpc.CallOption) (*SanitizeResponse, error) {
	out := new(SanitizeResponse)
	err := c.cc.Invoke(ctx, "/pluginextensionv2.Sanitizer/Sanitize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SanitizerServer is the server API for Sanitizer service.
type SanitizerServer interface {
	Sanitize(context.Context, *SanitizeRequest) (*SanitizeResponse, error)
}

// UnimplementedSanitizerServer can be embedded to have forward compatible implementations.
type UnimplementedSanitizerServer struct {
}

func (*UnimplementedSanitizerServer) Sanitize(context.Context, *SanitizeRequest) (*SanitizeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sanitize not implemented")
}

func RegisterSanitizerServer(s *grpc.Server, srv SanitizerServer) {
	s.RegisterService(&_Sanitizer_serviceDesc, srv)
}

func _Sanitizer_Sanitize_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(SanitizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SanitizerServer).Sanitize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pluginextensionv2.Sanitizer/Sanitize",
	}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(SanitizerServer).Sanitize(ctx, req.(*SanitizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Sanitizer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pluginextensionv2.Sanitizer",
	HandlerType: (*SanitizerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sanitize",
			Handler:    _Sanitizer_Sanitize_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sanitizer.proto",
}
