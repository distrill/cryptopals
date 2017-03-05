package set01

// EncryptRepeatingKeyXor - encrypt plaintext with repeating
// key xor and given key
func EncryptRepeatingKeyXor(plaintext []byte, key []byte) []byte {
	result := make([]byte, len(plaintext))
	keyLen := len(key)
	for i, b := range plaintext {
		result[i] = b ^ key[i%keyLen]
	}
	return result
}

// DecryptRepeatingKeyXor - decrypt ciphertext with repeating
// key xor and given key
func DecryptRepeatingKeyXor(ciphertext []byte, key []byte) []byte {
	plaintext := make([]byte, len(ciphertext))
	keylen := len(key)
	for i := range ciphertext {
		plaintext[i] = ciphertext[i] ^ key[i%keylen]
	}
	return plaintext
}
