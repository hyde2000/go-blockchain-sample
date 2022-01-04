package main

import (
	"fmt"
	"go-blockchain-sample/models"
)

func main() {
	w := models.NewWallet()
	fmt.Println(w.PrivateKeyStr())
	fmt.Println(w.PublicKeyStr())
	fmt.Println(w.BlockchainAddress())

	wt := models.NewWalletTransaction(w.PrivateKey(), w.PublicKey(), w.BlockchainAddress(), "B", 1.0)
	signature, _ := wt.GenerateSignature()
	fmt.Printf("signature %s \n", signature)
}
