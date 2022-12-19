// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: proto/vehicle.proto

package pb

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

type NewRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token  string            `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Type   string            `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Config map[string]string `protobuf:"bytes,3,rep,name=config,proto3" json:"config,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *NewRequest) Reset() {
	*x = NewRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_vehicle_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewRequest) ProtoMessage() {}

func (x *NewRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vehicle_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewRequest.ProtoReflect.Descriptor instead.
func (*NewRequest) Descriptor() ([]byte, []int) {
	return file_proto_vehicle_proto_rawDescGZIP(), []int{0}
}

func (x *NewRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *NewRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *NewRequest) GetConfig() map[string]string {
	if x != nil {
		return x.Config
	}
	return nil
}

type NewReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VehicleId int64 `protobuf:"varint,1,opt,name=vehicle_id,json=vehicleId,proto3" json:"vehicle_id,omitempty"`
}

func (x *NewReply) Reset() {
	*x = NewReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_vehicle_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewReply) ProtoMessage() {}

func (x *NewReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vehicle_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewReply.ProtoReflect.Descriptor instead.
func (*NewReply) Descriptor() ([]byte, []int) {
	return file_proto_vehicle_proto_rawDescGZIP(), []int{1}
}

func (x *NewReply) GetVehicleId() int64 {
	if x != nil {
		return x.VehicleId
	}
	return 0
}

type SocRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token     string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	VehicleId int64  `protobuf:"varint,2,opt,name=vehicle_id,json=vehicleId,proto3" json:"vehicle_id,omitempty"`
}

func (x *SocRequest) Reset() {
	*x = SocRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_vehicle_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SocRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SocRequest) ProtoMessage() {}

func (x *SocRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vehicle_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SocRequest.ProtoReflect.Descriptor instead.
func (*SocRequest) Descriptor() ([]byte, []int) {
	return file_proto_vehicle_proto_rawDescGZIP(), []int{2}
}

func (x *SocRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *SocRequest) GetVehicleId() int64 {
	if x != nil {
		return x.VehicleId
	}
	return 0
}

type SocReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Soc float64 `protobuf:"fixed64,1,opt,name=soc,proto3" json:"soc,omitempty"`
}

func (x *SocReply) Reset() {
	*x = SocReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_vehicle_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SocReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SocReply) ProtoMessage() {}

func (x *SocReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vehicle_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SocReply.ProtoReflect.Descriptor instead.
func (*SocReply) Descriptor() ([]byte, []int) {
	return file_proto_vehicle_proto_rawDescGZIP(), []int{3}
}

func (x *SocReply) GetSoc() float64 {
	if x != nil {
		return x.Soc
	}
	return 0
}

var File_proto_vehicle_proto protoreflect.FileDescriptor

var file_proto_vehicle_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa2, 0x01, 0x0a, 0x0a, 0x4e, 0x65, 0x77, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x2f,
	0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x4e, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a,
	0x39, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x29, 0x0a, 0x08, 0x4e, 0x65,
	0x77, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x76, 0x65, 0x68, 0x69,
	0x63, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x41, 0x0a, 0x0a, 0x53, 0x6f, 0x43, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x65, 0x68,
	0x69, 0x63, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x76,
	0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x1c, 0x0a, 0x08, 0x53, 0x6f, 0x43, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x6f, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x03, 0x73, 0x6f, 0x63, 0x32, 0x4b, 0x0a, 0x07, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c,
	0x65, 0x12, 0x1f, 0x0a, 0x03, 0x4e, 0x65, 0x77, 0x12, 0x0b, 0x2e, 0x4e, 0x65, 0x77, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x4e, 0x65, 0x77, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x12, 0x1f, 0x0a, 0x03, 0x53, 0x6f, 0x43, 0x12, 0x0b, 0x2e, 0x53, 0x6f, 0x43, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x53, 0x6f, 0x43, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_vehicle_proto_rawDescOnce sync.Once
	file_proto_vehicle_proto_rawDescData = file_proto_vehicle_proto_rawDesc
)

func file_proto_vehicle_proto_rawDescGZIP() []byte {
	file_proto_vehicle_proto_rawDescOnce.Do(func() {
		file_proto_vehicle_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_vehicle_proto_rawDescData)
	})
	return file_proto_vehicle_proto_rawDescData
}

var file_proto_vehicle_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_vehicle_proto_goTypes = []interface{}{
	(*NewRequest)(nil), // 0: NewRequest
	(*NewReply)(nil),   // 1: NewReply
	(*SocRequest)(nil), // 2: SocRequest
	(*SocReply)(nil),   // 3: SocReply
	nil,                // 4: NewRequest.ConfigEntry
}
var file_proto_vehicle_proto_depIdxs = []int32{
	4, // 0: NewRequest.config:type_name -> NewRequest.ConfigEntry
	0, // 1: Vehicle.New:input_type -> NewRequest
	2, // 2: Vehicle.Soc:input_type -> SocRequest
	1, // 3: Vehicle.New:output_type -> NewReply
	3, // 4: Vehicle.Soc:output_type -> SocReply
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_vehicle_proto_init() }
func file_proto_vehicle_proto_init() {
	if File_proto_vehicle_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_vehicle_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewRequest); i {
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
		file_proto_vehicle_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewReply); i {
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
		file_proto_vehicle_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SocRequest); i {
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
		file_proto_vehicle_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SocReply); i {
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
			RawDescriptor: file_proto_vehicle_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_vehicle_proto_goTypes,
		DependencyIndexes: file_proto_vehicle_proto_depIdxs,
		MessageInfos:      file_proto_vehicle_proto_msgTypes,
	}.Build()
	File_proto_vehicle_proto = out.File
	file_proto_vehicle_proto_rawDesc = nil
	file_proto_vehicle_proto_goTypes = nil
	file_proto_vehicle_proto_depIdxs = nil
}
