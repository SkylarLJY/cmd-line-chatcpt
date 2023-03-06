package main

import "testing"

func TestWrapStr(t *testing.T) {
	testcases := []struct {
		name  string
		input string
		width int
		exp   string
	}{
		{"emptyStirng", "", 1, ""},
		{"noWrap", "short string", 100, "short string"},
		{"wrapAtSpace", "123 123", 3, "123\n123"},
		{"wrapInWord", "123 123", 5, "123\n123"},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			res := wrapStr(tc.input, tc.width)
			if res != tc.exp {
				t.Errorf("expected %q but got %q\n", tc.exp, res)
			}
		})
	}
}
