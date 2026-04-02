package main

import "testing"

func TestFirstWord(t *testing.T) {
	testCases := []struct {
		in   string
		want string
	}{
		{"", "\n"},
		{"             a as", "a\n"},
		{"   f     d", "f\n"},
		{"   asd    ad", "asd\n"},
		{"   salut !!! ", "salut\n"},
		{" salut ! ! !", "salut\n"},
		{"salut ! !", "salut\n"},
	}

	for _, tc := range testCases {
		got := FirstWord(tc.in)
		if got != tc.want {
			t.Errorf("FirstWord(%q) = %q; want %q", tc.in, got, tc.want)
		}
	}
}
