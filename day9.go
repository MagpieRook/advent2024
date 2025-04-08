package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
)

// func printFileSystem(filesystem []int) {
// 	for i := range filesystem {
// 		if filesystem[i] == -1 {
// 			fmt.Print(".")
// 		} else {
// 			fmt.Print(filesystem[i])
// 		}
// 	}
// 	fmt.Print("\n")
// }

func fs2ToFS1(filesystem2 [][2]int) []int {
	filesystem := []int{}
	for i := range filesystem2 {
		fIndex, num := filesystem2[i][0], filesystem2[i][1]
		for range num {
			filesystem = append(filesystem, fIndex)
		}
	}
	return filesystem
}

// instead of returning a list that mimics the storage
// returns an ordered list of (index, len) pairs, with -1 being free space
func parseInput2(input string) [][2]int {
	fIndex := 0
	filesystem := [][2]int{}
	nextFile := true
	for i := range input {
		num, err := strconv.ParseInt(string(input[i]), 10, 0)
		if err != nil {
			panic(err)
		}
		var newInt int
		if nextFile {
			// parse the next file index
			newInt = fIndex
			fIndex++
		} else {
			// parse the free space
			newInt = -1
		}
		if num > 0 {
			filesystem = append(filesystem, [2]int{newInt, int(num)})
		}
		// flip from file to free
		nextFile = !nextFile
	}
	return filesystem
}

// each file also has a 0-indexed ID number based on the order of the files
// so a 1 block file that appears first would be represented as "0"
// if the next file is 3 blocks, it would be "111" (etc.)
// returns an ordered list of ints representing what block of file is there (-1 is free space)
func parseInput(input string) []int {
	fIndex := 0
	filesystem := []int{}
	nextFile := true
	for i := range input {
		num, err := strconv.ParseInt(string(input[i]), 10, 0)
		if err != nil {
			panic(err)
		}
		var newInt int
		if nextFile {
			// parse the next file index
			newInt = fIndex
			fIndex++
		} else {
			// parse the free space
			newInt = -1
		}
		// add number of characters based on size
		for range num {
			filesystem = append(filesystem, newInt)
		}
		// flip from file to free
		nextFile = !nextFile
	}
	return filesystem
}

func Day9() {
	// input is a series of integers in one line
	// integers represent blocks of space used for files and free space
	// first is blocks of file, then blocks of free, then file, then free, etc.
	input, err := os.ReadFile("day9input.txt")
	if err != nil {
		panic(err)
	}

	filesystem := parseInput(string(input))

	// part 1
	// compact the file by moving file blocks into free space
	// move one block of file at a time from end to beginning
	// replace -1 with index
	for i := len(filesystem) - 1; i > 0; i-- {
		if filesystem[i] != -1 {
			newIndex := slices.Index(filesystem, -1)
			if newIndex == -1 || newIndex > i {
				break
			}
			filesystem[newIndex], filesystem[i] = filesystem[i], filesystem[newIndex]
		}
	}

	// calculate filesystem checksum
	// multiply each position with index (position is 0-indexed) (skip -1)
	checksum := 0
	for i := range filesystem {
		fIndex := filesystem[i]
		if fIndex == -1 {
			break
		}
		checksum += fIndex * i
	}
	fmt.Println(checksum)

	// part 2
	// file system fragmentation is bad, instead move entire files
	// starting at the highest file id, move each file to a free space
	// free space must be able to fit the file
	filesystem2 := parseInput2(string(input))
	// printFileSystem(fs2ToFS1(filesystem2))
	for i := len(filesystem2) - 1; i >= 0; i-- {
		f := filesystem2[i]
		fIndex, fLength := f[0], f[1]
		if fIndex == -1 {
			continue
		}
		for j := range i {
			if filesystem2[j][0] == -1 && filesystem2[j][1] >= fLength {
				filesystem2[j] = [2]int{-1, filesystem2[j][1] - fLength}
				// add space to nearby free space, if it exists
				// then remove the moved file
				if i+1 < len(filesystem2) && filesystem2[i][0] == -1 {
					filesystem2[i+1] = [2]int{-1, filesystem2[i+1][1] + fLength}
					filesystem2 = slices.Delete(filesystem2, i, i+1)
				} else if filesystem2[i-1][0] == -1 {
					filesystem2[i-1] = [2]int{-1, filesystem2[i-1][1] + fLength}
					filesystem2 = slices.Delete(filesystem2, i, i+1)
				} else {
					// otherwise, replace file with new free space
					filesystem2[i] = [2]int{-1, fLength}
				}
				// add file to found free space
				filesystem2 = slices.Insert(filesystem2, j, f)
				// printFileSystem(fs2ToFS1(filesystem2))
				// in case we deleted where we were, check again
				i++
				break
			}
		}
	}

	// calculate filesystem checksum
	// multiply each position with index (position is 0-indexed) (skip -1)
	checksum2 := 0
	fs2 := fs2ToFS1(filesystem2)
	for i := range fs2 {
		fIndex := fs2[i]
		if fIndex == -1 {
			continue
		}
		checksum2 += fIndex * i
	}

	fmt.Println(checksum2)
}
