package parsinglogfiles

import (
	"fmt"
	"regexp"
)

var (
	line      = regexp.MustCompile(`^(\[TRC]|\[DBG]|\[INF]|\[WRN]|\[ERR]|\[FTL])`)
	password  = regexp.MustCompile(`(?i)".*password.*"`)
	remove    = regexp.MustCompile(`end-of-line[0-9]+`)
	separator = regexp.MustCompile(`<[~*=-]*>`)
	username  = regexp.MustCompile(`User\s+(\w+)\s+`)
)

func IsValidLine(text string) bool {
	return line.MatchString(text)
}

func SplitLogLine(text string) []string {
	return separator.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	count := 0
	for _, l := range lines {
		if password.MatchString(l) {
			count++
		}
	}

	return count
}

func RemoveEndOfLineText(text string) string {
	return remove.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	tagged := make([]string, 0, len(lines))

	for _, l := range lines {
		captures := username.FindStringSubmatch(l)

		if len(captures) > 0 {
			prefix := fmt.Sprintf("[USR] %s ", captures[1])

			l = prefix + l
		}

		tagged = append(tagged, l)
	}

	return tagged
}
