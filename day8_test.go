package main

import (
	"testing"
)

func TestDay8(t *testing.T) {
	pt1, pt2 := Day8("day8test.txt")
	// https://adventofcode.com/2024/day/8
	if pt1 != 14 || pt2 != 34 {
		t.Errorf(`Day8("day8test.txt") = (%d, %d), want (14, 34)`, pt1, pt2)
	}
}
