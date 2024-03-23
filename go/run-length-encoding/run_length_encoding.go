package encode

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func RunLengthEncode(input string) string {
	// Append a stop token
	input += "!"

	var b strings.Builder

	var active rune = 0x00
	var count = 0

	for _, letter := range input {
		if active == 0x00 {
			active = letter
			count = 1
			continue
		}

		if letter == active {
			count++
		} else {
			if count == 1 {
				_, _ = fmt.Fprintf(&b, "%c", active)
			} else {
				_, _ = fmt.Fprintf(&b, "%d%c", count, active)
			}
			count = 1
		}

		active = letter
	}

	return b.String()
}

func RunLengthDecode(input string) string {
	var b strings.Builder

	for left := 0; left < len(input); {
		right := left

		// Preserve whitespace
		if unicode.IsSpace(rune(input[left])) {
			b.WriteByte(input[left])
			left++
			continue
		}

		var n strings.Builder
		// Look for numeric characters to coalesce into a character count
		for ; right < len(input); right++ {
			if unicode.IsDigit(rune(input[right])) {
				n.WriteByte(input[right])
			} else {
				break
			}
		}

		var count int
		if n.Len() == 0 {
			count = 1
		} else {
			count, _ = strconv.Atoi(n.String())
		}

		// Get the letter to print
		letter := rune(input[right])

		// Emit the letter one or more times
		for j := 0; j < count; j++ {
			_, _ = fmt.Fprintf(&b, "%c", letter)
		}

		// Sync the left cursor to consume the token
		left = right + 1
	}

	return b.String()
}
