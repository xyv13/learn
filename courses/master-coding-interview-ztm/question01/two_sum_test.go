package question01

import (
	"testing"
)

type FindTwoSumFunc func(nums []int, targetSum int) (int, int)

func testFindTwoSum(t *testing.T, findTwoSumFunc FindTwoSumFunc) {
	t.Helper()

	type test struct {
		nums       []int
		target     int
		start, end int
	}

	nums := []int{1, 3, 7, 9, 2}

	tests := []test{
		{nums, 11, 3, 4},
		{nums, 25, -1, -1},
		{[]int{}, 1, -1, -1},
		{[]int{1}, 1, -1, -1},
	}

	for _, tc := range tests {
		start, end := findTwoSumFunc(tc.nums, tc.target)
		if start != tc.start || end != tc.end {
			t.Fatalf("Invalid computed pair (%d, %d), expected (%d, %d)",
				start, end, tc.start, tc.end)
		}
	}
}

func TestFindTwoSum(t *testing.T) {
	t.Run("Testing two sum brute force algorithm", func(t *testing.T) {
		testFindTwoSum(t, FindTwoSumBruteForce)
	})
	t.Run("Testing two sum algorithm", func(t *testing.T) {
		testFindTwoSum(t, FindTwoSum)
	})
}
