package main

import "testing"

func TestSubtract(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{5, 3, 2},
		{10, 5, 5},
		{0, 0, 0},
		{-1, -1, 0},
	}

	for _, test := range tests {
		result := Subtract(test.a, test.b)
		if result != test.expected {
			t.Errorf("Subtract(%d, %d) = %d; expected %d", test.a, test.b, result, test.expected)
		}
	}
}
