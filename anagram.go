package anagram

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var dictionary []Word

// Expensive init function that loads the word dictionary
func init() {
	f, err := os.Open("dictionary.txt")
	if err != nil {
		panic("dictionary.txt missing")
	}
	defer f.Close()
	buf := bufio.NewReader(f)
	var i = 0
	for {
		line, isPrefix, err := buf.ReadLine()
		if err != nil {
			return
		}
		if isPrefix {
			dictionary[i] = dictionary[i] + Word(line)
			continue
		}
		dictionary = append(dictionary, Word(line))
		i++
	}
}

// Word is a single word
type Word string

// Anagrams of word
func (w Word) Anagrams() []Word {
	return Dictionary(w.Occurences())
}

// Sentence is a collection of words
type Sentence []Word

// Occurences is the count of each character in the sentence
func (s Sentence) Occurences() Occurences {
	return s.Word().Occurences()
}

// Anagrams of the sentence
func (s Sentence) Anagrams() []Sentence {
	return anagrams(s.Occurences())
}

func anagrams(occurences Occurences) []Sentence {
	acc := make([]Sentence, 0, 1)
	if len(occurences) == 0 {
		acc = append(acc, Sentence{})
		return acc
	}

	for _, occurence := range occurences.Combinations() {
		for _, word := range Dictionary(occurence) {
			sentences := []Sentence{}
			for _, tail := range anagrams(occurences.Subtract(occurence)) {
				sentences = append(sentences, append(Sentence{word}, tail...))
			}
			acc = append(acc, sentences...)
		}
	}
	return acc
}

type occurence struct {
	Char rune
	Freq int
}

func (o occurence) String() string {
	return fmt.Sprintf("%c:%d", o.Char, o.Freq)
}

// Occurences is a list of character and frequency, it must be sorted alphabetically
type Occurences []occurence

func (o Occurences) Len() int {
	return len(o)
}

func (o Occurences) Swap(i, j int) {
	o[i], o[j] = o[j], o[i]
}

func (o Occurences) Less(i, j int) bool {
	return o[i].Char < o[j].Char
}

func (o Occurences) key() string {
	acc := make([]string, len(o))
	for i := range o {
		acc[i] = o[i].String()
	}
	return strings.Join(acc, " ")
}

// Subtract removes the occurences and returns the result
func (o Occurences) Subtract(sub Occurences) Occurences {
	remove := make(map[rune]int)
	for i := range sub {
		remove[sub[i].Char] = sub[i].Freq
	}

	acc := make(Occurences, 0, len(o))
	for i := range o {
		if n, ok := remove[o[i].Char]; ok {
			if frq := o[i].Freq - n; frq > 0 {
				acc = append(acc, occurence{o[i].Char, frq})
			}
			continue
		}
		acc = append(acc, o[i])
	}

	return acc
}

// Combinations returns all permutations of the occurences
func (o Occurences) Combinations() []Occurences {
	if len(o) == 0 {
		return []Occurences{{}}
	}

	var acc []Occurences
	head, tail := o[0], o[1:]
	for _, next := range tail.Combinations() {
		for i := 0; i <= head.Freq; i++ {
			var occ Occurences
			if i == 0 {
				occ = next
			} else {
				occ = append(Occurences{{head.Char, i}}, next...)
			}
			acc = append(acc, occ)
		}
	}
	return acc
}

// Occurences is the count of each character in the word
func (w Word) Occurences() Occurences {
	s := strings.ToLower(string(w))
	acc := make(map[rune]int)
	for _, r := range s {
		acc[r]++
	}

	out := make(Occurences, 0, len(acc))
	for k, v := range acc {
		out = append(out, occurence{k, v})
	}
	sort.Sort(out)
	return out
}

// Word converts the Sentence to a Word by concatenating with no separator
func (s Sentence) Word() Word {
	str := make([]string, len(s))
	// If only you could cast or copy a slice of string to slice of Word
	for i := range s {
		str[i] = string(s[i])
	}
	return Word(strings.Join(str, ""))
}

var lazyDictionary map[string][]Word

// Dictionary for the occurences
func Dictionary(o Occurences) []Word {
	if len(lazyDictionary) == 0 {
		loadDictionary()
	}
	if v, ok := lazyDictionary[o.key()]; ok {
		return v
	}
	return []Word{}
}

func loadDictionary() {
	lazyDictionary = make(map[string][]Word)
	for i := range dictionary {
		w := Word(dictionary[i])
		k := w.Occurences().key()
		lazyDictionary[k] = append(lazyDictionary[k], w)
	}
}
