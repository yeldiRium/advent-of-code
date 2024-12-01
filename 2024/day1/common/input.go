package common

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

func ParseInputLists(input io.Reader) (leftList, rightList []int, err error) {
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
