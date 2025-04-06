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
		case 1:
			total += ys[i]
		case 2:
			total *= ys[i]
		case 3:
			newTotal := fmt.Sprintf("%d%d", total, ys[i])
			total64, err := strconv.ParseInt(newTotal, 10, 0)
			if err != nil {
				panic(err)
			}
			total = int(total64)
		}
	}
	return total == x
}

// Unfortunately currently duplicates, but I got the right answer so we skip
// Better solution memory and time-wise: a tree! alas
// Part 2: add a third operator (HAH I PREPARED FOR THIS GOTTEM)
func generateAllOperators(allOperators [][]uint, baseOperators []uint) [][]uint {
	operators1 := slices.Clone(baseOperators)
	operators2 := slices.Clone(baseOperators)
	operators3 := slices.Clone(baseOperators)
	for i := range baseOperators {
		if baseOperators[i] == 0 {
			operators1[i] = 1
			operators2[i] = 2
			operators3[i] = 3
			return slices.Concat(generateAllOperators(allOperators, operators1), generateAllOperators(allOperators, operators2), generateAllOperators(allOperators, operators3))
		}
	}
	allOperators = append(allOperators, operators1, operators2, operators3)
	return allOperators
}

// Recursive function to check all possible variations of operators
func checkAllOperators(x int, ys []int) bool {
	baseOperators := make([]uint, len(ys)-1)
	allOperators := generateAllOperators([][]uint{}, baseOperators)
	for _, operators := range allOperators {
		if checkEquation(x, ys, operators) {
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
		nums := strings.Split(row, ": ")
		x64, err := strconv.ParseInt(nums[0], 10, 0)
		if err != nil {
			panic(err)
		}

		for y := range strings.SplitSeq(nums[1], " ") {
			y64, err := strconv.ParseInt(y, 10, 0)
			if err != nil {
				panic(err)
			}
			testEquations[int(x64)] = append(testEquations[int(x64)], int(y64))
		}
	}

	total := 0
	for x, ys := range testEquations {
		if checkAllOperators(x, ys) {
			total += x
		}
	}

	fmt.Println(total)
}
