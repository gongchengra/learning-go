package Factorial

import (
	"testing"
)

func TestFactorial(t *testing.T) {
	t.Run("Factorial of 0", func(t *testing.T) {
		got := Factorial(0)
		want := 1
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("Factorial of 5", func(t *testing.T) {
		got := Factorial(5)
		want := 120
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("Factorial of 8", func(t *testing.T) {
		got := Factorial(8)
		want := 40320
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("Factorial of 10", func(t *testing.T) {
		got := Factorial(10)
		want := 3628800
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
