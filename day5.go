package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func updateStringToInts(updateString string) []int64 {
	if updateString == "" {
		fmt.Println("Empty updateString")
		return []int64{}
	}
	update := strings.Split(updateString, ",")
	updateNums := []int64{}
	for _, num := range update {
		parsedNum, err := strconv.ParseInt(num, 10, 64)
		if err != nil {
			panic(err)
		}
		updateNums = append(updateNums, parsedNum)
	}

	return updateNums
}

func fixIncorrectUpdate(updateString string, ordering map[int64][]int64) int64 {
	update := updateStringToInts(updateString)
	if len(update) == 0 {
		return 0
	}
	fmt.Println("fix incorrect", update)
	for i := 0; i < len(update); {
		num := update[i]
		for j := i + 1; j < len(update); {
			laterNum := update[j]
			if slices.Contains(ordering[laterNum], num) {
				update[i] = laterNum
				update[j] = num
				fmt.Println(update)
				break
			}
			j++
		}
		if correct, _ := checkUpdateCorrect(update, ordering); correct {
			break
		} else if i+1 >= len(update) {
			i = 0
		} else {
			i++
		}
	}
	return update[len(update)/2]
}

func checkUpdateCorrect(update []int64, ordering map[int64][]int64) (bool, int64) {
	seen := []int64{}
	for _, num := range update {
		// fmt.Println(num, ordering[num])
		// fmt.Println(seen)
		for _, laterNum := range ordering[num] {
			// if we've seen a number in the rule for this number
			// the rule has been broken: it is before when it should be later
			if slices.Contains(seen, laterNum) {
				// fmt.Println(laterNum, "found")
				return false, 0
			}
			// fmt.Println(laterNum)
		}
		seen = append(seen, num)
	}

	// Finale part 1: get middle number
	return true, update[len(update)/2]
}

func Day5() {
	// Input
	// Section 1: XX|YY -- page XX is before page YY if both are in update
	// Section 2: XX,YY,ZZ,... -- print pages XX, YY, ZZ for update
	input, err := os.ReadFile("day5input.txt")
	if err != nil {
		panic(err)
	}

	rows := strings.Split(string(input), "\n")

	ordering := make(map[int64][]int64)
	endIndex := 0
	for i, row := range rows {
		updates := strings.Split(row, "|")
		if len(updates) != 2 {
			endIndex = i
			break
		}
		key, err := strconv.ParseInt(updates[0], 10, 64)
		if err != nil {
			panic(err)
		}
		value, err := strconv.ParseInt(updates[1], 10, 64)
		if err != nil {
			panic(err)
		}

		ordering[key] = append(ordering[key], value)
	}

	if rows[endIndex] == "\n" {
		endIndex++
	}

	// Part 1: Get total of correct
	var total int64 = 0
	incorrectUpdates := []string{}
	updates := rows[endIndex:]
	for _, update := range updates {
		updateSlice := updateStringToInts(update)
		if len(updateSlice) == 0 {
			continue
		}
		correct, middleItem := checkUpdateCorrect(updateSlice, ordering)
		fmt.Println(correct, middleItem)
		if correct {
			total += middleItem
		} else {
			incorrectUpdates = append(incorrectUpdates, update)
		}
	}
	fmt.Println(total)

	// Part 2: Fix incorrect and get their total
	total = 0
	for _, incorrect := range incorrectUpdates {
		middleItem := fixIncorrectUpdate(incorrect, ordering)
		fmt.Println("middle", middleItem)
		total += middleItem
	}

	fmt.Println(total)
}
