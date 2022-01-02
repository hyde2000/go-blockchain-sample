package main

import (
	"fmt"
	"strings"
	"time"
)

type Block struct {
	nonce        int
	previousHash string
	timestamp    int64
	transactions []string
}

func NewBlock(nonce int, previousHash string) *Block {
	return &Block{
		nonce:        nonce,
		previousHash: previousHash,
		timestamp:    time.Now().UnixNano(),
	}
}

func (b *Block) Print() {
	fmt.Printf("timestamp: %d \n", b.timestamp)
	fmt.Printf("nonce %d \n", b.nonce)
	fmt.Printf("previous_hash %s \n", b.previousHash)
	fmt.Printf("transactions %s \n", b.transactions)
}

type Blockchain struct {
	transactionPool []string
	chain           []*Block
}

func NewBlockchain() *Blockchain {
	bc := &Blockchain{}
	bc.CreateBlock(0, "init hash")

	return bc
}

func (bc *Blockchain) CreateBlock(nonce int, previousHash string) *Block {
	b := NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)

	return b
}

func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s Chain %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
	}
}

func main() {
	blockchain := NewBlockchain()
	blockchain.Print()
	blockchain.CreateBlock(5, "hash 1")
	blockchain.Print()
	blockchain.CreateBlock(2, "hash 2")
	blockchain.Print()
}
