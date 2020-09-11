package main

import (
	"fmt"
	"testing"
)

var ribbontests = []struct {
	in  Box
	out int
}{
	{Box{length: 2, width: 3, height: 4}, 34},
}

func TestRibbonCalc(t *testing.T) {
	for _, tt := range ribbontests {
		t.Run(fmt.Sprintf("%d%dx%d", tt.in.length, tt.in.width, tt.in.height), func(t *testing.T) {
			ribbon := calculateNeededRibbon(tt.in)
			if ribbon != tt.out {
				t.Errorf("got %d, want %d", ribbon, tt.out)
			}
		})
	}
}
