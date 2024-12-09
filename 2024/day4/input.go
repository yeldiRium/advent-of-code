package day4

import (
	"io"
	"strings"
)

func ParseInput(input io.Reader) (Grid, error) {
	inputText, err := io.ReadAll(input)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(inputText), "\n")
	if len(lines) == 0 {
		return NewGrid(0, 0), nil
	}

	grid := NewGrid(uint(len(lines[0])), uint(len(lines)))
	for y, line := range lines {
		for x, char := range line {
			grid[x][y] = char
		}
	}

	return grid, nil
}
