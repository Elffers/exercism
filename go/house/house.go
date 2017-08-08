package house

import (
	"fmt"
	"strings"
)

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
	song := make([]string, len(members))
	for i := 1; i < 13; i++ {
		song[i-1] = Verse(i)
	}
	return strings.Join(song, "\n\n")
}

func Verse(n int) (verse string) {
	verse = "This is the "
	for i := n - 1; i > 0; i-- {
		m := members[i]
		verse += fmt.Sprintf("%v\nthat %v the ", m.subj, m.verb)
	}
	verse += members[0].subj
	return
}
