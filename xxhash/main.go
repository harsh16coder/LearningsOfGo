package main

import (
	"fmt"

	"github.com/cespare/xxhash/v2"
	xxh128 "github.com/harsh16coder/xxhash"
	"github.com/zeebo/xxh3"
)

func main() {
	//num := xxhash.Sum64([]byte("harsh"))
	//fmt.Println(xxhash.Sum64([]byte("harsh")))
	fmt.Println(xxh3.Hash128([]byte("harsh")))
	fmt.Println(xxh128.Sum128([]byte("harsh")))
	h := xxh3.HashString128Seed("a", 0)
	fmt.Printf("%016x%016x", h.Hi, h.Lo)
	//fmt.Println(xxhash.Sum64([]byte("")))
	//fmt.Println(fmt.Sprintf("%016x", num))
	//fmt.Printf("%x", num)
	//num2 := xxh.Checksum64([]byte("harsh"))
	//fmt.Println(fmt.Sprintf("%016x", num2))
}

func returnDiget(s string) {
	var b = []byte(s)
	fmt.Println(Digest64(b))
}

func Digest64(b []byte) uint64 {
	return xxhash.Sum64(b)
}
