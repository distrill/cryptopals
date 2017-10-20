package main

import (
	"bh/cryptopals/set02"
	"fmt"
)

func main() {
	input := make([]byte, 1000)
	mode := set02.DetectionOracle(input, set02.EncryptionOracle)
	fmt.Printf("guessed: %s\n", mode)
}
