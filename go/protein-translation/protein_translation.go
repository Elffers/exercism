package protein

const testVersion = 1

var proteins = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

func FromCodon(in string) string {
	protein, ok := proteins[in]
	if ok {
		return protein
	}
	return "not found"
}

func FromRNA(in string) (proteins []string) {
	for i := 0; i < len(in); i += 3 {
		codon := in[i : i+3]
		protein := FromCodon(codon)
		if protein == "STOP" {
			return proteins
		}
		proteins = append(proteins, protein)
	}
	return
}
