package main

import "go-blockchain-sample/models"

func main() {
	blockchain := models.NewBlockchain()
	blockchain.Print()

	blockchain.AddTransaction("A", "B", 1.0)

	previousHash := blockchain.LastBlock().Hash()
	blockchain.CreateBlock(5, previousHash)
	blockchain.Print()
	blockchain.CreateBlock(2, previousHash)
	blockchain.Print()
}
