package funcluigi

import (
	"testing"
)

func TestSumTwoNumber(t *testing.T) {
	casi := []struct {
		InputTry1     int
		InputTry2     int
		OutputAspcted int
	}{
		{2, 4, 6},
		{5, 9, 14},
		{3, 3, 6},
		{8, 7, 15},
		{5, 5, 10},
		{0, 0, 0},
	}
	for _, c := range casi {
		got := SumTwoNumber(c.InputTry1, c.InputTry2)
		if got != c.OutputAspcted {
			t.Errorf("SumTwoNumber(%q and %q) == %q want (%q)", c.InputTry1, c.InputTry2, got, c.OutputAspcted)
		}
	}
}
