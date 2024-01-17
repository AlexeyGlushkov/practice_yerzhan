// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.2
// source: pkg/grpc/proto/emppos.proto

package emppos_proto

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

type EmployeeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EmployeeId string `protobuf:"bytes,1,opt,name=employee_id,json=employeeId,proto3" json:"employee_id,omitempty"`
}

func (x *EmployeeRequest) Reset() {
	*x = EmployeeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_proto_emppos_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmployeeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmployeeRequest) ProtoMessage() {}

func (x *EmployeeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_proto_emppos_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmployeeRequest.ProtoReflect.Descriptor instead.
func (*EmployeeRequest) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_proto_emppos_proto_rawDescGZIP(), []int{0}
}

func (x *EmployeeRequest) GetEmployeeId() string {
	if x != nil {
		return x.EmployeeId
	}
	return ""
}

type EmployeeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EmployeeId string `protobuf:"bytes,1,opt,name=employee_id,json=employeeId,proto3" json:"employee_id,omitempty"`
	FirstName  string `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName   string `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
}

func (x *EmployeeResponse) Reset() {
	*x = EmployeeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_proto_emppos_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmployeeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmployeeResponse) ProtoMessage() {}

func (x *EmployeeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_proto_emppos_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmployeeResponse.ProtoReflect.Descriptor instead.
func (*EmployeeResponse) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_proto_emppos_proto_rawDescGZIP(), []int{1}
}

func (x *EmployeeResponse) GetEmployeeId() string {
	if x != nil {
		return x.EmployeeId
	}
	return ""
}

func (x *EmployeeResponse) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *EmployeeResponse) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

type CreateEmployeeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName    string `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName     string `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	PositionName string `protobuf:"bytes,3,opt,name=position_name,json=positionName,proto3" json:"position_name,omitempty"`
	Salary       int32  `protobuf:"varint,4,opt,name=salary,proto3" json:"salary,omitempty"`
}

func (x *CreateEmployeeRequest) Reset() {
	*x = CreateEmployeeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_proto_emppos_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEmployeeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEmployeeRequest) ProtoMessage() {}

func (x *CreateEmployeeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_proto_emppos_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEmployeeRequest.ProtoReflect.Descriptor instead.
func (*CreateEmployeeRequest) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_proto_emppos_proto_rawDescGZIP(), []int{2}
}

func (x *CreateEmployeeRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *CreateEmployeeRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *CreateEmployeeRequest) GetPositionName() string {
	if x != nil {
		return x.PositionName
	}
	return ""
}

func (x *CreateEmployeeRequest) GetSalary() int32 {
	if x != nil {
		return x.Salary
	}
	return 0
}

type CreateEmployeeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateEmployeeResponse) Reset() {
	*x = CreateEmployeeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_proto_emppos_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEmployeeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEmployeeResponse) ProtoMessage() {}

func (x *CreateEmployeeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_proto_emppos_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEmployeeResponse.ProtoReflect.Descriptor instead.
func (*CreateEmployeeResponse) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_proto_emppos_proto_rawDescGZIP(), []int{3}
}

type UpdateEmployeeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EmployeeId string `protobuf:"bytes,1,opt,name=employee_id,json=employeeId,proto3" json:"employee_id,omitempty"`
	FirstName  string `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName   string `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
}

func (x *UpdateEmployeeRequest) Reset() {
	*x = UpdateEmployeeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_proto_emppos_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEmployeeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEmployeeRequest) ProtoMessage() {}

func (x *UpdateEmployeeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_proto_emppos_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEmployeeRequest.ProtoReflect.Descriptor instead.
func (*UpdateEmployeeRequest) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_proto_emppos_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateEmployeeRequest) GetEmployeeId() string {
	if x != nil {
		return x.EmployeeId
	}
	return ""
}

func (x *UpdateEmployeeRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *UpdateEmployeeRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

type UpdateEmployeeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateEmployeeResponse) Reset() {
	*x = UpdateEmployeeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_proto_emppos_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEmployeeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEmployeeResponse) ProtoMessage() {}

func (x *UpdateEmployeeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_proto_emppos_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEmployeeResponse.ProtoReflect.Descriptor instead.
func (*UpdateEmployeeResponse) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_proto_emppos_proto_rawDescGZIP(), []int{5}
}

type DeleteEmployeeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EmployeeId string `protobuf:"bytes,1,opt,name=employee_id,json=employeeId,proto3" json:"employee_id,omitempty"`
}

func (x *DeleteEmployeeRequest) Reset() {
	*x = DeleteEmployeeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_proto_emppos_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEmployeeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEmployeeRequest) ProtoMessage() {}

func (x *DeleteEmployeeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_proto_emppos_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEmployeeRequest.ProtoReflect.Descriptor instead.
func (*DeleteEmployeeRequest) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_proto_emppos_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteEmployeeRequest) GetEmployeeId() string {
	if x != nil {
		return x.EmployeeId
	}
	return ""
}

type DeleteEmployeeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteEmployeeResponse) Reset() {
	*x = DeleteEmployeeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_proto_emppos_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEmployeeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEmployeeResponse) ProtoMessage() {}

func (x *DeleteEmployeeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_proto_emppos_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEmployeeResponse.ProtoReflect.Descriptor instead.
func (*DeleteEmployeeResponse) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_proto_emppos_proto_rawDescGZIP(), []int{7}
}

var File_pkg_grpc_proto_emppos_proto protoreflect.FileDescriptor

var file_pkg_grpc_proto_emppos_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x65, 0x6d, 0x70, 0x70, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x32, 0x0a,
	0x0f, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x49,
	0x64, 0x22, 0x6f, 0x0a, 0x10, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x6d, 0x70, 0x6c,
	0x6f, 0x79, 0x65, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x22, 0x90, 0x01, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x70,
	0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c,
	0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73,
	0x61, 0x6c, 0x61, 0x72, 0x79, 0x22, 0x18, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45,
	0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x74, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6d, 0x70, 0x6c,
	0x6f, 0x79, 0x65, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65,
	0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x18, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45,
	0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x38, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6d, 0x70, 0x6c,
	0x6f, 0x79, 0x65, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65,
	0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x49, 0x64, 0x22, 0x18, 0x0a, 0x16, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x32, 0x92, 0x02, 0x0a, 0x0f, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x45, 0x6d,
	0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x42, 0x79, 0x49, 0x44, 0x12, 0x10, 0x2e, 0x45, 0x6d, 0x70,
	0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x45,
	0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x41, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65,
	0x65, 0x12, 0x16, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79,
	0x65, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x41, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x70, 0x6c,
	0x6f, 0x79, 0x65, 0x65, 0x12, 0x16, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x70,
	0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45,
	0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x12, 0x16, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2f, 0x5a, 0x2d, 0x73, 0x71, 0x6c, 0x5f,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x6d, 0x70,
	0x70, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_pkg_grpc_proto_emppos_proto_rawDescOnce sync.Once
	file_pkg_grpc_proto_emppos_proto_rawDescData = file_pkg_grpc_proto_emppos_proto_rawDesc
)

func file_pkg_grpc_proto_emppos_proto_rawDescGZIP() []byte {
	file_pkg_grpc_proto_emppos_proto_rawDescOnce.Do(func() {
		file_pkg_grpc_proto_emppos_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_grpc_proto_emppos_proto_rawDescData)
	})
	return file_pkg_grpc_proto_emppos_proto_rawDescData
}

var file_pkg_grpc_proto_emppos_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_pkg_grpc_proto_emppos_proto_goTypes = []interface{}{
	(*EmployeeRequest)(nil),        // 0: EmployeeRequest
	(*EmployeeResponse)(nil),       // 1: EmployeeResponse
	(*CreateEmployeeRequest)(nil),  // 2: CreateEmployeeRequest
	(*CreateEmployeeResponse)(nil), // 3: CreateEmployeeResponse
	(*UpdateEmployeeRequest)(nil),  // 4: UpdateEmployeeRequest
	(*UpdateEmployeeResponse)(nil), // 5: UpdateEmployeeResponse
	(*DeleteEmployeeRequest)(nil),  // 6: DeleteEmployeeRequest
	(*DeleteEmployeeResponse)(nil), // 7: DeleteEmployeeResponse
}
var file_pkg_grpc_proto_emppos_proto_depIdxs = []int32{
	0, // 0: EmployeeService.GetEmployeeByID:input_type -> EmployeeRequest
	2, // 1: EmployeeService.CreateEmployee:input_type -> CreateEmployeeRequest
	4, // 2: EmployeeService.UpdateEmployee:input_type -> UpdateEmployeeRequest
	6, // 3: EmployeeService.DeleteEmployee:input_type -> DeleteEmployeeRequest
	1, // 4: EmployeeService.GetEmployeeByID:output_type -> EmployeeResponse
	3, // 5: EmployeeService.CreateEmployee:output_type -> CreateEmployeeResponse
	5, // 6: EmployeeService.UpdateEmployee:output_type -> UpdateEmployeeResponse
	7, // 7: EmployeeService.DeleteEmployee:output_type -> DeleteEmployeeResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_grpc_proto_emppos_proto_init() }
func file_pkg_grpc_proto_emppos_proto_init() {
	if File_pkg_grpc_proto_emppos_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_grpc_proto_emppos_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmployeeRequest); i {
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
		file_pkg_grpc_proto_emppos_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmployeeResponse); i {
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
		file_pkg_grpc_proto_emppos_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEmployeeRequest); i {
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
		file_pkg_grpc_proto_emppos_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEmployeeResponse); i {
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
		file_pkg_grpc_proto_emppos_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEmployeeRequest); i {
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
		file_pkg_grpc_proto_emppos_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEmployeeResponse); i {
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
		file_pkg_grpc_proto_emppos_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEmployeeRequest); i {
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
		file_pkg_grpc_proto_emppos_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEmployeeResponse); i {
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
			RawDescriptor: file_pkg_grpc_proto_emppos_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_grpc_proto_emppos_proto_goTypes,
		DependencyIndexes: file_pkg_grpc_proto_emppos_proto_depIdxs,
		MessageInfos:      file_pkg_grpc_proto_emppos_proto_msgTypes,
	}.Build()
	File_pkg_grpc_proto_emppos_proto = out.File
	file_pkg_grpc_proto_emppos_proto_rawDesc = nil
	file_pkg_grpc_proto_emppos_proto_goTypes = nil
	file_pkg_grpc_proto_emppos_proto_depIdxs = nil
}
