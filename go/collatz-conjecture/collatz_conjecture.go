package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	steps := 0

	switch {
	case n == 0:
		return 0, errors.New("zero is not allowed")
	case n < 0:
		return 0, errors.New("negative numbers are not allowed")
	}

	for n > 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}

		steps++
	}

	return steps, nil
}
