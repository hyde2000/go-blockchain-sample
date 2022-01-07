package utils

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
)

type Signature struct {
	R *big.Int
	S *big.Int
}

func (s *Signature) String() string {
	return fmt.Sprintf("%064x%064x", s.R, s.S)
}

func String2BigIntTuple(s string) (big.Int, big.Int) {
	var bix, biy big.Int
	bx, err := hex.DecodeString(s[:64])
	if err != nil {
		log.Println("ERROR: decode failed", err)
	}
	by, err := hex.DecodeString(s[64:])
	if err != nil {
		log.Println("ERROR: decode failed", err)
	}
	_ = bix.SetBytes(bx)
	_ = biy.SetBytes(by)

	return bix, biy
}

func PublicKeyFromString(s string) *ecdsa.PublicKey {
	x, y := String2BigIntTuple(s)

	return &ecdsa.PublicKey{Curve: elliptic.P256(), X: &x, Y: &y}
}

func PrivateKeyFromString(s string, publicKey *ecdsa.PublicKey) *ecdsa.PrivateKey {
	var bi big.Int

	b, _ := hex.DecodeString(s[:])
	_ = bi.SetBytes(b)

	return &ecdsa.PrivateKey{PublicKey: *publicKey, D: &bi}
}

func SignatureFromString(str string) *Signature {
	r, s := String2BigIntTuple(str)

	return &Signature{R: &r, S: &s}
}
