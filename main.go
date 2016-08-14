package main

import (
	"math/rand"
	"time"
	"bufio"
	"os"
)

const (
	duplicateVowel bool = true
	removeVowel bool = false
)

func randBool() bool {
	return rand.Intn(2) == 0
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)
}
