package question02

func computeArea(heights []int, i, j int) int {
	var minHeight int
	if heights[i] < heights[j] {
		minHeight = heights[i]
	} else {
		minHeight = heights[j]
	}
	return minHeight * (j - i)
}

// MaxWaterContainerBruteForce computes the area of the maximum water container using a
// brute-force algorithm.
// Time complexity is O(n^2), space complexity is O(1).
func MaxWaterContainerBruteForce(heights []int) int {
	maxArea := 0
	for i := range heights {
		for j := i + 1; j < len(heights); j++ {
			area := computeArea(heights, i, j)
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

// MaxWaterContainer computes the area of the maximum water container using an optimized algorithm
// (two-pointers technique).
// Time complexity is O(n), space complexity is O(1).
func MaxWaterContainer(heights []int) int {
	i, j := 0, len(heights)-1
	maxArea := 0

	for i < j {
		area := computeArea(heights, i, j)
		if area > maxArea {
			maxArea = area
		}
		// Shrink the window by moving the pointer to the minimum element.
		if heights[i] < heights[j] {
			i++
		} else {
			j--
		}
	}
	return maxArea
}
