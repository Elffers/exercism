package twelve

import "fmt"

const testVersion = 1

var gifts = []struct {
	day  string
	gift string
}{
	{"first", "a Partridge in a Pear Tree."},
	{"second", "two Turtle Doves"},
	{"third", "three French Hens"},
	{"fourth", "four Calling Birds"},
	{"fifth", "five Gold Rings"},
	{"sixth", "six Geese-a-Laying"},
	{"seventh", "seven Swans-a-Swimming"},
	{"eighth", "eight Maids-a-Milking"},
	{"ninth", "nine Ladies Dancing"},
	{"tenth", "ten Lords-a-Leaping"},
	{"eleventh", "eleven Pipers Piping"},
	{"twelfth", "twelve Drummers Drumming"},
}

func Song() string {
	song := ""
	for i := range gifts {
		verse := Verse(i+1) + "\n"
		song += verse
	}
	return song
}

func Verse(n int) string {
	gift := gifts[n-1]
	list := ""

	for i := n - 1; i >= 0; i-- {
		g := gifts[i].gift
		if n > 1 && i == 0 {
			g = "and " + g
		} else if n > 1 {
			g = g + ", "
		}
		list += g
	}

	return fmt.Sprintf("On the %s day of Christmas my true love gave to me, %s", gift.day, list)
}
