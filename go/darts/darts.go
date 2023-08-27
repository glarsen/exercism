package darts

import "math"

func Score(x, y float64) (points int) {
	radius := math.Hypot(x, y)

	switch {
	case radius <= 1.0:
		points = 10
	case radius <= 5.0:
		points = 5
	case radius <= 10:
		points = 1
	default:
		points = 0
	}

	return
}
