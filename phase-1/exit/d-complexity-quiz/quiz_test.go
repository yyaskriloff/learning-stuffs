package complexityquiz_test

import (
	"testing"

	complexityquiz "github.com/yyaskriloff/learning-stuffs/phase-1/exit/d-complexity-quiz"
)

// Ensures the quiz package compiles. Answers are graded orally, not by this test.
func TestQuizPackageCompiles(t *testing.T) {
	n := 0
	complexityquiz.Q1(n, func() {})
	complexityquiz.Q2(n, func() {})
	complexityquiz.Q3(n, func() {})
	complexityquiz.Q4(n, func() {})
	complexityquiz.Q5(n, func() {})
	_ = complexityquiz.Q6(nil)
	_ = complexityquiz.Q7(nil)
	_ = complexityquiz.Q8(nil)
	_ = complexityquiz.Q9(nil)
	complexityquiz.Q10(nil, nil, func() {})
}
