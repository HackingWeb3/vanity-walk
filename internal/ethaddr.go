package internal

import (
	"golang.org/x/crypto/sha3"
)

// DeriveAddress returns a 20-byte Ethereum address from the public key
func DeriveAddress(pubKey []byte) []byte {
	hash := sha3.NewLegacyKeccak256()
	hash.Write(pubKey)
	fullHash := hash.Sum(nil)

	// Ethereum address = last 20 bytes
	return fullHash[12:]
}
