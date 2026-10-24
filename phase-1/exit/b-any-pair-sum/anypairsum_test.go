package anypairsum_test

import (
	"math/rand"
	"reflect"
	"testing"

	anypairsum "github.com/yyaskriloff/learning-stuffs/phase-1/exit/b-any-pair-sum"
)

func TestAnyPairSumFixed(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   bool
	}{
		{name: "empty slice", nums: []int{}, target: 0, want: false},
		{name: "nil slice", nums: nil, target: 0, want: false},
		{name: "single element", nums: []int{5}, target: 10, want: false},
		{name: "two elements yes", nums: []int{3, 7}, target: 10, want: true},
		{name: "two elements no", nums: []int{3, 7}, target: 11, want: false},
		{name: "same value distinct indexes", nums: []int{5, 5}, target: 10, want: true},
		{name: "cannot use same index twice", nums: []int{5}, target: 10, want: false},
		{name: "negatives", nums: []int{-3, 1, 4}, target: 1, want: true},
		{name: "zeros", nums: []int{0, 0}, target: 0, want: true},
		{name: "zero with nonzero", nums: []int{0, 4}, target: 4, want: true},
		{name: "no solution", nums: []int{1, 2, 3, 9}, target: 15, want: false},
		{name: "non-adjacent pair", nums: []int{1, 8, 2, 7}, target: 9, want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := anypairsum.AnyPairSum(tt.nums, tt.target)
			if got != tt.want {
				t.Errorf("AnyPairSum(%v, %d) = %v; want %v", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}

func TestAnyPairSumDoesNotMutate(t *testing.T) {
	nums := []int{1, 2, 3}
	original := append([]int(nil), nums...)
	_ = anypairsum.AnyPairSum(nums, 5)
	if !reflect.DeepEqual(nums, original) {
		t.Errorf("AnyPairSum mutated input: got %v; want %v", nums, original)
	}
}

func bruteAnyPairSum(nums []int, target int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return true
			}
		}
	}
	return false
}

func TestAnyPairSumRandomized(t *testing.T) {
	const seed = 20261024
	rng := rand.New(rand.NewSource(seed))
	t.Logf("random seed = %d", seed)

	for trial := 0; trial < 200; trial++ {
		n := rng.Intn(25)
		nums := make([]int, n)
		for i := range nums {
			nums[i] = rng.Intn(21) - 10
		}
		target := rng.Intn(41) - 20

		got := anypairsum.AnyPairSum(nums, target)
		want := bruteAnyPairSum(nums, target)
		if got != want {
			t.Fatalf("trial %d: AnyPairSum(%v, %d) = %v; want %v (seed=%d)",
				trial, nums, target, got, want, seed)
		}
	}
}
