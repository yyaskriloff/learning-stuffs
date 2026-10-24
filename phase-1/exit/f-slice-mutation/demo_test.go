package slicemutation_test

import (
	"testing"

	slicemutation "github.com/yyaskriloff/learning-stuffs/phase-1/exit/f-slice-mutation"
)

// Demonstration tests: predict behavior before running (see README).
func TestDemoRotateLeftKeepsBackingArray(t *testing.T) {
	nums := make([]int, 4, 8)
	copy(nums, []int{10, 20, 30, 40})
	before := &nums[0]
	slicemutation.RotateLeft(nums)
	after := &nums[0]
	t.Logf("before=%p after=%p len=%d cap=%d values=%v", before, after, len(nums), cap(nums), nums)
	if before != after {
		t.Fatalf("expected same backing array; before=%p after=%p", before, after)
	}
}

func TestDemoDropWhileResultIndependent(t *testing.T) {
	nums := []int{-1, -2, 5, 6}
	got := slicemutation.DropWhile(nums, func(n int) bool { return n < 0 })
	t.Logf("input=%v result=%v", nums, got)
	if len(got) < 1 {
		t.Fatal("expected remaining elements")
	}
	got[0] = 999
	if nums[2] == 999 {
		t.Fatal("result aliased input")
	}
	if nums[2] != 5 {
		t.Fatalf("input changed unexpectedly: %v", nums)
	}
}
