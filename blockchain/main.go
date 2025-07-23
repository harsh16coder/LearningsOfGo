package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	Timestamp int64
	Hash      []byte
	prevHash  []byte
	Data      []byte
}

func (b *Block) setHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.prevHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

func NewBlock(data string, prevHash []byte) *Block {
	initialBlock := &Block{time.Now().Unix(), []byte{}, prevHash, []byte(data)}
	initialBlock.setHash()
	return initialBlock
}

type BlockChain struct {
	blocks []*Block
}

func (bc *BlockChain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

func NewGenesisBlock() *Block {
	return NewBlock("Genenis block", []byte{})
}

func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}

func main() {
	bc := NewBlockChain()
	bc.AddBlock("Harsh send 1 rs")
	bc.AddBlock("Prachiti send 3 rs to Harsh")
	bc.AddBlock("Harsh got 5 Rs")
	for _, bc := range bc.blocks {
		fmt.Printf("Prev. hash: %x\n", bc.prevHash)
		fmt.Printf("Data: %s\n", bc.Data)
		fmt.Printf("Hash: %x\n", bc.Hash)
		fmt.Println()
	}
}
