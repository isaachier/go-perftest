.PHONY: all benchmark clean messages

all: benchmark

benchmark:
	go test -bench=. -benchmem -cpuprofile=cpu.out -memprofile=mem.out ./serialization

clean:
	rm -f mem.out

messages:
	protoc --go_out=. serialization/schema/protobuf/employee.proto
	protoc --gofast_out=. serialization/schema/gogoprotobuf/employee.proto
