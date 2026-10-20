package aliasing

// This package is demonstration-only. There is nothing to implement.
// Predict the test outcomes orally, then run the tests to confirm.

// bumpFirst sets nums[0] to n when nums is non-empty.
func bumpFirst(nums []int, n int) {
	if len(nums) > 0 {
		nums[0] = n
	}
}
