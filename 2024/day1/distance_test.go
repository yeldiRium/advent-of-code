package day1_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/yeldiRium/advent-of-code/2024/day1"
)

func TestDistance(t *testing.T) {
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
				actualDistance := day1.Distance(testCase.leftList, testCase.rightList)

				assert.Equal(t, testCase.expectedDistance, actualDistance)
			})
		}
	})

	t.Run("panics if the input lists aren't the same length.", func(t *testing.T) {
		assert.Panics(t, func() {
			day1.Distance([]int{1, 2}, []int{1, 2, 3})
		})
	})
}
