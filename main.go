package main

import (
	"fmt"
	"go-blockchain-sample/models"
)

func main() {
	walletM := models.NewWallet()
	walletA := models.NewWallet()
	walletB := models.NewWallet()

	wt := models.NewWalletTransaction(walletA.PrivateKey(), walletA.PublicKey(), walletA.BlockchainAddress(), walletB.BlockchainAddress(), 1.0)

	blockchain := models.NewBlockchain(walletM.BlockchainAddress())
	signature, _ := wt.GenerateSignature()
	isAdded := blockchain.AddTransaction(walletA.BlockchainAddress(), walletB.BlockchainAddress(), 1.0, walletA.PublicKey(), signature)
	fmt.Println(isAdded)

	blockchain.Mining()
	// blockchain.Print()

	fmt.Printf("A %.1f\n", blockchain.CalculateTotalAmount(walletA.BlockchainAddress()))
}
