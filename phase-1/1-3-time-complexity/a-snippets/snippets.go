package snippets

// Each Snippet* runs op according to a different loop pattern.
// Analyze the time complexity of these patterns; do not change the loops.

// SnippetA — single loop over n.
func SnippetA(n int, op func()) {
	for i := 0; i < n; i++ {
		op()
	}
}

// SnippetB — nested loops, each bound by n.
func SnippetB(n int, op func()) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			op()
		}
	}
}

// SnippetC — two consecutive loops, each bound by n.
func SnippetC(n int, op func()) {
	for i := 0; i < n; i++ {
		op()
	}
	for j := 0; j < n; j++ {
		op()
	}
}

// SnippetD — loop that halves its control value each iteration.
func SnippetD(n int, op func()) {
	for i := n; i > 0; i /= 2 {
		op()
	}
}

// SnippetE — outer loop over n; inner loop runs only when a nested condition holds.
func SnippetE(n int, op func()) {
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			for j := 0; j < n; j++ {
				op()
			}
		}
	}
}
