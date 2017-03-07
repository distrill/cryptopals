package set01

import (
	"crypto/aes"
	"log"

	"bh/cryptopals/util"
)

// DecryptAesEcb - use openssl's AES-128-ECB to decrypt given contents with given key
func DecryptAesEcb(ciphertext []byte, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}

	if len(ciphertext)%aes.BlockSize != 0 {
		log.Fatal("Ciphertext is not a multiple of the block size")
	}
	if len(ciphertext) < aes.BlockSize {
		log.Fatal("Ciphertext too short")
	}

	plaintext := make([]byte, len(ciphertext))
	mode := util.NewECBDecrypter(block)
	mode.CryptBlocks(plaintext, ciphertext)

	return plaintext
}

// EncryptAesEcb - use openssl's AES-128-ECB to encrypt given contents with given key
func EncryptAesEcb(ciphertext []byte, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}

	if len(ciphertext)%aes.BlockSize != 0 {
		log.Fatal("Ciphertext is not a multiple of the block size")
	}
	if len(ciphertext) < aes.BlockSize {
		log.Fatal("Ciphertext too short")
	}

	plaintext := make([]byte, len(ciphertext))
	mode := util.NewECBEncrypter(block)
	mode.CryptBlocks(plaintext, ciphertext)

	return plaintext
}
