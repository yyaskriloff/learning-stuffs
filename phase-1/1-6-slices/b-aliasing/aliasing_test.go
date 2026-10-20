package aliasing

import (
	"testing"
	"unsafe"
)

func header(t *testing.T, name string, s []int) {
	t.Helper()
	var ptr uintptr
	if len(s) > 0 {
		ptr = uintptr(unsafe.Pointer(&s[0]))
	}
	t.Logf("%s: len=%d cap=%d ptr=%#x values=%v", name, len(s), cap(s), ptr, s)
}

func TestLenAndCap(t *testing.T) {
	s := make([]int, 2, 5)
	s[0], s[1] = 10, 20
	header(t, "s", s)

	if len(s) != 2 {
		t.Fatalf("len(s) = %d; want 2", len(s))
	}
	if cap(s) != 5 {
		t.Fatalf("cap(s) = %d; want 5", cap(s))
	}

	part := s[1:2]
	header(t, "part=s[1:2]", part)
	if len(part) != 1 {
		t.Fatalf("len(part) = %d; want 1", len(part))
	}
	// From index 1 to end of backing array: 4 slots remain.
	if cap(part) != 4 {
		t.Fatalf("cap(part) = %d; want 4", cap(part))
	}
}

func TestAppendReusesBacking(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := a[:2]
	header(t, "a before", a)
	header(t, "b before", b)

	b = append(b, 99)
	header(t, "a after append on b", a)
	header(t, "b after append", b)

	if a[2] != 99 {
		t.Fatalf("a[2] = %d; want 99 (append reused a's array)", a[2])
	}
	if &a[0] != &b[0] {
		t.Fatal("expected a and b to share the same backing array start")
	}
}

func TestAppendReplacesBacking(t *testing.T) {
	a := make([]int, 2, 2)
	a[0], a[1] = 1, 2
	orig0 := &a[0]
	header(t, "a before", a)

	b := append(a, 3)
	header(t, "a after", a)
	header(t, "b after append", b)

	if a[0] != 1 || a[1] != 2 || len(a) != 2 {
		t.Fatalf("a changed unexpectedly: %v", a)
	}
	if len(b) != 3 || b[2] != 3 {
		t.Fatalf("b = %v; want [1 2 3]", b)
	}
	if &b[0] == orig0 {
		t.Fatal("expected append to allocate a new backing array for b")
	}
}

func TestAliasingAcrossVariables(t *testing.T) {
	x := []int{1, 2, 3}
	y := x
	header(t, "x before", x)
	header(t, "y before", y)

	y[1] = 9
	header(t, "x after y[1]=9", x)
	header(t, "y after", y)

	if x[1] != 9 {
		t.Fatalf("x[1] = %d; want 9 (y aliases x)", x[1])
	}
}

func TestAliasingThroughFunction(t *testing.T) {
	nums := []int{1, 2, 3}
	header(t, "nums before", nums)
	bumpFirst(nums, 7)
	header(t, "nums after bumpFirst", nums)

	if nums[0] != 7 {
		t.Fatalf("nums[0] = %d; want 7 (function mutated shared array)", nums[0])
	}
}
