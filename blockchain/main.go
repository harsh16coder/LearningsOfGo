package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
	"math"
	"math/big"
	"time"

	"github.com/boltdb/bolt"
)

const targetbits = 15
const dbFile = "blockchain.db"
const blocksBucket = "blocks"

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
	tip []byte
	db  *bolt.DB
}

func (bc *Blockhain) AddBlock(data string) {
	var lastHash []byte

	err := bc.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		lastHash = b.Get([]byte("l"))

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	newBlock := NewBlock(data, lastHash)

	err = bc.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		err := b.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			log.Panic(err)
		}

		err = b.Put([]byte("l"), newBlock.Hash)
		if err != nil {
			log.Panic(err)
		}

		bc.tip = newBlock.Hash

		return nil
	})
}

func NewGenesisBlock() *Block {
	return NewBlock("Genenis block", []byte{})
}

func NewBlockChain() *BlockChain {
	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))

		if b == nil {
			genesis := NewGenesisBlock()
			b, err := tx.CreateBucket([]byte(blocksBucket))
			if err != nil {
				return err
			}
			err = b.Put(genesis.Hash, genesis.Serialize())
			err = b.Put([]byte("l"), genesis.Hash)
			tip = genesis.Hash
		} else {
			tip = b.Get([]byte("l"))
		}

		return nil
	})
	bc := BlockChain{tip, db}

	return &bc
}

func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1

	return isValid
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

func deserilize(b []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(b))
	err := decoder.Decode(&block)
	if err != nil {
		fmt.Println(err)
	}
	return &block
}

func main() {
	bc := NewBlockChain()
	// bc.AddBlock("Harsh send 1 rs")
	// bc.AddBlock("Prachiti send 3 rs to Harsh")
	// bc.AddBlock("Harsh got 5 Rs")
	// for _, bc := range bc.blocks {
	// 	fmt.Printf("Prev. hash: %x\n", bc.prevHash)
	// 	fmt.Printf("Data: %s\n", bc.Data)
	// 	fmt.Printf("Hash: %x\n", bc.Hash)
	// 	pow := NewProofOfWork(bc)
	// 	fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
	// 	fmt.Println()
	// 	fmt.Println()
	// }
}
