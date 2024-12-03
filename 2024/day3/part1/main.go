package main

import (
	"fmt"
	"io"
	"os"

	"github.com/yeldiRium/advent-of-code/2024/day3"
)

func main() {
	inputFile, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	inputText, err := io.ReadAll(inputFile)
	if err != nil {
		panic(err)
	}

	instructions := day3.ParseInstructions(string(inputText))
	multiplicationInstructions := day3.GetAllMultiplicationInstructions(instructions)
	sum := day3.SumProducts(multiplicationInstructions)

	fmt.Printf("%d", sum)
}