package series

// All returns a list of all substrings of s with length n
func All(n int, s string) []string {
	var series []string
	for i := 0; i+n <= len(s); i++ {
		series = append(series, s[i:i+n])
	}
	return series
}

// UnsafeFirst returns the first substring of s with length n
func UnsafeFirst(n int, s string) string {
	return s[:n] // This can panic if n > len(s)
}
