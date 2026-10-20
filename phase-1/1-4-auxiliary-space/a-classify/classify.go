package classify

// Examples A–E are short algorithms to classify for auxiliary space.
// Do not change the implementations; analyze them as written.

// ExampleA scans nums and returns their sum.
func ExampleA(nums []int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}

// ExampleB returns a new slice containing each element of nums doubled.
func ExampleB(nums []int) []int {
	out := make([]int, len(nums))
	for i, v := range nums {
		out[i] = v * 2
	}
	return out
}

// ExampleC returns how many distinct values appear in nums.
func ExampleC(nums []int) int {
	seen := make(map[int]struct{}, len(nums))
	for _, v := range nums {
		seen[v] = struct{}{}
	}
	return len(seen)
}

// ExampleD returns the sum of nums using recursion over the remaining suffix.
func ExampleD(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	return nums[0] + ExampleD(nums[1:])
}

// ExampleE copies nums into a new slice, then returns the sum of that copy.
func ExampleE(nums []int) int {
	copyNums := append([]int(nil), nums...)
	total := 0
	for _, v := range copyNums {
		total += v
	}
	return total
}
