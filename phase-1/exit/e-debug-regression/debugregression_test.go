package debugregression_test

import (
	"reflect"
	"testing"

	debugregression "github.com/yyaskriloff/learning-stuffs/phase-1/exit/e-debug-regression"
)

func TestCumulativeMax(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{name: "empty slice", nums: []int{}, want: []int{}},
		{name: "nil slice", nums: nil, want: []int{}},
		{name: "single element", nums: []int{7}, want: []int{7}},
		{name: "strictly increasing", nums: []int{1, 2, 3, 4}, want: []int{1, 2, 3, 4}},
		{name: "strictly decreasing", nums: []int{4, 3, 2, 1}, want: []int{4, 4, 4, 4}},
		{name: "all equal", nums: []int{2, 2, 2}, want: []int{2, 2, 2}},
		{name: "mixed", nums: []int{3, 1, 4, 1, 5}, want: []int{3, 3, 4, 4, 5}},
		{name: "all negative", nums: []int{-5, -2, -8}, want: []int{-5, -2, -2}},
		{name: "starts negative then positive", nums: []int{-3, -1, 2}, want: []int{-3, -1, 2}},
		{name: "zeros", nums: []int{0, 0, 0}, want: []int{0, 0, 0}},
		{name: "single negative", nums: []int{-4}, want: []int{-4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := debugregression.CumulativeMax(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CumulativeMax(%v) = %v; want %v", tt.nums, got, tt.want)
			}
		})
	}
}

func TestCumulativeMaxDoesNotMutate(t *testing.T) {
	nums := []int{1, 3, 2}
	original := append([]int(nil), nums...)
	_ = debugregression.CumulativeMax(nums)
	if !reflect.DeepEqual(nums, original) {
		t.Errorf("CumulativeMax mutated input: got %v; want %v", nums, original)
	}
}
