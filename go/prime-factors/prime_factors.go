package prime

func Factors(n int64) []int64 {
	var factors []int64

	for f := int64(2); f <= n; f++ {
		for n%f == 0 {
			factors = append(factors, f)
			n /= f
		}
	}

	return factors
}
