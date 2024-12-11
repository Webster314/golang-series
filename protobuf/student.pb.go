// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.29.1
// source: student.proto

package main

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

type Student_Gender int32

const (
	Student_FEMALE     Student_Gender = 0
	Student_MALE       Student_Gender = 1
	Student_HELICOPTER Student_Gender = 1
)

// Enum value maps for Student_Gender.
var (
	Student_Gender_name = map[int32]string{
		0: "FEMALE",
		1: "MALE",
		// Duplicate value: 1: "HELICOPTER",
	}
	Student_Gender_value = map[string]int32{
		"FEMALE":     0,
		"MALE":       1,
		"HELICOPTER": 1,
	}
)

func (x Student_Gender) Enum() *Student_Gender {
	p := new(Student_Gender)
	*p = x
	return p
}

func (x Student_Gender) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Student_Gender) Descriptor() protoreflect.EnumDescriptor {
	return file_student_proto_enumTypes[0].Descriptor()
}

func (Student_Gender) Type() protoreflect.EnumType {
	return &file_student_proto_enumTypes[0]
}

func (x Student_Gender) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Student_Gender.Descriptor instead.
func (Student_Gender) EnumDescriptor() ([]byte, []int) {
	return file_student_proto_rawDescGZIP(), []int{0, 0}
}

type Student struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Male   bool           `protobuf:"varint,2,opt,name=male,proto3" json:"male,omitempty"`
	Score  []int32        `protobuf:"varint,3,rep,packed,name=score,proto3" json:"score,omitempty"`
	Gender Student_Gender `protobuf:"varint,6,opt,name=gender,proto3,enum=main.Student_Gender" json:"gender,omitempty"`
}

func (x *Student) Reset() {
	*x = Student{}
	mi := &file_student_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Student) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Student) ProtoMessage() {}

func (x *Student) ProtoReflect() protoreflect.Message {
	mi := &file_student_proto_msgTypes[0]
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
	return file_student_proto_rawDescGZIP(), []int{0}
}

func (x *Student) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Student) GetMale() bool {
	if x != nil {
		return x.Male
	}
	return false
}

func (x *Student) GetScore() []int32 {
	if x != nil {
		return x.Score
	}
	return nil
}

func (x *Student) GetGender() Student_Gender {
	if x != nil {
		return x.Gender
	}
	return Student_FEMALE
}

var File_student_proto protoreflect.FileDescriptor

var file_student_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x04, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0xc5, 0x01, 0x0a, 0x07, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x61, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x04, 0x6d, 0x61, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x05, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12,
	0x2c, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x14, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x47,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x22, 0x32, 0x0a,
	0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x45, 0x4d, 0x41, 0x4c,
	0x45, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x0e, 0x0a,
	0x0a, 0x48, 0x45, 0x4c, 0x49, 0x43, 0x4f, 0x50, 0x54, 0x45, 0x52, 0x10, 0x01, 0x1a, 0x02, 0x10,
	0x01, 0x4a, 0x04, 0x08, 0x04, 0x10, 0x05, 0x4a, 0x04, 0x08, 0x0f, 0x10, 0x10, 0x4a, 0x04, 0x08,
	0x09, 0x10, 0x0c, 0x52, 0x03, 0x66, 0x6f, 0x6f, 0x52, 0x03, 0x62, 0x61, 0x72, 0x42, 0x09, 0x5a,
	0x07, 0x2e, 0x2f, 0x3b, 0x6d, 0x61, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_student_proto_rawDescOnce sync.Once
	file_student_proto_rawDescData = file_student_proto_rawDesc
)

func file_student_proto_rawDescGZIP() []byte {
	file_student_proto_rawDescOnce.Do(func() {
		file_student_proto_rawDescData = protoimpl.X.CompressGZIP(file_student_proto_rawDescData)
	})
	return file_student_proto_rawDescData
}

var file_student_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_student_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_student_proto_goTypes = []any{
	(Student_Gender)(0), // 0: main.student.Gender
	(*Student)(nil),     // 1: main.student
}
var file_student_proto_depIdxs = []int32{
	0, // 0: main.student.gender:type_name -> main.student.Gender
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_student_proto_init() }
func file_student_proto_init() {
	if File_student_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_student_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_student_proto_goTypes,
		DependencyIndexes: file_student_proto_depIdxs,
		EnumInfos:         file_student_proto_enumTypes,
		MessageInfos:      file_student_proto_msgTypes,
	}.Build()
	File_student_proto = out.File
	file_student_proto_rawDesc = nil
	file_student_proto_goTypes = nil
	file_student_proto_depIdxs = nil
}
