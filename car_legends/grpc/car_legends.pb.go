// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.1
// source: car_legends.proto

package grpc

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

type Car struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Manufacturer string `protobuf:"bytes,1,opt,name=manufacturer,proto3" json:"manufacturer,omitempty"`
	Model        string `protobuf:"bytes,2,opt,name=model,proto3" json:"model,omitempty"`
	Year         string `protobuf:"bytes,3,opt,name=year,proto3" json:"year,omitempty"`
	Notes        string `protobuf:"bytes,4,opt,name=notes,proto3" json:"notes,omitempty"`
}

func (x *Car) Reset() {
	*x = Car{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_legends_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Car) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Car) ProtoMessage() {}

func (x *Car) ProtoReflect() protoreflect.Message {
	mi := &file_car_legends_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Car.ProtoReflect.Descriptor instead.
func (*Car) Descriptor() ([]byte, []int) {
	return file_car_legends_proto_rawDescGZIP(), []int{0}
}

func (x *Car) GetManufacturer() string {
	if x != nil {
		return x.Manufacturer
	}
	return ""
}

func (x *Car) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *Car) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *Car) GetNotes() string {
	if x != nil {
		return x.Notes
	}
	return ""
}

type CarList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cars []*Car `protobuf:"bytes,1,rep,name=cars,proto3" json:"cars,omitempty"`
}

func (x *CarList) Reset() {
	*x = CarList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_legends_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarList) ProtoMessage() {}

func (x *CarList) ProtoReflect() protoreflect.Message {
	mi := &file_car_legends_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CarList.ProtoReflect.Descriptor instead.
func (*CarList) Descriptor() ([]byte, []int) {
	return file_car_legends_proto_rawDescGZIP(), []int{1}
}

func (x *CarList) GetCars() []*Car {
	if x != nil {
		return x.Cars
	}
	return nil
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_legends_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_car_legends_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_car_legends_proto_rawDescGZIP(), []int{2}
}

var File_car_legends_proto protoreflect.FileDescriptor

var file_car_legends_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x61, 0x72, 0x5f, 0x6c, 0x65, 0x67, 0x65, 0x6e, 0x64, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x69, 0x0a, 0x03, 0x43, 0x61, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x61,
	0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x12, 0x14,
	0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x74, 0x65,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x22, 0x23,
	0x0a, 0x07, 0x43, 0x61, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x04, 0x63, 0x61, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x04, 0x2e, 0x43, 0x61, 0x72, 0x52, 0x04, 0x63,
	0x61, 0x72, 0x73, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0x62, 0x0a, 0x0a,
	0x43, 0x61, 0x72, 0x4c, 0x65, 0x67, 0x65, 0x6e, 0x64, 0x73, 0x12, 0x18, 0x0a, 0x06, 0x41, 0x64,
	0x64, 0x43, 0x61, 0x72, 0x12, 0x04, 0x2e, 0x43, 0x61, 0x72, 0x1a, 0x06, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x12, 0x1d, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x73, 0x12,
	0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x08, 0x2e, 0x43, 0x61, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x22, 0x00, 0x12, 0x1b, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x43, 0x61, 0x72, 0x73, 0x12, 0x04,
	0x2e, 0x43, 0x61, 0x72, 0x1a, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x28, 0x01,
	0x42, 0x12, 0x5a, 0x10, 0x63, 0x61, 0x72, 0x5f, 0x6c, 0x65, 0x67, 0x65, 0x6e, 0x64, 0x73, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_car_legends_proto_rawDescOnce sync.Once
	file_car_legends_proto_rawDescData = file_car_legends_proto_rawDesc
)

func file_car_legends_proto_rawDescGZIP() []byte {
	file_car_legends_proto_rawDescOnce.Do(func() {
		file_car_legends_proto_rawDescData = protoimpl.X.CompressGZIP(file_car_legends_proto_rawDescData)
	})
	return file_car_legends_proto_rawDescData
}

var file_car_legends_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_car_legends_proto_goTypes = []interface{}{
	(*Car)(nil),     // 0: Car
	(*CarList)(nil), // 1: CarList
	(*Empty)(nil),   // 2: Empty
}
var file_car_legends_proto_depIdxs = []int32{
	0, // 0: CarList.cars:type_name -> Car
	0, // 1: CarLegends.AddCar:input_type -> Car
	2, // 2: CarLegends.GetCars:input_type -> Empty
	0, // 3: CarLegends.AddCars:input_type -> Car
	2, // 4: CarLegends.AddCar:output_type -> Empty
	1, // 5: CarLegends.GetCars:output_type -> CarList
	2, // 6: CarLegends.AddCars:output_type -> Empty
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_car_legends_proto_init() }
func file_car_legends_proto_init() {
	if File_car_legends_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_car_legends_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Car); i {
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
		file_car_legends_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarList); i {
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
		file_car_legends_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_car_legends_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_car_legends_proto_goTypes,
		DependencyIndexes: file_car_legends_proto_depIdxs,
		MessageInfos:      file_car_legends_proto_msgTypes,
	}.Build()
	File_car_legends_proto = out.File
	file_car_legends_proto_rawDesc = nil
	file_car_legends_proto_goTypes = nil
	file_car_legends_proto_depIdxs = nil
}
