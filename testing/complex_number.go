// Package testing provides a simple demo how testing works in go
package testing

type ComplexNumber struct {
	Real, Imaginary float64
}

func (c ComplexNumber) Add(other ComplexNumber) ComplexNumber {
	return ComplexNumber{c.Real + other.Real, c.Imaginary + other.Imaginary}
}

func (c ComplexNumber) Multiply(other ComplexNumber) ComplexNumber {
	return ComplexNumber{c.Real*other.Real - c.Imaginary*other.Imaginary, c.Real*other.Imaginary + c.Imaginary*other.Real}
}
