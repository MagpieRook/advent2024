package main

import (
	"os"
	"strings"
)

func calculateDistance(start, end [2]int) [2]int {
	dist := [2]int{start[0] - end[0], start[1] - end[1]}
	return dist
}

func Day8(filename string) (int, int) {
	if filename == "" {
		filename = "day8input.txt"
	}
	// input is a grid of "." characters and any digits/letters
	// non-period digits represent antennas of a certain frequency
	input, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	// antennas is a map of frequencies to locations
	antennas := map[rune][][2]int{}
	rows := strings.Split(string(input), "\n")
	width := len(rows[0])
	height := len(rows)
	for y, row := range rows {
		for x, val := range row {
			if val != '.' {
				antennas[val] = append(antennas[val], [2]int{y, x})
			}
		}
	}

	// part 1: if any two antennas with the same frequency (letter)
	// are 1x and 2x (x being some distance in any direction, including diagonals)
	// they create an "antinode".
	// how many unique locations (possibly including antenna locations) have antinodes?
	antinodes := map[[2]int]bool{}
	for _, locations := range antennas {
		if len(locations) < 2 {
			continue
		}
		for i := range len(locations) - 1 {
			for j := i + 1; j < len(locations); j++ {
				start := locations[i]
				end := locations[j]

				// brute force: for every location in the grid, including antennas
				// check against every pair of same-frequency antennas
				// if they are the same distance away, log as an antinode
				for y := range height {
					for x := range width {
						antinode := [2]int{y, x}
						if antinodes[antinode] {
							continue
						}
						distStart := calculateDistance(start, antinode)
						doubleDistStart := [2]int{2 * distStart[0], 2 * distStart[1]}
						distEnd := calculateDistance(end, antinode)
						doubleDistEnd := [2]int{2 * distEnd[0], 2 * distEnd[1]}
						if distStart == doubleDistEnd || distEnd == doubleDistStart {
							antinodes[antinode] = true
						}
					}
				}
			}
		}
	}

	// part 2: antinodes actually occur at all locations in line with 2 antennas
	antinodes2 := map[[2]int]bool{}
	for _, locations := range antennas {
		if len(locations) < 2 {
			continue
		}
		for i := range locations {
			for j := range locations {
				if i == j {
					continue
				}
				start := locations[i]
				end := locations[j]
				distance := calculateDistance(start, end)
				antinode := [2]int{end[0] + distance[0], end[1] + distance[1]}
				// the distance between antinodes can also be used as the slope of antinodes
				// since we're running through both positive and negative slopes (i.e. i = 4, j = 1 && i = 1, j = 4)
				// we just need to add the slope until we're out of bounds to find new antinodes
				for antinode[0] >= 0 && antinode[0] < height && antinode[1] >= 0 && antinode[1] < width {
					antinodes2[antinode] = true
					antinode = [2]int{antinode[0] + distance[0], antinode[1] + distance[1]}
				}
			}
		}
	}

	// antinodes should just be a list of keys with true values at this point, so count keys for answer
	return len(antinodes), len(antinodes2)
}
