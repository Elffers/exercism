package strand

import "strings"

const testVersion = 3

func ToRNA(dna string) string {
	r := strings.NewReplacer(
		"C", "G",
		"G", "C",
		"A", "U",
		"T", "A",
	)
	return r.Replace(dna)
}
