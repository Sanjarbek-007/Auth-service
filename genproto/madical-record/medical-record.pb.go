// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: medical-record.proto

package madical_record

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

type CreateMedicalReportReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      string        `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	RecordType  string        `protobuf:"bytes,2,opt,name=record_type,json=recordType,proto3" json:"record_type,omitempty"`
	RecordDate  string        `protobuf:"bytes,3,opt,name=record_date,json=recordDate,proto3" json:"record_date,omitempty"`
	Description string        `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	DoctorId    string        `protobuf:"bytes,5,opt,name=doctor_id,json=doctorId,proto3" json:"doctor_id,omitempty"`
	Attachments []*Attachment `protobuf:"bytes,6,rep,name=attachments,proto3" json:"attachments,omitempty"`
}

func (x *CreateMedicalReportReq) Reset() {
	*x = CreateMedicalReportReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_medical_record_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMedicalReportReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMedicalReportReq) ProtoMessage() {}

func (x *CreateMedicalReportReq) ProtoReflect() protoreflect.Message {
	mi := &file_medical_record_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMedicalReportReq.ProtoReflect.Descriptor instead.
func (*CreateMedicalReportReq) Descriptor() ([]byte, []int) {
	return file_medical_record_proto_rawDescGZIP(), []int{0}
}

func (x *CreateMedicalReportReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateMedicalReportReq) GetRecordType() string {
	if x != nil {
		return x.RecordType
	}
	return ""
}

func (x *CreateMedicalReportReq) GetRecordDate() string {
	if x != nil {
		return x.RecordDate
	}
	return ""
}

func (x *CreateMedicalReportReq) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateMedicalReportReq) GetDoctorId() string {
	if x != nil {
		return x.DoctorId
	}
	return ""
}

func (x *CreateMedicalReportReq) GetAttachments() []*Attachment {
	if x != nil {
		return x.Attachments
	}
	return nil
}

type Attachment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileUrl string `protobuf:"bytes,1,opt,name=file_url,json=fileUrl,proto3" json:"file_url,omitempty"`
}

func (x *Attachment) Reset() {
	*x = Attachment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_medical_record_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Attachment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Attachment) ProtoMessage() {}

func (x *Attachment) ProtoReflect() protoreflect.Message {
	mi := &file_medical_record_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Attachment.ProtoReflect.Descriptor instead.
func (*Attachment) Descriptor() ([]byte, []int) {
	return file_medical_record_proto_rawDescGZIP(), []int{1}
}

func (x *Attachment) GetFileUrl() string {
	if x != nil {
		return x.FileUrl
	}
	return ""
}

type CreateMedicalReportRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CreateMedicalReportRes) Reset() {
	*x = CreateMedicalReportRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_medical_record_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMedicalReportRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMedicalReportRes) ProtoMessage() {}

func (x *CreateMedicalReportRes) ProtoReflect() protoreflect.Message {
	mi := &file_medical_record_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMedicalReportRes.ProtoReflect.Descriptor instead.
func (*CreateMedicalReportRes) Descriptor() ([]byte, []int) {
	return file_medical_record_proto_rawDescGZIP(), []int{2}
}

func (x *CreateMedicalReportRes) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetMedicalReportReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetMedicalReportReq) Reset() {
	*x = GetMedicalReportReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_medical_record_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMedicalReportReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMedicalReportReq) ProtoMessage() {}

func (x *GetMedicalReportReq) ProtoReflect() protoreflect.Message {
	mi := &file_medical_record_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMedicalReportReq.ProtoReflect.Descriptor instead.
func (*GetMedicalReportReq) Descriptor() ([]byte, []int) {
	return file_medical_record_proto_rawDescGZIP(), []int{3}
}

func (x *GetMedicalReportReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetMedicalReportRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MedicalReport *MedicalReport `protobuf:"bytes,1,opt,name=medical_report,json=medicalReport,proto3" json:"medical_report,omitempty"`
}

func (x *GetMedicalReportRes) Reset() {
	*x = GetMedicalReportRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_medical_record_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMedicalReportRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMedicalReportRes) ProtoMessage() {}

func (x *GetMedicalReportRes) ProtoReflect() protoreflect.Message {
	mi := &file_medical_record_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMedicalReportRes.ProtoReflect.Descriptor instead.
func (*GetMedicalReportRes) Descriptor() ([]byte, []int) {
	return file_medical_record_proto_rawDescGZIP(), []int{4}
}

func (x *GetMedicalReportRes) GetMedicalReport() *MedicalReport {
	if x != nil {
		return x.MedicalReport
	}
	return nil
}

type MedicalReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MedicalReport) Reset() {
	*x = MedicalReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_medical_record_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MedicalReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MedicalReport) ProtoMessage() {}

func (x *MedicalReport) ProtoReflect() protoreflect.Message {
	mi := &file_medical_record_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MedicalReport.ProtoReflect.Descriptor instead.
func (*MedicalReport) Descriptor() ([]byte, []int) {
	return file_medical_record_proto_rawDescGZIP(), []int{5}
}

type GetMedicalReportByIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetMedicalReportByIdReq) Reset() {
	*x = GetMedicalReportByIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_medical_record_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMedicalReportByIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMedicalReportByIdReq) ProtoMessage() {}

func (x *GetMedicalReportByIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_medical_record_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMedicalReportByIdReq.ProtoReflect.Descriptor instead.
func (*GetMedicalReportByIdReq) Descriptor() ([]byte, []int) {
	return file_medical_record_proto_rawDescGZIP(), []int{6}
}

type GetMedicalReportByIdRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetMedicalReportByIdRes) Reset() {
	*x = GetMedicalReportByIdRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_medical_record_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMedicalReportByIdRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMedicalReportByIdRes) ProtoMessage() {}

func (x *GetMedicalReportByIdRes) ProtoReflect() protoreflect.Message {
	mi := &file_medical_record_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMedicalReportByIdRes.ProtoReflect.Descriptor instead.
func (*GetMedicalReportByIdRes) Descriptor() ([]byte, []int) {
	return file_medical_record_proto_rawDescGZIP(), []int{7}
}

type UpdateMedicalReportReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateMedicalReportReq) Reset() {
	*x = UpdateMedicalReportReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_medical_record_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMedicalReportReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMedicalReportReq) ProtoMessage() {}

func (x *UpdateMedicalReportReq) ProtoReflect() protoreflect.Message {
	mi := &file_medical_record_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMedicalReportReq.ProtoReflect.Descriptor instead.
func (*UpdateMedicalReportReq) Descriptor() ([]byte, []int) {
	return file_medical_record_proto_rawDescGZIP(), []int{8}
}

type UpdateMedicalReportRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateMedicalReportRes) Reset() {
	*x = UpdateMedicalReportRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_medical_record_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMedicalReportRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMedicalReportRes) ProtoMessage() {}

func (x *UpdateMedicalReportRes) ProtoReflect() protoreflect.Message {
	mi := &file_medical_record_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMedicalReportRes.ProtoReflect.Descriptor instead.
func (*UpdateMedicalReportRes) Descriptor() ([]byte, []int) {
	return file_medical_record_proto_rawDescGZIP(), []int{9}
}

type DeleteMedicalReportReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteMedicalReportReq) Reset() {
	*x = DeleteMedicalReportReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_medical_record_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMedicalReportReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMedicalReportReq) ProtoMessage() {}

func (x *DeleteMedicalReportReq) ProtoReflect() protoreflect.Message {
	mi := &file_medical_record_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMedicalReportReq.ProtoReflect.Descriptor instead.
func (*DeleteMedicalReportReq) Descriptor() ([]byte, []int) {
	return file_medical_record_proto_rawDescGZIP(), []int{10}
}

type DeleteMedicalReportRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteMedicalReportRes) Reset() {
	*x = DeleteMedicalReportRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_medical_record_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMedicalReportRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMedicalReportRes) ProtoMessage() {}

func (x *DeleteMedicalReportRes) ProtoReflect() protoreflect.Message {
	mi := &file_medical_record_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMedicalReportRes.ProtoReflect.Descriptor instead.
func (*DeleteMedicalReportRes) Descriptor() ([]byte, []int) {
	return file_medical_record_proto_rawDescGZIP(), []int{11}
}

var File_medical_record_proto protoreflect.FileDescriptor

var file_medical_record_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6d, 0x65, 0x64, 0x69, 0x63, 0x61, 0x6c, 0x2d, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x22, 0xe8,
	0x01, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x64, 0x69, 0x63, 0x61, 0x6c,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x44, 0x61, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x63, 0x74, 0x6f,
	0x72, 0x49, 0x64, 0x12, 0x34, 0x0a, 0x0b, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x61, 0x74,
	0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x27, 0x0a, 0x0a, 0x41, 0x74, 0x74,
	0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x55,
	0x72, 0x6c, 0x22, 0x32, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x64, 0x69,
	0x63, 0x61, 0x6c, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2e, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x64,
	0x69, 0x63, 0x61, 0x6c, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x53, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x64,
	0x69, 0x63, 0x61, 0x6c, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x12, 0x3c, 0x0a,
	0x0e, 0x6d, 0x65, 0x64, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x4d,
	0x65, 0x64, 0x69, 0x63, 0x61, 0x6c, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x0d, 0x6d, 0x65,
	0x64, 0x69, 0x63, 0x61, 0x6c, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x0f, 0x0a, 0x0d, 0x4d,
	0x65, 0x64, 0x69, 0x63, 0x61, 0x6c, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x19, 0x0a, 0x17,
	0x47, 0x65, 0x74, 0x4d, 0x65, 0x64, 0x69, 0x63, 0x61, 0x6c, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x22, 0x19, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x4d, 0x65,
	0x64, 0x69, 0x63, 0x61, 0x6c, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52,
	0x65, 0x73, 0x22, 0x18, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x64, 0x69,
	0x63, 0x61, 0x6c, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x22, 0x18, 0x0a, 0x16,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x64, 0x69, 0x63, 0x61, 0x6c, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x22, 0x18, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x4d, 0x65, 0x64, 0x69, 0x63, 0x61, 0x6c, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71,
	0x22, 0x18, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x64, 0x69, 0x63, 0x61,
	0x6c, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x32, 0xbc, 0x03, 0x0a, 0x0d, 0x4d,
	0x65, 0x64, 0x69, 0x63, 0x61, 0x6c, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x55, 0x0a, 0x13,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x64, 0x69, 0x63, 0x61, 0x6c, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x12, 0x1e, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4d, 0x65, 0x64, 0x69, 0x63, 0x61, 0x6c, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4d, 0x65, 0x64, 0x69, 0x63, 0x61, 0x6c, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x52, 0x65, 0x73, 0x12, 0x4c, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x64, 0x69, 0x63, 0x61,
	0x6c, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1b, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x64, 0x69, 0x63, 0x61, 0x6c, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x47, 0x65,
	0x74, 0x4d, 0x65, 0x64, 0x69, 0x63, 0x61, 0x6c, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65,
	0x73, 0x12, 0x58, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x64, 0x69, 0x63, 0x61, 0x6c, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x1f, 0x2e, 0x68, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x64, 0x69, 0x63, 0x61, 0x6c, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x1f, 0x2e, 0x68, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x64, 0x69, 0x63, 0x61, 0x6c, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x12, 0x55, 0x0a, 0x13, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x64, 0x69, 0x63, 0x61, 0x6c, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x12, 0x1e, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4d, 0x65, 0x64, 0x69, 0x63, 0x61, 0x6c, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52,
	0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4d, 0x65, 0x64, 0x69, 0x63, 0x61, 0x6c, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52,
	0x65, 0x73, 0x12, 0x55, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x64, 0x69,
	0x63, 0x61, 0x6c, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1e, 0x2e, 0x68, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x64, 0x69, 0x63, 0x61, 0x6c,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x68, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x64, 0x69, 0x63, 0x61, 0x6c,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x42, 0x19, 0x5a, 0x17, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x61, 0x64, 0x69, 0x63, 0x61, 0x6c, 0x2d, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_medical_record_proto_rawDescOnce sync.Once
	file_medical_record_proto_rawDescData = file_medical_record_proto_rawDesc
)

func file_medical_record_proto_rawDescGZIP() []byte {
	file_medical_record_proto_rawDescOnce.Do(func() {
		file_medical_record_proto_rawDescData = protoimpl.X.CompressGZIP(file_medical_record_proto_rawDescData)
	})
	return file_medical_record_proto_rawDescData
}

var file_medical_record_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_medical_record_proto_goTypes = []any{
	(*CreateMedicalReportReq)(nil),  // 0: health.CreateMedicalReportReq
	(*Attachment)(nil),              // 1: health.Attachment
	(*CreateMedicalReportRes)(nil),  // 2: health.CreateMedicalReportRes
	(*GetMedicalReportReq)(nil),     // 3: health.GetMedicalReportReq
	(*GetMedicalReportRes)(nil),     // 4: health.GetMedicalReportRes
	(*MedicalReport)(nil),           // 5: health.MedicalReport
	(*GetMedicalReportByIdReq)(nil), // 6: health.GetMedicalReportByIdReq
	(*GetMedicalReportByIdRes)(nil), // 7: health.GetMedicalReportByIdRes
	(*UpdateMedicalReportReq)(nil),  // 8: health.UpdateMedicalReportReq
	(*UpdateMedicalReportRes)(nil),  // 9: health.UpdateMedicalReportRes
	(*DeleteMedicalReportReq)(nil),  // 10: health.DeleteMedicalReportReq
	(*DeleteMedicalReportRes)(nil),  // 11: health.DeleteMedicalReportRes
}
var file_medical_record_proto_depIdxs = []int32{
	1,  // 0: health.CreateMedicalReportReq.attachments:type_name -> health.Attachment
	5,  // 1: health.GetMedicalReportRes.medical_report:type_name -> health.MedicalReport
	0,  // 2: health.MedicalRecord.CreateMedicalReport:input_type -> health.CreateMedicalReportReq
	3,  // 3: health.MedicalRecord.GetMedicalReport:input_type -> health.GetMedicalReportReq
	6,  // 4: health.MedicalRecord.GetMedicalReportById:input_type -> health.GetMedicalReportByIdReq
	8,  // 5: health.MedicalRecord.UpdateMedicalReport:input_type -> health.UpdateMedicalReportReq
	10, // 6: health.MedicalRecord.DeleteMedicalReport:input_type -> health.DeleteMedicalReportReq
	2,  // 7: health.MedicalRecord.CreateMedicalReport:output_type -> health.CreateMedicalReportRes
	4,  // 8: health.MedicalRecord.GetMedicalReport:output_type -> health.GetMedicalReportRes
	7,  // 9: health.MedicalRecord.GetMedicalReportById:output_type -> health.GetMedicalReportByIdRes
	9,  // 10: health.MedicalRecord.UpdateMedicalReport:output_type -> health.UpdateMedicalReportRes
	11, // 11: health.MedicalRecord.DeleteMedicalReport:output_type -> health.DeleteMedicalReportRes
	7,  // [7:12] is the sub-list for method output_type
	2,  // [2:7] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_medical_record_proto_init() }
func file_medical_record_proto_init() {
	if File_medical_record_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_medical_record_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateMedicalReportReq); i {
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
		file_medical_record_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Attachment); i {
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
		file_medical_record_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CreateMedicalReportRes); i {
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
		file_medical_record_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetMedicalReportReq); i {
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
		file_medical_record_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetMedicalReportRes); i {
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
		file_medical_record_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*MedicalReport); i {
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
		file_medical_record_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*GetMedicalReportByIdReq); i {
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
		file_medical_record_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*GetMedicalReportByIdRes); i {
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
		file_medical_record_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateMedicalReportReq); i {
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
		file_medical_record_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateMedicalReportRes); i {
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
		file_medical_record_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteMedicalReportReq); i {
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
		file_medical_record_proto_msgTypes[11].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteMedicalReportRes); i {
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
			RawDescriptor: file_medical_record_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_medical_record_proto_goTypes,
		DependencyIndexes: file_medical_record_proto_depIdxs,
		MessageInfos:      file_medical_record_proto_msgTypes,
	}.Build()
	File_medical_record_proto = out.File
	file_medical_record_proto_rawDesc = nil
	file_medical_record_proto_goTypes = nil
	file_medical_record_proto_depIdxs = nil
}
