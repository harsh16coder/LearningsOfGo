package main

import (
	"fmt"

	"github.com/cespare/xxhash/v2"
)

func main() {
	var name = "Roger federar is not a tennis player"
	returnDiget(name)
	fmt.Println(xxhash.Sum64([]byte("harsh")))
	var num uint64 = xxhash.Sum64([]byte("1234567890"))
	fmt.Printf("%016x\n", num)
	fmt.Printf("%x", num)
}

func returnDiget(s string) {
	var b = []byte(s)
	fmt.Println(Digest64(b))
}

func Digest64(b []byte) uint64 {
	return xxhash.Sum64(b)
}
