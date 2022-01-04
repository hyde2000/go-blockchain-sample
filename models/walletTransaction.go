package models

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"go-blockchain-sample/utils"
	"golang.org/x/xerrors"
)

type WalletTransaction struct {
	senderPrivateKey           *ecdsa.PrivateKey
	senderPublicKey            *ecdsa.PublicKey
	senderBlockchainAddress    string
	recipientBlockchainAddress string
	value                      float32
}

func NewWalletTransaction(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey, sender string, recipient string, value float32) *WalletTransaction {
	return &WalletTransaction{
		senderPrivateKey:           privateKey,
		senderPublicKey:            publicKey,
		senderBlockchainAddress:    sender,
		recipientBlockchainAddress: recipient,
		value:                      value,
	}
}

func (wt *WalletTransaction) GenerateSignature() (*utils.Signature, error) {
	m, _ := json.Marshal(wt)
	h := sha256.Sum256(m)

	r, s, err := ecdsa.Sign(rand.Reader, wt.senderPrivateKey, h[:])
	if err != nil {
		return &utils.Signature{}, xerrors.Errorf("transaction signature sign error %w:", err)
	}

	return &utils.Signature{R: r, S: s}, nil
}

func (wt *WalletTransaction) MarshallJSON() ([]byte, error) {
	return json.Marshal(struct {
		Sender    string  `json:"sender_blockchain_address"`
		Recipient string  `json:"recipient_blockchain_address"`
		Value     float32 `json:"value"`
	}{
		Sender:    wt.senderBlockchainAddress,
		Recipient: wt.recipientBlockchainAddress,
		Value:     wt.value,
	})
}
