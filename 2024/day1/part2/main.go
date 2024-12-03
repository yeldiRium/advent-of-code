package main

import (
	"fmt"
	"os"

	"github.com/yeldiRium/advent-of-code/2024/day1"
)


func main() {
	inputFile, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	leftList, rightList, err := day1.ParseInputLists(inputFile)
	if err != nil {
		panic(err)
	}

	similarity := day1.Similarity(leftList, rightList)

	fmt.Printf("%d", similarity)
}