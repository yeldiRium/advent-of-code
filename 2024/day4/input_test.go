package day4_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yeldiRium/advent-of-code/2024/day4"
)

func TestParseInput(t *testing.T) {
	t.Run("parses an input file correctly.", func(t *testing.T) {
		type testCase struct {
			input        string
			expectedGrid day4.Grid
		}

		testCases := []testCase{
			{
				input:        "",
				expectedGrid: day4.NewGrid(0, 0),
			},
			{
				input: "ABC",
				expectedGrid: day4.Grid{
					{'A'}, {'B'}, {'C'},
				},
			},
			{
				input: "AB\nCD\nEF",
				expectedGrid: day4.Grid{
					{'A', 'C', 'E'}, {'B', 'D', 'F'},
				},
			},
		}
		for i, testCase := range testCases {
			t.Run(fmt.Sprintf("Test case: %d", i), func(t *testing.T) {
				inputReader := bytes.NewBuffer([]byte(testCase.input))
				actualGrid, err := day4.ParseInput(inputReader)

				assert.NoError(t, err)
				assert.Equal(t, testCase.expectedGrid, actualGrid)
			})
		}
	})
}
