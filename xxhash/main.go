package main

import (
	"fmt"
	"unsafe"

	"github.com/cespare/xxhash"
)

func main() {
	var name = "harsh"
	returnDiget(name)
	fmt.Println(xxhash.Sum64String(name))
}

func returnDiget(s string) {
	var b []byte
	b = unsafe.Slice(unsafe.StringData(s), len(s))
	fmt.Println(b)
	fmt.Println([]byte(s))
	fmt.Println(Digest64(b))
}

func Digest64(b []byte) uint64 {
	return xxhash.Sum64(b)
}
