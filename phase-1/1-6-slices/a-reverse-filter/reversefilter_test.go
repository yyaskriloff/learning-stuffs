package reversefilter_test

import (
	"slices"
	"testing"

	"github.com/yyaskriloff/learning-stuffs/phase-1/1-6-slices/a-reverse-filter"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{name: "empty slice", in: []int{}, want: []int{}},
		{name: "nil slice", in: nil, want: nil},
		{name: "single element", in: []int{7}, want: []int{7}},
		{name: "two elements", in: []int{1, 2}, want: []int{2, 1}},
		{name: "duplicates", in: []int{1, 2, 2, 3}, want: []int{3, 2, 2, 1}},
		{name: "typical", in: []int{1, 2, 3, 4, 5}, want: []int{5, 4, 3, 2, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var nums []int
			if tt.in != nil {
				nums = slices.Clone(tt.in)
			}
			reversefilter.Reverse(nums)
			if !slices.Equal(nums, tt.want) {
				t.Errorf("Reverse(%v) => %v; want %v", tt.in, nums, tt.want)
			}
		})
	}
}

func TestReverseInPlace(t *testing.T) {
	nums := []int{1, 2, 3}
	header := nums
	reversefilter.Reverse(nums)
	if &nums[0] != &header[0] {
		t.Fatal("Reverse must mutate in place; backing array address changed")
	}
	if !slices.Equal(nums, []int{3, 2, 1}) {
		t.Errorf("after Reverse got %v; want [3 2 1]", nums)
	}
}

func TestFilter(t *testing.T) {
	even := func(n int) bool { return n%2 == 0 }
	positive := func(n int) bool { return n > 0 }

	tests := []struct {
		name string
		in   []int
		keep func(int) bool
		want []int
	}{
		{name: "empty slice", in: []int{}, keep: even, want: []int{}},
		{name: "nil slice", in: nil, keep: even, want: []int{}},
		{name: "single element kept", in: []int{4}, keep: even, want: []int{4}},
		{name: "single element dropped", in: []int{3}, keep: even, want: []int{}},
		{name: "duplicates kept", in: []int{2, 2, 3, 2}, keep: even, want: []int{2, 2, 2}},
		{name: "all dropped", in: []int{1, 3, 5}, keep: even, want: []int{}},
		{name: "all kept", in: []int{2, 4, 6}, keep: even, want: []int{2, 4, 6}},
		{name: "mixed positives", in: []int{-2, 0, 3, -1, 4}, keep: positive, want: []int{3, 4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var nums []int
			if tt.in != nil {
				nums = slices.Clone(tt.in)
			}
			original := slices.Clone(nums)
			got := reversefilter.Filter(nums, tt.keep)
			if !slices.Equal(got, tt.want) {
				t.Errorf("Filter(%v) = %v; want %v", tt.in, got, tt.want)
			}
			if !slices.Equal(nums, original) {
				t.Errorf("Filter mutated input: got %v; want unchanged %v", nums, original)
			}
		})
	}
}

func TestFilterDoesNotAliasInput(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	got := reversefilter.Filter(nums, func(n int) bool { return n%2 == 0 })
	if len(got) == 0 {
		t.Fatal("expected non-empty filter result")
	}
	got[0] = 99
	if nums[1] == 99 {
		t.Fatal("Filter result aliases input; changing result mutated nums")
	}
}
