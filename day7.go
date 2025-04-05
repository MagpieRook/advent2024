package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func checkEquation(x int, ys []int, operators []uint) bool {
	total := ys[0]
	for i := 1; i < len(ys); i++ {
		switch operators[i-1] {
		case 0:
			total += ys[i]
		case 1:
			total *= ys[i]
		}
	}
	return total == x
}

func checkAllEquations(x int, ys []int, operators []uint) bool {
	if len(operators) == 0 {
		operators = make([]uint, len(ys)-1)
	}
	for i := range operators {
		operators0 := slices.Clone(operators)
		operators0[i] = 0
		operators1 := slices.Clone(operators)
		operators1[i] = 1

		if checkEquation(x, ys, operators0) || checkEquation(x, ys, operators1) {
			return true
		}
	}
	return false
}

func Day7() {
	// Input is a list of possible equations without operators in the format
	// `x: y1 y2 y3 ...` where x is the final value and yN are the values to operate
	// operators are addition and multiplication
	input, err := os.ReadFile("day7input.txt")
	if err != nil {
		panic(err)
	}

	// part 1: find the equations which can be solved with the values provided
	// sum the values (x) to get the answer
	testEquations := map[int][]int{}
	rows := strings.SplitSeq(string(input), "\n")
	for row := range rows {
		nums := strings.Split(row, ":")
		x64, err := strconv.ParseInt(nums[0], 10, 0)
		if err != nil {
			panic(err)
		}

		for y := range strings.SplitSeq(nums[1], " ") {
			if y == "" {
				continue
			}
			y64, err := strconv.ParseInt(y, 10, 0)
			if err != nil {
				panic(err)
			}
			testEquations[int(x64)] = append(testEquations[int(x64)], int(y64))
		}
	}

	total := 0
	for x, ys := range testEquations {
		if checkAllEquations(x, ys, nil) {
			total += x
		}
	}

	fmt.Println(total)
}
