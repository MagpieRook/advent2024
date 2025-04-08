package main

import (
	"fmt"
	"os"
	"strings"
)

// Part 1: count how many times XMAS appears
func checkNeighborsXMAS(data []string, rowIndex, columnIndex int) int {
	fmt.Println(string(data[rowIndex][columnIndex]))
	if data[rowIndex][columnIndex] != byte('X') {
		fmt.Println("Not an X")
		return 0
	}

	checkWord := "XMAS"

	// Can be refactored to do all these checks at once, but idk simplicity

	count := 0
	// Check right
	for rightIndex := 1; rightIndex < 4; rightIndex++ {
		if columnIndex+rightIndex >= len(data[rowIndex]) || data[rowIndex][columnIndex+rightIndex] != checkWord[rightIndex] {
			count -= 1
			break
		}
	}
	count += 1

	// Check left
	for leftIndex := 1; leftIndex < 4; leftIndex++ {
		if columnIndex-leftIndex < 0 || data[rowIndex][columnIndex-leftIndex] != checkWord[leftIndex] {
			count -= 1
			break
		}
	}
	count += 1

	// Check down
	for downIndex := 1; downIndex < 4; downIndex++ {
		if rowIndex+downIndex >= len(data) || data[rowIndex+downIndex][columnIndex] != checkWord[downIndex] {
			count -= 1
			break
		}
	}
	count += 1

	// Check up
	for upIndex := 1; upIndex < 4; upIndex++ {
		if rowIndex-upIndex < 0 || data[rowIndex-upIndex][columnIndex] != checkWord[upIndex] {
			count -= 1
			break
		}
	}
	count += 1

	// Check diagonal down + right
	for diagonalIndex := 1; diagonalIndex < 4; diagonalIndex++ {
		if rowIndex+diagonalIndex >= len(data) || columnIndex+diagonalIndex >= len(data[rowIndex+diagonalIndex]) || data[rowIndex+diagonalIndex][columnIndex+diagonalIndex] != checkWord[diagonalIndex] {
			count -= 1
			break
		}
	}
	count += 1

	// Check diagonal up + right
	for diagonalIndex := 1; diagonalIndex < 4; diagonalIndex++ {
		if rowIndex-diagonalIndex < 0 || columnIndex+diagonalIndex >= len(data[rowIndex-diagonalIndex]) || data[rowIndex-diagonalIndex][columnIndex+diagonalIndex] != checkWord[diagonalIndex] {
			count -= 1
			break
		}
	}
	count += 1

	// Check diagonal down + left
	for diagonalIndex := 1; diagonalIndex < 4; diagonalIndex++ {
		if rowIndex+diagonalIndex >= len(data) || columnIndex-diagonalIndex < 0 || data[rowIndex+diagonalIndex][columnIndex-diagonalIndex] != checkWord[diagonalIndex] {
			count -= 1
			break
		}
	}
	count += 1

	// Check diagonal up + left
	for diagonalIndex := 1; diagonalIndex < 4; diagonalIndex++ {
		if rowIndex-diagonalIndex < 0 || columnIndex-diagonalIndex < 0 || data[rowIndex-diagonalIndex][columnIndex-diagonalIndex] != checkWord[diagonalIndex] {
			count -= 1
			break
		}
	}
	count += 1

	return count
}

// Part 2: count how many times "MAS" makes an "X", for example:
//   - M _ S
//   - _ A _
//   - M _ S
func checkNeighbors(data []string, rowIndex, columnIndex int) int {
	if data[rowIndex][columnIndex] != byte('A') || rowIndex-1 < 0 || rowIndex+1 >= len(data) || columnIndex-1 < 0 || columnIndex+1 >= len(data[rowIndex]) {
		return 0
	}

	diagonals := 0

	topLeft := data[rowIndex-1][columnIndex-1]
	topRight := data[rowIndex-1][columnIndex+1]
	bottomLeft := data[rowIndex+1][columnIndex-1]
	bottomRight := data[rowIndex+1][columnIndex+1]

	// Check top left to bottom right
	if topLeft == byte('M') && bottomRight == byte('S') {
		diagonals += 1
	}

	// Check top right to bottom left
	if topRight == byte('M') && bottomLeft == byte('S') {
		diagonals += 1
	}

	// Check bottom left to top right
	if bottomLeft == byte('M') && topRight == byte('S') {
		diagonals += 1
	}

	// Check bottom right to top left
	if bottomRight == byte('M') && topLeft == byte('S') {
		diagonals += 1
	}

	// If there are two matching diagonals (a cross/X) return true/1
	var ret int
	if diagonals == 2 {
		ret = 1
	} else {
		ret = 0
	}
	return ret
}

func Day4() {
	inputBytes, err := os.ReadFile("day4input.txt")
	if err != nil {
		panic(err)
	}

	input := string(inputBytes)
	rows := strings.Split(input, "\n")
	total := 0
	for rowIndex, row := range rows {
		for columnIndex := range row {
			total += checkNeighborsXMAS(rows, rowIndex, columnIndex)
		}
	}

	fmt.Println(total)

	total = 0
	for rowIndex, row := range rows {
		for columnIndex := range row {
			total += checkNeighbors(rows, rowIndex, columnIndex)
		}
	}

	fmt.Println(total)
}
