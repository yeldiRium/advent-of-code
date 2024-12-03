package day2_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yeldiRium/advent-of-code/2024/day2"
)

func TestReportSafety(t *testing.T) {
	t.Run("succeeds for various examples.", func(t *testing.T) {
		type testCase struct {
			report         day2.Report
			expectedResult bool
		}

		testCases := []testCase{
			{
				report:         day2.Report{1, 2, 3, 4, 5},
				expectedResult: true,
			},
			{
				report:         day2.Report{1, 3, 2},
				expectedResult: false,
			},
			{
				report:         day2.Report{1, 6},
				expectedResult: false,
			},
			{
				report:         day2.Report{5, 1},
				expectedResult: false,
			},
			{
				report:         day2.Report{1, 3, 4, 6},
				expectedResult: true,
			},
			{
				report:         day2.Report{7, 6, 4, 2, 1},
				expectedResult: true,
			},
			{
				report:         day2.Report{1, 2, 7, 8, 9},
				expectedResult: false,
			},
			{
				report:         day2.Report{9, 7, 6, 2, 1},
				expectedResult: false,
			},
			{
				report:         day2.Report{1, 3, 2, 4, 5},
				expectedResult: false,
			},
			{
				report:         day2.Report{8, 6, 4, 4, 1},
				expectedResult: false,
			},
			{
				report:         day2.Report{1, 3, 6, 7, 9},
				expectedResult: true,
			},
		}

		for i, testCase := range testCases {
			t.Run(fmt.Sprintf("Test case %d:", i), func(t *testing.T) {
				result := day2.ReportSafety(testCase.report)

				assert.Equal(t, testCase.expectedResult, result)
			})
		}
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
