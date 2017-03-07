package main

import (
	"bh/cryptopals/set02"
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

func main() {
	key := []byte("YELLOW SUBMARINE")
	blocksize := 16
	iv := make([]byte, blocksize)
	// contents, _ := ioutil.ReadFile("data/02_10.txt")
	// decoded, _ := base64.StdEncoding.DecodeString(string(contents))
	// result := set02.DecryptCbc(decoded, key, iv, blocksize)
	// fmt.Println(string(result))
	// fmt.Println(result)
	contents, _ := ioutil.ReadFile("data/02_10_plain.txt")
	ciphertext := set02.EncryptCbc(contents, key, iv, blocksize)
	encoded := base64.StdEncoding.EncodeToString(ciphertext)
	fmt.Println(encoded)
}
