package strand

const testVersion = 3

func ToRNA(dna string) string {
	rna := ""
	for _, r := range dna {
		switch r {
		case 'C':
			rna += "G"
		case 'G':
			rna += "C"
		case 'A':
			rna += "U"
		case 'T':
			rna += "A"
		}
	}
	return rna
}
