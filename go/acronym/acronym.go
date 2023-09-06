// Package acronym provides abbreviation services.
package acronym

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// Abbreviate returns an acronym for the provided string.
func Abbreviate(s string) string {
	words := strings.FieldsFunc(s, func(r rune) bool {
		return r == ' ' || r == '-' || r == '_'
	})

	var acronym strings.Builder

	for _, word := range words {
		first, _ := utf8.DecodeRuneInString(word)
		acronym.WriteRune(unicode.ToUpper(first))
	}

	return acronym.String()
}
