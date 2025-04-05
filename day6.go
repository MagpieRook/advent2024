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
	visited := map[[2]int]int{}
	for checkPos[0] < width && checkPos[0] > 0 && checkPos[1] < height && checkPos[1] > 0 {
		checkPos[0] = currentPos[0] + facing[0]
		checkPos[1] = currentPos[1] + facing[1]

		// this is > 3 to avoid somehow duplicating things or whatever idk
		if visited[currentPos] > 3 && visited[checkPos] > 3 {
			return true
		}

		visited[currentPos] += 1

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

		// fmt.Println("visited:", visited)
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

		// fmt.Println("update current position")
		// if !slices.Contains(visited, checkPos) {
		// 	visited = append(visited, checkPos)
		// }
		currentPos[0], currentPos[1] = checkPos[0], checkPos[1]
		// fmt.Println("visited:", visited)
	}

	fmt.Println(len(visited))
	fmt.Println(len(possibleObstructions))

	// part 2: count number of loops possible from starting position with one added obstacle
	// basically, somewhere where there's 3 obstacles:
	//   - in the direction of facing (i.e. [0,3] when facing is up)
	//   - in the direction of changeFacing (i.e. [1,6], which would be hit after facing turns right)
	//   - in the direction of the second changeFacing (i.e. [4,5], as above)
	//   - and a fourth possible obstacle that can be added in the final facing to complete the loop (i.e. [3,2])
	//   the example would give us (# for obstacle, arrows for facing, O for added):
	// 		...#......
	//		...>.v#...
	// 		...^......
	// 		..#^.<....
	// 		.....O....
	// therefore, we need a square that fits the following criteria, where
	// highest means lowest obstacle[0], lowest means highest obstacle[0],
	// furthest means highest obstacle[1], and closest means lowest obstacle[1]:
	//   - the top obstacle needs to be 1 further than the closest obstacle    (obstacle{y,x+1})
	//   - the bottom obstacle needs to be 1 closer than the furthest obstacle (obstacle{y,x-1})
	//   - the left obstacle needs to be 1 higher than the lowest obstacle     (obstacle{y-1,x})
	//   - the right obstacle needs to be 1 lower than the highest obstacle    (obstacle{y+1,x})
	// possibleChanges stores the indexes of obstacles to prevent duplicates
	// possibleChanges := [][]int{}
	// for i, obstacle := range obstacles {
	// 	for j := i + 1; j < len(obstacles); j++ {
	// 		// could the i obstacle go on the left, and j be the bottom?
	// 		if compareSidesLeftBottom(obstacle, obstacles[j]) {
	// 			fmt.Println("checking left + bottom")
	// 			possibleChange := checkSquare(obstacles, -1, -1, j, i)
	// 			fmt.Println("checked left + bottom", len(possibleChange))
	// 			if len(possibleChange) == 4 {
	// 				possibleChanges = append(possibleChanges, possibleChange)
	// 			}
	// 		}
	// 		// could the i obstacle go on the right, and j be the top?
	// 		if compareSidesRightTop(obstacle, obstacles[j]) {
	// 			fmt.Println("checking right + top")
	// 			possibleChange := checkSquare(obstacles, j, i, -1, -1)
	// 			fmt.Println("checked right + top", len(possibleChange))
	// 			if len(possibleChange) == 4 {
	// 				possibleChanges = append(possibleChanges, possibleChange)
	// 			}
	// 		}
	// 		// could the i obstacle go on the top, and j be the left?
	// 		if compareSidesLeftTop(obstacles[j], obstacle) {
	// 			fmt.Println("checking left + top")
	// 			possibleChange := checkSquare(obstacles, i, -1, -1, j)
	// 			fmt.Println("checked left + top", len(possibleChange))
	// 			if len(possibleChange) == 4 {
	// 				possibleChanges = append(possibleChanges, possibleChange)
	// 			}
	// 		}
	// 		// could the i obstacle go on the bottom, and j be the right?
	// 		if compareSidesRightBottom(obstacles[j], obstacle) {
	// 			fmt.Println("checking right + bottom")
	// 			possibleChange := checkSquare(obstacles, -1, i, j, -1)
	// 			fmt.Println("checking right + bottom", len(possibleChange))
	// 			if len(possibleChange) == 4 {
	// 				possibleChanges = append(possibleChanges, possibleChange)
	// 			}
	// 		}
	// 	}
	// }
	// fmt.Println(len(possibleChanges))
}
