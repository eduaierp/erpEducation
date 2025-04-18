// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: proto/student.proto

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

type Student struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email         string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Course        string                 `protobuf:"bytes,4,opt,name=course,proto3" json:"course,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Student) Reset() {
	*x = Student{}
	mi := &file_proto_student_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Student) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Student) ProtoMessage() {}

func (x *Student) ProtoReflect() protoreflect.Message {
	mi := &file_proto_student_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Student.ProtoReflect.Descriptor instead.
func (*Student) Descriptor() ([]byte, []int) {
	return file_proto_student_proto_rawDescGZIP(), []int{0}
}

func (x *Student) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Student) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Student) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Student) GetCourse() string {
	if x != nil {
		return x.Course
	}
	return ""
}

type CreateStudentRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email         string                 `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Course        string                 `protobuf:"bytes,3,opt,name=course,proto3" json:"course,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateStudentRequest) Reset() {
	*x = CreateStudentRequest{}
	mi := &file_proto_student_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateStudentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStudentRequest) ProtoMessage() {}

func (x *CreateStudentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_student_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStudentRequest.ProtoReflect.Descriptor instead.
func (*CreateStudentRequest) Descriptor() ([]byte, []int) {
	return file_proto_student_proto_rawDescGZIP(), []int{1}
}

func (x *CreateStudentRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateStudentRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateStudentRequest) GetCourse() string {
	if x != nil {
		return x.Course
	}
	return ""
}

type GetStudentRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetStudentRequest) Reset() {
	*x = GetStudentRequest{}
	mi := &file_proto_student_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetStudentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStudentRequest) ProtoMessage() {}

func (x *GetStudentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_student_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStudentRequest.ProtoReflect.Descriptor instead.
func (*GetStudentRequest) Descriptor() ([]byte, []int) {
	return file_proto_student_proto_rawDescGZIP(), []int{2}
}

func (x *GetStudentRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UpdateStudentRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email         string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Course        string                 `protobuf:"bytes,4,opt,name=course,proto3" json:"course,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateStudentRequest) Reset() {
	*x = UpdateStudentRequest{}
	mi := &file_proto_student_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateStudentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStudentRequest) ProtoMessage() {}

func (x *UpdateStudentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_student_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStudentRequest.ProtoReflect.Descriptor instead.
func (*UpdateStudentRequest) Descriptor() ([]byte, []int) {
	return file_proto_student_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateStudentRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateStudentRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateStudentRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UpdateStudentRequest) GetCourse() string {
	if x != nil {
		return x.Course
	}
	return ""
}

type DeleteStudentRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteStudentRequest) Reset() {
	*x = DeleteStudentRequest{}
	mi := &file_proto_student_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteStudentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteStudentRequest) ProtoMessage() {}

func (x *DeleteStudentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_student_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteStudentRequest.ProtoReflect.Descriptor instead.
func (*DeleteStudentRequest) Descriptor() ([]byte, []int) {
	return file_proto_student_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteStudentRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteStudentResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteStudentResponse) Reset() {
	*x = DeleteStudentResponse{}
	mi := &file_proto_student_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteStudentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteStudentResponse) ProtoMessage() {}

func (x *DeleteStudentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_student_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteStudentResponse.ProtoReflect.Descriptor instead.
func (*DeleteStudentResponse) Descriptor() ([]byte, []int) {
	return file_proto_student_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteStudentResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type StudentResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Student       *Student               `protobuf:"bytes,1,opt,name=student,proto3" json:"student,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StudentResponse) Reset() {
	*x = StudentResponse{}
	mi := &file_proto_student_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StudentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StudentResponse) ProtoMessage() {}

func (x *StudentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_student_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StudentResponse.ProtoReflect.Descriptor instead.
func (*StudentResponse) Descriptor() ([]byte, []int) {
	return file_proto_student_proto_rawDescGZIP(), []int{6}
}

func (x *StudentResponse) GetStudent() *Student {
	if x != nil {
		return x.Student
	}
	return nil
}

type ListStudentsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListStudentsRequest) Reset() {
	*x = ListStudentsRequest{}
	mi := &file_proto_student_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListStudentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListStudentsRequest) ProtoMessage() {}

func (x *ListStudentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_student_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListStudentsRequest.ProtoReflect.Descriptor instead.
func (*ListStudentsRequest) Descriptor() ([]byte, []int) {
	return file_proto_student_proto_rawDescGZIP(), []int{7}
}

type ListStudentsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Students      []*Student             `protobuf:"bytes,1,rep,name=students,proto3" json:"students,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListStudentsResponse) Reset() {
	*x = ListStudentsResponse{}
	mi := &file_proto_student_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListStudentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListStudentsResponse) ProtoMessage() {}

func (x *ListStudentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_student_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListStudentsResponse.ProtoReflect.Descriptor instead.
func (*ListStudentsResponse) Descriptor() ([]byte, []int) {
	return file_proto_student_proto_rawDescGZIP(), []int{8}
}

func (x *ListStudentsResponse) GetStudents() []*Student {
	if x != nil {
		return x.Students
	}
	return nil
}

var File_proto_student_proto protoreflect.FileDescriptor

const file_proto_student_proto_rawDesc = "" +
	"\n" +
	"\x13proto/student.proto\x12\astudent\"[\n" +
	"\aStudent\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x14\n" +
	"\x05email\x18\x03 \x01(\tR\x05email\x12\x16\n" +
	"\x06course\x18\x04 \x01(\tR\x06course\"X\n" +
	"\x14CreateStudentRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x14\n" +
	"\x05email\x18\x02 \x01(\tR\x05email\x12\x16\n" +
	"\x06course\x18\x03 \x01(\tR\x06course\"#\n" +
	"\x11GetStudentRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\"h\n" +
	"\x14UpdateStudentRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x14\n" +
	"\x05email\x18\x03 \x01(\tR\x05email\x12\x16\n" +
	"\x06course\x18\x04 \x01(\tR\x06course\"&\n" +
	"\x14DeleteStudentRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\"1\n" +
	"\x15DeleteStudentResponse\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage\"=\n" +
	"\x0fStudentResponse\x12*\n" +
	"\astudent\x18\x01 \x01(\v2\x10.student.StudentR\astudent\"\x15\n" +
	"\x13ListStudentsRequest\"D\n" +
	"\x14ListStudentsResponse\x12,\n" +
	"\bstudents\x18\x01 \x03(\v2\x10.student.StudentR\bstudents2\x85\x03\n" +
	"\x0eStudentService\x12H\n" +
	"\rCreateStudent\x12\x1d.student.CreateStudentRequest\x1a\x18.student.StudentResponse\x12B\n" +
	"\n" +
	"GetStudent\x12\x1a.student.GetStudentRequest\x1a\x18.student.StudentResponse\x12H\n" +
	"\rUpdateStudent\x12\x1d.student.UpdateStudentRequest\x1a\x18.student.StudentResponse\x12N\n" +
	"\rDeleteStudent\x12\x1d.student.DeleteStudentRequest\x1a\x1e.student.DeleteStudentResponse\x12K\n" +
	"\fListStudents\x12\x1c.student.ListStudentsRequest\x1a\x1d.student.ListStudentsResponseB\tZ\a./protob\x06proto3"

var (
	file_proto_student_proto_rawDescOnce sync.Once
	file_proto_student_proto_rawDescData []byte
)

func file_proto_student_proto_rawDescGZIP() []byte {
	file_proto_student_proto_rawDescOnce.Do(func() {
		file_proto_student_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_student_proto_rawDesc), len(file_proto_student_proto_rawDesc)))
	})
	return file_proto_student_proto_rawDescData
}

var file_proto_student_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_student_proto_goTypes = []any{
	(*Student)(nil),               // 0: student.Student
	(*CreateStudentRequest)(nil),  // 1: student.CreateStudentRequest
	(*GetStudentRequest)(nil),     // 2: student.GetStudentRequest
	(*UpdateStudentRequest)(nil),  // 3: student.UpdateStudentRequest
	(*DeleteStudentRequest)(nil),  // 4: student.DeleteStudentRequest
	(*DeleteStudentResponse)(nil), // 5: student.DeleteStudentResponse
	(*StudentResponse)(nil),       // 6: student.StudentResponse
	(*ListStudentsRequest)(nil),   // 7: student.ListStudentsRequest
	(*ListStudentsResponse)(nil),  // 8: student.ListStudentsResponse
}
var file_proto_student_proto_depIdxs = []int32{
	0, // 0: student.StudentResponse.student:type_name -> student.Student
	0, // 1: student.ListStudentsResponse.students:type_name -> student.Student
	1, // 2: student.StudentService.CreateStudent:input_type -> student.CreateStudentRequest
	2, // 3: student.StudentService.GetStudent:input_type -> student.GetStudentRequest
	3, // 4: student.StudentService.UpdateStudent:input_type -> student.UpdateStudentRequest
	4, // 5: student.StudentService.DeleteStudent:input_type -> student.DeleteStudentRequest
	7, // 6: student.StudentService.ListStudents:input_type -> student.ListStudentsRequest
	6, // 7: student.StudentService.CreateStudent:output_type -> student.StudentResponse
	6, // 8: student.StudentService.GetStudent:output_type -> student.StudentResponse
	6, // 9: student.StudentService.UpdateStudent:output_type -> student.StudentResponse
	5, // 10: student.StudentService.DeleteStudent:output_type -> student.DeleteStudentResponse
	8, // 11: student.StudentService.ListStudents:output_type -> student.ListStudentsResponse
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_student_proto_init() }
func file_proto_student_proto_init() {
	if File_proto_student_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_student_proto_rawDesc), len(file_proto_student_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_student_proto_goTypes,
		DependencyIndexes: file_proto_student_proto_depIdxs,
		MessageInfos:      file_proto_student_proto_msgTypes,
	}.Build()
	File_proto_student_proto = out.File
	file_proto_student_proto_goTypes = nil
	file_proto_student_proto_depIdxs = nil
}
