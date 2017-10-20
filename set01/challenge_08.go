package set01

import (
	"encoding/hex"
	"log"
)

// DetectAesInEcbMode - find (but don't do anything with)
// the single ciphertext that has been encrypted with AES-128-ECB
func DetectAesInEcbMode(hexLines [][]byte) int {
	blockSize := 16
	lines := getBytesFromHexStrings(hexLines)
	for i, line := range lines {
		if IsEcb(line, blockSize) {
			return i
		}
	}
	return -1
}

// IsEcb - is the given input (with given blocksize) encrypted with ECB or not?
func IsEcb(line []byte, blockSize int) bool {
	found := make(map[string]bool)
	// ecb was used if there are any repeating blocks
	blocks := splitLineIntoBlocks(line, blockSize)
	for _, block := range blocks {
		if found[string(block)] {
			return true
		}
		found[string(block)] = true
	}
	return false
}

func getBytesFromHexStrings(hexLines [][]byte) [][]byte {
	byteLines := make([][]byte, 0)
	for _, line := range hexLines {
		decoded, err := hex.DecodeString(string(line))
		if err != nil {
			log.Fatal(err)
		}
		byteLines = append(byteLines, decoded)
	}
	return byteLines
}

func splitLineIntoBlocks(line []byte, blockSize int) [][]byte {
	lines := make([][]byte, 0)
	for i := 0; (i+1)*blockSize < len(line); i++ {
		lines = append(lines, line[i*blockSize:(i+1)*blockSize])
	}
	return lines
}
