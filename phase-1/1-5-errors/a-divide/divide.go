package divide

import "errors"

// ErrDivideByZero is returned when the divisor is zero.
var ErrDivideByZero = errors.New("divide: division by zero")

// Divide returns a divided by b using Go integer division (truncated toward zero).
// If b is 0, it returns ErrDivideByZero.
//
// Contract: decide and document behavior for math.MinInt / -1.
func Divide(a, b int) (int, error) {
	panic("unimplemented")
}
