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

	xmasCounter := 0
	for window := range grid.Windows(4) {
		stringInWindow := ""
		for _, coordinate := range window {
			stringInWindow += string(grid.RuneAt(coordinate.X, coordinate.Y))
		}
		if stringInWindow == "XMAS" || stringInWindow == "SAMX" {
			xmasCounter += 1
		}
	}

	fmt.Printf("%d", xmasCounter)
}
