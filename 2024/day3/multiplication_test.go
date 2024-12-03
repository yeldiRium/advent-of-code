package day3_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yeldiRium/advent-of-code/2024/day3"
)

func TestParseMulInstructions(t *testing.T) {
	t.Run("correctly identifies mul instructions and discards the rest.", func(t *testing.T) {
		type testCase struct {
			input                string
			expectedInstructions []day3.MulInstruction
		}

		testCases := []testCase{
			{
				input:                "",
				expectedInstructions: []day3.MulInstruction{},
			},
			{
				input: "mul(5,17)",
				expectedInstructions: []day3.MulInstruction{
					{X: 5, Y: 17},
				},
			},
			{
				input: "mul(5, 17)",
				expectedInstructions: []day3.MulInstruction{},
			},
			{
				input: "mul(5,17",
				expectedInstructions: []day3.MulInstruction{},
			},
			{
				input: "mul(5,17)\nexqlmulnqex98mul(2,8)8392e:(}=‚mul[517]mul(12)89e32nmul(\n15,200)",
				expectedInstructions: []day3.MulInstruction{
					{X:5, Y:17},
					{X:2, Y:8},
				},
			},
			{
				input: "mul(5,1017)",
				expectedInstructions: []day3.MulInstruction{},
			},
		}
		for i, testCase := range testCases {
			t.Run(fmt.Sprintf("Test case %d:", i), func(t *testing.T) {
				result := day3.ParseMulInstructions(testCase.input)
				assert.Equal(t, testCase.expectedInstructions, result)
			})
		}
	})
}

func TestSumProducts(t *testing.T) {
	t.Run("calculates the sum of the products of multiplication instructions.", func(t *testing.T) {
		type testCase struct {
			instructions   []day3.MulInstruction
			expectedResult int
		}

		testCases := []testCase{
			{
				instructions:   []day3.MulInstruction{},
				expectedResult: 0,
			},
			{
				instructions: []day3.MulInstruction{
					{X: 1, Y: 1},
				},
				expectedResult: 1,
			},
			{
				instructions: []day3.MulInstruction{
					{X: 205, Y: 12},
					{X: 33, Y: 777},
					{X: 100, Y: 100},
				},
				expectedResult: 38_101,
			},
		}
		for i, testCase := range testCases {
			t.Run(fmt.Sprintf("Test case %d:", i), func(t *testing.T) {
				result := day3.SumProducts(testCase.instructions)

				assert.Equal(t, testCase.expectedResult, result)
			})
		}
	})
}
