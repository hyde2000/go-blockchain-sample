package main

import "go-blockchain-sample/models"

func main() {
	myAddress := "my_blockchain_address"
	blockchain := models.NewBlockchain(myAddress)

	blockchain.AddTransaction("A", "B", 1.0)
	blockchain.Mining()
	blockchain.Print()

	blockchain.AddTransaction("C", "D", 2.0)
	blockchain.AddTransaction("X", "Y", 3.0)
	blockchain.Mining()
	blockchain.Print()
}
