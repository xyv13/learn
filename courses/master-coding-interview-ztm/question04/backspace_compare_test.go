package question04

import (
	"testing"
)

type backspaceCompareFunc func(s, t string) bool

func testBackspaceCompare(t *testing.T, backspaceCompare backspaceCompareFunc) {
	t.Helper()

	type test struct {
		s, t   string
		result bool
	}

	tests := []test{
		/*
			{"ab#z", "az#z", true},
			{"abc#d", "acc#c", false},
			{"x#y#z#", "a#", true},
			{"a###b", "b", true},
			{"Ab#z", "ab#z", false},
			{"gtc#uz#", "gtcm##uz#", true},
			{"bxj##tw", "bxj###tw", false},
		*/
		{"bbbextm", "bbb#extm", false},
	}

	for _, tc := range tests {
		result := backspaceCompare(tc.s, tc.t)
		if result != tc.result {
			t.Fatalf("Expected backspace compare result for strings %s and %s to be %t, got %t",
				tc.s, tc.t, tc.result, result)
		}
	}
}

func TestBackspaceCompare(t *testing.T) {
	t.Run("Testing backspace comparing strings using the brute force algorithm", func(t *testing.T) {
		testBackspaceCompare(t, BackspaceCompareBruteForce)
	})
	t.Run("Testing backspace comparing strings using the optimized algorithm", func(t *testing.T) {
		testBackspaceCompare(t, BackspaceCompare)
	})
}
