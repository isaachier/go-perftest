package serialization

import (
	"bytes"
	"encoding/binary"
	"io/ioutil"
	"os"
	"testing"

	"github.com/linkedin/goavro"
)

var schema string

func init() {
	schemaBytes, err := ioutil.ReadFile("schema/employee.avsc")
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

func BenchmarkAvroEncoding(b *testing.B) {
	codec, err := goavro.NewCodec(schema)
	if err != nil {
		b.Fatalf("Failed to create codec from schema: %v", err)
	}

	binaryFile, err := ioutil.TempFile("testdata", "avro")
	if err != nil {
		b.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(binaryFile.Name())

	var buffer bytes.Buffer
	for _, employee := range dataset {
		binaryData, err := codec.BinaryFromNative(nil, employee.Map())
		if err != nil {
			b.Fatalf("Failed to encode employee record: %v", err)
		}

		sizeBytes := make([]byte, 8)
		binary.BigEndian.PutUint64(sizeBytes[0:], uint64(len(binaryData)))
		_, err = buffer.Write(sizeBytes)
		if err != nil {
			b.Fatalf("Failed to write binary data to buffer: %v", err)
		}

		_, err = buffer.Write(binaryData)
		if err != nil {
			b.Fatalf("Failed to write binary data to buffer: %v", err)
		}

		_, err = buffer.WriteTo(binaryFile)
		if err != nil {
			b.Fatalf("Failed to write employee to file: %v", err)
		}

		buffer.Reset()
	}
}
