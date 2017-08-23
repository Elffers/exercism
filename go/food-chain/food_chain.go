package foodchain

import (
	"fmt"
)

const testVersion = 3

var foodChain = []struct {
	food    string
	comment string
}{
	{"fly", ""},
	{"spider", "It wriggled and jiggled and tickled inside her."},
	{"bird", "How absurd to swallow a bird!"},
	{"cat", "Imagine that, to swallow a cat!"},
	{"dog", "What a hog, to swallow a dog!"},
	{"goat", "Just opened her throat and swallowed a goat!"},
	{"cow", "I don't know how she swallowed a cow!"},
	{"horse", "She's dead, of course!"},
}

var coda = "I don't know why she swallowed the fly. Perhaps she'll die."

func Song() string {
	return Verses(1, 8)
}

func Verse(n int) (verse string) {
	current := foodChain[n-1]
	verse += fmt.Sprintf("I know an old lady who swallowed a %v.\n", current.food)
	if n == 1 {
		return verse + coda
	}

	if n == 8 {
		return verse + current.comment
	}

	verse += fmt.Sprintf("%v\n", current.comment)

	index := n - 2

	for index >= 0 {
		next := foodChain[index]
		nextFood := next.food
		if index == 1 {
			nextFood += " that wriggled and jiggled and tickled inside her"
		}
		verse += fmt.Sprintf("She swallowed the %v to catch the %v.\n", current.food, nextFood)
		index--
		current = next
	}

	verse += coda
	return verse
}

func Verses(start, end int) string {
	out := Verse(start)
	for start < end {
		start++
		out += "\n\n" + Verse(start)
	}
	return out
}
