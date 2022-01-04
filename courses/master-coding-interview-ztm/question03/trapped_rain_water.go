package question03

// TrappedRainWaterBruteForce implements a brute-force algorithm for calculating
// the quantity of trapped water.
// Time complexity is O(n^2), space complexity is O(1).
func TrappedRainWaterBruteForce(heights []int) int {
	totalWater := 0

	for i := range heights {
		maxLeft := 0
		maxRight := 0

		// Search the max height on the left side of the current element.
		for j := 0; j < i; j++ {
			if heights[j] > maxLeft {
				maxLeft = heights[j]
			}
		}

		// Search the max height on the right side of the current element.
		for j := i + 1; j < len(heights); j++ {
			if heights[j] > maxRight {
				maxRight = heights[j]
			}
		}

		var waterLevel int
		if maxLeft < maxRight {
			waterLevel = maxLeft
		} else {
			waterLevel = maxRight
		}

		if waterLevel > heights[i] {
			totalWater += waterLevel - heights[i]
		}
	}

	return totalWater
}

// TrappedRainWater efficiently solves the quantity of trapped water using the two pointers technique.
func TrappedRainWater(heights []int) int {
	// Keeps the maximum heights seen by the left and right pointers.
	maxLeft, maxRight := 0, 0
	totalWater := 0

	i, j := 0, len(heights)-1

	for i < j {
		if heights[i] < heights[j] {
			if heights[i] < maxLeft {
				totalWater += maxLeft - heights[i]
			} else {
				maxLeft = heights[i]
			}
			i++
		} else {
			if heights[j] < maxRight {
				totalWater += maxRight - heights[j]
			} else {
				maxRight = heights[j]
			}
			j--
		}
	}

	return totalWater
}
