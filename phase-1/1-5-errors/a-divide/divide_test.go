package divide_test

import (
	"errors"
	"math"
	"testing"

	"github.com/yyaskriloff/learning-stuffs/phase-1/1-5-errors/a-divide"
)

func TestDivideSuccess(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{name: "positive values", a: 10, b: 2, want: 5},
		{name: "zero dividend", a: 0, b: 5, want: 0},
		{name: "negative dividend", a: -10, b: 2, want: -5},
		{name: "negative divisor", a: 10, b: -2, want: -5},
		{name: "both negative", a: -10, b: -2, want: 5},
		{name: "truncate toward zero positive", a: 7, b: 2, want: 3},
		{name: "truncate toward zero negative dividend", a: -7, b: 2, want: -3},
		{name: "truncate toward zero negative divisor", a: 7, b: -2, want: -3},
		{name: "truncate toward zero both negative", a: -7, b: -2, want: 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := divide.Divide(tt.a, tt.b)
			if err != nil {
				t.Fatalf("Divide(%d, %d) unexpected error: %v", tt.a, tt.b, err)
			}
			if got != tt.want {
				t.Errorf("Divide(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestDivideByZero(t *testing.T) {
	tests := []struct {
		name string
		a    int
	}{
		{name: "zero divisor positive dividend", a: 10},
		{name: "zero divisor zero dividend", a: 0},
		{name: "zero divisor negative dividend", a: -10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := divide.Divide(tt.a, 0)
			if !errors.Is(err, divide.ErrDivideByZero) {
				t.Fatalf("Divide(%d, 0) error = %v; want errors.Is(..., ErrDivideByZero)", tt.a, err)
			}
			if got != 0 {
				t.Errorf("Divide(%d, 0) = %d; want 0 on error", tt.a, got)
			}
		})
	}
}

func TestDivideMinIntByMinusOne(t *testing.T) {
	got, err := divide.Divide(math.MinInt, -1)
	if !errors.Is(err, divide.ErrIntOverflow) {
		t.Fatalf("Divide(math.MinInt, -1) error = %v; want errors.Is(..., ErrIntOverflow)", err)
	}
	if got != 0 {
		t.Errorf("Divide(math.MinInt, -1) = %d; want 0 on error", got)
	}
}
