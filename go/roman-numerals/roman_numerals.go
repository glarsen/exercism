package romannumerals

import (
	"fmt"
)

func valid(number int) error {
	if number <= 0 {
		return fmt.Errorf("%d cannot be represented", number)
	}

	if number > 3999 {
		return fmt.Errorf("%d is too large", number)
	}

	return nil
}

func ToRomanNumeral(input int) (string, error) {
	if err := valid(input); err != nil {
		return "", err
	}

	var numerals string
	for input > 0 {
		switch {
		case input >= 1000:
			numerals += "M"
			input -= 1000
		case input >= 900:
			numerals += "CM"
			input -= 900
		case input >= 500:
			numerals += "D"
			input -= 500
		case input >= 400:
			numerals += "CD"
			input -= 400
		case input >= 100:
			numerals += "C"
			input -= 100
		case input >= 90:
			numerals += "XC"
			input -= 90
		case input >= 50:
			numerals += "L"
			input -= 50
		case input >= 40:
			numerals += "XL"
			input -= 40
		case input >= 10:
			numerals += "X"
			input -= 10
		case input >= 9:
			numerals += "IX"
			input -= 9
		case input >= 5:
			numerals += "V"
			input -= 5
		case input >= 4:
			numerals += "IV"
			input -= 4
		case input >= 1:
			numerals += "I"
			input -= 1
		default:
			break
		}
	}

	return numerals, nil
}
