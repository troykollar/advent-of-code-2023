package main

import (
	days "advent-of-code-2023/days/day_1"
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	if len(args) < 2 {
		panic("--day is a required argument.")
	}

	day, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("--day is required and must be a number.")
	}

	switch day {
	case 1:
		days.Day1()
	default:
		panic(fmt.Sprintf("Day %v not yet implemented.", day))
	}
}
