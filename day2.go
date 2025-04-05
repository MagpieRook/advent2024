package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

// Returns true if:
//   - levels are all increasing or all decreasing
//   - adjacent levels differ by at least one and at most three
func safetyTest(levels []int64) bool {
	increasing := levels[0] < levels[1]
	decreasing := levels[0] > levels[1]
	// println("increasing", increasing, "decreasing", decreasing)
	if !increasing && !decreasing {
		return false
	}
	for i := range len(levels) - 1 {
		// println(levels[i], levels[i+1])
		// Make sure our trend line is consistent (safe)
		currentIncreasing := levels[i] < levels[i+1]
		currentDecreasing := levels[i] > levels[i+1]
		if increasing != currentIncreasing || decreasing != currentDecreasing {
			// println("increasing", currentIncreasing, "decreasing", currentDecreasing)
			return false
		}

		// Make sure our amount of change is safe
		level := int64(math.Abs(float64(levels[i] - levels[i+1])))
		// println(level)
		if 1 <= level && level <= 3 {
			continue
		}
		return false
	}
	return true
}

func Day2() {
	// Input is rows of five integers separated by spaces
	input, err := os.ReadFile("day2input.txt")
	if err != nil {
		panic(err)
	}

	safeReports := 0
	for row := range strings.SplitSeq(string(input), "\n") {
		// println(row)
		levelStrings := strings.Split(row, " ")
		levels := []int64{}
		for _, s := range levelStrings {
			if s == "" {
				continue
			}
			level, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				panic(err)
			}
			levels = append(levels, level)
		}

		if len(levels) == 0 {
			continue
		}
		fmt.Println(levels)
		if safetyTest(levels) {
			safeReports += 1
			// Part 2: Problem Dampener can tolerate a single bad level
		} else {
			for i := range len(levels) {
				newLevels := slices.Clone(levels)
				newNewLevels := append(newLevels[:i], newLevels[i+1:]...)
				fmt.Println(levels, newNewLevels)
				if safetyTest(newNewLevels) {
					safeReports += 1
					break
				}
			}
		}
	}

	fmt.Println(safeReports)
}
