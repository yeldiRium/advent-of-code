package day2

import (
	"io"
	"strconv"
	"strings"
)

func ParseReports(input io.Reader) (reports []Report, err error) {
	rawLists, err := io.ReadAll(input)
	if err != nil {
		return nil, err
	}

	inputWithoutTrailingNewline := strings.TrimSpace(string(rawLists))
	lines := strings.Split(inputWithoutTrailingNewline, "\n")

	reports = make([]Report, len(lines))

	for i, line := range lines {
		if len(line) == 0 {
			break
		}

		parts := strings.Split(line, " ")
		report := make([]int, len(parts))

		for i, part := range parts {
			value, err := strconv.Atoi(part)
			if err != nil {
				return nil, err
			}

			report[i] = value
		}

		reports[i] = report
	}

	return reports, nil
}
