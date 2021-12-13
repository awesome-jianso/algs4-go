package strings_test

import (
	"testing"

	. "github.com/youngzhu/algs4-go/strings"
)

var testCases = []struct {
	pattern, text string
	index         int
}{
	{"abracadabra", "abacadabrabracabracadabrabrabracad", 14},
	{"rab", "abacadabrabracabracadabrabrabracad", 8},
	{"rabrabracad", "abacadabrabracabracadabrabrabracad", 23},
	{"bcara", "abacadabrabracabracadabrabrabracad", -1},
	{"abacad", "abacadabrabracabracadabrabrabracad", 0},
}

func TestBruteSearch1(t *testing.T) {
	for _, tc := range testCases {
		got := BruteSearch1(tc.pattern, tc.text)
		want := tc.index

		if got != want {
			t.Errorf("BruteSearch1(%q, %q) got %d, want %d", tc.pattern, tc.text, got, want)
		}
	}
}

func TestBruteSearch2(t *testing.T) {
	for _, tc := range testCases {
		got := BruteSearch2(tc.pattern, tc.text)
		want := tc.index

		if got != want {
			t.Errorf("BruteSearch1(%q, %q) got %d, want %d", tc.pattern, tc.text, got, want)
		}
	}
}