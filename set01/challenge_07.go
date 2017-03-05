package set01

import (
	"crypto/aes"
	"encoding/base64"
	"log"

	"bh/cryptopals/util"
)

// DecryptAesEcb - use openssl's AES-128-ECB to decript given contents with given key
func DecryptAesEcb(contents []byte, key []byte) []byte {
	ciphertext, err := base64.StdEncoding.DecodeString(string(contents))
	if err != nil {
		log.Fatal(err)
	}
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
