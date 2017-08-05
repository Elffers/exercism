package secret

import (
	"strconv"
)

const testVersion = 2

var Power = map[int]string{
	0: "wink",
	1: "double blink",
	2: "close your eyes",
	3: "jump",
}

func Handshake(code uint) []string {
	handshake := []string{}
	bin := strconv.FormatInt(int64(code), 2)
	pow := 0

	for i := len(bin) - 1; i >= 0; i-- {
		if i == 4 {
			defer func() {
				handshake = Reverse(handshake)
			}()
		}

		if bin[i] == '1' && pow < 4 {
			handshake = append(handshake, Power[pow])
		}

		pow++
	}
	return handshake
}

func Reverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
