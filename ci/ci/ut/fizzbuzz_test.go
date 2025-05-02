package fizzbuzz

import (
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	cases := []struct {
		name     string
		input    []int
		expected []string
	}{
		{
			name:     "normal case",
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			expected: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"},
		},
		{
			name:     "zero or negative numbers",
			input:    []int{0, -1, -2, -3, -4, -5},
			expected: []string{"0", "-1", "-2", "-3", "-4", "-5"},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := FizzBuzz(c.input)
			if len(got) != len(c.expected) {
				t.Errorf("length mismatch: got %v, want %v", got, c.expected)
				return
			}
			for i := range got {
				if got[i] != c.expected[i] {
					t.Errorf("at index %d: got %v, want %v", i, got[i], c.expected[i])
				}
			}
		})
	}
}
