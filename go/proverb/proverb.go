// Package proverb emits proverbs.
package proverb

import "fmt"

const (
	stanza = "For want of a %s the %s was lost."
	last   = "And all for the want of a %s."
)

// Proverb creates a proverb from the provided rhyme.
func Proverb(rhyme []string) []string {
	if len(rhyme) == 0 {
		return nil
	}

	var sentences []string
	for i := 0; i+1 < len(rhyme); i++ {
		sentences = append(sentences, fmt.Sprintf(stanza, rhyme[i], rhyme[i+1]))
	}
	sentences = append(sentences, fmt.Sprintf(last, rhyme[0]))

	return sentences
}
