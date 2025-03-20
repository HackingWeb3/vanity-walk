package internal

import (
	"crypto/rand"
	"log"

	secp256k1 "github.com/decred/dcrd/dcrec/secp256k1/v4"
)

// GenerateKeyPair returns the private key bytes and the uncompressed public key (no prefix)
func GenerateKeyPair() ([]byte, []byte) {
	privKeyBytes := make([]byte, 32)
	_, err := rand.Read(privKeyBytes)
	if err != nil {
		log.Fatal(err)
	}

	privKey := secp256k1.PrivKeyFromBytes(privKeyBytes)
	pubKey := privKey.PubKey()
	pubKeyBytes := pubKey.SerializeUncompressed()[1:] // remove 0x04 prefix. Ethereum uses the uncompressed public key, but without the prefix byte.

	return privKeyBytes, pubKeyBytes
}
