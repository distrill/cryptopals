package set01

import "bh/cryptopals/util"

// SingleByteXor - detect the single character key used to xor input.
// decrypt and return plaintext, score, and key
func SingleByteXor(decoded []byte) ([]byte, int, int) {
	winner := make([]byte, len(decoded))
	score := 0
	key := 0

	for i := 0; i < 256; i++ {
		var tempResult = make([]byte, len(decoded))
		for j := 0; j < len(decoded); j++ {
			tempResult[j] = decoded[j] ^ byte(i)
		}
		tempScore := util.GetStringScore(tempResult)
		if tempScore > score {
			winner = tempResult
			score = tempScore
			key = i
		}
	}

	return winner, score, key
}
