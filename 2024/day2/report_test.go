package day2_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yeldiRium/advent-of-code/2024/day2"
)

func TestCountUnsafeLevels(t *testing.T) {
	t.Run("succeeds for various examples.", func(t *testing.T) {
		type testCase struct {
			report         day2.Report
			expectedResult int
		}

		testCases := []testCase{
			{
				report:         day2.Report{1, 2, 3, 4, 5},
				expectedResult: 0,
			},
			{
				report:         day2.Report{1, 3, 2},
				expectedResult: 1,
			},
			{
				report:         day2.Report{1, 6},
				expectedResult: 1,
			},
			{
				report:         day2.Report{5, 1},
				expectedResult: 1,
			},
			{
				report:         day2.Report{1, 3, 4, 6},
				expectedResult: 0,
			},
			{
				report:         day2.Report{7, 6, 4, 2, 1},
				expectedResult: 0,
			},
			{
				report:         day2.Report{1, 2, 7, 8, 9},
				expectedResult: 1,
			},
			{
				report:         day2.Report{9, 7, 6, 2, 1},
				expectedResult: 1,
			},
			{
				report:         day2.Report{1, 3, 2, 4, 5},
				expectedResult: 1,
			},
			{
				report:         day2.Report{8, 6, 4, 4, 1},
				expectedResult: 1,
			},
			{
				report:         day2.Report{1, 3, 6, 7, 9},
				expectedResult: 0,
			},
		}

		for i, testCase := range testCases {
			t.Run(fmt.Sprintf("Test case %d:", i), func(t *testing.T) {
				result := day2.CountUnsafeLevels(testCase.report)

				assert.Equal(t, testCase.expectedResult, result)
			})
		}
	})
}

func TestReportSafety(t *testing.T) {
	t.Run("reports unsafe test for one unsafe level.", func(t *testing.T) {
		report := day2.Report{1, 6}

		isSafe := day2.ReportSafety(report)

		assert.False(t, isSafe)
	})

	t.Run("reports safe test for zero unsafe levels.", func(t *testing.T) {
		report := day2.Report{1, 3, 5, 7}

		isSafe := day2.ReportSafety(report)

		assert.True(t, isSafe)
	})
}

func TestCountSafeReports(t *testing.T) {
	t.Run("succeeds for various examples.", func(t *testing.T) {
		type testCase struct {
			reports        []day2.Report
			expectedResult int
		}

		testCases := []testCase{
			{
				reports:        []day2.Report{},
				expectedResult: 0,
			},
			{
				reports: []day2.Report{
					{7, 6, 4, 2, 1},
					{1, 2, 7, 8, 9},
					{9, 7, 6, 2, 1},
					{1, 3, 2, 4, 5},
					{8, 6, 4, 4, 1},
					{1, 3, 6, 7, 9},
				},
				expectedResult: 2,
			},
		}

		for i, testCase := range testCases {
			t.Run(fmt.Sprintf("Test case %d:", i), func(t *testing.T) {
				result := day2.CountSafeReports(testCase.reports)

				assert.Equal(t, testCase.expectedResult, result)
			})
		}
	})
}
