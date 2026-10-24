package countinrange

// CountInRange returns how many elements of nums fall in the closed interval [lo, hi].
// If lo > hi, it returns 0. A nil or empty nums returns 0. The input is not modified.
func CountInRange(nums []int, lo, hi int) int {
	if lo > hi {
		return 0
	}

	runningCount := 0

	for _, v := range nums {
		if v >= lo && v <= hi {
			runningCount++
		}
	}

	return runningCount
}
