package countinrange_test

import (
	"reflect"
	"testing"

	countinrange "github.com/yyaskriloff/learning-stuffs/phase-1/exit/a-count-in-range"
)

func TestCountInRange(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		lo   int
		hi   int
		want int
	}{
		{name: "empty slice", nums: []int{}, lo: 0, hi: 10, want: 0},
		{name: "nil slice", nums: nil, lo: 0, hi: 10, want: 0},
		{name: "lo greater than hi", nums: []int{1, 2, 3}, lo: 5, hi: 1, want: 0},
		{name: "single element in range", nums: []int{7}, lo: 0, hi: 10, want: 1},
		{name: "single element out of range", nums: []int{7}, lo: 0, hi: 5, want: 0},
		{name: "all in range", nums: []int{2, 3, 4}, lo: 2, hi: 4, want: 3},
		{name: "none in range", nums: []int{1, 2, 3}, lo: 10, hi: 20, want: 0},
		{name: "mixed some in range", nums: []int{-2, 0, 3, 8}, lo: 0, hi: 5, want: 2},
		{name: "negatives in range", nums: []int{-5, -3, -1}, lo: -4, hi: -2, want: 1},
		{name: "lo equals hi exact match", nums: []int{1, 4, 4, 2}, lo: 4, hi: 4, want: 2},
		{name: "inclusive bounds", nums: []int{1, 2, 3, 4, 5}, lo: 2, hi: 4, want: 3},
		{name: "zero only", nums: []int{-1, 0, 1}, lo: 0, hi: 0, want: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countinrange.CountInRange(tt.nums, tt.lo, tt.hi)
			if got != tt.want {
				t.Errorf("CountInRange(%v, %d, %d) = %d; want %d",
					tt.nums, tt.lo, tt.hi, got, tt.want)
			}
		})
	}
}

func TestCountInRangeDoesNotMutate(t *testing.T) {
	nums := []int{1, 5, 9}
	original := append([]int(nil), nums...)
	_ = countinrange.CountInRange(nums, 0, 10)
	if !reflect.DeepEqual(nums, original) {
		t.Errorf("CountInRange mutated input: got %v; want %v", nums, original)
	}
}
