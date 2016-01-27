package anagram

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"sync"
)

var defaltDictionary = "dictionary.txt"

var dictionary []Word

type index struct {
	once sync.Once
	m    map[string][]Word
}

var dictIndex index

func (i *index) load() *index {
	i.once.Do(func() {
		if len(dictionary) == 0 {
			Source(defaltDictionary)
		}
		i.m = make(map[string][]Word)
		defer func() {
			dictionary = []Word(nil)
		}()
		for _, word := range dictionary {
			key := word.Occurences().key()
			i.m[key] = append(i.m[key], word)
		}
	})
	return i
}

// Source loads the word dictionary
func Source(dict string) {
	f, err := os.Open(dict)
	if err != nil {
		panic("dictionary.txt missing")
	}
	defer f.Close()
	buf := bufio.NewReader(f)
	for {
		line, isPrefix, err := buf.ReadLine()
		if err != nil {
			return
		}
		// Our dictionary words are shorter than the defaultBufSize so isPrefix
		// should always be false for the current state.
		if isPrefix {
			panic("insufficient buffer for line: " + string(line))
		}
		dictionary = append(dictionary, Word(line))
	}
}

// Word is a single word
type Word string

// Anagrams of word
func (w Word) Anagrams() []Word {
	return Dictionary(w.Occurences())
}

// lower case the word
func (w Word) lower() Word {
	return Word(strings.ToLower(string(w)))
}

// Occurences is the count of each character in the word
func (w Word) Occurences() Occurences {
	freqs := make(map[rune]int)
	for _, r := range w.lower() {
		freqs[r]++
	}

	occurences := make(Occurences, 0, len(freqs))
	for char, freq := range freqs {
		occurences = append(occurences, occurence{char, freq})
	}
	sort.Sort(occurences)
	return occurences
}

// Sentence is a collection of words
type Sentence []Word

// word converts the Sentence to a Word by concatenating with no separator
func (s Sentence) word() Word {
	var w Word
	for _, word := range s {
		w += word
	}
	return w
}

// Occurences is the count of each character in the sentence
func (s Sentence) Occurences() Occurences {
	return s.word().Occurences()
}

// Anagrams of the sentence
func (s Sentence) Anagrams() []Sentence {
	return anagrams(s.Occurences())
}

func (s Sentence) String() string {
	var str []string
	for _, w := range s {
		str = append(str, string(w))
	}
	return strings.Join(str, " ")
}

func anagrams(occurences Occurences) []Sentence {
	anags := make([]Sentence, 0, 1)
	if len(occurences) == 0 {
		// An empty sentence must be returned for occurences of length 0
		anags = append(anags, Sentence{})
		return anags
	}

	for _, occurence := range occurences.Combinations() {
		for _, word := range Dictionary(occurence) {
			sentences := []Sentence{}
			for _, tail := range anagrams(occurences.Subtract(occurence)) {
				sentences = append(sentences, append(Sentence{word}, tail...))
			}
			anags = append(anags, sentences...)
		}
	}
	return anags
}

type occurence struct {
	Char rune
	Freq int
}

func (o occurence) String() string {
	return fmt.Sprintf("%c:%d", o.Char, o.Freq)
}

// Occurences is a list of character and frequency, it must be sorted alphabetically
// therefore it can not be a map
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
	return fmt.Sprintf("%v", o)
}

// Subtract removes the occurences and returns the result
func (o Occurences) Subtract(sub Occurences) Occurences {
	amount := make(map[rune]int)
	for _, occ := range sub {
		amount[occ.Char] = occ.Freq
	}

	acc := make(Occurences, 0, len(o))
	for _, occ := range o {
		if n, ok := amount[occ.Char]; ok {
			if frq := occ.Freq - n; frq > 0 {
				acc = append(acc, occurence{occ.Char, frq})
			}
			continue
		}
		acc = append(acc, occ)
	}
	return acc
}

// Combinations returns all permutations of the occurences
func (o Occurences) Combinations() []Occurences {
	combs := make([]Occurences, 0, 1)
	if len(o) == 0 {
		combs = append(combs, Occurences{})
		return combs
	}

	head, tail := o[0], o[1:]
	for _, next := range tail.Combinations() {
		for i := 0; i <= head.Freq; i++ {
			if i == 0 {
				combs = append(combs, next)
				continue
			}
			combs = append(combs, append(Occurences{{head.Char, i}}, next...))
		}
	}
	return combs
}

// Dictionary of words that match the given occurence
// Example Dictionary(Occurences{{a, 1}, {e, 1}, {t, 1}}) = [ate eat tea]
func Dictionary(o Occurences) []Word {
	if v, ok := dictIndex.load().m[o.key()]; ok {
		return v
	}
	return []Word(nil)
}
