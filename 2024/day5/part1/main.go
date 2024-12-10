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

	correctUpdates := make([][]int, 0)

	for _, update := range updates {
		if day5.IsInOrder(update, ruleBook) {
			correctUpdates = append(correctUpdates, update)
		}
	}

	middleNumberSum := 0
	for _, update := range correctUpdates {
		middleIndex := len(update) / 2
		middleNumberSum += update[middleIndex]
	}

	fmt.Printf("%d", middleNumberSum)
}
