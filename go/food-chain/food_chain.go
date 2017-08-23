package foodchain

import (
	"fmt"
)

const testVersion = 3

var foodChain = []struct {
	food    string
	comment string
}{
	{"", ""},
	{"fly", "I don't know why she swallowed the fly. Perhaps she'll die."},
	{"spider", "It wriggled and jiggled and tickled inside her."},
	{"bird", "How absurd to swallow a bird!"},
	{"cat", "Imagine that, to swallow a cat!"},
	{"dog", "What a hog, to swallow a dog!"},
	{"goat", "Just opened her throat and swallowed a goat!"},
	{"cow", "I don't know how she swallowed a cow!"},
	{"horse", "She's dead, of course!"},
}

func Song() string {
	return Verses(1, 8)
}

func Verse(n int) (verse string) {
	current := foodChain[n]
	verse += fmt.Sprintf("I know an old lady who swallowed a %v.\n%v", current.food, current.comment)
	if n == 1 || n == 8 {
		return verse
	}

	for i := n - 1; i > 0; i-- {
		next := foodChain[i]
		nextFood := next.food
		if i == 2 {
			nextFood += " that wriggled and jiggled and tickled inside her"
		}
		verse += fmt.Sprintf("\nShe swallowed the %v to catch the %v.", current.food, nextFood)
		current = next
	}

	verse += "\n" + foodChain[1].comment
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
