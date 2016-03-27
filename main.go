package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
	"strings"
)

func main() {
	numWordsPtr := flag.Int("words", 6, "number of words to generate")
	flag.Parse()

	var words []string
	wordCnt := big.NewInt(int64(len(Words)))
	for i := 0; i < *numWordsPtr; i++ {
		r, _ := rand.Int(rand.Reader, wordCnt)
		words = append(words, Words[int(r.Int64())])
	}
	fmt.Println(strings.Join(words, " "))
}
