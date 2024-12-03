package day3_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yeldiRium/advent-of-code/2024/day3"
)

func TestParseInstructions(t *testing.T) {
	t.Run("correctly identifies mul/do/don't instructions and discards the rest.", func(t *testing.T) {
		type testCase struct {
			input                string
			expectedInstructions []interface{}
		}

		testCases := []testCase{
			{
				input:                "",
				expectedInstructions: []interface{}{},
			},
			{
				input: "mul(5,17)",
				expectedInstructions: []interface{}{
					day3.MulInstruction{X: 5, Y: 17},
				},
			},
			{
				input: "mul(5, 17)",
				expectedInstructions: []interface{}{},
			},
			{
				input: "mul(5,17",
				expectedInstructions: []interface{}{},
			},
			{
				input: "mul(5,17)\nexqlmulnqex98mul(2,8)8392e:(}=‚mul[517]mul(12)89e32nmul(\n15,200)",
				expectedInstructions: []interface{}{
					day3.MulInstruction{X:5, Y:17},
					day3.MulInstruction{X:2, Y:8},
				},
			},
			{
				input: "mul(5,1017)",
				expectedInstructions: []interface{}{},
			},
			{
				input: "do()don't()mul(5,8)",
				expectedInstructions: []interface{}{
					day3.DoInstruction{},
					day3.DontInstruction{},
					day3.MulInstruction{
						X: 5, Y: 8,
					},
				},
			},
		}
		for i, testCase := range testCases {
			t.Run(fmt.Sprintf("Test case %d:", i), func(t *testing.T) {
				result := day3.ParseInstructions(testCase.input)
				assert.Equal(t, testCase.expectedInstructions, result)
			})
		}
	})
}

func TestGetAllMultiplicationInstructions(t *testing.T) {
	t.Run("returns all contained multiplications instructions.", func(t *testing.T) {
		input := []interface{}{
			day3.MulInstruction{
				X: 5, Y: 7,
			},
			day3.DoInstruction{},
			day3.MulInstruction{
				X: 312, Y: 18,
			},
			day3.DontInstruction{},
			day3.DontInstruction{},
			day3.DoInstruction{},
			day3.MulInstruction{
				X: 312, Y: 18,
			},
		}

		result := day3.GetAllMultiplicationInstructions(input)
		assert.Equal(t, []day3.MulInstruction{
			day3.MulInstruction{
				X: 5, Y: 7,
			},
			day3.MulInstruction{
				X: 312, Y: 18,
			},
			day3.MulInstruction{
				X: 312, Y: 18,
			},
		}, result)
	})
}

func TestGetAllEnabledMultiplicationInstructions(t *testing.T) {
	t.Run("returns all contained multiplication instructions that are enable - i.e. they do not follow a dont instruction without another do instruction.", func(t *testing.T) {
		input := []interface{}{
			day3.MulInstruction{
				X: 2, Y: 4,
			},
			day3.DontInstruction{},
			day3.MulInstruction{
				X: 5, Y: 5,
			},
			day3.MulInstruction{
				X: 11, Y: 8,
			},
			day3.DoInstruction{},
			day3.MulInstruction{
				X: 8, Y: 5,
			},
		}

		result := day3.GetAllEnabledMultiplicationInstructions(input)
		assert.Equal(t, []day3.MulInstruction{
			day3.MulInstruction{
				X: 2, Y: 4,
			},
			day3.MulInstruction{
				X: 8, Y: 5,
			},
		}, result)
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
