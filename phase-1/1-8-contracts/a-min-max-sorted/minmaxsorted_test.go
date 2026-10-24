package minmaxsorted_test

import (
	"testing"

	minmaxsorted "github.com/yyaskriloff/learning-stuffs/phase-1/1-8-contracts/a-min-max-sorted"
)

func TestMinMax(t *testing.T) {
	tests := []struct {
		name    string
		nums    []int
		wantMin int
		wantMax int
		wantOK  bool
	}{
		{name: "single element", nums: []int{7}, wantMin: 7, wantMax: 7, wantOK: true},
		{name: "two elements", nums: []int{3, 8}, wantMin: 3, wantMax: 8, wantOK: true},
		{name: "duplicates", nums: []int{4, 1, 4, 2, 4}, wantMin: 1, wantMax: 4, wantOK: true},
		{name: "negative values", nums: []int{-5, -1, -9}, wantMin: -9, wantMax: -1, wantOK: true},
		{name: "mixed signs", nums: []int{-2, 0, 5, -7}, wantMin: -7, wantMax: 5, wantOK: true},
		{name: "already sorted", nums: []int{1, 2, 3, 4}, wantMin: 1, wantMax: 4, wantOK: true},
		{name: "reverse sorted", nums: []int{9, 6, 3, 1}, wantMin: 1, wantMax: 9, wantOK: true},
		{name: "all equal", nums: []int{5, 5, 5}, wantMin: 5, wantMax: 5, wantOK: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMin, gotMax, gotOK := minmaxsorted.MinMax(tt.nums)
			if gotOK != tt.wantOK || gotMin != tt.wantMin || gotMax != tt.wantMax {
				t.Errorf("MinMax(%v) = (%d, %d, %v); want (%d, %d, %v)",
					tt.nums, gotMin, gotMax, gotOK, tt.wantMin, tt.wantMax, tt.wantOK)
			}
		})
	}
}

func TestMinMaxEmptyContract(t *testing.T) {
	// Contract: empty/nil → (0, 0, false)
	for _, tt := range []struct {
		name string
		nums []int
	}{
		{name: "empty slice", nums: []int{}},
		{name: "nil slice", nums: nil},
	} {
		t.Run(tt.name, func(t *testing.T) {
			gotMin, gotMax, gotOK := minmaxsorted.MinMax(tt.nums)
			if gotOK || gotMin != 0 || gotMax != 0 {
				t.Errorf("MinMax(%v) = (%d, %d, %v); want (0, 0, false)",
					tt.nums, gotMin, gotMax, gotOK)
			}
		})
	}
}

func TestIsSortedNonDecreasing(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{name: "empty input", nums: []int{}, want: true},
		{name: "nil input", nums: nil, want: true},
		{name: "single element", nums: []int{3}, want: true},
		{name: "already sorted", nums: []int{1, 2, 3, 4}, want: true},
		{name: "sorted with duplicates", nums: []int{1, 2, 2, 5}, want: true},
		{name: "reverse sorted", nums: []int{4, 3, 2, 1}, want: false},
		{name: "unsorted middle dip", nums: []int{1, 3, 2, 4}, want: false},
		{name: "negative values sorted", nums: []int{-5, -3, -1, 0}, want: true},
		{name: "negative values unsorted", nums: []int{-1, -5, 0}, want: false},
		{name: "all equal", nums: []int{7, 7, 7}, want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minmaxsorted.IsSortedNonDecreasing(tt.nums)
			if got != tt.want {
				t.Errorf("IsSortedNonDecreasing(%v) = %v; want %v", tt.nums, got, tt.want)
			}
		})
	}
}

func TestIndexOfMax(t *testing.T) {
	tests := []struct {
		name      string
		nums      []int
		wantIndex int
		wantOK    bool
	}{
		{name: "single element", nums: []int{7}, wantIndex: 0, wantOK: true},
		{name: "unique max at start", nums: []int{9, 1, 2}, wantIndex: 0, wantOK: true},
		{name: "unique max at end", nums: []int{1, 2, 9}, wantIndex: 2, wantOK: true},
		{name: "duplicates first occurrence", nums: []int{3, 8, 8, 2}, wantIndex: 1, wantOK: true},
		{name: "all equal", nums: []int{4, 4, 4}, wantIndex: 0, wantOK: true},
		{name: "negative values", nums: []int{-5, -1, -9}, wantIndex: 1, wantOK: true},
		{name: "already sorted", nums: []int{1, 2, 3, 4}, wantIndex: 3, wantOK: true},
		{name: "reverse sorted", nums: []int{9, 6, 3, 1}, wantIndex: 0, wantOK: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIndex, gotOK := minmaxsorted.IndexOfMax(tt.nums)
			if gotOK != tt.wantOK || gotIndex != tt.wantIndex {
				t.Errorf("IndexOfMax(%v) = (%d, %v); want (%d, %v)",
					tt.nums, gotIndex, gotOK, tt.wantIndex, tt.wantOK)
			}
		})
	}
}

func TestIndexOfMaxEmptyContract(t *testing.T) {
	// Contract: empty/nil → (0, false)
	for _, tt := range []struct {
		name string
		nums []int
	}{
		{name: "empty slice", nums: []int{}},
		{name: "nil slice", nums: nil},
	} {
		t.Run(tt.name, func(t *testing.T) {
			gotIndex, gotOK := minmaxsorted.IndexOfMax(tt.nums)
			if gotOK || gotIndex != 0 {
				t.Errorf("IndexOfMax(%v) = (%d, %v); want (0, false)",
					tt.nums, gotIndex, gotOK)
			}
		})
	}
}
