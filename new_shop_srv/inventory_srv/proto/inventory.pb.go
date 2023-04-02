// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.5
// source: inventory.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type InvInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GoodsId int32 `protobuf:"varint,1,opt,name=goodsId,proto3" json:"goodsId,omitempty"`
	Num     int32 `protobuf:"varint,2,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *InvInfo) Reset() {
	*x = InvInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inventory_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvInfo) ProtoMessage() {}

func (x *InvInfo) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvInfo.ProtoReflect.Descriptor instead.
func (*InvInfo) Descriptor() ([]byte, []int) {
	return file_inventory_proto_rawDescGZIP(), []int{0}
}

func (x *InvInfo) GetGoodsId() int32 {
	if x != nil {
		return x.GoodsId
	}
	return 0
}

func (x *InvInfo) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

type SellInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GoodsInfo []*InvInfo `protobuf:"bytes,1,rep,name=goodsInfo,proto3" json:"goodsInfo,omitempty"`
}

func (x *SellInfo) Reset() {
	*x = SellInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inventory_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SellInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SellInfo) ProtoMessage() {}

func (x *SellInfo) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SellInfo.ProtoReflect.Descriptor instead.
func (*SellInfo) Descriptor() ([]byte, []int) {
	return file_inventory_proto_rawDescGZIP(), []int{1}
}

func (x *SellInfo) GetGoodsInfo() []*InvInfo {
	if x != nil {
		return x.GoodsInfo
	}
	return nil
}

var File_inventory_proto protoreflect.FileDescriptor

var file_inventory_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x35,
	0x0a, 0x07, 0x49, 0x6e, 0x76, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x6f, 0x6f,
	0x64, 0x73, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x67, 0x6f, 0x6f, 0x64,
	0x73, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x03, 0x6e, 0x75, 0x6d, 0x22, 0x32, 0x0a, 0x08, 0x53, 0x65, 0x6c, 0x6c, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x26, 0x0a, 0x09, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x49, 0x6e, 0x76, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09,
	0x67, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x32, 0xb0, 0x01, 0x0a, 0x09, 0x49, 0x6e,
	0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x2a, 0x0a, 0x06, 0x53, 0x65, 0x74, 0x49, 0x6e,
	0x76, 0x12, 0x08, 0x2e, 0x49, 0x6e, 0x76, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x12, 0x1f, 0x0a, 0x09, 0x49, 0x6e, 0x76, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x12, 0x08, 0x2e, 0x49, 0x6e, 0x76, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x08, 0x2e, 0x49, 0x6e, 0x76,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x29, 0x0a, 0x04, 0x53, 0x65, 0x6c, 0x6c, 0x12, 0x09, 0x2e, 0x53,
	0x65, 0x6c, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12,
	0x2b, 0x0a, 0x06, 0x52, 0x65, 0x42, 0x61, 0x63, 0x6b, 0x12, 0x09, 0x2e, 0x53, 0x65, 0x6c, 0x6c,
	0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x0a, 0x5a, 0x08,
	0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_inventory_proto_rawDescOnce sync.Once
	file_inventory_proto_rawDescData = file_inventory_proto_rawDesc
)

func file_inventory_proto_rawDescGZIP() []byte {
	file_inventory_proto_rawDescOnce.Do(func() {
		file_inventory_proto_rawDescData = protoimpl.X.CompressGZIP(file_inventory_proto_rawDescData)
	})
	return file_inventory_proto_rawDescData
}

var file_inventory_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_inventory_proto_goTypes = []interface{}{
	(*InvInfo)(nil),       // 0: InvInfo
	(*SellInfo)(nil),      // 1: SellInfo
	(*emptypb.Empty)(nil), // 2: google.protobuf.Empty
}
var file_inventory_proto_depIdxs = []int32{
	0, // 0: SellInfo.goodsInfo:type_name -> InvInfo
	0, // 1: Inventory.SetInv:input_type -> InvInfo
	0, // 2: Inventory.InvDetail:input_type -> InvInfo
	1, // 3: Inventory.Sell:input_type -> SellInfo
	1, // 4: Inventory.ReBack:input_type -> SellInfo
	2, // 5: Inventory.SetInv:output_type -> google.protobuf.Empty
	0, // 6: Inventory.InvDetail:output_type -> InvInfo
	2, // 7: Inventory.Sell:output_type -> google.protobuf.Empty
	2, // 8: Inventory.ReBack:output_type -> google.protobuf.Empty
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_inventory_proto_init() }
func file_inventory_proto_init() {
	if File_inventory_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_inventory_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InvInfo); i {
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
		file_inventory_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SellInfo); i {
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
			RawDescriptor: file_inventory_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_inventory_proto_goTypes,
		DependencyIndexes: file_inventory_proto_depIdxs,
		MessageInfos:      file_inventory_proto_msgTypes,
	}.Build()
	File_inventory_proto = out.File
	file_inventory_proto_rawDesc = nil
	file_inventory_proto_goTypes = nil
	file_inventory_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// InventoryClient is the client API for Inventory service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InventoryClient interface {
	SetInv(ctx context.Context, in *InvInfo, opts ...grpc.CallOption) (*emptypb.Empty, error)
	InvDetail(ctx context.Context, in *InvInfo, opts ...grpc.CallOption) (*InvInfo, error)
	Sell(ctx context.Context, in *SellInfo, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ReBack(ctx context.Context, in *SellInfo, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type inventoryClient struct {
	cc grpc.ClientConnInterface
}

func NewInventoryClient(cc grpc.ClientConnInterface) InventoryClient {
	return &inventoryClient{cc}
}

func (c *inventoryClient) SetInv(ctx context.Context, in *InvInfo, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/Inventory/SetInv", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryClient) InvDetail(ctx context.Context, in *InvInfo, opts ...grpc.CallOption) (*InvInfo, error) {
	out := new(InvInfo)
	err := c.cc.Invoke(ctx, "/Inventory/InvDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryClient) Sell(ctx context.Context, in *SellInfo, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/Inventory/Sell", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryClient) ReBack(ctx context.Context, in *SellInfo, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/Inventory/ReBack", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InventoryServer is the server API for Inventory service.
type InventoryServer interface {
	SetInv(context.Context, *InvInfo) (*emptypb.Empty, error)
	InvDetail(context.Context, *InvInfo) (*InvInfo, error)
	Sell(context.Context, *SellInfo) (*emptypb.Empty, error)
	ReBack(context.Context, *SellInfo) (*emptypb.Empty, error)
}

// UnimplementedInventoryServer can be embedded to have forward compatible implementations.
type UnimplementedInventoryServer struct {
}

func (*UnimplementedInventoryServer) SetInv(context.Context, *InvInfo) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetInv not implemented")
}
func (*UnimplementedInventoryServer) InvDetail(context.Context, *InvInfo) (*InvInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InvDetail not implemented")
}
func (*UnimplementedInventoryServer) Sell(context.Context, *SellInfo) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sell not implemented")
}
func (*UnimplementedInventoryServer) ReBack(context.Context, *SellInfo) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReBack not implemented")
}

func RegisterInventoryServer(s *grpc.Server, srv InventoryServer) {
	s.RegisterService(&_Inventory_serviceDesc, srv)
}

func _Inventory_SetInv_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServer).SetInv(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Inventory/SetInv",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServer).SetInv(ctx, req.(*InvInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Inventory_InvDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServer).InvDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Inventory/InvDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServer).InvDetail(ctx, req.(*InvInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Inventory_Sell_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SellInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServer).Sell(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Inventory/Sell",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServer).Sell(ctx, req.(*SellInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Inventory_ReBack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SellInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServer).ReBack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Inventory/ReBack",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServer).ReBack(ctx, req.(*SellInfo))
	}
	return interceptor(ctx, in, info, handler)
}

var _Inventory_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Inventory",
	HandlerType: (*InventoryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetInv",
			Handler:    _Inventory_SetInv_Handler,
		},
		{
			MethodName: "InvDetail",
			Handler:    _Inventory_InvDetail_Handler,
		},
		{
			MethodName: "Sell",
			Handler:    _Inventory_Sell_Handler,
		},
		{
			MethodName: "ReBack",
			Handler:    _Inventory_ReBack_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "inventory.proto",
}
