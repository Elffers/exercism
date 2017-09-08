package allergies

const testVersion = 1

type Allergen int

const (
	Eggs Allergen = 1 << iota
	Peanuts
	Shellfish
	Strawberries
	Tomatoes
	Chocolate
	Pollen
	Cats
)

var allergens = map[string]Allergen{
	"eggs":         Eggs,
	"peanuts":      Peanuts,
	"shellfish":    Shellfish,
	"strawberries": Strawberries,
	"tomatoes":     Tomatoes,
	"chocolate":    Chocolate,
	"pollen":       Pollen,
	"cats":         Cats,
}

func AllergicTo(score uint, a string) bool {
	allergen := allergens[a]
	return score&uint(allergen) != 0
}

func Allergies(score uint) []string {
	allergies := []string{}
	for k, v := range allergens {
		bit := score & uint(v)
		if bit != 0 {
			allergies = append(allergies, k)
		}
	}
	return allergies
}
