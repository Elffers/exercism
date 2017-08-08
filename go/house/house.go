package house

import "fmt"

const testVersion = 1

var members = []struct {
	subj string
	verb string
}{
	{"house that Jack built.", ""},
	{"malt", "lay in"},
	{"rat", "ate"},
	{"cat", "killed"},
	{"dog", "worried"},
	{"cow with the crumpled horn", "tossed"},
	{"maiden all forlorn", "milked"},
	{"man all tattered and torn", "kissed"},
	{"priest all shaven and shorn", "married"},
	{"rooster that crowed in the morn", "woke"},
	{"farmer sowing his corn", "kept"},
	{"horse and the hound and the horn", "belonged to"},
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

func Verse(n int) (verse string) {
	verse = "This is the "
	base := members[0].subj

	for i := n - 1; i > 0; i-- {
		m := members[i]
		verse += fmt.Sprintf("%v\nthat %v the ", m.subj, m.verb)
	}
	verse += base
	return
}
