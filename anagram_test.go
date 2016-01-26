package anagram_test

import (
	"reflect"
	"testing"

	. "github.com/jamiecuthill/anagrams"
)

func TestWordOccurences(t *testing.T) {
	tests := []struct {
		in  Word
		out Occurences
	}{
		{
			"abcd",
			Occurences{
				{toRune("a"), 1},
				{toRune("b"), 1},
				{toRune("c"), 1},
				{toRune("d"), 1},
			},
		},
		{
			"Robert",
			Occurences{
				{toRune("b"), 1},
				{toRune("e"), 1},
				{toRune("o"), 1},
				{toRune("r"), 2},
				{toRune("t"), 1},
			},
		},
	}
	for _, test := range tests {
		o := test.in.Occurences()
		if !reflect.DeepEqual(o, test.out) {
			t.Errorf("unexpected occurences %v", o)
		}
	}

}

func TestSentenceOccurences(t *testing.T) {
	o := Sentence{"abcd", "e"}.Occurences()
	if !reflect.DeepEqual(o, Occurences{
		{toRune("a"), 1},
		{toRune("b"), 1},
		{toRune("c"), 1},
		{toRune("d"), 1},
		{toRune("e"), 1},
	}) {
		t.Errorf("unepxected occurences %v", o)
	}
}

func TestDictionaryEat(t *testing.T) {
	words, ok := Dictionary(Occurences{
		{toRune("a"), 1},
		{toRune("e"), 1},
		{toRune("t"), 1},
	})
	if !ok {
		t.Fatal("expected occurences to exist in the dictionary")
	}
	if !reflect.DeepEqual(words, []Word{"ate", "eat", "tea"}) {
		t.Errorf("unexpected word set %v", words)
	}
}

func TestWordAnagrams(t *testing.T) {
	tests := []struct {
		in  Word
		out []Word
	}{
		{
			"married",
			[]Word{"admirer", "married"},
		},
		{
			"player",
			[]Word{"parley", "pearly", "player", "replay"},
		},
	}
	for _, test := range tests {
		anagrams := test.in.Anagrams()
		if !reflect.DeepEqual(anagrams, test.out) {
			t.Errorf("unexpected anagrams: %v", anagrams)
		}
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		in  Word
		sub Occurences
		out Occurences
	}{
		{
			"lard",
			Occurences{{toRune("r"), 1}},
			Occurences{
				{toRune("a"), 1},
				{toRune("d"), 1},
				{toRune("l"), 1},
			},
		},
		{
			"helloworld",
			Occurences{{toRune("l"), 1}},
			Occurences{
				{toRune("d"), 1},
				{toRune("e"), 1},
				{toRune("h"), 1},
				{toRune("l"), 2},
				{toRune("o"), 2},
				{toRune("r"), 1},
				{toRune("w"), 1},
			},
		},
	}
	for _, test := range tests {
		occ := test.in.Occurences()
		res := occ.Subtract(test.sub)
		if !reflect.DeepEqual(res, test.out) {
			t.Errorf("unexpected occurences: %v", res)
		}
	}
}

func TestCombinations(t *testing.T) {
	var w Word = "abba"
	expect := []Occurences{
		{},
		{{toRune("a"), 1}},
		{{toRune("a"), 2}},
		{{toRune("b"), 1}},
		{{toRune("a"), 1}, {toRune("b"), 1}},
		{{toRune("a"), 2}, {toRune("b"), 1}},
		{{toRune("b"), 2}},
		{{toRune("a"), 1}, {toRune("b"), 2}},
		{{toRune("a"), 2}, {toRune("b"), 2}},
	}
	comb := w.Occurences().Combinations()
	if !reflect.DeepEqual(comb, expect) {
		t.Errorf("unexpected combinations: %v", comb)
	}
}

func toRune(c string) rune {
	if len(c) == 0 {
		return 0
	}
	return []rune(c)[0]
}
