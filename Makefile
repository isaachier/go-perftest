.PHONY: all benchmark clean

all: benchmark

benchmark:
	go test -bench . -benchmem -memprofile mem.out ./serialization

clean:
	rm -f mem.out
