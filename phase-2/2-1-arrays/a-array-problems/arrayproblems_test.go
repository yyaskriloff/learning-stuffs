package arrayproblems_test

import (
	"slices"
	"testing"

	arrayproblems "github.com/yyaskriloff/learning-stuffs/phase-2/2-1-arrays/a-array-problems"
)

func TestSecondLargestDistinct(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
		ok   bool
	}{
		{name: "typical", nums: []int{3, 1, 4, 1, 5}, want: 4, ok: true},
		{name: "two distinct", nums: []int{10, 3}, want: 3, ok: true},
		{name: "second at end", nums: []int{1, 2, 3}, want: 2, ok: true},
		{name: "second at start", nums: []int{8, 9, 9}, want: 8, ok: true},
		{name: "negative numbers", nums: []int{-1, -5, -3}, want: -3, ok: true},
		{name: "mixed signs", nums: []int{-10, 0, 5, -3}, want: 0, ok: true},
		{name: "duplicates of largest", nums: []int{7, 7, 2, 7}, want: 2, ok: true},
		{name: "many values", nums: []int{4, 9, 1, 9, 4, 8}, want: 8, ok: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := arrayproblems.SecondLargestDistinct(tt.nums)
			if ok != tt.ok || got != tt.want {
				t.Errorf("SecondLargestDistinct(%v) = (%d, %v); want (%d, %v)",
					tt.nums, got, ok, tt.want, tt.ok)
			}
		})
	}
}

func TestSecondLargestDistinctFewerThanTwoDistinctContract(t *testing.T) {
	t.Skip("decide the contract, then ask the instructor to complete this case")

	for _, tt := range []struct {
		name string
		nums []int
	}{
		{name: "empty input", nums: []int{}},
		{name: "nil input", nums: nil},
		{name: "single element", nums: []int{5}},
		{name: "all duplicates", nums: []int{4, 4, 4}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			_, _ = arrayproblems.SecondLargestDistinct(tt.nums)
			// Instructor fills expectations after the learner documents the contract.
		})
	}
}

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{name: "empty input", in: []int{}, want: []int{}},
		{name: "nil input", in: nil, want: nil},
		{name: "no zeroes", in: []int{1, 2, 3}, want: []int{1, 2, 3}},
		{name: "all zeroes", in: []int{0, 0, 0}, want: []int{0, 0, 0}},
		{name: "typical", in: []int{0, 1, 0, 3, 12}, want: []int{1, 3, 12, 0, 0}},
		{name: "zeroes already at end", in: []int{1, 2, 0, 0}, want: []int{1, 2, 0, 0}},
		{name: "zeroes at start", in: []int{0, 0, 1, 2}, want: []int{1, 2, 0, 0}},
		{name: "preserve nonzero order", in: []int{0, 4, 0, 2, 0, 1}, want: []int{4, 2, 1, 0, 0, 0}},
		{name: "negative numbers", in: []int{0, -1, 0, -2, 3}, want: []int{-1, -2, 3, 0, 0}},
		{name: "single zero", in: []int{0}, want: []int{0}},
		{name: "single nonzero", in: []int{9}, want: []int{9}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var nums []int
			if tt.in != nil {
				nums = slices.Clone(tt.in)
			}
			arrayproblems.MoveZeroes(nums)
			if !slices.Equal(nums, tt.want) {
				t.Errorf("MoveZeroes(%v) => %v; want %v", tt.in, nums, tt.want)
			}
		})
	}
}

func TestMoveZeroesInPlace(t *testing.T) {
	nums := []int{0, 1, 0, 2}
	header := nums
	arrayproblems.MoveZeroes(nums)
	if len(nums) == 0 {
		t.Fatal("unexpected empty result")
	}
	if &nums[0] != &header[0] {
		t.Fatal("MoveZeroes must mutate in place; backing array address changed")
	}
	if !slices.Equal(nums, []int{1, 2, 0, 0}) {
		t.Errorf("after MoveZeroes got %v; want [1 2 0 0]", nums)
	}
}

func TestMergeSorted(t *testing.T) {
	tests := []struct {
		name string
		a    []int
		b    []int
		want []int
	}{
		{name: "both empty", a: []int{}, b: []int{}, want: []int{}},
		{name: "nil and empty", a: nil, b: []int{}, want: []int{}},
		{name: "one empty unequal lengths", a: []int{1, 2, 3}, b: []int{}, want: []int{1, 2, 3}},
		{name: "other empty", a: []int{}, b: []int{4, 5}, want: []int{4, 5}},
		{name: "unequal slice lengths", a: []int{1, 3}, b: []int{2, 4, 5, 6}, want: []int{1, 2, 3, 4, 5, 6}},
		{name: "equal lengths", a: []int{1, 4}, b: []int{2, 3}, want: []int{1, 2, 3, 4}},
		{name: "duplicates across inputs", a: []int{1, 2, 2}, b: []int{2, 3}, want: []int{1, 2, 2, 2, 3}},
		{name: "negative numbers", a: []int{-5, -1, 0}, b: []int{-3, 2}, want: []int{-5, -3, -1, 0, 2}},
		{name: "all of a before b", a: []int{1, 2}, b: []int{3, 4}, want: []int{1, 2, 3, 4}},
		{name: "all of b before a", a: []int{5, 6}, b: []int{1, 2}, want: []int{1, 2, 5, 6}},
		{name: "single elements", a: []int{2}, b: []int{1}, want: []int{1, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var a, b []int
			if tt.a != nil {
				a = slices.Clone(tt.a)
			}
			if tt.b != nil {
				b = slices.Clone(tt.b)
			}
			origA := slices.Clone(a)
			origB := slices.Clone(b)

			got := arrayproblems.MergeSorted(a, b)
			if !slices.Equal(got, tt.want) {
				t.Errorf("MergeSorted(%v, %v) = %v; want %v", tt.a, tt.b, got, tt.want)
			}
			if !slices.Equal(a, origA) {
				t.Errorf("MergeSorted mutated a: got %v; want %v", a, origA)
			}
			if !slices.Equal(b, origB) {
				t.Errorf("MergeSorted mutated b: got %v; want %v", b, origB)
			}
		})
	}
}

func TestRemoveDuplicatesSorted(t *testing.T) {
	tests := []struct {
		name     string
		in       []int
		wantLen  int
		wantHead []int
	}{
		{name: "empty input", in: []int{}, wantLen: 0, wantHead: []int{}},
		{name: "nil input", in: nil, wantLen: 0, wantHead: nil},
		{name: "no duplicates", in: []int{1, 2, 3}, wantLen: 3, wantHead: []int{1, 2, 3}},
		{name: "all duplicates", in: []int{5, 5, 5, 5}, wantLen: 1, wantHead: []int{5}},
		{name: "typical", in: []int{1, 1, 2}, wantLen: 2, wantHead: []int{1, 2}},
		{name: "multiple groups", in: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, wantLen: 5, wantHead: []int{0, 1, 2, 3, 4}},
		{name: "negative numbers", in: []int{-3, -3, -1, 0, 0, 2}, wantLen: 4, wantHead: []int{-3, -1, 0, 2}},
		{name: "single element", in: []int{9}, wantLen: 1, wantHead: []int{9}},
		{name: "two equal", in: []int{4, 4}, wantLen: 1, wantHead: []int{4}},
		{name: "two distinct", in: []int{4, 7}, wantLen: 2, wantHead: []int{4, 7}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var nums []int
			if tt.in != nil {
				nums = slices.Clone(tt.in)
			}
			gotLen := arrayproblems.RemoveDuplicatesSorted(nums)
			if gotLen != tt.wantLen {
				t.Fatalf("RemoveDuplicatesSorted(%v) length = %d; want %d", tt.in, gotLen, tt.wantLen)
			}
			if tt.wantHead == nil {
				if nums != nil && gotLen != 0 {
					t.Fatalf("expected nil or empty logical result; got %v len=%d", nums, gotLen)
				}
				return
			}
			if gotLen > len(nums) {
				t.Fatalf("returned length %d exceeds slice length %d", gotLen, len(nums))
			}
			gotHead := nums[:gotLen]
			if !slices.Equal(gotHead, tt.wantHead) {
				t.Errorf("RemoveDuplicatesSorted(%v) head = %v; want %v", tt.in, gotHead, tt.wantHead)
			}
		})
	}
}

func TestRemoveDuplicatesSortedInPlace(t *testing.T) {
	nums := []int{1, 1, 2, 3, 3}
	header := nums
	n := arrayproblems.RemoveDuplicatesSorted(nums)
	if n != 3 {
		t.Fatalf("length = %d; want 3", n)
	}
	if &nums[0] != &header[0] {
		t.Fatal("RemoveDuplicatesSorted must mutate in place; backing array address changed")
	}
	if !slices.Equal(nums[:n], []int{1, 2, 3}) {
		t.Errorf("head = %v; want [1 2 3]", nums[:n])
	}
}
