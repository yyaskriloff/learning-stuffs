package add

import "testing"

func TestAdd(t *testing.T) {
	t.Run("positive values", func(t *testing.T) {
		got := Add(2, 4)
		want := 6
		if got != want {
			t.Errorf("Add(2, 4) = %d; want %d", got, want)
		}
	})
	t.Run("negative values", func(t *testing.T) {
		got := Add(-1, -3)
		want := -4
		if got != want {
			t.Errorf("Add(1, 3) = %d; want %d", got, want)
		}
	})
	t.Run("zero values", func(t *testing.T) {
		got := Add(0, 0)
		want := 0
		if got != want {
			t.Errorf("Add(0, 0) = %d; want %d", got, want)
		}
	})
	t.Run("mixed signs", func(t *testing.T) {
		got := Add(-2, 4)
		want := 2
		if got != want {
			t.Errorf("Add(2, 4) = %d; want %d", got, want)
		}
	})

}
