package util

import (
	"bufio"
	"log"
	"os"
)

// Hamming - get hamming distance between two strings
func Hamming(a string, b string) int {
	if len(a) != len(b) {
		log.Fatal("Strings must be of same length")
	}

	distance := 0
	for i := 0; i < len(a); i++ {
		distance += singleByteHamming(uint(a[i]), uint(b[i]))
	}

	return distance
}

func singleByteHamming(a uint, b uint) int {
	distance := 0
	val := uint(a) ^ uint(b)
	for val != 0 {
		distance++
		val &= val - 1
	}
	return distance
}

// GetStringScore - score a given text on english character frequency
// (plus some secret sause)
func GetStringScore(result []byte) int {
	etaoin := map[byte]int{
		' ': 13,
		'e': 12, 't': 11, 'a': 10, 'o': 9,
		'i': 8, 'n': 7, 's': 6, 'h': 5,
		'r': 4, 'd': 3, 'l': 2, 'u': 1,
	}
	score := 0
	for i := 0; i < len(result); i++ {
		score += etaoin[result[i]]
	}
	return score
}

// GetLinesFromFile - read a file and return array of strings,
// one string for each line in the file
func GetLinesFromFile(filename string) [][]byte {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	var result = make([][]byte, 0)
	for scanner.Scan() {
		result = append(result, []byte(scanner.Text()))
	}
	return result
}
