package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

var numWords int = 6

func main() {
	var words []string
	wordCnt := big.NewInt(int64(len(Words)))
	for i := 0; i < numWords; i++ {
		r, _ := rand.Int(rand.Reader, wordCnt)
		words = append(words, Words[int(r.Int64())])
	}
	fmt.Println(strings.Join(words, " "))
}
