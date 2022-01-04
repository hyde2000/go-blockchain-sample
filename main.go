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
}
