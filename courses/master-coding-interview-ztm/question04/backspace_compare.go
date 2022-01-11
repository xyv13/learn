package question04

func computeBackspacedStringBruteForce(str string) []rune {
	var stack []rune
	for _, c := range str {
		if c != '#' {
			// Push to stack.
			stack = append(stack, c)
		} else if len(stack) > 0 {
			// Pop from stack.
			stack = stack[:len(stack)-1]
		}
	}
	return stack
}

// BackspaceCompareBruteForce compares two strings after removing the typed-out characters
// using a brute-force algorithm.
// Time complexity is O(|s| + |t|), space complexity is O(|s| + |t|).
func BackspaceCompareBruteForce(s, t string) bool {
	stackS := computeBackspacedStringBruteForce(s)
	stackT := computeBackspacedStringBruteForce(t)

	if len(stackS) != len(stackT) {
		return false
	}

	for i := 0; i < len(stackS); i++ {
		if stackS[i] != stackT[i] {
			return false
		}
	}

	return true
}

// BackspaceCompareBruteForce compares two strings after removing the typed-out characters
// using an optimized algorithm.
// Time complexity is O(|s| + |t|), space complexity is O(1).
func BackspaceCompare(s, t string) bool {
	i := len(s) - 1
	j := len(t) - 1

	for i >= 0 || j >= 0 {
		// Skip the typed-out characters.
		if i >= 0 && s[i] == '#' {
			for skip := 2; skip > 0; skip-- {
				i--
				if i < 0 {
					break
				}
				if s[i] == '#' {
					skip += 2
				}
			}
		}

		if j >= 0 && t[j] == '#' {
			for skip := 2; skip > 0; skip-- {
				j--
				if j < 0 {
					break
				}
				if t[j] == '#' {
					skip += 2
				}
			}
		}

		// Empty strings are equal.
		if i == -1 && j == -1 {
			return true
		}

		// If one string become empty and the other not, or if characters don't match
		// then return false.
		if i^j < 0 || s[i] != t[j] {
			return false
		}

		i--
		j--
	}

	return true
}
