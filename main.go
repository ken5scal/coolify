package main

import (
	"math/rand"
	"time"
	"bufio"
	"os"
	"fmt"
)

const (
	duplicateVowelCase bool = true
	removeVowelCase bool = false
)

func randBool() bool {
	return rand.Intn(2) == 0
}

func duplicateVowel(word []byte, i int) []byte {
	return append(word[:i+1], word[i:]...)
}

func removeVowel(word []byte, i int) []byte {
	return append(word[:i], word[i+1:]...)
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		word := []byte(s.Text())
		if randBool() {
			var vI int = -1
			for i, char := range word {
				switch char {
				case 'a', 'i', 'u', 'e', 'o', 'A', 'I', 'U', 'E', 'O':
					if randBool() {
						vI = i
					}
				}
			}

			if vI >= 0 {
				switch randBool() {
				case duplicateVowelCase:
					word = duplicateVowel(word, vI)
				case removeVowelCase:
					word = removeVowel(word, vI)
				}
			}
		}
		fmt.Println(string(word))
	}
}
