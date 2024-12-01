package main

import (
	"bufio"
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompareLocationIdLists(t *testing.T) {
	t.Run("succeeds for various inputs", func(t *testing.T) {
		type testCase struct {
			leftList         []int
			rightList        []int
			expectedDistance int
		}

		testCases := []testCase{
			{
				leftList:         []int{},
				rightList:        []int{},
				expectedDistance: 0,
			},
			{
				leftList:         []int{3, 4, 2, 1, 3, 3},
				rightList:        []int{4, 3, 5, 3, 9, 3},
				expectedDistance: 11,
			},
		}
		for i, testCase := range testCases {
			t.Run(fmt.Sprintf("Test case %d", i), func(t *testing.T) {
				actualDistance := compareLocationIdLists(testCase.leftList, testCase.rightList)

				assert.Equal(t, testCase.expectedDistance, actualDistance)
			})
		}
	})

	t.Run("panics if the input lists aren't the same length.", func(t *testing.T) {
		assert.Panics(t, func() {
			compareLocationIdLists([]int{1, 2}, []int{1, 2, 3})
		})
	})
}

func TestParseInputLists(t *testing.T) {
	t.Run("succeeds for a valid input file.", func(t *testing.T) {
		inputFile := bytes.NewBuffer([]byte("1   2\n3   4\n"))

		leftList, rightList, err := parseInputLists(bufio.NewReader(inputFile))
		assert.NoError(t, err)
		assert.Equal(t, []int{1, 3}, leftList)
		assert.Equal(t, []int{2, 4}, rightList)
	})
}
