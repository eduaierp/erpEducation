// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: proto/academic.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Program struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Program) Reset() {
	*x = Program{}
	mi := &file_proto_academic_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Program) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Program) ProtoMessage() {}

func (x *Program) ProtoReflect() protoreflect.Message {
	mi := &file_proto_academic_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Program.ProtoReflect.Descriptor instead.
func (*Program) Descriptor() ([]byte, []int) {
	return file_proto_academic_proto_rawDescGZIP(), []int{0}
}

func (x *Program) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Program) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Program) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type Course struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Code          string                 `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	ProgramId     int32                  `protobuf:"varint,4,opt,name=program_id,json=programId,proto3" json:"program_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Course) Reset() {
	*x = Course{}
	mi := &file_proto_academic_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Course) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Course) ProtoMessage() {}

func (x *Course) ProtoReflect() protoreflect.Message {
	mi := &file_proto_academic_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Course.ProtoReflect.Descriptor instead.
func (*Course) Descriptor() ([]byte, []int) {
	return file_proto_academic_proto_rawDescGZIP(), []int{1}
}

func (x *Course) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Course) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Course) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Course) GetProgramId() int32 {
	if x != nil {
		return x.ProgramId
	}
	return 0
}

type CreateProgramRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateProgramRequest) Reset() {
	*x = CreateProgramRequest{}
	mi := &file_proto_academic_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateProgramRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProgramRequest) ProtoMessage() {}

func (x *CreateProgramRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_academic_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProgramRequest.ProtoReflect.Descriptor instead.
func (*CreateProgramRequest) Descriptor() ([]byte, []int) {
	return file_proto_academic_proto_rawDescGZIP(), []int{2}
}

func (x *CreateProgramRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateProgramRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type GetProgramRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetProgramRequest) Reset() {
	*x = GetProgramRequest{}
	mi := &file_proto_academic_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProgramRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProgramRequest) ProtoMessage() {}

func (x *GetProgramRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_academic_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProgramRequest.ProtoReflect.Descriptor instead.
func (*GetProgramRequest) Descriptor() ([]byte, []int) {
	return file_proto_academic_proto_rawDescGZIP(), []int{3}
}

func (x *GetProgramRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UpdateProgramRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateProgramRequest) Reset() {
	*x = UpdateProgramRequest{}
	mi := &file_proto_academic_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateProgramRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProgramRequest) ProtoMessage() {}

func (x *UpdateProgramRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_academic_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProgramRequest.ProtoReflect.Descriptor instead.
func (*UpdateProgramRequest) Descriptor() ([]byte, []int) {
	return file_proto_academic_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateProgramRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateProgramRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateProgramRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type DeleteProgramRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteProgramRequest) Reset() {
	*x = DeleteProgramRequest{}
	mi := &file_proto_academic_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteProgramRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProgramRequest) ProtoMessage() {}

func (x *DeleteProgramRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_academic_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProgramRequest.ProtoReflect.Descriptor instead.
func (*DeleteProgramRequest) Descriptor() ([]byte, []int) {
	return file_proto_academic_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteProgramRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteProgramResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteProgramResponse) Reset() {
	*x = DeleteProgramResponse{}
	mi := &file_proto_academic_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteProgramResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProgramResponse) ProtoMessage() {}

func (x *DeleteProgramResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_academic_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProgramResponse.ProtoReflect.Descriptor instead.
func (*DeleteProgramResponse) Descriptor() ([]byte, []int) {
	return file_proto_academic_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteProgramResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ProgramResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Program       *Program               `protobuf:"bytes,1,opt,name=program,proto3" json:"program,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProgramResponse) Reset() {
	*x = ProgramResponse{}
	mi := &file_proto_academic_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProgramResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProgramResponse) ProtoMessage() {}

func (x *ProgramResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_academic_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProgramResponse.ProtoReflect.Descriptor instead.
func (*ProgramResponse) Descriptor() ([]byte, []int) {
	return file_proto_academic_proto_rawDescGZIP(), []int{7}
}

func (x *ProgramResponse) GetProgram() *Program {
	if x != nil {
		return x.Program
	}
	return nil
}

type ListProgramsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListProgramsRequest) Reset() {
	*x = ListProgramsRequest{}
	mi := &file_proto_academic_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListProgramsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProgramsRequest) ProtoMessage() {}

func (x *ListProgramsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_academic_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProgramsRequest.ProtoReflect.Descriptor instead.
func (*ListProgramsRequest) Descriptor() ([]byte, []int) {
	return file_proto_academic_proto_rawDescGZIP(), []int{8}
}

type ListProgramsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Programs      []*Program             `protobuf:"bytes,1,rep,name=programs,proto3" json:"programs,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListProgramsResponse) Reset() {
	*x = ListProgramsResponse{}
	mi := &file_proto_academic_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListProgramsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProgramsResponse) ProtoMessage() {}

func (x *ListProgramsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_academic_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProgramsResponse.ProtoReflect.Descriptor instead.
func (*ListProgramsResponse) Descriptor() ([]byte, []int) {
	return file_proto_academic_proto_rawDescGZIP(), []int{9}
}

func (x *ListProgramsResponse) GetPrograms() []*Program {
	if x != nil {
		return x.Programs
	}
	return nil
}

type CreateCourseRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Title         string                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Code          string                 `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	ProgramId     int32                  `protobuf:"varint,3,opt,name=program_id,json=programId,proto3" json:"program_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateCourseRequest) Reset() {
	*x = CreateCourseRequest{}
	mi := &file_proto_academic_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateCourseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCourseRequest) ProtoMessage() {}

func (x *CreateCourseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_academic_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCourseRequest.ProtoReflect.Descriptor instead.
func (*CreateCourseRequest) Descriptor() ([]byte, []int) {
	return file_proto_academic_proto_rawDescGZIP(), []int{10}
}

func (x *CreateCourseRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateCourseRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *CreateCourseRequest) GetProgramId() int32 {
	if x != nil {
		return x.ProgramId
	}
	return 0
}

type GetCourseRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCourseRequest) Reset() {
	*x = GetCourseRequest{}
	mi := &file_proto_academic_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCourseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCourseRequest) ProtoMessage() {}

func (x *GetCourseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_academic_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCourseRequest.ProtoReflect.Descriptor instead.
func (*GetCourseRequest) Descriptor() ([]byte, []int) {
	return file_proto_academic_proto_rawDescGZIP(), []int{11}
}

func (x *GetCourseRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UpdateCourseRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Code          string                 `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	ProgramId     int32                  `protobuf:"varint,4,opt,name=program_id,json=programId,proto3" json:"program_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateCourseRequest) Reset() {
	*x = UpdateCourseRequest{}
	mi := &file_proto_academic_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateCourseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCourseRequest) ProtoMessage() {}

func (x *UpdateCourseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_academic_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCourseRequest.ProtoReflect.Descriptor instead.
func (*UpdateCourseRequest) Descriptor() ([]byte, []int) {
	return file_proto_academic_proto_rawDescGZIP(), []int{12}
}

func (x *UpdateCourseRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateCourseRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateCourseRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *UpdateCourseRequest) GetProgramId() int32 {
	if x != nil {
		return x.ProgramId
	}
	return 0
}

type DeleteCourseRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteCourseRequest) Reset() {
	*x = DeleteCourseRequest{}
	mi := &file_proto_academic_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteCourseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCourseRequest) ProtoMessage() {}

func (x *DeleteCourseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_academic_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCourseRequest.ProtoReflect.Descriptor instead.
func (*DeleteCourseRequest) Descriptor() ([]byte, []int) {
	return file_proto_academic_proto_rawDescGZIP(), []int{13}
}

func (x *DeleteCourseRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteCourseResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteCourseResponse) Reset() {
	*x = DeleteCourseResponse{}
	mi := &file_proto_academic_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteCourseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCourseResponse) ProtoMessage() {}

func (x *DeleteCourseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_academic_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCourseResponse.ProtoReflect.Descriptor instead.
func (*DeleteCourseResponse) Descriptor() ([]byte, []int) {
	return file_proto_academic_proto_rawDescGZIP(), []int{14}
}

func (x *DeleteCourseResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type CourseResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Course        *Course                `protobuf:"bytes,1,opt,name=course,proto3" json:"course,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CourseResponse) Reset() {
	*x = CourseResponse{}
	mi := &file_proto_academic_proto_msgTypes[15]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CourseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CourseResponse) ProtoMessage() {}

func (x *CourseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_academic_proto_msgTypes[15]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CourseResponse.ProtoReflect.Descriptor instead.
func (*CourseResponse) Descriptor() ([]byte, []int) {
	return file_proto_academic_proto_rawDescGZIP(), []int{15}
}

func (x *CourseResponse) GetCourse() *Course {
	if x != nil {
		return x.Course
	}
	return nil
}

type ListCoursesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListCoursesRequest) Reset() {
	*x = ListCoursesRequest{}
	mi := &file_proto_academic_proto_msgTypes[16]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListCoursesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCoursesRequest) ProtoMessage() {}

func (x *ListCoursesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_academic_proto_msgTypes[16]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCoursesRequest.ProtoReflect.Descriptor instead.
func (*ListCoursesRequest) Descriptor() ([]byte, []int) {
	return file_proto_academic_proto_rawDescGZIP(), []int{16}
}

type ListCoursesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Courses       []*Course              `protobuf:"bytes,1,rep,name=courses,proto3" json:"courses,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListCoursesResponse) Reset() {
	*x = ListCoursesResponse{}
	mi := &file_proto_academic_proto_msgTypes[17]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListCoursesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCoursesResponse) ProtoMessage() {}

func (x *ListCoursesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_academic_proto_msgTypes[17]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCoursesResponse.ProtoReflect.Descriptor instead.
func (*ListCoursesResponse) Descriptor() ([]byte, []int) {
	return file_proto_academic_proto_rawDescGZIP(), []int{17}
}

func (x *ListCoursesResponse) GetCourses() []*Course {
	if x != nil {
		return x.Courses
	}
	return nil
}

var File_proto_academic_proto protoreflect.FileDescriptor

const file_proto_academic_proto_rawDesc = "" +
	"\n" +
	"\x14proto/academic.proto\x12\bacademic\"O\n" +
	"\aProgram\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\"a\n" +
	"\x06Course\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x12\x14\n" +
	"\x05title\x18\x02 \x01(\tR\x05title\x12\x12\n" +
	"\x04code\x18\x03 \x01(\tR\x04code\x12\x1d\n" +
	"\n" +
	"program_id\x18\x04 \x01(\x05R\tprogramId\"L\n" +
	"\x14CreateProgramRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x02 \x01(\tR\vdescription\"#\n" +
	"\x11GetProgramRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\"\\\n" +
	"\x14UpdateProgramRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\"&\n" +
	"\x14DeleteProgramRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\"1\n" +
	"\x15DeleteProgramResponse\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage\">\n" +
	"\x0fProgramResponse\x12+\n" +
	"\aprogram\x18\x01 \x01(\v2\x11.academic.ProgramR\aprogram\"\x15\n" +
	"\x13ListProgramsRequest\"E\n" +
	"\x14ListProgramsResponse\x12-\n" +
	"\bprograms\x18\x01 \x03(\v2\x11.academic.ProgramR\bprograms\"^\n" +
	"\x13CreateCourseRequest\x12\x14\n" +
	"\x05title\x18\x01 \x01(\tR\x05title\x12\x12\n" +
	"\x04code\x18\x02 \x01(\tR\x04code\x12\x1d\n" +
	"\n" +
	"program_id\x18\x03 \x01(\x05R\tprogramId\"\"\n" +
	"\x10GetCourseRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\"n\n" +
	"\x13UpdateCourseRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x12\x14\n" +
	"\x05title\x18\x02 \x01(\tR\x05title\x12\x12\n" +
	"\x04code\x18\x03 \x01(\tR\x04code\x12\x1d\n" +
	"\n" +
	"program_id\x18\x04 \x01(\x05R\tprogramId\"%\n" +
	"\x13DeleteCourseRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\"0\n" +
	"\x14DeleteCourseResponse\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage\":\n" +
	"\x0eCourseResponse\x12(\n" +
	"\x06course\x18\x01 \x01(\v2\x10.academic.CourseR\x06course\"\x14\n" +
	"\x12ListCoursesRequest\"A\n" +
	"\x13ListCoursesResponse\x12*\n" +
	"\acourses\x18\x01 \x03(\v2\x10.academic.CourseR\acourses2\x80\x06\n" +
	"\x0fAcademicService\x12J\n" +
	"\rCreateProgram\x12\x1e.academic.CreateProgramRequest\x1a\x19.academic.ProgramResponse\x12D\n" +
	"\n" +
	"GetProgram\x12\x1b.academic.GetProgramRequest\x1a\x19.academic.ProgramResponse\x12J\n" +
	"\rUpdateProgram\x12\x1e.academic.UpdateProgramRequest\x1a\x19.academic.ProgramResponse\x12P\n" +
	"\rDeleteProgram\x12\x1e.academic.DeleteProgramRequest\x1a\x1f.academic.DeleteProgramResponse\x12M\n" +
	"\fListPrograms\x12\x1d.academic.ListProgramsRequest\x1a\x1e.academic.ListProgramsResponse\x12G\n" +
	"\fCreateCourse\x12\x1d.academic.CreateCourseRequest\x1a\x18.academic.CourseResponse\x12A\n" +
	"\tGetCourse\x12\x1a.academic.GetCourseRequest\x1a\x18.academic.CourseResponse\x12G\n" +
	"\fUpdateCourse\x12\x1d.academic.UpdateCourseRequest\x1a\x18.academic.CourseResponse\x12M\n" +
	"\fDeleteCourse\x12\x1d.academic.DeleteCourseRequest\x1a\x1e.academic.DeleteCourseResponse\x12J\n" +
	"\vListCourses\x12\x1c.academic.ListCoursesRequest\x1a\x1d.academic.ListCoursesResponseB\tZ\a./protob\x06proto3"

var (
	file_proto_academic_proto_rawDescOnce sync.Once
	file_proto_academic_proto_rawDescData []byte
)

func file_proto_academic_proto_rawDescGZIP() []byte {
	file_proto_academic_proto_rawDescOnce.Do(func() {
		file_proto_academic_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_academic_proto_rawDesc), len(file_proto_academic_proto_rawDesc)))
	})
	return file_proto_academic_proto_rawDescData
}

var file_proto_academic_proto_msgTypes = make([]protoimpl.MessageInfo, 18)
var file_proto_academic_proto_goTypes = []any{
	(*Program)(nil),               // 0: academic.Program
	(*Course)(nil),                // 1: academic.Course
	(*CreateProgramRequest)(nil),  // 2: academic.CreateProgramRequest
	(*GetProgramRequest)(nil),     // 3: academic.GetProgramRequest
	(*UpdateProgramRequest)(nil),  // 4: academic.UpdateProgramRequest
	(*DeleteProgramRequest)(nil),  // 5: academic.DeleteProgramRequest
	(*DeleteProgramResponse)(nil), // 6: academic.DeleteProgramResponse
	(*ProgramResponse)(nil),       // 7: academic.ProgramResponse
	(*ListProgramsRequest)(nil),   // 8: academic.ListProgramsRequest
	(*ListProgramsResponse)(nil),  // 9: academic.ListProgramsResponse
	(*CreateCourseRequest)(nil),   // 10: academic.CreateCourseRequest
	(*GetCourseRequest)(nil),      // 11: academic.GetCourseRequest
	(*UpdateCourseRequest)(nil),   // 12: academic.UpdateCourseRequest
	(*DeleteCourseRequest)(nil),   // 13: academic.DeleteCourseRequest
	(*DeleteCourseResponse)(nil),  // 14: academic.DeleteCourseResponse
	(*CourseResponse)(nil),        // 15: academic.CourseResponse
	(*ListCoursesRequest)(nil),    // 16: academic.ListCoursesRequest
	(*ListCoursesResponse)(nil),   // 17: academic.ListCoursesResponse
}
var file_proto_academic_proto_depIdxs = []int32{
	0,  // 0: academic.ProgramResponse.program:type_name -> academic.Program
	0,  // 1: academic.ListProgramsResponse.programs:type_name -> academic.Program
	1,  // 2: academic.CourseResponse.course:type_name -> academic.Course
	1,  // 3: academic.ListCoursesResponse.courses:type_name -> academic.Course
	2,  // 4: academic.AcademicService.CreateProgram:input_type -> academic.CreateProgramRequest
	3,  // 5: academic.AcademicService.GetProgram:input_type -> academic.GetProgramRequest
	4,  // 6: academic.AcademicService.UpdateProgram:input_type -> academic.UpdateProgramRequest
	5,  // 7: academic.AcademicService.DeleteProgram:input_type -> academic.DeleteProgramRequest
	8,  // 8: academic.AcademicService.ListPrograms:input_type -> academic.ListProgramsRequest
	10, // 9: academic.AcademicService.CreateCourse:input_type -> academic.CreateCourseRequest
	11, // 10: academic.AcademicService.GetCourse:input_type -> academic.GetCourseRequest
	12, // 11: academic.AcademicService.UpdateCourse:input_type -> academic.UpdateCourseRequest
	13, // 12: academic.AcademicService.DeleteCourse:input_type -> academic.DeleteCourseRequest
	16, // 13: academic.AcademicService.ListCourses:input_type -> academic.ListCoursesRequest
	7,  // 14: academic.AcademicService.CreateProgram:output_type -> academic.ProgramResponse
	7,  // 15: academic.AcademicService.GetProgram:output_type -> academic.ProgramResponse
	7,  // 16: academic.AcademicService.UpdateProgram:output_type -> academic.ProgramResponse
	6,  // 17: academic.AcademicService.DeleteProgram:output_type -> academic.DeleteProgramResponse
	9,  // 18: academic.AcademicService.ListPrograms:output_type -> academic.ListProgramsResponse
	15, // 19: academic.AcademicService.CreateCourse:output_type -> academic.CourseResponse
	15, // 20: academic.AcademicService.GetCourse:output_type -> academic.CourseResponse
	15, // 21: academic.AcademicService.UpdateCourse:output_type -> academic.CourseResponse
	14, // 22: academic.AcademicService.DeleteCourse:output_type -> academic.DeleteCourseResponse
	17, // 23: academic.AcademicService.ListCourses:output_type -> academic.ListCoursesResponse
	14, // [14:24] is the sub-list for method output_type
	4,  // [4:14] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_proto_academic_proto_init() }
func file_proto_academic_proto_init() {
	if File_proto_academic_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_academic_proto_rawDesc), len(file_proto_academic_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   18,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_academic_proto_goTypes,
		DependencyIndexes: file_proto_academic_proto_depIdxs,
		MessageInfos:      file_proto_academic_proto_msgTypes,
	}.Build()
	File_proto_academic_proto = out.File
	file_proto_academic_proto_goTypes = nil
	file_proto_academic_proto_depIdxs = nil
}
