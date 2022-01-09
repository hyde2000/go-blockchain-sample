package models

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ripemd160"
)

type Wallet struct {
	privateKey        *ecdsa.PrivateKey
	publicKey         *ecdsa.PublicKey
	blockchainAddress string
}

// NewWallet ウォレット作成（アドレス生成アルゴリズムはBitcoin仕様）
func NewWallet() *Wallet {
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	publicKey := &privateKey.PublicKey
	// Hashing on the public key by SHA-256
	h2 := sha256.New()
	h2.Write(publicKey.X.Bytes())
	h2.Write(publicKey.Y.Bytes())
	digest2 := h2.Sum(nil)
	// Hashing on the result of SHA-256 by RIPEMD-160
	h3 := ripemd160.New()
	h3.Write(digest2)
	digest3 := h3.Sum(nil)
	// Add version byte
	vd4 := make([]byte, 21)
	vd4[0] = 0x00
	copy(vd4[1:], digest3[:])
	// Hashing on the extended RIPEMD-160 by SHA-256
	h5 := sha256.New()
	h5.Write(vd4)
	digest5 := h5.Sum(nil)
	// ReHashing on the SHA-256 by SHA-256
	h6 := sha256.New()
	h6.Write(digest5)
	digest6 := h6.Sum(nil)
	// Take the first 4 bytes
	checkSum := digest6[:4]
	dc8 := make([]byte, 25)
	// Add the checksum at the end of extended RIPEMD-160 hash (21 + 4 = 25)
	copy(dc8[:21], vd4[:])
	copy(dc8[21:], checkSum)

	address := base58.Encode(dc8)

	return &Wallet{
		privateKey:        privateKey,
		publicKey:         publicKey,
		blockchainAddress: address,
	}
}

func (w *Wallet) PrivateKey() *ecdsa.PrivateKey {
	return w.privateKey
}

func (w *Wallet) PrivateKeyStr() string {
	return fmt.Sprintf("%x", w.privateKey.D.Bytes())
}

func (w *Wallet) PublicKey() *ecdsa.PublicKey {
	return w.publicKey
}

func (w *Wallet) PublicKeyStr() string {
	return fmt.Sprintf("%x%x", w.publicKey.X.Bytes(), w.publicKey.Y.Bytes())
}

func (w *Wallet) BlockchainAddress() string {
	return w.blockchainAddress
}

func (w *Wallet) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		PrivateKey        string `json:"private_key"`
		PublicKey         string `json:"public_key"`
		BlockchainAddress string `json:"blockchain_address"`
	}{
		PublicKey:         w.PublicKeyStr(),
		PrivateKey:        w.PrivateKeyStr(),
		BlockchainAddress: w.BlockchainAddress(),
	})
}
