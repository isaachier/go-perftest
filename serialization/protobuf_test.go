package serialization

import (
	"testing"

	"github.com/golang/protobuf/proto"

	"github.com/isaachier/go-perftest/serialization/schema/protobuf"
)

type protobufCodec struct{}

func (c *protobufCodec) encode(e *employee) ([]byte, error) {
	employeePB := protobuf.Employee{}
	employeePB.FirstName = e.firstName
	employeePB.LastName = e.lastName
	employeePB.Position = e.position
	employeePB.Salary = e.salary
	employeePB.Id = e.id
	if e.managerID != nil {
		employeePB.ManagerId = *e.managerID
	}
	employeePB.UpdateTime = e.updateTime
	return proto.Marshal(&employeePB)
}

func (c *protobufCodec) decode(b []byte) (*employee, error) {
	employeePB := protobuf.Employee{}
	if err := proto.Unmarshal(b, &employeePB); err != nil {
		return nil, err
	}
	e := employee{}
	e.firstName = employeePB.GetFirstName()
	e.lastName = employeePB.GetLastName()
	e.position = employeePB.GetPosition()
	e.salary = employeePB.GetSalary()
	e.id = employeePB.GetId()
	if managerID := employeePB.GetManagerId(); managerID != 0 {
		e.managerID = &managerID
	}
	e.updateTime = employeePB.GetUpdateTime()
	return &e, nil
}

func (c *protobufCodec) name() string {
	return "protobuf"
}

func BenchmarkProtobufEncoding(b *testing.B) {
	c := &protobufCodec{}
	runEncodingBenchmark(c, b)
}

func BenchmarkProtobufDecoding(b *testing.B) {
	c := &protobufCodec{}
	runDecodingBenchmark(c, b)
}
