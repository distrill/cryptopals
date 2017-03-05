package set01

import (
	"fmt"
)

// DetectSingleCharacterXor - detect the single character key used to xor a
// line in a file's contents. decrypt and return winning plaintext, score, and key
func DetectSingleCharacterXor(lines [][]byte) ([]byte, int, int) {
	var winner = make([]byte, 0)
	var score = 0
	var key = 0
	fmt.Println(len(lines))
	for _, line := range lines {
		result, tempScore, i := SingleByteXor([]byte(line))
		if tempScore > score {
			score = tempScore
			winner = result
			key = i
		}
	}
	return winner, score, key
}
