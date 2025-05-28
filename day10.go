package main

import (
	"fmt"
	"maps"
	"os"
	"strconv"
	"strings"
)

// Takes topomap and index (row,col) of position
// Returns how many paths (0-9 sequences) exist starting at trail head (incl. 0)
// Counts ALL paths to ALL nines, including duplicate nines
func checkTrailheadPt2(topomap [][]uint8, row, col int) map[int][]int {
	curr := topomap[row][col]
	if curr == 9 {
		return map[int][]int{row: {col}}
	}

	total := map[int][]int{}
	if row > 0 && topomap[row-1][col] == curr+1 {
		temp := checkTrailheadPt2(topomap, row-1, col)
		for k, v := range temp {
			total[k] = append(total[k], v...)
		}
	}
	if row < len(topomap)-1 && topomap[row+1][col] == curr+1 {
		temp := checkTrailheadPt2(topomap, row+1, col)
		for k, v := range temp {
			total[k] = append(total[k], v...)
		}
	}
	if col > 0 && topomap[row][col-1] == curr+1 {
		temp := checkTrailheadPt2(topomap, row, col-1)
		for k, v := range temp {
			total[k] = append(total[k], v...)
		}
	}
	if col < len(topomap[row])-1 && topomap[row][col+1] == curr+1 {
		temp := checkTrailheadPt2(topomap, row, col+1)
		for k, v := range temp {
			total[k] = append(total[k], v...)
		}
	}
	return total
}

// Takes topomap and index (row,col) of position
// Returns how many paths (0-9 sequences) exist starting at trail head (incl. 0)
// Only counts each 9 once
func checkTrailheadPt1(topomap [][]uint8, row, col int) map[int]map[int]bool {
	curr := topomap[row][col]
	if curr == 9 {
		return map[int]map[int]bool{row: {col: true}}
	}

	total := map[int]map[int]bool{}
	if row > 0 && topomap[row-1][col] == curr+1 {
		temp := checkTrailheadPt1(topomap, row-1, col)
		for k, v := range temp {
			if total[k] == nil {
				total[k] = map[int]bool{}
			}
			maps.Copy(total[k], v)
		}
	}
	if row < len(topomap)-1 && topomap[row+1][col] == curr+1 {
		temp := checkTrailheadPt1(topomap, row+1, col)
		for k, v := range temp {
			if total[k] == nil {
				total[k] = map[int]bool{}
			}
			maps.Copy(total[k], v)
		}
	}
	if col > 0 && topomap[row][col-1] == curr+1 {
		temp := checkTrailheadPt1(topomap, row, col-1)
		for k, v := range temp {
			if total[k] == nil {
				total[k] = map[int]bool{}
			}
			maps.Copy(total[k], v)
		}
	}
	if col < len(topomap[row])-1 && topomap[row][col+1] == curr+1 {
		temp := checkTrailheadPt1(topomap, row, col+1)
		for k, v := range temp {
			if total[k] == nil {
				total[k] = map[int]bool{}
			}
			maps.Copy(total[k], v)
		}
	}
	return total
}

func Day10() {
	// input is a grid of integers
	// part 1: find all sequences of 0-9 and count how many there are
	// any integer can be used in multiple sequences
	input, err := os.ReadFile("day10input.txt")
	if err != nil {
		panic(err)
	}

	rows := strings.Split(string(input), "\n")
	topomap := [][]uint8{}
	for y, row := range rows {
		topomap = append(topomap, []uint8{})
		for _, col := range row {
			n, err := strconv.ParseInt(string(col), 10, 16)
			if err != nil {
				panic(err)
			}
			topomap[y] = append(topomap[y], uint8(n))
		}
	}
	score := 0
	score2 := 0
	for y, row := range topomap {
		for x, col := range row {
			if col == 0 {
				temp := checkTrailheadPt1(topomap, y, x)
				for k := range temp {
					score += len(temp[k])
				}

				// Part 2: now undo that work you did to deduplicate nines :laugh:
				// Trailheads are now rated by the number of ways to get to any number of 9s
				// Return this (likely larger) score
				temp2 := checkTrailheadPt2(topomap, y, x)
				for k := range temp2 {
					score2 += len(temp2[k])
				}
			}
		}
	}
	fmt.Println(score, score2)
}
