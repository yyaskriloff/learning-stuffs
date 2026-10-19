package sumcounteven_test

import (
	"testing"

	"github.com/yyaskriloff/learning-stuffs/phase-1/1-2-loops/a-sum-count-even"
)

func TestSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "empty slice", nums: nil, want: 0},
		{name: "one element", nums: []int{7}, want: 7},
		{name: "all even", nums: []int{2, 4, 6}, want: 12},
		{name: "all odd", nums: []int{1, 3, 5}, want: 9},
		{name: "mixed values", nums: []int{1, 2, 3, 4}, want: 10},
		{name: "negative values", nums: []int{-3, -2, 5}, want: 0},
		{name: "typical", nums: []int{10, -4, 3}, want: 9},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sumcounteven.Sum(tt.nums)
			if got != tt.want {
				t.Errorf("Sum(%v) = %d; want %d", tt.nums, got, tt.want)
			}
		})
	}
}

func TestCountEven(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "empty slice", nums: nil, want: 0},
		{name: "one element even", nums: []int{8}, want: 1},
		{name: "one element odd", nums: []int{7}, want: 0},
		{name: "all even", nums: []int{2, 4, 6}, want: 3},
		{name: "all odd", nums: []int{1, 3, 5}, want: 0},
		{name: "mixed values", nums: []int{1, 2, 3, 4}, want: 2},
		{name: "negative values", nums: []int{-3, -2, 5, -4}, want: 2},
		{name: "zero is even", nums: []int{0, 1}, want: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sumcounteven.CountEven(tt.nums)
			if got != tt.want {
				t.Errorf("CountEven(%v) = %d; want %d", tt.nums, got, tt.want)
			}
		})
	}
}
