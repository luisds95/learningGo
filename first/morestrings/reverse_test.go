package morestrings

import "testing"

func TestReverseRunes(t *testing.T) {
	cases := []struct {
		str, expected string
	}{
		{"Hello, World!", "!dlroW ,olleH"},
		{"", ""},
		{"Hello, 世界", "界世 ,olleH"},
	}
	for _, c := range cases {
		result := ReverseRunes(c.str)
		if result != c.expected {
			t.Errorf("ReverseRunes(%q) == %q, but expected %q", c.str, result, c.expected)
		}
	}
}
