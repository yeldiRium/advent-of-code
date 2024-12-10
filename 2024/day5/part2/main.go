package main

import (
	"fmt"
	"os"

	"github.com/yeldiRium/advent-of-code/2024/day5"
)

func main() {
	inputFile, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	ruleBook, updates, err := day5.ParseInput(inputFile)
	if err != nil {
		panic(err)
	}

	incorrectUpdates := make([][]int, 0)
	for _, update := range updates {
		if !day5.IsInOrder(update, ruleBook) {
			incorrectUpdates = append(incorrectUpdates, update)
		}
	}

	sortedUpdates := make([][]int, len(incorrectUpdates))
	for i, update := range incorrectUpdates {
		sortedUpdate, err := day5.Sort(update, ruleBook)
		if err != nil {
			panic(err)
		}
		sortedUpdates[i] = sortedUpdate
	}

	middleNumberSum := 0
	for _, update := range sortedUpdates {
		middleIndex := len(update) / 2
		middleNumberSum += update[middleIndex]
	}

	fmt.Printf("%d", middleNumberSum)
}
