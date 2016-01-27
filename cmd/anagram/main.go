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
	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("Usage: anagram hello world")
		os.Exit(0)
	}

	if len(args) == 1 {
		anagrams := anagram.Word(args[0]).Anagrams()
		for _, word := range anagrams {
			fmt.Println(word)
		}
		os.Exit(0)
	}

	var words anagram.Sentence
	for _, w := range flag.Args() {
		words = append(words, anagram.Word(w))
	}
	for _, out := range words.Anagrams() {
		fmt.Println(out)
	}
}
