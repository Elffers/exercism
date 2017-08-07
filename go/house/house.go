package house

import "fmt"

const testVersion = 1

var Members = []Member{
	{"the house that Jack built.", ""},
	{"the malt", "lay in"},
	{"the rat", "ate"},
	{"the cat", "killed"},
	{"the dog", "worried"},
	{"the cow with the crumpled horn", "tossed"},
	{"the maiden all forlorn", "milked"},
	{"the man all tattered and torn", "kissed"},
	{"the priest all shaven and shorn", "married"},
	{"the rooster that crowed in the morn", "woke"},
	{"the farmer sowing his corn", "kept"},
	{"the horse and the hound and the horn", "belonged to"},
}

type Member struct {
	subj string
	verb string
}

func Song() string {
	var song string

	for i := 1; i < 13; i++ {
		song += Verse(i)
		if i != 12 {
			song += "\n\n"
		}
	}
	return song
}

func Verse(n int) string {
	var verse string
	base := Members[0].subj

	for i := n - 1; i > 0; i-- {
		m := Members[i]
		verse += fmt.Sprintf("%v\nthat %v", m.subj, m.verb)
		verse += " "

	}
	verse += base
	return fmt.Sprintf("This is %v", verse)
}
