# Anagram

This package is an Go implementation of the Scala forcomp exercise in the Martin Odersky functional programming course https://www.coursera.org/course/progfun

The Scala version: https://github.com/jamiecuthill/forcomp

Including the package will cause the entire dictionary to be loaded into memory.

## Package Usage

Anagrams of a single word
```go
eat := anagram.Word("eat")
eat.Anagrams() // [ate eat tea]

```

Anagrams of a sentence
```go
eatMe := anagram.Sentence{"eat", "me"}
eatMe.Anagrams() // [[me tea], [me ate], ...]
```

Specify a different dictionary file
```go
// Before first use of anagram package
anagram.Source("/tmp/dictionary.txt")
greeting := anagram.Sentence{"hello", "world"}
greeting.Anagrams() // [...]
```

## CLI Command Usage

To specify a different dictionary file

```bash
anagram -dict /path/to/dictionary.txt hello world
```
