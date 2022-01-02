package question01

import (
	"testing"
)

type FindTwoSumFunc func(nums []int, targetSum int) []int

func testFindTwoSum(t *testing.T, findTwoSumFunc FindTwoSumFunc) {
	t.Helper()

	type test struct {
		nums   []int
		target int
		result []int
	}

	nums := []int{1, 3, 7, 9, 2}

	tests := []test{
		{nums, 11, []int{3, 4}},
		{nums, 25, nil},
		{[]int{}, 1, nil},
		{[]int{1}, 1, nil},
	}

	for _, tc := range tests {
		result := findTwoSumFunc(tc.nums, tc.target)
		if len(result) == 0 && len(tc.result) == 0 {
			continue
		}
		if len(result) != 0 && (result[0] != tc.result[0] || result[1] != tc.result[1]) {
			t.Fatalf("Invalid computed pair (%d, %d), expected (%d, %d)",
				result[0], result[1], tc.result[0], tc.result[1])
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
