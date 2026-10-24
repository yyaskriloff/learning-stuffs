package anypairsum

// AnyPairSum reports whether two distinct indexes i and j satisfy nums[i]+nums[j] == target.
// Nil, empty, and single-element inputs return false. The input is not modified.
func AnyPairSum(nums []int, target int) bool {

	for i, iv := range nums {
		for j, jv := range nums {
			if i == j {
				continue
			}

			if iv+jv == target {
				return true
			}

		}
	}

	return false
}
