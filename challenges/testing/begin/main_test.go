// challenges/testing/begin/main_test.go
package main

import "testing"

// write a test for letterCounter.count
func TestLetterCount(t *testing.T) {
	var tests = []struct {
		name  string
		input string
		want  int
	}{
		{"empty", "", 0},
		{"first", "a2 32 ^ &/)", 1},
		{"second", "812 %6ab//", 2},
	}

	lc := letterCounter{}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := lc.count(tc.input)
			if got != tc.want {
				t.Errorf("got %v want %v", got, tc.want)
			}
		})
	}
}

// write a test for numberCounter.count
func TestNumberCount(t *testing.T) {
	var tests = []struct {
		name  string
		input string
		want  int
	}{
		{"empty", "", 0},
		{"first", "abc 1,?/", 1},
		{"second", "abc 12&8 ^", 3},
	}

	nc := numberCounter{}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := nc.count(tc.input)
			if got != tc.want {
				t.Errorf("got %v want %v", got, tc.want)
			}
		})
	}
}

// write a test for symbolCounter.count
func TestSymbolCount(t *testing.T) {
	var tests = []struct {
		name  string
		input string
		want  int
	}{
		{"empty", "", 0},
		{"first", "abc 1,?/", 4},
		{"second", "abc 12&8 ^_", 5},
	}

	sc := symbolCounter{}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := sc.count(tc.input)
			if got != tc.want {
				t.Errorf("got %v want %v", got, tc.want)
			}
		})
	}
}
