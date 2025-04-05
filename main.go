package main

import (
	"flag"
	"fmt"
)

func main() {
	dayPtr := flag.Int("day", 0, "the day to run (1-25)")
	flag.Parse()
	if *dayPtr > 25 || *dayPtr <= 0 {
		fmt.Println("day flag must be an integer between 1 and 25 inclusive for the days of the advent calendar")
		return
	}

	switch *dayPtr {
	case 1:
		Day1()
	case 2:
		Day2()
	case 3:
		Day3()
	case 4:
		Day4()
	case 5:
		Day5()
	case 6:
		Day6()
	case 7:
		Day7()
	default:
		fmt.Println("Not completed")
	}
}
