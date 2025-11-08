package resistorcolortrio

import (
	"fmt"
)

func value(colors []string) int {
	lookup := map[string]int{
		"black":  0,
		"brown":  1,
		"red":    2,
		"orange": 3,
		"yellow": 4,
		"green":  5,
		"blue":   6,
		"violet": 7,
		"grey":   8,
		"white":  9,
	}

	first, second, third := colors[0], colors[1], colors[2]

	v := lookup[first]*10 + lookup[second]
	zeroes := lookup[third]

	for range zeroes {
		v *= 10
	}

	return v
}

// Label describes the resistance value given the colors of a resistor.
// The label is a string with a resistance value with an unit appended
// (e.g. "33 ohms", "470 kiloohms").
func Label(colors []string) string {
	resistance := value(colors)

	var unit string
	var scale int

	switch {
	case resistance < 1_000:
		unit = ""
		scale = 1
	case resistance < 1_000_000:
		unit = "kilo"
		scale = 1_000
	case resistance < 1_000_000_000:
		unit = "mega"
		scale = 1_000_000
	default:
		unit = "giga"
		scale = 1_000_000_000
	}

	return fmt.Sprintf("%d %s%s", resistance/scale, unit, "ohms")
}
