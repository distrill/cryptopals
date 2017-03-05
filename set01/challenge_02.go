package set01

import (
	"log"
)

// FixedXor - xor two strings against one another and return the result
func FixedXor(one []byte, two []byte) []byte {
	// they need to be the same length for this operation
	if len(one) != len(two) {
		log.Fatal("Starting strings must be of same length")
	}

	// one is picked arbitrarily, one and two have same length
	length := len(one)
	var result = make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = one[i] ^ two[i]
	}

	return result
}
