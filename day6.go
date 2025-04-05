package main

import (
	"fmt"
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

// honestly at this point it might be easier just to brute force it via walking, even if it takes longer to run
// too tired for this problem i guess
// ah butts i just realized i also need to keep in mind that if there's an obstacle in the way it won't work
// yeah i think maybe a refactor is warranted idk

func compareSidesLeftTop(leftObstacle [2]int, topObstacle [2]int) bool {
	return (topObstacle[0] != -1 && leftObstacle[0] != -1 &&
		topObstacle[0] < leftObstacle[0] &&
		topObstacle[1]-1 == leftObstacle[1])
}

func compareSidesLeftBottom(leftObstacle [2]int, bottomObstacle [2]int) bool {
	return (bottomObstacle[0] != -1 && leftObstacle[0] != -1 &&
		bottomObstacle[0]-1 == leftObstacle[0] &&
		bottomObstacle[1] > leftObstacle[1])
}

func compareSidesRightTop(rightObstacle [2]int, topObstacle [2]int) bool {
	return (topObstacle[0] != -1 && rightObstacle[0] != -1 &&
		topObstacle[0]+1 == rightObstacle[0] &&
		topObstacle[1] < rightObstacle[1])
}

func compareSidesRightBottom(rightObstacle [2]int, bottomObstacle [2]int) bool {
	return (bottomObstacle[0] != 1 && rightObstacle[0] != 1 &&
		bottomObstacle[0] > rightObstacle[0] &&
		bottomObstacle[1]-1 == rightObstacle[1])
}

// takes optional indexes (-1 for nil) and returns an empty list if incorrect
// given >1 indexes, look through the obstacles list to match other indexes
// to make the "corners" of the square
func checkSquare(obstacles [][2]int, topIndex, rightIndex, bottomIndex, leftIndex int) []int {
	originalSolution := []int{topIndex, rightIndex, bottomIndex, leftIndex}
	// attemptedSolutions: obstaclesIndex to list of tried top/right/bottom/left
	attemptedSolutions := []int{}
	previousSolutionLen := len(attemptedSolutions)
startCheckOver:
	solution := []int{topIndex, rightIndex, bottomIndex, leftIndex}
	for i, obstacle := range obstacles {
		if slices.Contains(solution, i) || slices.Contains(attemptedSolutions, i) {
			continue
		}
		if topIndex != -1 || bottomIndex != -1 {
			topObstacle := [2]int{-1, -1}
			bottomObstacle := [2]int{-1, -1}
			if topIndex != -1 {
				topObstacle = obstacles[topIndex]
			}
			if bottomIndex != -1 {
				bottomObstacle = obstacles[bottomIndex]
			}

			if leftIndex == -1 {
				leftObstacle := obstacle
				// basic logic for these kinds of if statements:
				// if topIndex == -1, left side and right side will be false, so we can continue
				// if topIndex != -1, left side will be true, so we only continue if right side is also true (which we want)
				// either way, the expression returns true in cases we want to continue the check
				// so by stringing two together, the && will only compare both if both are > -1
				// otherwise, false == false will be true, so we can compare against the index > -1
				if (topIndex != -1) == (compareSidesLeftTop(leftObstacle, topObstacle)) &&
					(bottomIndex != -1) == (compareSidesLeftBottom(leftObstacle, bottomObstacle)) {
					leftIndex = i
					solution[3] = leftIndex
				}
			}
			if rightIndex == -1 {
				rightObstacle := obstacle
				if (topIndex != -1) == (compareSidesRightTop(rightObstacle, topObstacle)) &&
					(bottomIndex != -1) == (compareSidesRightBottom(rightObstacle, bottomObstacle)) {
					rightIndex = i
					solution[1] = rightIndex
				}
			}
		}

		if rightIndex != -1 || leftIndex != -1 {
			rightObstacle := [2]int{-1, -1}
			leftObstacle := [2]int{-1, -1}
			if rightIndex != -1 {
				rightObstacle = obstacles[rightIndex]
			}
			if leftIndex != -1 {
				leftObstacle = obstacles[leftIndex]
			}

			if topIndex == -1 {
				topObstacle := obstacle
				if (leftIndex != -1) == (compareSidesLeftTop(leftObstacle, topObstacle)) &&
					(rightIndex != -1) == (compareSidesRightTop(rightObstacle, topObstacle)) {
					topIndex = i
					solution[0] = leftIndex
				}
			}
			if bottomIndex == -1 {
				bottomObstacle := obstacle
				if (leftIndex != -1) == (compareSidesLeftBottom(leftObstacle, bottomObstacle)) &&
					(rightIndex != -1) == (compareSidesRightBottom(rightObstacle, bottomObstacle)) {
					bottomIndex = i
					solution[2] = bottomIndex
				}
			}
		}
	}

	if slices.Contains(solution, -1) {
		for i, index := range solution {
			if index != -1 && originalSolution[i] != index {
				attemptedSolutions = append(attemptedSolutions, index)
			}
		}
		if len(attemptedSolutions) == previousSolutionLen || len(attemptedSolutions)+2 == len(obstacles) {
			return []int{}
		}
		previousSolutionLen = len(attemptedSolutions)
		goto startCheckOver
	}
	return solution
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
	obstacles := [][2]int{}
	rows := strings.Split(string(input), "\n")
	for y, row := range rows {
		for x, col := range row {
			if col == '^' {
				startingPos[0] = y
				startingPos[1] = x
			}
			if col == '#' {
				obstacles = append(obstacles, [2]int{y, x})
			}
		}
	}
	currentPos := [2]int{}
	currentPos[0], currentPos[1] = startingPos[0], startingPos[1]
	facing := [2]int{-1, 0}
	visited := [][2]int{startingPos}
	possibleObstructions := [][2]int{}
mLoop:
	for {
		if !slices.Contains(visited, currentPos) {
			visited = append(visited, currentPos)
		}
		// fmt.Println("visited:", visited)
		checkPos := [2]int{}
		checkPos[0] = currentPos[0] + facing[0]
		checkPos[1] = currentPos[1] + facing[1]
		// fmt.Println("checking", checkPos)
		if checkPos[0] >= len(rows[currentPos[1]]) || checkPos[0] < 0 || checkPos[1] >= len(rows) || checkPos[1] < 0 {
			break
		}

		for i := range obstacles {
			// fmt.Println("obstacle:", obstacles[i])
			if checkPos == obstacles[i] {
				// fmt.Println("facing:", facing)
				facing = changeFacing(facing)
				// fmt.Println("changed facing to", facing)
				continue mLoop
			}
		}
		// fmt.Println("update current position")
		if !slices.Contains(visited, checkPos) {
			visited = append(visited, checkPos)
		} else {
			// part 2: check for added obstacles that would make a loop
			// if we're on a position that has been visited, that might be enough?
			possibleObstructions = append(possibleObstructions, checkPos)
		}
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
