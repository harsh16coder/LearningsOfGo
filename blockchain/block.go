package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"time"
)

type Block struct {
	Timestamp int64
	Hash      []byte
	PrevHash  []byte
	Data      []byte
	Nonce     int
}

func NewBlock(data string, PrevHash []byte) *Block {
	initialBlock := &Block{time.Now().Unix(), []byte{}, PrevHash, []byte(data), 0}
	pow := NewProofOfWork(initialBlock)
	Nonce, hash := pow.Run()
	initialBlock.Hash = hash[:]
	initialBlock.Nonce = Nonce
	return initialBlock
}

func NewGenesisBlock() *Block {
	return NewBlock("Genenis block", []byte{})
}

func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(b)
	if err != nil {
		fmt.Println(err)
	}
	return result.Bytes()
}

func Deserilize(b []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(b))
	err := decoder.Decode(&block)
	if err != nil {
		fmt.Println(err)
	}
	return &block
}
