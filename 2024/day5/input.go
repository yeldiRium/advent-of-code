package day5

import (
	"io"
	"strconv"
	"strings"
)

func ParseInput(input io.Reader) (RuleBook, [][]int, error) {
	rawInput, err := io.ReadAll(input)
	if err != nil {
		return nil, nil, err
	}

	ruleBook := make(RuleBook, 0)
	elements := make([][]int, 0)

	lines := strings.Split(string(rawInput), "\n")

	index := 0
	for {
		line := lines[index]
		index += 1

		// The rules and elements are separated by an empty line.
		if line == "" {
			break
		}

		parts := strings.SplitN(line, "|", 2)
		earlier, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, nil, err
		}
		later, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, nil, err
		}
		ruleBook = append(ruleBook, OrderRule{ Earlier: earlier, Later: later })
	}
	for {
		if index == len(lines) {
			break
		}

		line := lines[index]
		index += 1

		// There might be a newline at the end, which makes the last line empty.
		if line == "" {
			break
		}

		parsedLine, err := parseLine(line)
		if err != nil {
			return nil, nil, err
		}

		elements = append(elements, parsedLine)
	}

	return ruleBook, elements, nil
}

func parseLine(line string) ([]int, error) {
	segments := strings.Split(line, ",")
	parsedSegments := make([]int, len(segments))

	for i, segment := range segments {
		parsedSegment, err := strconv.Atoi(segment)
		if err != nil {
			return nil, err
		}
		parsedSegments[i] = parsedSegment
	}

	return parsedSegments, nil
}
