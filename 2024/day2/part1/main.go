package main

import (
	"fmt"
	"os"

	"github.com/yeldiRium/advent-of-code/2024/day2"
)

func main() {
	inputFile, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	reports, err := day2.ParseReports(inputFile)
	if err != nil {
		panic(err)
	}

	count := day2.CountSafeReports(reports)
	fmt.Printf("%d", count)
}
