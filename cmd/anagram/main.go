package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jamiecuthill/anagram"
)

var dict = flag.String("dict", "dictionary.txt", "path to the dictionary file")

func init() {
	flag.Parse()
	anagram.Source(*dict)
}

func main() {
	var words anagram.Sentence
	for _, w := range os.Args[1:] {
		words = append(words, anagram.Word(w))
	}
	for _, out := range words.Anagrams() {
		fmt.Println(out)
	}
}
