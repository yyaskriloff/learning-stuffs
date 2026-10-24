package duplicates_test

import (
	"testing"

	duplicates "github.com/yyaskriloff/learning-stuffs/phase-1/1-9-brute-force/a-duplicates"
)

// Sizes chosen so pairwise cost grows noticeably while map stays flatter.
var benchSizes = []int{10, 100, 500, 1000}

func makeUnique(n int) []int {
	nums := make([]int, n)
	for i := range nums {
		nums[i] = i
	}
	return nums
}

func makeWithDuplicate(n int) []int {
	nums := makeUnique(n)
	if n >= 2 {
		nums[n-1] = nums[0]
	}
	return nums
}

func BenchmarkHasDuplicatePairwiseUnique(b *testing.B) {
	for _, n := range benchSizes {
		nums := makeUnique(n)
		b.Run(itoa(n), func(b *testing.B) {
			for b.Loop() {
				_ = duplicates.HasDuplicatePairwise(nums)
			}
		})
	}
}

func BenchmarkHasDuplicateMapUnique(b *testing.B) {
	for _, n := range benchSizes {
		nums := makeUnique(n)
		b.Run(itoa(n), func(b *testing.B) {
			for b.Loop() {
				_ = duplicates.HasDuplicateMap(nums)
			}
		})
	}
}

func BenchmarkHasDuplicatePairwiseDuplicateAtEnd(b *testing.B) {
	for _, n := range benchSizes {
		nums := makeWithDuplicate(n)
		b.Run(itoa(n), func(b *testing.B) {
			for b.Loop() {
				_ = duplicates.HasDuplicatePairwise(nums)
			}
		})
	}
}

func BenchmarkHasDuplicateMapDuplicateAtEnd(b *testing.B) {
	for _, n := range benchSizes {
		nums := makeWithDuplicate(n)
		b.Run(itoa(n), func(b *testing.B) {
			for b.Loop() {
				_ = duplicates.HasDuplicateMap(nums)
			}
		})
	}
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	var buf [16]byte
	i := len(buf)
	for n > 0 {
		i--
		buf[i] = byte('0' + n%10)
		n /= 10
	}
	return string(buf[i:])
}
