package main

import "testing"

var tcs = []struct {
	in  string
	out int
}{
	{"abcdef", 609043},
	{"pqrstuv", 1048970},
}

func TestPartOne(t *testing.T) {
	for _, tt := range tcs {
		t.Run(tt.in, func(t *testing.T) {
			result := partOne(tt.in)
			if result != tt.out {
				t.Errorf("got %d want %d", result, tt.out)
			}
		})
	}
}
