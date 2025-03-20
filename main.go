package main

import (
	"encoding/hex"
	"fmt"

	"vanity-walk/internal"
)

func main() {
	target := "1337" // 👈 your vanity pattern (no 0x)

	result := internal.FindVanityAddress(target)

	fmt.Println("🔥 Match Found!")
	fmt.Println("Address:    ", result.Address)
	fmt.Println("Private Key:", hex.EncodeToString(result.PrivateKey))
}