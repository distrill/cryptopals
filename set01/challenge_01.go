package set01

import (
	"encoding/base64"
	"encoding/hex"
	"log"
)

// HexToBase64 - turn hex string into base64 string
func HexToBase64(start string) string {
	decoded, err := hex.DecodeString(start)
	if err != nil {
		log.Fatal(err)
	}
	return base64.StdEncoding.EncodeToString(decoded)
}
