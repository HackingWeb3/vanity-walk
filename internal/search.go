package internal

import (
	"encoding/hex"
	"strings"
)

type VanityResult struct {
	PrivateKey []byte
	Address    string
}

// FindVanityAddress runs a brute-force loop to find an ETH address with the given hex prefix (after 0x)
func FindVanityAddress(prefix string) VanityResult {
	prefix = strings.ToLower(prefix)
	for {
		priv, pub := GenerateKeyPair()
		addrBytes := DeriveAddress(pub)
		addr := "0x" + hex.EncodeToString(addrBytes)

		if strings.HasPrefix(addr, "0x"+prefix) {
			return VanityResult{
				PrivateKey: priv,
				Address:    addr,
			}
		}
	}
}
