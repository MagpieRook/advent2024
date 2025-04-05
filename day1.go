package main

import (
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Day1() {
	// Input is 2 lists (in file) separated by whitespace
	input, err := os.ReadFile("day1input.txt")
	if err != nil {
		panic(err)
	}

	listStrings := string(input)
	leftList := []int64{}
	rightList := []int64{}
	for row := range strings.SplitSeq(listStrings, "\n") {
		newValues := strings.Split(row, "   ")
		if newValues[0] == "" {
			continue
		}
		leftValue, err := strconv.ParseInt(newValues[0], 10, 64)
		if err != nil {
			panic(err)
		}
		leftList = append(leftList, leftValue)
		rightValue, err := strconv.ParseInt(newValues[1], 0, 64)
		if err != nil {
			panic(err)
		}
		rightList = append(rightList, rightValue)
	}

	// Sort both lists
	slices.Sort(leftList)
	slices.Sort(rightList)

	// Part 1: What is the total distance (i.e. sum of differences) of 2 lists
	totalDistance := float64(0)
	for i := range leftList {
		totalDistance += math.Abs(float64(leftList[i] - rightList[i]))
	}

	println(int(totalDistance))

	// Part 2: What is the similarity of the two lists, defined as:
	// sum(left list value * # times left list value appears in right list)

	leftIndex, rightIndex := 0, 0
	similarity := int64(0)
	for leftIndex < len(leftList) && rightIndex < len(rightList) {
		current := leftList[leftIndex]
		count := int64(0)
		// Right list is sorted, so move index up to next number and count
		// (Could also take rightIndex - startIndex)
		for rightIndex < len(rightList) {
			if rightList[rightIndex] > current {
				break
			}
			if rightList[rightIndex] == current {
				count += 1
			}
			rightIndex++
		}

		// similarity score is defined as value * right list appearances
		similarity += current * count

		// List is sorted, so avoid recalculating doubles
		for leftIndex = leftIndex + 1; leftIndex < len(leftList); {
			if leftList[leftIndex] != current {
				break
			}
			similarity += current * count
			leftIndex++
		}
	}

	println(similarity)
}
