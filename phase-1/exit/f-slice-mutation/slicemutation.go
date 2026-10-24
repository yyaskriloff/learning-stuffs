package slicemutation

// RotateLeft rotates nums one position to the left in place.
// Empty and single-element slices are unchanged.
func RotateLeft(nums []int) {
	if len(nums) < 2 {
		return
	}

	for i := 0; i < len(nums)-1; i++ {
		nums[i], nums[i+1] = nums[i+1], nums[i]
	}

}

// DropWhile returns a new slice equal to nums after removing the longest prefix
// where every element satisfies drop. It must not modify nums.
func DropWhile(nums []int, drop func(int) bool) []int {
	newNums := []int{}

	donePrefix := false

	for _, v := range nums {

		if donePrefix {
			newNums = append(newNums, v)
		} else if !drop(v) {
			newNums = append(newNums, v)
			donePrefix = true
		}
	}

	return newNums
}
