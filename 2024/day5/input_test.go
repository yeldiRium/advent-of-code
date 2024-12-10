package day5_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yeldiRium/advent-of-code/2024/day5"
)

func TestParseInput(t *testing.T) {
	t.Run("parses valid input correctly.", func(t *testing.T) {
		input := `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47
`

		expectedRuleBook := day5.RuleBook{
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
		expectedElements := [][]int{
			{75, 47, 61, 53, 29},
			{97, 61, 53, 29, 13},
			{75, 29, 13},
			{75, 97, 47, 61, 53},
			{61, 13, 29},
			{97, 13, 75, 29, 47},
		}

		actualRuleBook, actualElements, err := day5.ParseInput(bytes.NewBuffer([]byte(input)))

		assert.NoError(t, err)
		assert.Equal(t, expectedRuleBook, actualRuleBook)
		assert.Equal(t, expectedElements, actualElements)
	})
}
