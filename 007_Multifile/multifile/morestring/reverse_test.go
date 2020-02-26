package morestrings

import "testing"

func TestReverseRunes(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}
	for _, c := range cases {
		got := ReverseRunes(c.in)
		if got != c.want {
			t.Errorf("ReverseRunes(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestOneplusOne(t *testing.T) {
	cases1 := []struct {
		in1, want1 int
	}{
		{2, 4},
		{4, 8},
		{6, 12},
	}

	for _, z := range cases1 {
		got1 := OnePlusOne(z.in1)
		if got1 != z.want1 {
			t.Errorf("OnePlusOne (%v) == %v, want %v", z.in1, got1, z.want1)
		}
	}
}
