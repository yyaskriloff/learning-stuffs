package duplicates

// HasDuplicatePairwise reports whether any value appears more than once in nums,
// using pairwise comparisons (nested loops). It must not use a map.
// Empty and single-element inputs return false. The input is not modified.
func HasDuplicatePairwise(nums []int) bool {

	for i, iv := range nums {
		for j, jv := range nums {
			if i == j {
				continue
			}

			if iv == jv {
				return true
			}

		}

	}

	return false
}

// HasDuplicateMap reports whether any value appears more than once in nums,
// using a map to track values seen so far.
// Empty and single-element inputs return false. The input is not modified.
func HasDuplicateMap(nums []int) bool {

	seen := make(map[int]struct{})

	for _, v := range nums {

		if _, exists := seen[v]; exists {
			return true
		}
		seen[v] = struct{}{}

	}

	return false
}
