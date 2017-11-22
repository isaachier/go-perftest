package serialization

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/Pallinder/go-randomdata"
)

type employee struct {
	firstName  string
	lastName   string
	position   string
	salary     float64
	id         int64
	managerID  *int64
	updateTime int64
}

var dataset []employee
var positions = [...]string{"engineer", "janitor", "manager"}

func init() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 1000; i++ {
		dataset = append(dataset, employee{
			firstName:  randomdata.FirstName(randomdata.RandomGender),
			lastName:   randomdata.LastName(),
			position:   positions[rand.Intn(len(positions))],
			salary:     float64(rand.Intn(1000000)),
			id:         rand.Int63(),
			updateTime: time.Now().Unix(),
		})
		if randomdata.Boolean() {
			managerID := rand.Int63()
			dataset[len(dataset)-1].managerID = &managerID
		}
	}
}

type codec interface {
	encode(e *employee) ([]byte, error)
	decode(b []byte) (*employee, error)
	name() string
}

func runEncodingBenchmark(c codec, b *testing.B) {
	binaryFile, err := ioutil.TempFile("testdata", c.name())
	if err != nil {
		b.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(binaryFile.Name())

	var buffer bytes.Buffer
	for _, e := range dataset {
		binaryData, err := c.encode(&e)
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

func runDecodingBenchmark(c codec, b *testing.B) {
	for i := 1; i <= 3; i++ {
		runDecodingBenchmarkHelper(c, b, i)
	}
}

func runDecodingBenchmarkHelper(c codec, b *testing.B, index int) {
	file, err := os.Open(fmt.Sprintf("testdata/%s%d.data", c.name(), index))
	if err != nil {
		b.Fatalf("Failed to open data file: %v", err)
	}
	defer file.Close()

	for {
		data := make([]byte, 8)
		if numBytes, err := file.Read(data); numBytes == 0 {
			break
		} else if err != nil {
			b.Fatalf("Failed to decode number of bytes for the record: %v", err)
		}

		numRecordBytes := binary.BigEndian.Uint64(data)
		data = make([]byte, numRecordBytes)
		if numBytes, err := file.Read(data); err != nil {
			b.Fatalf("Failed to read record data: %v", err)
		} else if uint64(numBytes) != numRecordBytes {
			b.Fatalf("Failed to read required number of bytes"+
				", expected %d, actual %d", numRecordBytes, numBytes)
		}
		_, err := c.decode(data[0:])
		if err != nil {
			b.Fatalf("Failed to decode employee record: %v", err)
		}
	}
}
