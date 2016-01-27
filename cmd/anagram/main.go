package main

import (
	"flag"
	"fmt"

	"github.com/jamiecuthill/anagram"
)

var dict = flag.String("dict", "dictionary.txt", "path to the dictionary file")

func init() {
	flag.Parse()
	anagram.Source(*dict)
}

func main() {
	var words anagram.Sentence
	for _, w := range flag.Args() {
		words = append(words, anagram.Word(w))
	}
	for _, out := range words.Anagrams() {
		fmt.Println(out)
	}
}
