package main

import (
	"fmt"
	"os"

	"github.com/yeldiRium/advent-of-code/2024/day4"
)

func main() {
	inputFile, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	grid, err := day4.ParseInput(inputFile)
	if err != nil {
		panic(err)
	}

	xmasCount := 0
	for coordinate := range grid.X3() {
		at := func(x, y int) rune {
			return grid.RuneAt(coordinate.X + uint(x), coordinate.Y + uint(y))
		}
		if at(1, 1) == 'A' && ((at(0, 0) == 'M' && at(2, 2) == 'S') || (at(0, 0) == 'S' && at(2, 2) == 'M')) && ((at(0, 2) == 'M' && at(2, 0) == 'S') || (at(0, 2) == 'S' && at(2, 0) == 'M')) {
			xmasCount += 1
		}
	}

	fmt.Printf("%d", xmasCount)
}
