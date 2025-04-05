package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Day3() {
	input, err := os.ReadFile("day3input.txt")
	if err != nil {
		panic(err)
	}

	// Part 2: remove all mul(x,y) statements between don't() and do()
	regexReplace, err := regexp.Compile(`don\'t\(\)[\s\S]*?do\(\)`)
	if err != nil {
		panic(err)
	}
	fmt.Println("Start:")
	fmt.Println(string(input))
	replacedInput := regexReplace.ReplaceAll(input, []byte("replaceme"))
	fmt.Println("Replaced:")
	fmt.Println(string(replacedInput))

	regexDontEnd, err := regexp.Compile(`don\'t\(\)`)
	if err != nil {
		panic(err)
	}
	inputLocation := regexDontEnd.FindIndex(replacedInput)
	var finalInput []byte
	if inputLocation != nil {
		finalInput = replacedInput[:inputLocation[0]]
	} else {
		finalInput = replacedInput
	}
	fmt.Println("Last Don't:")
	fmt.Println(string(finalInput))
	// Calculate only specifically-formatted mul(x,y) statements
	regex, err := regexp.Compile(`mul\((\d+,\d+)\)`)
	if err != nil {
		panic(err)
	}
	matches := regex.FindAll(finalInput, -1)
	total := int64(0)
	for _, match := range matches {
		nums := strings.Split(string(match), ",")
		nums[0] = strings.Replace(nums[0], "mul(", "", -1)
		nums[1] = strings.Replace(nums[1], ")", "", -1)
		left, err := strconv.ParseInt(nums[0], 10, 64)
		if err != nil {
			panic(err)
		}
		right, err := strconv.ParseInt(nums[1], 10, 64)
		if err != nil {
			panic(err)
		}
		total += left * right
	}

	fmt.Println(total)
}
