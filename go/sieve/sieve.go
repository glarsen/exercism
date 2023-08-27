package sieve

// Sieve of Eratosthenes
// Ref: https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes
func Sieve(limit int) []int {
	// let A be an array of Boolean values indexed by integers 2 to n, initially all set to true.
	A := make([]bool, limit+1)
	for i := 2; i <= limit; i++ {
		A[i] = true
	}

	// for i = 2, 3, 4, ..., not exceeding √n do
	for i := 2; i*i <= limit; i++ {
		// if A[i] is true
		if A[i] == true {
			// for j = i2, i2+i, i2+2i, i2+3i, ..., not exceeding n do
			for j := i * 2; j <= limit; j += i {
				// set A[j] := false
				A[j] = false
			}
		}
	}

	// return all i such that A[i] is true.
	var primes []int
	for i := 2; i <= limit; i++ {
		if A[i] == true {
			primes = append(primes, i)
		}
	}

	return primes
}
