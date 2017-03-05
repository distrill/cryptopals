package main

import (
	"bh/cryptopals/set01"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	contents, err := ioutil.ReadFile("data/01_06.txt")
	if err != nil {
		log.Fatal(err)
	}
	ciphertext, err := base64.StdEncoding.DecodeString(string(contents))
	if err != nil {
		log.Fatal(err)
	}

	plaintext, key := set01.BreakRepeatingKeyXor(ciphertext)
	fmt.Printf("%s\n\n%s\n", plaintext, key)
}
