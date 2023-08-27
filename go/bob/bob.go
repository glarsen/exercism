// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"regexp"
	"strings"
)

var (
	// clean is used to remove everything but letters and ?!
	clean = regexp.MustCompile(`[^a-zA-Z?!]+`)
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	empty := strings.TrimSpace(remark) == ""

	r := clean.ReplaceAllString(remark, "")

	yelling := strings.ToUpper(r) == r && strings.ToLower(r) != r
	question := strings.HasSuffix(r, "?")

	switch {
	case !yelling && question:
		return "Sure."
	case yelling && question:
		return "Calm down, I know what I'm doing!"
	case yelling:
		return "Whoa, chill out!"
	case empty:
		return "Fine. Be that way!"
	default:
		return "Whatever."
	}
}
