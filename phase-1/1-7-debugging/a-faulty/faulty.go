package faulty

// RunningSum returns a new slice where out[i] is the sum of nums[0] through nums[i].
// A nil or empty input returns a slice with length 0. The input is not modified.
func RunningSum(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	out := make([]int, len(nums))

	workingSum := 0

	for i, v := range nums {
		workingSum += v
		out[i] = workingSum
	}

	return out
}
