package partition

// Partition returns two new slices: values strictly less than pivot, then values
// greater than or equal to pivot, each preserving relative order from nums.
// Nil or empty nums returns two empty slices. The input is not modified.
func Partition(nums []int, pivot int) (less, ge []int) {
	runningLess := []int{}
	runningGe := []int{}

	for _, v := range nums {
		if v < pivot {
			runningLess = append(runningLess, v)
		}

		if v >= pivot {
			runningGe = append(runningGe, v)
		}
	}

	return runningLess, runningGe

}
