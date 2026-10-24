package faulty

// RunningSum returns a new slice where out[i] is the sum of nums[0] through nums[i].
// A nil or empty input returns a slice with length 0. The input is not modified.
func RunningSum(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	out := make([]int, len(nums))
	out[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		out[i] = out[i-1] + nums[i-1]
	}
	return out
}
