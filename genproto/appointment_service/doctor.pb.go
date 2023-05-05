// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: doctor.proto

package appointment_service

import (
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

type Doctor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DoctorId   string `protobuf:"bytes,1,opt,name=doctor_id,json=doctorId,proto3" json:"doctor_id"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	Age        int32  `protobuf:"varint,3,opt,name=age,proto3" json:"age"`
	Role       string `protobuf:"bytes,4,opt,name=role,proto3" json:"role"`
	Polyclinic string `protobuf:"bytes,5,opt,name=polyclinic,proto3" json:"polyclinic"`
	Gender     int32  `protobuf:"varint,6,opt,name=gender,proto3" json:"gender"`
}

func (x *Doctor) Reset() {
	*x = Doctor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doctor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Doctor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Doctor) ProtoMessage() {}

func (x *Doctor) ProtoReflect() protoreflect.Message {
	mi := &file_doctor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Doctor.ProtoReflect.Descriptor instead.
func (*Doctor) Descriptor() ([]byte, []int) {
	return file_doctor_proto_rawDescGZIP(), []int{0}
}

func (x *Doctor) GetDoctorId() string {
	if x != nil {
		return x.DoctorId
	}
	return ""
}

func (x *Doctor) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Doctor) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *Doctor) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *Doctor) GetPolyclinic() string {
	if x != nil {
		return x.Polyclinic
	}
	return ""
}

func (x *Doctor) GetGender() int32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

type GetListDoctorsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  int32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit"`
	Offset int32 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset"`
}

func (x *GetListDoctorsRequest) Reset() {
	*x = GetListDoctorsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doctor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListDoctorsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListDoctorsRequest) ProtoMessage() {}

func (x *GetListDoctorsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_doctor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListDoctorsRequest.ProtoReflect.Descriptor instead.
func (*GetListDoctorsRequest) Descriptor() ([]byte, []int) {
	return file_doctor_proto_rawDescGZIP(), []int{1}
}

func (x *GetListDoctorsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetListDoctorsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type GetListDoctorsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Doctors []*Doctor `protobuf:"bytes,1,rep,name=doctors,proto3" json:"doctors"`
	Count   int32     `protobuf:"varint,2,opt,name=count,proto3" json:"count"`
}

func (x *GetListDoctorsResponse) Reset() {
	*x = GetListDoctorsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doctor_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListDoctorsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListDoctorsResponse) ProtoMessage() {}

func (x *GetListDoctorsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_doctor_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListDoctorsResponse.ProtoReflect.Descriptor instead.
func (*GetListDoctorsResponse) Descriptor() ([]byte, []int) {
	return file_doctor_proto_rawDescGZIP(), []int{2}
}

func (x *GetListDoctorsResponse) GetDoctors() []*Doctor {
	if x != nil {
		return x.Doctors
	}
	return nil
}

func (x *GetListDoctorsResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type DoctorId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DoctorId string `protobuf:"bytes,1,opt,name=doctor_id,json=doctorId,proto3" json:"doctor_id"`
}

func (x *DoctorId) Reset() {
	*x = DoctorId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doctor_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoctorId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoctorId) ProtoMessage() {}

func (x *DoctorId) ProtoReflect() protoreflect.Message {
	mi := &file_doctor_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoctorId.ProtoReflect.Descriptor instead.
func (*DoctorId) Descriptor() ([]byte, []int) {
	return file_doctor_proto_rawDescGZIP(), []int{3}
}

func (x *DoctorId) GetDoctorId() string {
	if x != nil {
		return x.DoctorId
	}
	return ""
}

var File_doctor_proto protoreflect.FileDescriptor

var file_doctor_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13,
	0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x97, 0x01, 0x0a, 0x06, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x64,
	0x6f, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f,
	0x6c, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x6f, 0x6c, 0x79, 0x63, 0x6c, 0x69, 0x6e, 0x69, 0x63,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6f, 0x6c, 0x79, 0x63, 0x6c, 0x69, 0x6e,
	0x69, 0x63, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x22, 0x45, 0x0a, 0x15, 0x47, 0x65,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x22, 0x65, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x6f, 0x63, 0x74,
	0x6f, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x07, 0x64,
	0x6f, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x61,
	0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x07, 0x64, 0x6f, 0x63, 0x74, 0x6f,
	0x72, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x27, 0x0a, 0x08, 0x44, 0x6f, 0x63, 0x74,
	0x6f, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x49,
	0x64, 0x32, 0xff, 0x02, 0x0a, 0x0d, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x2e,
	0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x12, 0x64, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x2a, 0x2e, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x6f, 0x63,
	0x74, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x61, 0x70,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x03, 0x47, 0x65,
	0x74, 0x12, 0x1d, 0x2e, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x64,
	0x1a, 0x1b, 0x2e, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x22, 0x00, 0x12,
	0x3f, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x61, 0x70, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x12, 0x41, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x61, 0x70, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x00, 0x42, 0x1e, 0x5a, 0x1c, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_doctor_proto_rawDescOnce sync.Once
	file_doctor_proto_rawDescData = file_doctor_proto_rawDesc
)

func file_doctor_proto_rawDescGZIP() []byte {
	file_doctor_proto_rawDescOnce.Do(func() {
		file_doctor_proto_rawDescData = protoimpl.X.CompressGZIP(file_doctor_proto_rawDescData)
	})
	return file_doctor_proto_rawDescData
}

var file_doctor_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_doctor_proto_goTypes = []interface{}{
	(*Doctor)(nil),                 // 0: appointment_service.Doctor
	(*GetListDoctorsRequest)(nil),  // 1: appointment_service.GetListDoctorsRequest
	(*GetListDoctorsResponse)(nil), // 2: appointment_service.GetListDoctorsResponse
	(*DoctorId)(nil),               // 3: appointment_service.DoctorId
	(*emptypb.Empty)(nil),          // 4: google.protobuf.Empty
}
var file_doctor_proto_depIdxs = []int32{
	0, // 0: appointment_service.GetListDoctorsResponse.doctors:type_name -> appointment_service.Doctor
	0, // 1: appointment_service.DoctorService.Create:input_type -> appointment_service.Doctor
	1, // 2: appointment_service.DoctorService.GetList:input_type -> appointment_service.GetListDoctorsRequest
	3, // 3: appointment_service.DoctorService.Get:input_type -> appointment_service.DoctorId
	0, // 4: appointment_service.DoctorService.Update:input_type -> appointment_service.Doctor
	3, // 5: appointment_service.DoctorService.Delete:input_type -> appointment_service.DoctorId
	4, // 6: appointment_service.DoctorService.Create:output_type -> google.protobuf.Empty
	2, // 7: appointment_service.DoctorService.GetList:output_type -> appointment_service.GetListDoctorsResponse
	0, // 8: appointment_service.DoctorService.Get:output_type -> appointment_service.Doctor
	4, // 9: appointment_service.DoctorService.Update:output_type -> google.protobuf.Empty
	4, // 10: appointment_service.DoctorService.Delete:output_type -> google.protobuf.Empty
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_doctor_proto_init() }
func file_doctor_proto_init() {
	if File_doctor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_doctor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Doctor); i {
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
		file_doctor_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListDoctorsRequest); i {
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
		file_doctor_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListDoctorsResponse); i {
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
		file_doctor_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoctorId); i {
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
			RawDescriptor: file_doctor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_doctor_proto_goTypes,
		DependencyIndexes: file_doctor_proto_depIdxs,
		MessageInfos:      file_doctor_proto_msgTypes,
	}.Build()
	File_doctor_proto = out.File
	file_doctor_proto_rawDesc = nil
	file_doctor_proto_goTypes = nil
	file_doctor_proto_depIdxs = nil
}