package set02

import "math"

// PkcsPadding - pad a string out to blocksize length using PKCS#7
func PkcsPadding(unpadded []byte, blockSize int) []byte {
	startLength := len(unpadded)
	targetLength := int(math.Ceil(float64(startLength/blockSize))) * blockSize
	padLength := targetLength - len(unpadded)
	padded := make([]byte, targetLength)
	copy(padded, unpadded)
	for i := targetLength - 1; i >= startLength; i-- {
		padded[i] = byte(padLength)
	}
	return padded
}
