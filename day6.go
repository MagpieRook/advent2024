package main

import (
	"fmt"
	"maps"
	"os"
	"slices"
	"strings"
)

func changeFacing(facing [2]int) [2]int {
	if facing[0] == -1 {
		return [2]int{0, 1}
	} else if facing[0] == 1 {
		return [2]int{0, -1}
	}

	if facing[1] == -1 {
		return [2]int{-1, 0}
	} else if facing[1] == 1 {
		return [2]int{1, 0}
	}
	panic("Invalid input")
}

func checkForLoop(currentPos, facing [2]int, obstacles map[int][]int, width, height int) bool {
	checkPos := [2]int{currentPos[0] + facing[0], currentPos[1] + facing[1]}
	visited := map[[2]int][][2]int{}
	for currentPos[0] < height && currentPos[0] >= 0 && currentPos[1] < width && currentPos[1] >= 0 {
		if slices.Contains(visited[currentPos], facing) {
			return true
		}

		visited[currentPos] = append(visited[currentPos], facing)

		checkPos[0] = currentPos[0] + facing[0]
		checkPos[1] = currentPos[1] + facing[1]
		if slices.Contains(obstacles[checkPos[0]], checkPos[1]) {
			facing = changeFacing(facing)
			continue
		}
		currentPos[0], currentPos[1] = checkPos[0], checkPos[1]
	}
	return false
}

func main() {
	// input is a grid with:
	//   - an arrow (^/>/</v) for guard starting position
	//   - a # for obstructions the guard can't see or move through
	//   - a . everywhere else
	input, err := os.ReadFile("day6input.txt")
	if err != nil {
		panic(err)
	}

	// part 1 goal: mark every tile the guard will stand on
	// guard steps forward each turn
	// when guard reaches an obstruction, turn right 90 degrees
	obstacleMap := map[int][]int{}
	newObstaclePositions := map[int][]int{}
	var startingPos [2]int
	rows := strings.Split(string(input), "\n")
	for y, row := range rows {
		for x, col := range row {
			if col == '^' {
				startingPos = [2]int{y, x}
			} else if col == '#' {
				obstacleMap[y] = append(obstacleMap[y], x)
			} else {
				newObstaclePositions[y] = append(newObstaclePositions[y], x)
			}
		}
	}
	currentPos := [2]int{startingPos[0], startingPos[1]}
	facing := [2]int{-1, 0}
	visited := [][2]int{startingPos}
	width, height := len(rows[currentPos[1]]), len(rows)
	checkPos := [2]int{}
	for currentPos[0] < height && currentPos[0] >= 0 && currentPos[1] < width && currentPos[1] >= 0 {
		checkPos[0] = currentPos[0] + facing[0]
		checkPos[1] = currentPos[1] + facing[1]
		onExistingPath := slices.Contains(visited, currentPos)
		if !onExistingPath {
			visited = append(visited, currentPos)
		}

		if slices.Contains(obstacleMap[checkPos[0]], checkPos[1]) {
			facing = changeFacing(facing)
			continue
		}

		currentPos[0], currentPos[1] = checkPos[0], checkPos[1]
	}

	fmt.Println(len(visited))

	// part 2: brute force edition II
	// just check every position, who cares if it takes 4 years if I get the right answer
	possibleObstructions := [][2]int{}
	facing = [2]int{-1, 0}
	for y, row := range newObstaclePositions {
		for _, col := range row {
			newObstacleMap := maps.Clone(obstacleMap)
			newObstacleMap[y] = append(newObstacleMap[y], col)
			if checkForLoop(startingPos, facing, newObstacleMap, width, height) {
				possibleObstructions = append(possibleObstructions, [2]int{y, col})
			}
		}
	}

	// for some reason, this solution gets 1836 (the correct answer) whereas the earlier version got 1913
	// something about following the path the person was taking rather than checking all positions, I guess
	// whatever we got there in the end
	fmt.Println(len(possibleObstructions))
}
