package question01

// FindTwoSumBruteForce will find a pair of numbers which sum up to 'target' and
// return their indices. If no such pair exists, returns (-1, -1).
// The time complexity is O(n^2), the space complexity is O(1).
func FindTwoSumBruteForce(nums []int, target int) []int {
	for i := range nums {
		numToFind := target - nums[i]
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == numToFind {
				return []int{i, j}
			}
		}
	}
	return nil
}

// FindTwoSum, compared to FindTwoSumBruteForce, uses a map to search the number pairs.
// The time complexity is O(n), the space complexity is O(n).
func FindTwoSum(nums []int, target int) []int {
	numsMap := make(map[int]int)
	for i := range nums {
		if pos, ok := numsMap[nums[i]]; ok {
			return []int{pos, i}
		}
		numToFind := target - nums[i]
		numsMap[numToFind] = i
	}
	return nil
}
