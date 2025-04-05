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
	return [2]int{0, 0}
}

// how does this brute force solution not work D:
func checkForLoop(currentPos, facing [2]int, obstacles map[int][]int, width, height int) bool {
	checkPos := [2]int{currentPos[0] + facing[0], currentPos[1] + facing[1]}
	newObstacles := maps.Clone(obstacles)
	newObstacles[checkPos[0]] = append(newObstacles[checkPos[0]], checkPos[1])
	visited := map[[2]int][][2]int{}
	for checkPos[0] < width && checkPos[0] >= 0 && checkPos[1] < height && checkPos[1] >= 0 {
		checkPos[0] = currentPos[0] + facing[0]
		checkPos[1] = currentPos[1] + facing[1]

		if slices.Contains(visited[currentPos], facing) {
			return true
		}

		visited[currentPos] = append(visited[currentPos], facing)

		if slices.Contains(newObstacles[checkPos[0]], checkPos[1]) {
			facing = changeFacing(facing)
			continue
		}
		currentPos[0], currentPos[1] = checkPos[0], checkPos[1]
	}
	return false
}

func main() {
	// input is a grid with:
	//   - an arrow (^/>/</v) for guards
	//   - a # for obstructions the guard can't see or move through
	//   - a . everywhere else
	// part 1 goal: mark every tile the guard will stand on
	// guard steps forward each turn
	// when guard reaches an obstruction, turn right 90 degrees
	input, err := os.ReadFile("day6input.txt")
	if err != nil {
		panic(err)
	}

	startingPos := [2]int{}
	// a map of obstacles by row and column
	obstacleMap := map[int][]int{}
	rows := strings.Split(string(input), "\n")
	for y, row := range rows {
		for x, col := range row {
			if col == '^' {
				startingPos[0] = y
				startingPos[1] = x
			}
			if col == '#' {
				obstacleMap[y] = append(obstacleMap[y], x)
			}
		}
	}
	currentPos := [2]int{}
	currentPos[0], currentPos[1] = startingPos[0], startingPos[1]
	facing := [2]int{-1, 0}
	visited := [][2]int{startingPos}
	possibleObstructions := [][2]int{}
	for {
		checkPos := [2]int{}
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

		// part 2: check for ways to make loops with new obstructions
		// we know the next position isn't an obstruction, so let's try adding one
		if checkForLoop(currentPos, facing, obstacleMap, len(rows[currentPos[1]]), len(rows)) {
			possibleObstructions = append(possibleObstructions, checkPos)
		}

		if checkPos[0] >= len(rows[currentPos[1]]) || checkPos[0] < 0 || checkPos[1] >= len(rows) || checkPos[1] < 0 {
			break
		}

		currentPos[0], currentPos[1] = checkPos[0], checkPos[1]
	}

	fmt.Println(len(visited))
	fmt.Println(len(possibleObstructions))
}
