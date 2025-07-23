package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
	"strconv"
	"time"
)

const targetbits = 15

type Block struct {
	Timestamp int64
	Hash      []byte
	prevHash  []byte
	Data      []byte
	nonce     int
}

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

func NewProofOfWork(block *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetbits))

	pow := &ProofOfWork{block, target}
	return pow
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join([][]byte{
		pow.block.prevHash,
		pow.block.Data,
		[]byte(fmt.Sprintf("%x", pow.block.Timestamp)),
		[]byte(fmt.Sprintf("%x", int64(targetbits))),
		[]byte(fmt.Sprintf("%x", int64(nonce))),
	},
		[]byte{})
	return data
}

const maxNonce = math.MaxInt64

func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)
	for nonce < maxNonce {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Print("\n\n")

	return nonce, hash[:]
}

func NewBlock(data string, prevHash []byte) *Block {
	initialBlock := &Block{time.Now().Unix(), []byte{}, prevHash, []byte(data), 0}
	pow := NewProofOfWork(initialBlock)
	nonce, hash := pow.Run()
	initialBlock.Hash = hash[:]
	initialBlock.nonce = nonce
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

func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1

	return isValid
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
		pow := NewProofOfWork(bc)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
		fmt.Println()
	}
}
