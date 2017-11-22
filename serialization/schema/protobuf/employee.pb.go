// Code generated by protoc-gen-go. DO NOT EDIT.
// source: serialization/schema/protobuf/employee.proto

/*
Package protobuf is a generated protocol buffer package.

It is generated from these files:
	serialization/schema/protobuf/employee.proto

It has these top-level messages:
	Employee
*/
package protobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Employee struct {
	FirstName  string  `protobuf:"bytes,1,opt,name=first_name,json=firstName" json:"first_name,omitempty"`
	LastName   string  `protobuf:"bytes,2,opt,name=last_name,json=lastName" json:"last_name,omitempty"`
	Position   string  `protobuf:"bytes,3,opt,name=position" json:"position,omitempty"`
	Salary     float64 `protobuf:"fixed64,4,opt,name=salary" json:"salary,omitempty"`
	Id         int64   `protobuf:"varint,5,opt,name=id" json:"id,omitempty"`
	ManagerId  int64   `protobuf:"varint,6,opt,name=manager_id,json=managerId" json:"manager_id,omitempty"`
	UpdateTime int64   `protobuf:"varint,7,opt,name=update_time,json=updateTime" json:"update_time,omitempty"`
}

func (m *Employee) Reset()                    { *m = Employee{} }
func (m *Employee) String() string            { return proto.CompactTextString(m) }
func (*Employee) ProtoMessage()               {}
func (*Employee) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Employee) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Employee) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *Employee) GetPosition() string {
	if m != nil {
		return m.Position
	}
	return ""
}

func (m *Employee) GetSalary() float64 {
	if m != nil {
		return m.Salary
	}
	return 0
}

func (m *Employee) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Employee) GetManagerId() int64 {
	if m != nil {
		return m.ManagerId
	}
	return 0
}

func (m *Employee) GetUpdateTime() int64 {
	if m != nil {
		return m.UpdateTime
	}
	return 0
}

func init() {
	proto.RegisterType((*Employee)(nil), "protobuf.Employee")
}

func init() { proto.RegisterFile("serialization/schema/protobuf/employee.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 212 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0xcf, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0xc6, 0x71, 0x39, 0x85, 0x60, 0x1f, 0x12, 0x83, 0x07, 0x64, 0x81, 0x10, 0x11, 0x53, 0x06,
	0x44, 0x06, 0x9e, 0x81, 0x81, 0x85, 0x21, 0x62, 0x8f, 0xae, 0xf8, 0x0a, 0x27, 0xd9, 0xb1, 0x65,
	0xbb, 0x43, 0x79, 0x44, 0x9e, 0x0a, 0xc5, 0x4d, 0xc6, 0xfb, 0xff, 0x86, 0x4f, 0x07, 0xcf, 0x99,
	0x12, 0xa3, 0xe3, 0x5f, 0x2c, 0x1c, 0xe6, 0x21, 0x7f, 0xfd, 0x90, 0xc7, 0x21, 0xa6, 0x50, 0xc2,
	0xfe, 0x78, 0x18, 0xc8, 0x47, 0x17, 0x4e, 0x44, 0x2f, 0xb5, 0x68, 0xb9, 0xc1, 0xd3, 0x9f, 0x00,
	0xf9, 0xb6, 0xa2, 0x7e, 0x00, 0x38, 0x70, 0xca, 0x65, 0x9a, 0xd1, 0x93, 0x11, 0x9d, 0xe8, 0xd5,
	0xa8, 0x6a, 0xf9, 0x40, 0x4f, 0xfa, 0x1e, 0x94, 0xc3, 0x4d, 0x9b, 0xaa, 0x72, 0x09, 0x15, 0xef,
	0x40, 0xc6, 0x90, 0x79, 0x59, 0x37, 0xbb, 0xb3, 0x6d, 0xb7, 0xbe, 0x85, 0x36, 0xa3, 0xc3, 0x74,
	0x32, 0x17, 0x9d, 0xe8, 0xc5, 0xb8, 0x5e, 0xfa, 0x06, 0x1a, 0xb6, 0xe6, 0xb2, 0x13, 0xfd, 0x6e,
	0x6c, 0xd8, 0x2e, 0xfb, 0x1e, 0x67, 0xfc, 0xa6, 0x34, 0xb1, 0x35, 0x6d, 0xed, 0x6a, 0x2d, 0xef,
	0x56, 0x3f, 0xc2, 0xf5, 0x31, 0x5a, 0x2c, 0x34, 0x15, 0xf6, 0x64, 0xae, 0xaa, 0xc3, 0x39, 0x7d,
	0xb2, 0xa7, 0x7d, 0x5b, 0xdf, 0x7a, 0xfd, 0x0f, 0x00, 0x00, 0xff, 0xff, 0xea, 0x21, 0x81, 0x49,
	0x0d, 0x01, 0x00, 0x00,
}
