package partition_test

import (
	"reflect"
	"slices"
	"testing"

	partition "github.com/yyaskriloff/learning-stuffs/phase-1/exit/c-partition"
)

func TestPartition(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		pivot    int
		wantLess []int
		wantGe   []int
	}{
		{name: "empty slice", nums: []int{}, pivot: 0, wantLess: []int{}, wantGe: []int{}},
		{name: "nil slice", nums: nil, pivot: 0, wantLess: []int{}, wantGe: []int{}},
		{name: "single less", nums: []int{3}, pivot: 5, wantLess: []int{3}, wantGe: []int{}},
		{name: "single ge", nums: []int{7}, pivot: 5, wantLess: []int{}, wantGe: []int{7}},
		{name: "single equal", nums: []int{5}, pivot: 5, wantLess: []int{}, wantGe: []int{5}},
		{name: "all less", nums: []int{1, 2, 3}, pivot: 10, wantLess: []int{1, 2, 3}, wantGe: []int{}},
		{name: "all ge", nums: []int{5, 6, 7}, pivot: 5, wantLess: []int{}, wantGe: []int{5, 6, 7}},
		{name: "mixed order preserved", nums: []int{4, 1, 5, 2, 5}, pivot: 4, wantLess: []int{1, 2}, wantGe: []int{4, 5, 5}},
		{name: "duplicates equal pivot", nums: []int{3, 3, 3}, pivot: 3, wantLess: []int{}, wantGe: []int{3, 3, 3}},
		{name: "negatives and zero", nums: []int{-2, 0, -1, 3}, pivot: 0, wantLess: []int{-2, -1}, wantGe: []int{0, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var nums []int
			if tt.nums != nil {
				nums = slices.Clone(tt.nums)
			}
			less, ge := partition.Partition(nums, tt.pivot)
			if !slices.Equal(less, tt.wantLess) {
				t.Errorf("Partition(%v, %d) less = %v; want %v", tt.nums, tt.pivot, less, tt.wantLess)
			}
			if !slices.Equal(ge, tt.wantGe) {
				t.Errorf("Partition(%v, %d) ge = %v; want %v", tt.nums, tt.pivot, ge, tt.wantGe)
			}
		})
	}
}

func TestPartitionDoesNotMutate(t *testing.T) {
	nums := []int{4, 1, 5, 2}
	original := append([]int(nil), nums...)
	_, _ = partition.Partition(nums, 3)
	if !reflect.DeepEqual(nums, original) {
		t.Errorf("Partition mutated input: got %v; want %v", nums, original)
	}
}

func TestPartitionDoesNotAliasInput(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	less, ge := partition.Partition(nums, 3)
	if len(less) == 0 && len(ge) == 0 {
		t.Fatal("expected non-empty partition results")
	}
	if len(less) > 0 {
		less[0] = 99
		if nums[0] == 99 {
			t.Fatal("less aliases input; changing less mutated nums")
		}
	}
	if len(ge) > 0 {
		ge[0] = 88
		for _, v := range nums {
			if v == 88 {
				t.Fatal("ge aliases input; changing ge mutated nums")
			}
		}
	}
}
