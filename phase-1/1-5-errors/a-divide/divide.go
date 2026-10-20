package divide

import (
	"errors"
	"math"
)

// ErrDivideByZero is returned when the divisor is zero.
var ErrDivideByZero = errors.New("divide: division by zero")
var ErrIntOverflow = errors.New("divide: int overflow")

// Divide returns a divided by b using Go integer division (truncated toward zero).
// If b is 0, it returns ErrDivideByZero.
//
// Contract: decide and document behavior for math.MinInt / -1.
// In this scenerio it returns and ErrIntOverflow
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}

	if a == math.MinInt && b == -1 {
		return 0, ErrIntOverflow
	}

	return a / b, nil
}
