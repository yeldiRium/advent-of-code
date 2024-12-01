package main

import (
	"fmt"
	"os"

	"github.com/yeldiRium/advent-of-code/2024/day1/common"
)


func main() {
	inputFile, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	leftList, rightList, err := common.ParseInputLists(inputFile)
	if err != nil {
		panic(err)
	}

	distance := common.CompareLocationIdLists(leftList, rightList)

	fmt.Printf("%d", distance)
}
