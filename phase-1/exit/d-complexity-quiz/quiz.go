package complexityquiz

// Quiz snippets for Phase 1 exit. Analyze only; do not change the loops or allocations.

// Q1 — what is the time complexity in terms of n?
func Q1(n int, op func()) {
	for i := 0; i < n; i++ {
		op()
		op()
	}
}

// Q2 — what is the time complexity in terms of n?
func Q2(n int, op func()) {
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			op()
		}
	}
}

// Q3 — what is the time complexity in terms of n?
func Q3(n int, op func()) {
	for i := 1; i < n; i *= 2 {
		for j := 0; j < n; j++ {
			op()
		}
	}
}

// Q4 — what is the time complexity in terms of n?
func Q4(n int, op func()) {
	for i := 0; i < n; i++ {
		op()
	}
	for i := 0; i < n; i++ {
		for j := 0; j < 3; j++ {
			op()
		}
	}
}

// Q5 — what is the time complexity in terms of n?
func Q5(n int, op func()) {
	i := n
	for i > 1 {
		op()
		i = i / 2
	}
}

// Q6 — what is the auxiliary space in terms of n = len(nums)?
func Q6(nums []int) int {
	max := 0
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	return max
}

// Q7 — what is the auxiliary space in terms of n = len(nums)?
func Q7(nums []int) []int {
	out := make([]int, 0, len(nums))
	for _, v := range nums {
		if v%2 == 0 {
			out = append(out, v)
		}
	}
	return out
}

// Q8 — what is the auxiliary space in terms of n = len(nums)?
func Q8(nums []int) int {
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
	}
	best := 0
	for _, c := range freq {
		if c > best {
			best = c
		}
	}
	return best
}

// Q9 — give time and auxiliary space in terms of n = len(nums).
func Q9(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	return nums[0] + Q9(nums[1:])
}

// Q10 — give time and auxiliary space in terms of n = len(a). Assume len(a) == len(b).
func Q10(a, b []int, op func()) {
	n := len(a)
	tmp := make([]int, n)
	for i := 0; i < n; i++ {
		tmp[i] = a[i] + b[i]
		op()
	}
	_ = tmp
}
