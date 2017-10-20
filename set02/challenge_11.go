package set02

import (
	"bh/cryptopals/set01"
	"bh/cryptopals/util"
	"fmt"
)

type blackBox func([]byte, int) []byte

func generateAesKey(keyLength int) []byte {
	return util.GetRandomBytes(keyLength)
}

func shouldEcb() bool {
	return util.GetRandomInt(2)%2 == 0
}

func generateCipherText(input []byte) []byte {
	cipher := append(util.GetRandomBytes(5+util.GetRandomInt(5)), input...)
	return append(cipher, util.GetRandomBytes(5+util.GetRandomInt(5))...)
}

// EncryptionOracle - black box that might be encrypting either ECB or CBC
func EncryptionOracle(input []byte, blocksize int) []byte {
	key := generateAesKey(blocksize)
	plaintext := generateCipherText(input)
	if shouldEcb() {
		fmt.Println("...using ECB")
		ciphertext := set01.EncryptAesEcb(PkcsPadding(plaintext, blocksize), key)
		return ciphertext
	}
	fmt.Println("...using CBC")
	iv := util.GetRandomBytes(blocksize)
	ciphertext := EncryptCbc(plaintext, key, iv, blocksize)
	return ciphertext
}

// DetectionOracle - can ECB or CBC if pointed at a black box that might be encrypting either
func DetectionOracle(input []byte, fn blackBox) string {
	blocksize := 16
	ciphertext := fn(input, blocksize)
	if set01.IsEcb(ciphertext, blocksize) {
		return "ECB"
	}
	return "CBC"
}
