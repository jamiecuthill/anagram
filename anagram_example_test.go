package anagram_test

import (
	"fmt"

	"github.com/jamiecuthill/anagram"
)

func ExampleSentence_Anagrams() {
	sentence := anagram.Sentence{"eat", "me"}

	fmt.Println(sentence.Anagrams())
	// Output: [[em ate] [em eat] [em tea] [me ate] [me eat] [me tea] [Mae et] [et Mae] [ate em] [ate me] [eat em] [eat me] [tea em] [tea me]]
}

func ExampleWord_Anagrams() {
	w := anagram.Word("eat")
	fmt.Println(w.Anagrams())
	// Output: [ate eat tea]
}
