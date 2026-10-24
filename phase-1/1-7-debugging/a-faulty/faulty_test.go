package faulty_test

import (
	"reflect"
	"testing"

	faulty "github.com/yyaskriloff/learning-stuffs/phase-1/1-7-debugging/a-faulty"
)

func TestRunningSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{name: "empty slice", nums: []int{}, want: []int{}},
		{name: "nil slice", nums: nil, want: []int{}},
		{name: "single element", nums: []int{7}, want: []int{7}},
		{name: "two elements", nums: []int{1, 2}, want: []int{1, 3}},
		{name: "typical", nums: []int{1, 2, 3, 4}, want: []int{1, 3, 6, 10}},
		{name: "mixed signs", nums: []int{5, -2, 3}, want: []int{5, 3, 6}},
		{name: "all negative", nums: []int{-1, -2, -3}, want: []int{-1, -3, -6}},
		{name: "zeros", nums: []int{0, 0, 0}, want: []int{0, 0, 0}},
		{name: "zero then values", nums: []int{0, 4, -1}, want: []int{0, 4, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := faulty.RunningSum(tt.nums)
			if len(got) != len(tt.want) {
				t.Fatalf("RunningSum(%v) len = %d; want %d (got %v)", tt.nums, len(got), len(tt.want), got)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RunningSum(%v) = %v; want %v", tt.nums, got, tt.want)
			}
		})
	}
}

func TestRunningSumDoesNotMutateInput(t *testing.T) {
	nums := []int{1, 2, 3}
	original := append([]int(nil), nums...)
	_ = faulty.RunningSum(nums)
	if !reflect.DeepEqual(nums, original) {
		t.Errorf("RunningSum mutated input: got %v; want %v", nums, original)
	}
}
