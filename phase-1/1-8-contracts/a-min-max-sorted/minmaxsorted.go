package minmaxsorted

// MinMax returns the minimum and maximum values in nums.
// Contract: decide and document what happens for an empty slice (min, max, and ok).
func MinMax(nums []int) (min, max int, ok bool) {
	if len(nums) == 0 {
		return 0, 0, false
	}

	runningMin := nums[0]
	runningMax := nums[0]

	// for loop would skip the first itteration which would always reassign both values
	for _, v := range nums {
		if v <= runningMin {
			runningMin = v
		}
		if v >= runningMax {
			runningMax = v
		}
	}

	// this or we can return early

	return runningMin, runningMax, true

}

// IsSortedNonDecreasing reports whether nums is sorted in non-decreasing order
// (each element is greater than or equal to the one before it).
func IsSortedNonDecreasing(nums []int) bool {
	if len(nums) == 0 {
		return true
	}

	previousValue := nums[0]

	for i := 1; i < len(nums); i++ {
		if previousValue > nums[i] {
			return false
		}
		previousValue = nums[i]
	}

	return true

}

// IndexOfMax returns the index of the maximum value in nums.
// When the maximum appears more than once, it returns the first occurrence.
// Contract: decide and document what happens for an empty slice (index and ok).
func IndexOfMax(nums []int) (int, bool) {
	if len(nums) == 0 {
		return 0, false
	}

	runningMaxIndex := 0
	runningMaxValue := nums[0]

	for i := 1; i < len(nums); i++ {

		if nums[i] > runningMaxValue {
			runningMaxIndex = i
			runningMaxValue = nums[i]

		}

	}

	return runningMaxIndex, true
}
