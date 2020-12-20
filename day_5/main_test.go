package main

import (
	"fmt"
	"testing"
)

func TestSeatID(t *testing.T) {
	var tests = []struct {
		letters string
		want    float64
	}{
		{"BFFFBBFRRR", 567},
		{"FFFBBBFRRR", 119},
		{"BBFFBBFRLL", 820},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s should yield %v", tt.letters, tt.want)

		t.Run(testname, func(t *testing.T) {
			got := GetSeatID(tt.letters)
			if got != tt.want {
				t.Errorf("got %v, but we want %v\n", got, tt.want)
			}
		})
	}
}
