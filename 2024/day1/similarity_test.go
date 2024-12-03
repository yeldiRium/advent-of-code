package day1_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yeldiRium/advent-of-code/2024/day1"
)

func TestCountOccurences(t *testing.T) {
	t.Run("returns an empty map if the input is empty.", func(t *testing.T) {
		result := day1.CountOccurences([]int{})

		assert.Equal(t, map[int]int{}, result)
	})

	t.Run("returns correct results for various inputs.", func(t *testing.T) {
		type testCase struct {
			input          []int
			expectedCounts map[int]int
		}

		testCases := []testCase{
			{
				input:          []int{5, 1, 2, 2},
				expectedCounts: map[int]int{1: 1, 2: 2, 5: 1},
			},
		}
		for i, testCase := range testCases {
			t.Run(fmt.Sprintf("Test case %d", i), func(t *testing.T) {
				result := day1.CountOccurences(testCase.input)

				assert.Equal(t, testCase.expectedCounts, result)
			})
		}
	})
}

func TestSimilarity(t *testing.T) {
	t.Run("returns correct results for various inputs.", func(t *testing.T) {
		type testCase struct {
			leftList []int
			rightList []int
			expectedSimilarity int
		}

		testCases := []testCase{
			{
				leftList: []int{3, 4, 2, 1, 3, 3},
				rightList: []int{4, 3, 5, 3, 9, 3},
				expectedSimilarity: 31,
			},
		}
		for i, testCase := range testCases {
			t.Run(fmt.Sprintf("Test case %d", i), func(t *testing.T) {
				result := day1.Similarity(testCase.leftList, testCase.rightList)

				assert.Equal(t, testCase.expectedSimilarity, result)
			})
		}
	})

	t.Run("panics if the input lists have different lengths.", func(t *testing.T) {
		assert.Panics(t, func() {
			day1.Similarity([]int{1, 2}, []int{1, 2, 3})
		})
	})
}
