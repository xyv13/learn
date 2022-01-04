package question03

import (
	"testing"
)

type trappedRainWaterFunc func(heights []int) int

func testTrappedRainWater(t *testing.T, trappedRainWater trappedRainWaterFunc) {
	t.Helper()

	type test struct {
		heights  []int
		maxWater int
	}

	tests := []test{
		{[]int{0, 1, 0, 2, 1, 0, 3, 1, 0, 1, 2}, 8},
		{[]int{}, 0},
		{[]int{1}, 0},
	}

	for _, tc := range tests {
		maxWater := trappedRainWater(tc.heights)
		if maxWater != tc.maxWater {
			t.Fatalf("Expected max water to be %d, got %d", tc.maxWater, maxWater)
		}
	}
}

func TestTrappedRainWater(t *testing.T) {
	t.Run("Testing finding max trapped water using the brute force algorithm", func(t *testing.T) {
		testTrappedRainWater(t, TrappedRainWaterBruteForce)
	})
	t.Run("Testing finding max trapped water using the optimized algorithm", func(t *testing.T) {
		testTrappedRainWater(t, TrappedRainWater)
	})
}
