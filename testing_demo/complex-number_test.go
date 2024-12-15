package testing_demo

import "testing"

func TestAdd(t *testing.T) {
	cone := ComplexNumber{1, 2}
	got := cone.Add(ComplexNumber{3, 4})
	want := ComplexNumber{4, 6}
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	multiplicand := ComplexNumber{1, 2}
	multiplier := ComplexNumber{3, 4}
	got = multiplicand.Multiply(multiplier)
	want = ComplexNumber{-5, 10}
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
