package reversefilter

import "slices"

// Reverse reverses nums in place.
func Reverse(nums []int) {
	slices.Reverse(nums)
}

// Filter returns a new slice of elements from nums for which keep returns true,
// in the same order. It must not modify nums.
func Filter(nums []int, keep func(int) bool) []int {
	newSlice := []int{}

	for _, num := range nums {

		if keep(num) {
			newSlice = append(newSlice, num)
		}
	}
	return newSlice
}
