package anagram

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

// Expensive init function that loads the word dictionary
func init() {
	f, err := os.Open("dictionary.txt")
	if err != nil {
		panic("dictionary.txt missing")
	}
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

var dictionary []Word

// Word is a single word
type Word string

// Anagrams of word
func (w Word) Anagrams() []Word {
	if words, ok := Dictionary(w.Occurences()); ok {
		return words
	}
	return nil
}

// Sentence is a collection of words
type Sentence []Word

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
	// sub could be converted to a map and used as a lookup
	// Current cost is ~N*M
	// Cost with subtract map could be N+M
	acc := make(Occurences, len(o))
	copy(acc, o)
	for _, v := range sub {
		for i := range acc {
			if acc[i].Char != v.Char {
				continue
			}
			if frq := acc[i].Freq - v.Freq; frq > 0 {
				acc[i].Freq = frq
				break
			}
			acc = append(acc[:i], acc[i+1:]...)
		}
	}
	return acc
}

// Combinations returns all permutations of the occurences
func (o Occurences) Combinations() []Occurences {
	acc := make([]Occurences, 0, len(o)*len(o))
	for j, occ := range o {
		for i := 0; i < occ.Freq; i++ {
			_ = o[j+1:].Combinations() // next
			if i == 0 {
				// append next
				continue
			}
			// append (char, i) :: next

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

// Occurences is the count of each character in the sentence
func (s Sentence) Occurences() Occurences {
	return s.Word().Occurences()
}

// Word converts the Sentence to a Word by concatenating with no separator
func (s Sentence) Word() Word {
	str := make([]string, len(s))
	for i := range s {
		str[i] = string(s[i])
	}
	return Word(strings.Join(str, ""))
}

var lazyDictionary map[string][]Word

// Dictionary for the occurences
func Dictionary(o Occurences) ([]Word, bool) {
	if len(lazyDictionary) == 0 {
		loadDictionary()
	}
	v, ok := lazyDictionary[o.key()]
	return v, ok
}

func loadDictionary() {
	lazyDictionary = make(map[string][]Word)
	for i := range dictionary {
		w := Word(dictionary[i])
		k := w.Occurences().key()
		lazyDictionary[k] = append(lazyDictionary[k], w)
	}
}
