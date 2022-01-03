package question02

import (
	"testing"
)

type maxWaterContainerFunc func(heights []int) int

func testMaxWaterContainer(t *testing.T, maxWaterContainer maxWaterContainerFunc) {
	t.Helper()

	type test struct {
		heights []int
		maxArea int
	}

	tests := []test{
		{[]int{7, 1, 2, 3, 9}, 28},
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1}, 1},
		{[]int{}, 0},
		{[]int{1}, 0},
	}

	for _, tc := range tests {
		maxArea := maxWaterContainer(tc.heights)
		if maxArea != tc.maxArea {
			t.Fatalf("Expected max area to be %d, got %d", tc.maxArea, maxArea)
		}
	}
}

func TestMaxWaterContainer(t *testing.T) {
	t.Run("Testing finding max area container using the brute force algorithm", func(t *testing.T) {
		testMaxWaterContainer(t, MaxWaterContainerBruteForce)
	})
	t.Run("Testing finding max area container using the optimized algorithm", func(t *testing.T) {
		testMaxWaterContainer(t, MaxWaterContainer)
	})
}
