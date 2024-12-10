package day5_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yeldiRium/advent-of-code/2024/day5"
)

func TestIsInOrder(t *testing.T) {
	commonRuleBook := day5.RuleBook{
		{Earlier: 47, Later: 53},
		{Earlier: 97, Later: 13},
		{Earlier: 97, Later: 61},
		{Earlier: 97, Later: 47},
		{Earlier: 75, Later: 29},
		{Earlier: 61, Later: 13},
		{Earlier: 75, Later: 53},
		{Earlier: 29, Later: 13},
		{Earlier: 97, Later: 29},
		{Earlier: 53, Later: 29},
		{Earlier: 61, Later: 53},
		{Earlier: 97, Later: 53},
		{Earlier: 61, Later: 29},
		{Earlier: 47, Later: 13},
		{Earlier: 75, Later: 47},
		{Earlier: 97, Later: 75},
		{Earlier: 47, Later: 61},
		{Earlier: 75, Later: 61},
		{Earlier: 47, Later: 29},
		{Earlier: 75, Later: 13},
		{Earlier: 53, Later: 13},
	}

	t.Run("returns true for an empty list of elements.", func(t *testing.T) {
		elements := []int{}
		ruleBook := day5.RuleBook{
			{Earlier: 5, Later: 7},
			{Earlier: 3, Later: 2},
		}

		result := day5.IsInOrder(elements, ruleBook)

		assert.True(t, result)
	})

	t.Run("returns true for an empty rule book.", func(t *testing.T) {
		elements := []int{1, 5, 3}
		ruleBook := day5.RuleBook{}

		result := day5.IsInOrder(elements, ruleBook)

		assert.True(t, result)
	})

	t.Run("returns true if all elements adhere to the ordering rules.", func(t *testing.T) {
		type testCase struct {
			elements []int
			ruleBook day5.RuleBook
		}

		testCases := []testCase{
			{
				elements: []int{75, 47, 61, 53, 29},
				ruleBook: commonRuleBook,
			},
			{
				elements: []int{97, 61, 53, 29, 13},
				ruleBook: commonRuleBook,
			},
			{
				elements: []int{75, 29, 13},
				ruleBook: commonRuleBook,
			},
		}
		for i, testCase := range testCases {
			t.Run(fmt.Sprintf("Test case %d should be true but is not:", i), func(t *testing.T) {
				result := day5.IsInOrder(testCase.elements, testCase.ruleBook)
				assert.True(t, result)
			})
		}
	})

	t.Run("returns false if some element does not adhere to the ordering rules.", func(t *testing.T) {
		type testCase struct {
			elements []int
			ruleBook day5.RuleBook
		}

		testCases := []testCase{
			{
				elements: []int{75, 97, 47, 61, 53},
				ruleBook: commonRuleBook,
			},
			{
				elements: []int{97, 13, 75, 29, 47},
				ruleBook: commonRuleBook,
			},
		}
		for i, testCase := range testCases {
			t.Run(fmt.Sprintf("Test case %d should be false but is not:", i), func(t *testing.T) {
				result := day5.IsInOrder(testCase.elements, testCase.ruleBook)
				assert.False(t, result)
			})
		}
	})
}
