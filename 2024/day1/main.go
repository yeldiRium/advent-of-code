package main

import (
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

func compareLocationIdLists(leftList, rightList []int) int {
	if len(leftList) != len(rightList) {
		panic("input lists must have the same length")
	}

	slices.Sort(leftList)
	slices.Sort(rightList)

	distance := 0
	for i, left := range leftList {
		right := rightList[i]

		difference := left - right
		if difference < 0 {
			difference = -difference
		}

		distance += difference
	}

	return distance
}

func parseInputLists(input io.Reader) (leftList, rightList []int, err error) {
	rawLists, err := io.ReadAll(input)
	if err != nil {
		return nil, nil, err
	}

	inputWithoutTrailingNewline := strings.TrimSpace(string(rawLists))
	lines := strings.Split(inputWithoutTrailingNewline, "\n")

	leftList = make([]int, len(lines))
	rightList = make([]int, len(lines))

	for i, line := range lines {
		if len(line) == 0 {
			break
		}

		parts := strings.Split(line, "   ")
		left, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, nil, fmt.Errorf("failed to parse left value in line %d: %w", i, err)
		}
		leftList[i] = left

		right, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, nil, fmt.Errorf("failed to parse right value in line %d: %w", i, err)
		}
		rightList[i] = right
	}

	return leftList, rightList, nil
}

func main() {
	inputFile, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	leftList, rightList, err := parseInputLists(inputFile)
	if err != nil {
		panic(err)
	}

	distance := compareLocationIdLists(leftList, rightList)

	fmt.Printf("%d", distance)
}
