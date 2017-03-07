package set02

import (
	"bh/cryptopals/set01"
)

// EncryptCbc - Encrypt plaintext using hand-rolled CBC, which
// uses AES-128-ECB under the hood.
func EncryptCbc(plaintext []byte, key []byte, iv []byte, blocksize int) []byte {
	ciphertext := make([]byte, 0)
	cipherblock := make([]byte, blocksize)
	padded := PkcsPadding(plaintext, blocksize)
	numBlocks := len(padded) / blocksize

	for i := 0; i < numBlocks; i++ {
		plainBlock := padded[i*blocksize : (i+1)*blocksize]
		if i == 0 {
			cipherblock = iv
		}
		plainXorBlock := set01.FixedXor(plainBlock, cipherblock)
		cipherblock = set01.EncryptAesEcb(plainXorBlock, key)
		ciphertext = append(ciphertext, cipherblock...)
	}

	return ciphertext
}

// DecryptCbc - Decrypt ciphertext using hand-rolled CBC, which
// uses AES-128-ECB under the hood.
func DecryptCbc(ciphertext []byte, key []byte, iv []byte, blocksize int) []byte {
	plaintext := make([]byte, 0)
	numBlocks := len(ciphertext) / blocksize

	for i := numBlocks - 1; i >= 0; i-- {
		currentCipherBlock := ciphertext[i*blocksize : (i+1)*blocksize]
		currentEcbBlock := set01.DecryptAesEcb(currentCipherBlock, key)
		var previousCipherBlock []byte
		if i == 0 {
			previousCipherBlock = iv
		} else {
			previousCipherBlock = ciphertext[(i-1)*blocksize : i*blocksize]
		}

		plainblock := set01.FixedXor(currentEcbBlock, previousCipherBlock)
		plaintext = append(plainblock, plaintext...)
	}

	return plaintext
}
