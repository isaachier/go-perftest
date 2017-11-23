// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: serialization/schema/gogoprotobuf/employee.proto

/*
	Package gogoprotobuf is a generated protocol buffer package.

	It is generated from these files:
		serialization/schema/gogoprotobuf/employee.proto

	It has these top-level messages:
		Employee
*/
package gogoprotobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import binary "encoding/binary"

import io "io"

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
	FirstName  string  `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName   string  `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Position   string  `protobuf:"bytes,3,opt,name=position,proto3" json:"position,omitempty"`
	Salary     float64 `protobuf:"fixed64,4,opt,name=salary,proto3" json:"salary,omitempty"`
	Id         int64   `protobuf:"varint,5,opt,name=id,proto3" json:"id,omitempty"`
	ManagerId  int64   `protobuf:"varint,6,opt,name=manager_id,json=managerId,proto3" json:"manager_id,omitempty"`
	UpdateTime int64   `protobuf:"varint,7,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
}

func (m *Employee) Reset()                    { *m = Employee{} }
func (m *Employee) String() string            { return proto.CompactTextString(m) }
func (*Employee) ProtoMessage()               {}
func (*Employee) Descriptor() ([]byte, []int) { return fileDescriptorEmployee, []int{0} }

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
	proto.RegisterType((*Employee)(nil), "gogoprotobuf.Employee")
}
func (m *Employee) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Employee) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.FirstName) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintEmployee(dAtA, i, uint64(len(m.FirstName)))
		i += copy(dAtA[i:], m.FirstName)
	}
	if len(m.LastName) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintEmployee(dAtA, i, uint64(len(m.LastName)))
		i += copy(dAtA[i:], m.LastName)
	}
	if len(m.Position) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintEmployee(dAtA, i, uint64(len(m.Position)))
		i += copy(dAtA[i:], m.Position)
	}
	if m.Salary != 0 {
		dAtA[i] = 0x21
		i++
		binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Salary))))
		i += 8
	}
	if m.Id != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintEmployee(dAtA, i, uint64(m.Id))
	}
	if m.ManagerId != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintEmployee(dAtA, i, uint64(m.ManagerId))
	}
	if m.UpdateTime != 0 {
		dAtA[i] = 0x38
		i++
		i = encodeVarintEmployee(dAtA, i, uint64(m.UpdateTime))
	}
	return i, nil
}

func encodeVarintEmployee(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Employee) Size() (n int) {
	var l int
	_ = l
	l = len(m.FirstName)
	if l > 0 {
		n += 1 + l + sovEmployee(uint64(l))
	}
	l = len(m.LastName)
	if l > 0 {
		n += 1 + l + sovEmployee(uint64(l))
	}
	l = len(m.Position)
	if l > 0 {
		n += 1 + l + sovEmployee(uint64(l))
	}
	if m.Salary != 0 {
		n += 9
	}
	if m.Id != 0 {
		n += 1 + sovEmployee(uint64(m.Id))
	}
	if m.ManagerId != 0 {
		n += 1 + sovEmployee(uint64(m.ManagerId))
	}
	if m.UpdateTime != 0 {
		n += 1 + sovEmployee(uint64(m.UpdateTime))
	}
	return n
}

func sovEmployee(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozEmployee(x uint64) (n int) {
	return sovEmployee(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Employee) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEmployee
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Employee: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Employee: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FirstName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEmployee
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEmployee
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FirstName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEmployee
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEmployee
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LastName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Position", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEmployee
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEmployee
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Position = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Salary", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Salary = float64(math.Float64frombits(v))
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEmployee
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ManagerId", wireType)
			}
			m.ManagerId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEmployee
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ManagerId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdateTime", wireType)
			}
			m.UpdateTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEmployee
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UpdateTime |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEmployee(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEmployee
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipEmployee(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEmployee
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEmployee
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEmployee
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthEmployee
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowEmployee
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipEmployee(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthEmployee = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEmployee   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("serialization/schema/gogoprotobuf/employee.proto", fileDescriptorEmployee)
}

var fileDescriptorEmployee = []byte{
	// 242 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0xcf, 0x41, 0x4a, 0xc4, 0x30,
	0x14, 0xc6, 0x71, 0xd3, 0xd1, 0xda, 0x3e, 0x45, 0x24, 0x0b, 0x09, 0x8a, 0xb5, 0xb8, 0xea, 0xca,
	0x0a, 0xde, 0x40, 0x70, 0xe1, 0xc6, 0x45, 0x71, 0x5f, 0xde, 0xd8, 0x37, 0xf5, 0x41, 0xd3, 0x94,
	0x24, 0xb3, 0x18, 0x4f, 0xe2, 0x91, 0xc4, 0x95, 0x47, 0x90, 0x7a, 0x11, 0x69, 0xa6, 0x03, 0xb3,
	0xfc, 0xfe, 0xbf, 0x40, 0x12, 0xb8, 0x77, 0x64, 0x19, 0x3b, 0xfe, 0x40, 0xcf, 0xa6, 0x2f, 0xdd,
	0xdb, 0x3b, 0x69, 0x2c, 0x5b, 0xd3, 0x9a, 0xc1, 0x1a, 0x6f, 0x96, 0xeb, 0x55, 0x49, 0x7a, 0xe8,
	0xcc, 0x86, 0xe8, 0x2e, 0x14, 0x79, 0xba, 0x8f, 0xb7, 0xdf, 0x02, 0x92, 0xa7, 0xf9, 0x80, 0xbc,
	0x06, 0x58, 0xb1, 0x75, 0xbe, 0xee, 0x51, 0x93, 0x12, 0xb9, 0x28, 0xd2, 0x2a, 0x0d, 0xe5, 0x05,
	0x35, 0xc9, 0x2b, 0x48, 0x3b, 0xdc, 0x69, 0x14, 0x34, 0x99, 0x42, 0xc0, 0x4b, 0x48, 0x06, 0xe3,
	0x78, 0x7a, 0x85, 0x5a, 0x6c, 0x6d, 0xb7, 0xe5, 0x05, 0xc4, 0x0e, 0x3b, 0xb4, 0x1b, 0x75, 0x98,
	0x8b, 0x42, 0x54, 0xf3, 0x92, 0x67, 0x10, 0x71, 0xa3, 0x8e, 0x72, 0x51, 0x2c, 0xaa, 0x88, 0x9b,
	0xe9, 0x7e, 0x8d, 0x3d, 0xb6, 0x64, 0x6b, 0x6e, 0x54, 0x1c, 0x7a, 0x3a, 0x97, 0xe7, 0x46, 0xde,
	0xc0, 0xc9, 0x7a, 0x68, 0xd0, 0x53, 0xed, 0x59, 0x93, 0x3a, 0x0e, 0x0e, 0xdb, 0xf4, 0xca, 0x9a,
	0x1e, 0xcf, 0xbf, 0xc6, 0x4c, 0xfc, 0x8c, 0x99, 0xf8, 0x1d, 0x33, 0xf1, 0xf9, 0x97, 0x1d, 0x2c,
	0xe3, 0xf0, 0xd1, 0x87, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9c, 0xac, 0xbd, 0xf2, 0x27, 0x01,
	0x00, 0x00,
}
