// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: patient.proto

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

type Patient struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PatientId    string `protobuf:"bytes,1,opt,name=patient_id,json=patientId,proto3" json:"patient_id"`
	Name         string `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	Age          int32  `protobuf:"varint,3,opt,name=age,proto3" json:"age"`
	Gender       int32  `protobuf:"varint,4,opt,name=gender,proto3" json:"gender"`
	Login        string `protobuf:"bytes,5,opt,name=login,proto3" json:"login"`
	Password     string `protobuf:"bytes,6,opt,name=password,proto3" json:"password"`
	AccessToken  string `protobuf:"bytes,7,opt,name=access_token,json=accessToken,proto3" json:"access_token"`
	RefreshToken string `protobuf:"bytes,8,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token"`
}

func (x *Patient) Reset() {
	*x = Patient{}
	if protoimpl.UnsafeEnabled {
		mi := &file_patient_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Patient) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Patient) ProtoMessage() {}

func (x *Patient) ProtoReflect() protoreflect.Message {
	mi := &file_patient_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Patient.ProtoReflect.Descriptor instead.
func (*Patient) Descriptor() ([]byte, []int) {
	return file_patient_proto_rawDescGZIP(), []int{0}
}

func (x *Patient) GetPatientId() string {
	if x != nil {
		return x.PatientId
	}
	return ""
}

func (x *Patient) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Patient) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *Patient) GetGender() int32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *Patient) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *Patient) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Patient) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *Patient) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type GetListPatientsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  int32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit"`
	Offset int32 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset"`
}

func (x *GetListPatientsRequest) Reset() {
	*x = GetListPatientsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_patient_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListPatientsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListPatientsRequest) ProtoMessage() {}

func (x *GetListPatientsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_patient_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListPatientsRequest.ProtoReflect.Descriptor instead.
func (*GetListPatientsRequest) Descriptor() ([]byte, []int) {
	return file_patient_proto_rawDescGZIP(), []int{1}
}

func (x *GetListPatientsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetListPatientsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type GetListPatientsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Patients []*Patient `protobuf:"bytes,1,rep,name=patients,proto3" json:"patients"`
	Count    int32      `protobuf:"varint,2,opt,name=count,proto3" json:"count"`
}

func (x *GetListPatientsResponse) Reset() {
	*x = GetListPatientsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_patient_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListPatientsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListPatientsResponse) ProtoMessage() {}

func (x *GetListPatientsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_patient_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListPatientsResponse.ProtoReflect.Descriptor instead.
func (*GetListPatientsResponse) Descriptor() ([]byte, []int) {
	return file_patient_proto_rawDescGZIP(), []int{2}
}

func (x *GetListPatientsResponse) GetPatients() []*Patient {
	if x != nil {
		return x.Patients
	}
	return nil
}

func (x *GetListPatientsResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type PatientId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PatientId string `protobuf:"bytes,1,opt,name=patient_id,json=patientId,proto3" json:"patient_id"`
}

func (x *PatientId) Reset() {
	*x = PatientId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_patient_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PatientId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatientId) ProtoMessage() {}

func (x *PatientId) ProtoReflect() protoreflect.Message {
	mi := &file_patient_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatientId.ProtoReflect.Descriptor instead.
func (*PatientId) Descriptor() ([]byte, []int) {
	return file_patient_proto_rawDescGZIP(), []int{3}
}

func (x *PatientId) GetPatientId() string {
	if x != nil {
		return x.PatientId
	}
	return ""
}

var File_patient_proto protoreflect.FileDescriptor

var file_patient_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x13, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xe0, 0x01, 0x0a, 0x07, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61,
	0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f,
	0x67, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x21, 0x0a, 0x0c,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x46, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x69, 0x0a, 0x17,
	0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x08, 0x70, 0x61, 0x74, 0x69, 0x65,
	0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x70, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x2a, 0x0a, 0x09, 0x50, 0x61, 0x74, 0x69, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x32, 0x87, 0x03, 0x0a, 0x0e, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x1c, 0x2e, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x66, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x2b, 0x2e, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2c, 0x2e, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61,
	0x74, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x45, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x61,
	0x74, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x61,
	0x74, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x1c, 0x2e, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x06, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x1e, 0x5a,
	0x1c, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_patient_proto_rawDescOnce sync.Once
	file_patient_proto_rawDescData = file_patient_proto_rawDesc
)

func file_patient_proto_rawDescGZIP() []byte {
	file_patient_proto_rawDescOnce.Do(func() {
		file_patient_proto_rawDescData = protoimpl.X.CompressGZIP(file_patient_proto_rawDescData)
	})
	return file_patient_proto_rawDescData
}

var file_patient_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_patient_proto_goTypes = []interface{}{
	(*Patient)(nil),                 // 0: appointment_service.Patient
	(*GetListPatientsRequest)(nil),  // 1: appointment_service.GetListPatientsRequest
	(*GetListPatientsResponse)(nil), // 2: appointment_service.GetListPatientsResponse
	(*PatientId)(nil),               // 3: appointment_service.PatientId
	(*emptypb.Empty)(nil),           // 4: google.protobuf.Empty
}
var file_patient_proto_depIdxs = []int32{
	0, // 0: appointment_service.GetListPatientsResponse.patients:type_name -> appointment_service.Patient
	0, // 1: appointment_service.PatientService.Create:input_type -> appointment_service.Patient
	1, // 2: appointment_service.PatientService.GetList:input_type -> appointment_service.GetListPatientsRequest
	3, // 3: appointment_service.PatientService.Get:input_type -> appointment_service.PatientId
	0, // 4: appointment_service.PatientService.Update:input_type -> appointment_service.Patient
	3, // 5: appointment_service.PatientService.Delete:input_type -> appointment_service.PatientId
	4, // 6: appointment_service.PatientService.Create:output_type -> google.protobuf.Empty
	2, // 7: appointment_service.PatientService.GetList:output_type -> appointment_service.GetListPatientsResponse
	0, // 8: appointment_service.PatientService.Get:output_type -> appointment_service.Patient
	4, // 9: appointment_service.PatientService.Update:output_type -> google.protobuf.Empty
	4, // 10: appointment_service.PatientService.Delete:output_type -> google.protobuf.Empty
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_patient_proto_init() }
func file_patient_proto_init() {
	if File_patient_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_patient_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Patient); i {
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
		file_patient_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListPatientsRequest); i {
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
		file_patient_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListPatientsResponse); i {
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
		file_patient_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PatientId); i {
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
			RawDescriptor: file_patient_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_patient_proto_goTypes,
		DependencyIndexes: file_patient_proto_depIdxs,
		MessageInfos:      file_patient_proto_msgTypes,
	}.Build()
	File_patient_proto = out.File
	file_patient_proto_rawDesc = nil
	file_patient_proto_goTypes = nil
	file_patient_proto_depIdxs = nil
}
