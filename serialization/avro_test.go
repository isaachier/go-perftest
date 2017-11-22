package serialization

import (
	"io/ioutil"
	"runtime"
	"testing"

	"github.com/linkedin/goavro"
)

var schema string

func init() {
	schemaBytes, err := ioutil.ReadFile("schema/avro/employee.avsc")
	if err != nil {
		panic(err.Error())
	}
	schema = string(schemaBytes)
}

func (e *employee) Map() map[string]interface{} {
	m := map[string]interface{}{
		"firstName":  e.firstName,
		"lastName":   e.lastName,
		"position":   e.position,
		"salary":     e.salary,
		"id":         e.id,
		"updateTime": e.updateTime,
	}
	if e.managerID != nil {
		m["managerID"] = map[string]interface{}{"long": *e.managerID}
	} else {
		m["managerID"] = map[string]interface{}{"null": nil}
	}
	return m
}

func (e *employee) FromMap(m map[string]interface{}) (err error) {
	defer func() {
		if r := recover(); r != nil {
			if _, ok := r.(runtime.Error); ok {
				panic(r)
			}
			err = r.(error)
		}
	}()

	e.firstName = m["firstName"].(string)
	e.lastName = m["lastName"].(string)
	e.position = m["position"].(string)
	e.salary = m["salary"].(float64)
	e.id = m["id"].(int64)
	e.updateTime = m["updateTime"].(int64)

	if managerIDEntry, ok := m["managerID"]; managerIDEntry != nil && ok {
		if managerID, ok := managerIDEntry.(map[string]interface{})["long"]; ok {
			managerIDLong := managerID.(int64)
			e.managerID = &managerIDLong
		}
	}

	return nil
}

type avroCodec struct {
	codec *goavro.Codec
}

func (c *avroCodec) encode(e *employee) ([]byte, error) {
	return c.codec.BinaryFromNative(nil, e.Map())
}

func (c *avroCodec) decode(b []byte) (*employee, error) {
	var result interface{}
	var err error
	var native map[string]interface{}
	if result, _, err = c.codec.NativeFromBinary(b); err != nil {
		return nil, err
	}
	native = result.(map[string]interface{})

	e := employee{}
	err = e.FromMap(native)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func (c *avroCodec) name() string {
	return "avro"
}

func BenchmarkAvroEncoding(b *testing.B) {
	codec, err := goavro.NewCodec(schema)
	if err != nil {
		b.Fatalf("Failed to create codec from schema: %v", err)
	}
	c := &avroCodec{codec: codec}
	runEncodingBenchmark(c, b)
}

func BenchmarkAvroDecoding(b *testing.B) {
	codec, err := goavro.NewCodec(schema)
	if err != nil {
		b.Fatalf("Failed to create codec from schema: %v", err)
	}
	c := &avroCodec{codec: codec}
	runDecodingBenchmark(c, b)
}
