package beer

import (
	"fmt"
	"strings"
)

const testVersion = 1

func Verse(count int) (string, error) {
	if count < 0 || count > 99 {
		return "", fmt.Errorf("Out of range")
	}
	line2 := fmt.Sprintf("Take one down and pass it around, %v of beer on the wall.\n", bottles(count-1))
	if count == 0 {
		line2 = "Go to the store and buy some more, 99 bottles of beer on the wall.\n"
	}

	if count == 1 {
		line2 = fmt.Sprintf("Take it down and pass it around, %v of beer on the wall.\n", strings.ToLower(bottles(count-1)))
	}

	line1 := fmt.Sprintf("%v of beer on the wall, %v of beer.\n", bottles(count), strings.ToLower(bottles(count)))
	out := line1 + line2
	return out, nil
}

func bottles(count int) string {
	unit := "bottles"
	if count == 1 {
		unit = "bottle"
	}

	if count == 0 {
		return fmt.Sprintf("No more %v", unit) // take into account capitalization
	}

	return fmt.Sprintf("%v %v", count, unit)
}

func Verses(max, min int) (string, error) {
	if max < min {
		return "", fmt.Errorf("Bad range")
	}
	out := ""
	// size := max - min + 1
	// out := make([]string, size)
	for n := max; n >= min; n-- {
		verse, err := Verse(n)
		if err != nil {
			return "", err
		}
		out += verse
		out += "\n"
	}
	return out, nil
}

func Song() string {
	song, _ := Verses(99, 0)
	return song
}
