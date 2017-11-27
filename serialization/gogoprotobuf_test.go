package serialization

import (
	"testing"

	"github.com/isaachier/go-perftest/serialization/schema/gogoprotobuf"
)

type gogoProtobufCodec struct{}

func (c *gogoProtobufCodec) encode(e *employee) ([]byte, error) {
	employeePB := gogoprotobuf.Employee{}
	employeePB.FirstName = e.firstName
	employeePB.LastName = e.lastName
	employeePB.Position = e.position
	employeePB.Salary = e.salary
	employeePB.Id = e.id
	if e.managerID != nil {
		employeePB.ManagerId = *e.managerID
	}
	employeePB.UpdateTime = e.updateTime
	return employeePB.Marshal()
}

func (c *gogoProtobufCodec) decode(b []byte) (*employee, error) {
	employeePB := gogoprotobuf.Employee{}
	if err := employeePB.Unmarshal(b); err != nil {
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

func (c *gogoProtobufCodec) name() string {
	return "gogoprotobuf"
}

func BenchmarkGogoProtobufEncoding(b *testing.B) {
	c := &gogoProtobufCodec{}
	runEncodingBenchmark(c, b)
}

func BenchmarkGogoProtobufDecoding(b *testing.B) {
	c := &gogoProtobufCodec{}
	runDecodingBenchmark(c, b)
}
