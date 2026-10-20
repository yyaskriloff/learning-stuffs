package sumcounteven

// Sum returns the sum of all elements in nums.
func Sum(nums []int) int {
	total := 0

	if len(nums) == 0 || nums == nil {
		return total
	}

	for _, num := range nums {
		total += num
	}

	return total
}

// CountEven returns how many elements in nums are even.
func CountEven(nums []int) int {
	total := 0

	if len(nums) == 0 || nums == nil {
		return total
	}

	for _, num := range nums {
		if num%2 == 0 {
			total++
		}
	}

	return total

}
