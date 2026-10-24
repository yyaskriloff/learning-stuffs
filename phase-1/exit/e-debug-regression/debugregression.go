package debugregression

// CumulativeMax returns a new slice where out[i] is the maximum of nums[0]..nums[i].
// A nil or empty input returns a slice with length 0. The input is not modified.
func CumulativeMax(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	out := make([]int, len(nums))
	maxSoFar := nums[0]
	for i, v := range nums {
		if v > maxSoFar {
			maxSoFar = v
		}
		out[i] = maxSoFar
	}
	return out
}
