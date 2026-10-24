package slicemutation_test

import (
	"reflect"
	"slices"
	"testing"

	slicemutation "github.com/yyaskriloff/learning-stuffs/phase-1/exit/f-slice-mutation"
)

func TestRotateLeft(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{name: "empty slice", in: []int{}, want: []int{}},
		{name: "nil slice", in: nil, want: nil},
		{name: "single element", in: []int{7}, want: []int{7}},
		{name: "two elements", in: []int{1, 2}, want: []int{2, 1}},
		{name: "typical", in: []int{1, 2, 3, 4}, want: []int{2, 3, 4, 1}},
		{name: "duplicates", in: []int{5, 5, 1}, want: []int{5, 1, 5}},
		{name: "negatives", in: []int{-1, 0, 1}, want: []int{0, 1, -1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var nums []int
			if tt.in != nil {
				nums = slices.Clone(tt.in)
			}
			slicemutation.RotateLeft(nums)
			if !slices.Equal(nums, tt.want) {
				t.Errorf("RotateLeft(%v) => %v; want %v", tt.in, nums, tt.want)
			}
		})
	}
}

func TestRotateLeftInPlace(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	header := nums
	slicemutation.RotateLeft(nums)
	if &nums[0] != &header[0] {
		t.Fatal("RotateLeft must mutate in place; backing array address changed")
	}
	if !slices.Equal(nums, []int{2, 3, 4, 1}) {
		t.Errorf("after RotateLeft got %v; want [2 3 4 1]", nums)
	}
}

func TestDropWhile(t *testing.T) {
	isNeg := func(n int) bool { return n < 0 }
	isZero := func(n int) bool { return n == 0 }

	tests := []struct {
		name string
		in   []int
		drop func(int) bool
		want []int
	}{
		{name: "empty slice", in: []int{}, drop: isNeg, want: []int{}},
		{name: "nil slice", in: nil, drop: isNeg, want: []int{}},
		{name: "drops nothing", in: []int{1, 2, 3}, drop: isNeg, want: []int{1, 2, 3}},
		{name: "drops all", in: []int{-1, -2}, drop: isNeg, want: []int{}},
		{name: "drops proper prefix", in: []int{-1, -2, 3, -4}, drop: isNeg, want: []int{3, -4}},
		{name: "single kept", in: []int{5}, drop: isNeg, want: []int{5}},
		{name: "single dropped", in: []int{-5}, drop: isNeg, want: []int{}},
		{name: "zeros prefix", in: []int{0, 0, 2, 0}, drop: isZero, want: []int{2, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var nums []int
			if tt.in != nil {
				nums = slices.Clone(tt.in)
			}
			original := slices.Clone(nums)
			got := slicemutation.DropWhile(nums, tt.drop)
			if !slices.Equal(got, tt.want) {
				t.Errorf("DropWhile(%v) = %v; want %v", tt.in, got, tt.want)
			}
			if !slices.Equal(nums, original) {
				t.Errorf("DropWhile mutated input: got %v; want unchanged %v", nums, original)
			}
		})
	}
}

func TestDropWhileDoesNotAliasInput(t *testing.T) {
	nums := []int{0, 0, 3, 4}
	got := slicemutation.DropWhile(nums, func(n int) bool { return n == 0 })
	if len(got) == 0 {
		t.Fatal("expected non-empty DropWhile result")
	}
	got[0] = 99
	if nums[2] == 99 {
		t.Fatal("DropWhile result aliases input; changing result mutated nums")
	}
}

func TestDropWhileEmptyResultIndependent(t *testing.T) {
	nums := []int{1, 2}
	got := slicemutation.DropWhile(nums, func(n int) bool { return true })
	if len(got) != 0 {
		t.Fatalf("expected empty result, got %v", got)
	}
	if !reflect.DeepEqual(nums, []int{1, 2}) {
		t.Errorf("DropWhile mutated input when dropping all: got %v", nums)
	}
}
