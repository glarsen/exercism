package allergies

var allergens = map[string]uint{
	"eggs":         1,
	"peanuts":      2,
	"shellfish":    4,
	"strawberries": 8,
	"tomatoes":     16,
	"chocolate":    32,
	"pollen":       64,
	"cats":         128,
}

func Allergies(allergies uint) []string {
	var matches []string

	for allergen := range allergens {
		if AllergicTo(allergies, allergen) {
			matches = append(matches, allergen)
		}
	}

	return matches
}

func AllergicTo(allergies uint, allergen string) bool {
	if value, ok := allergens[allergen]; ok {
		return (value & allergies) != 0
	}

	return false
}
