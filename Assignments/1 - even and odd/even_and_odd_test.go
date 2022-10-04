package main

import (
	"testing"
)

func TestCalculateEvenOrOdd(t *testing.T) {
	tc := []struct {
		value  int
		result string
	}{{1, "odd"}, {2, "even"}, {3, "odd"}, {4, "even"}}

	for _, test := range tc {
		got := CalculateEvenOrOdd(test.value)
		want := test.result

		if got != want {
			t.Errorf("Expected value %v to be %v", test.value, want)
		}
	}

}
