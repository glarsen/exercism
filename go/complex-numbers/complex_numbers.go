package complexnumbers

import "math"

// Number is an imaginary number.
type Number struct {
	a float64
	b float64
}

func (n Number) Real() float64 {
	return n.a
}

func (n Number) Imaginary() float64 {
	return n.b
}

func (n1 Number) Add(n2 Number) Number {
	return Number{a: n1.a + n2.a, b: n1.b + n2.b}
}

func (n1 Number) Subtract(n2 Number) Number {
	return Number{a: n1.a - n2.a, b: n1.b - n2.b}
}

func (n1 Number) Multiply(n2 Number) Number {
	return Number{
		a: n1.a*n2.a - n1.b*n2.b,
		b: n1.b*n2.a + n1.a*n2.b,
	}
}

func (n Number) Times(factor float64) Number {
	return n.Multiply(Number{a: factor, b: 0})
}

func (n1 Number) Divide(n2 Number) Number {
	divisor := n2.a*n2.a + n2.b*n2.b
	return Number{
		a: (n1.a*n2.a + n1.b*n2.b) / divisor,
		b: (n1.b*n2.a - n1.a*n2.b) / divisor,
	}
}

func (n Number) Conjugate() Number {
	return Number{a: n.a, b: -n.b}
}

func (n Number) Abs() float64 {
	return math.Sqrt(n.a*n.a + n.b*n.b)
}

func (n Number) Exp() Number {
	factor := math.Exp(n.a)
	return Number{
		a: factor * math.Cos(n.b),
		b: factor * math.Sin(n.b),
	}
}
