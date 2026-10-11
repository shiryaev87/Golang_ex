package evenorodd // ← должно совпадать с solution.go

import (
	"fmt"
	"testing"
)

func TestEvenOrOdd(t *testing.T) {
	if got := EvenOrOdd(3); got != "Odd" {
		t.Fatalf("EvenOrOdd(3) = %q, хотим %q", got, "Odd")
	}

	if got := EvenOrOdd(4); got != "Even" {
		t.Fatalf("EvenOrOdd(4) = %q, хотим %q", got, "Even")
	}
}

func TestEvenOrOddTable(t *testing.T) {
	cases := []struct {
		n    int
		want string
	}{
		{1, "Odd"},
		{2, "Even"},
		{0, "Even"},
		{-3, "Odd"},
		{-4, "Even"},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("n=%d", c.n), func(t *testing.T) {
			if got := EvenOrOdd(c.n); got != c.want {
				t.Errorf("EvenOrOdd(%d) = %q, хотим %q", c.n, got, c.want)
			}
		})
	}
}
