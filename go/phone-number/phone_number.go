package phonenumber

import (
	"fmt"
	"regexp"
	"unicode"
)

const testVersion = 3

func Number(in string) (string, error) {
	out := ""
	for _, num := range in {
		if unicode.IsDigit(num) {
			out += string(num)
		}
	}

	length := len(out)

	if length < 10 {
		return "", fmt.Errorf("Less than 10 digits")
	}
	if length > 11 {
		return "", fmt.Errorf("More than 11 digits")
	}

	if length == 11 {
		cc, number := out[:1], out[1:]
		if cc != "1" {
			return "", fmt.Errorf("bad country code")
		}
		out = number
	}

	validStart := regexp.MustCompile(`[2-9]`)

	if !validStart.MatchString(string(out[0])) || !validStart.MatchString(string(out[3])) {
		return "", fmt.Errorf("bad area code or exchange code")
	}

	return out, nil
}

func AreaCode(in string) (string, error) {
	number, err := Number(in)
	if err != nil {
		return "", err
	}
	out := number[:3]

	return out, nil
}

func Format(in string) (string, error) {
	number, err := Number(in)
	if err != nil {
		return "", err
	}
	areaCode := number[:3]
	exchangeCode := number[3:6]
	subscriberNum := number[6:]
	return fmt.Sprintf("(%v) %v-%v", areaCode, exchangeCode, subscriberNum), nil
}
