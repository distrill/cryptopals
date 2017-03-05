package set01

import (
	"bh/cryptopals/util"
)

// BreakRepeatingKeyXor - decode contents that has been base64'd max keylength of 40
func BreakRepeatingKeyXor(ciphertext []byte) ([]byte, []byte) {
	keySize := getKeySize(ciphertext)
	key := getKey(ciphertext, keySize)
	plaintext := DecryptRepeatingKeyXor(ciphertext, key)

	return plaintext, key
}

func getKey(contents []byte, keySize int) []byte {
	key := make([]byte, keySize)
	blocks := make([][]byte, keySize)
	for i, c := range contents {
		blocks[i%keySize] = append(blocks[i%keySize], c)
	}
	for i, block := range blocks {
		_, _, k := SingleByteXor(block)
		key[i] = byte(k)
	}
	return key
}

func getKeySize(contents []byte) int {
	var winningKeySize int
	var winningKeyScore float32
	for keySize := 2; keySize <= 40; keySize++ {
		numBlocks := len(contents) / keySize
		ham := float32(0)
		for i := 0; i < numBlocks-1; i += 2 {
			one := contents[keySize*i : keySize*(i+1)]
			two := contents[keySize*(i+1) : keySize*(i+2)]
			distnace := util.Hamming(string(one), string(two))
			ham += float32(distnace)
		}

		// normalize distance for number of blocks and key size
		ham /= float32(numBlocks)
		ham /= float32(keySize)

		if winningKeyScore == float32(0) || ham < winningKeyScore {
			winningKeyScore = ham
			winningKeySize = keySize
		}
	}
	return winningKeySize
}
