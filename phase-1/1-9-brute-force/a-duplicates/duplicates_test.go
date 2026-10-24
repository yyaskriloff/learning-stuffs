package duplicates_test

import (
	"math/rand"
	"reflect"
	"testing"

	duplicates "github.com/yyaskriloff/learning-stuffs/phase-1/1-9-brute-force/a-duplicates"
)

func TestHasDuplicateFixed(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{name: "empty slice", nums: []int{}, want: false},
		{name: "nil slice", nums: nil, want: false},
		{name: "single element", nums: []int{7}, want: false},
		{name: "no duplicates", nums: []int{1, 2, 3, 4}, want: false},
		{name: "adjacent duplicates", nums: []int{1, 2, 2, 3}, want: true},
		{name: "non-adjacent duplicates", nums: []int{1, 2, 3, 1}, want: true},
		{name: "all equal", nums: []int{5, 5, 5}, want: true},
		{name: "negative values with duplicate", nums: []int{-3, -1, -3}, want: true},
		{name: "negative values unique", nums: []int{-3, -1, 0}, want: false},
		{name: "two elements duplicate", nums: []int{9, 9}, want: true},
		{name: "two elements unique", nums: []int{9, 8}, want: false},
	}

	for _, tt := range tests {
		t.Run("pairwise/"+tt.name, func(t *testing.T) {
			got := duplicates.HasDuplicatePairwise(tt.nums)
			if got != tt.want {
				t.Errorf("HasDuplicatePairwise(%v) = %v; want %v", tt.nums, got, tt.want)
			}
		})
		t.Run("map/"+tt.name, func(t *testing.T) {
			got := duplicates.HasDuplicateMap(tt.nums)
			if got != tt.want {
				t.Errorf("HasDuplicateMap(%v) = %v; want %v", tt.nums, got, tt.want)
			}
		})
	}
}

func TestHasDuplicateDoesNotMutate(t *testing.T) {
	nums := []int{1, 2, 1}
	original := append([]int(nil), nums...)

	_ = duplicates.HasDuplicatePairwise(nums)
	if !reflect.DeepEqual(nums, original) {
		t.Errorf("HasDuplicatePairwise mutated input: got %v; want %v", nums, original)
	}

	_ = duplicates.HasDuplicateMap(nums)
	if !reflect.DeepEqual(nums, original) {
		t.Errorf("HasDuplicateMap mutated input: got %v; want %v", nums, original)
	}
}

func TestHasDuplicateRandomizedAgreement(t *testing.T) {
	const seed = 20261024
	rng := rand.New(rand.NewSource(seed))
	t.Logf("random seed = %d", seed)

	for trial := 0; trial < 200; trial++ {
		n := rng.Intn(40)
		nums := make([]int, n)
		for i := range nums {
			nums[i] = rng.Intn(25) - 5
		}

		pairwise := duplicates.HasDuplicatePairwise(nums)
		mapped := duplicates.HasDuplicateMap(nums)
		if pairwise != mapped {
			t.Fatalf("trial %d: disagreement on %v: pairwise=%v map=%v (seed=%d)",
				trial, nums, pairwise, mapped, seed)
		}
	}
}
