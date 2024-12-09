package day4_test

import (
	"fmt"
	"iter"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/yeldiRium/advent-of-code/2024/day4"
)

func collectWindows(iterator iter.Seq[[]day4.Coordinate]) [][]day4.Coordinate {
	windows := make([][]day4.Coordinate, 0)
	for window := range iterator {
		windows = append(windows, window)
	}
	return windows
}

func TestGridWindows(t *testing.T) {
	t.Run("generates no possible windows in a 0x0 grid.", func(t *testing.T) {
		grid := day4.NewGrid(0, 0)

		windows := collectWindows(grid.Windows(3))

		assert.Equal(t, [][]day4.Coordinate{}, windows)
	})

	t.Run("generates all possible windows.", func(t *testing.T) {
		type testCase struct {
			gridWidth       uint
			gridHeight      uint
			windowSize      uint
			expectedWindows [][]day4.Coordinate
		}

		testCases := []testCase{
			{
				gridWidth:   3,
				gridHeight: 4,
				windowSize: 3,
				expectedWindows: [][]day4.Coordinate{
					// horizontal
					{
						{X: 0, Y: 0},
						{X: 1, Y: 0},
						{X: 2, Y: 0},
					},
					{
						{X: 0, Y: 1},
						{X: 1, Y: 1},
						{X: 2, Y: 1},
					},
					{
						{X: 0, Y: 2},
						{X: 1, Y: 2},
						{X: 2, Y: 2},
					},
					{
						{X: 0, Y: 3},
						{X: 1, Y: 3},
						{X: 2, Y: 3},
					},
					// vertical
					{
						{X: 0, Y: 0},
						{X: 0, Y: 1},
						{X: 0, Y: 2},
					},
					{
						{X: 1, Y: 0},
						{X: 1, Y: 1},
						{X: 1, Y: 2},
					},
					{
						{X: 2, Y: 0},
						{X: 2, Y: 1},
						{X: 2, Y: 2},
					},
					{
						{X: 0, Y: 1},
						{X: 0, Y: 2},
						{X: 0, Y: 3},
					},
					{
						{X: 1, Y: 1},
						{X: 1, Y: 2},
						{X: 1, Y: 3},
					},
					{
						{X: 2, Y: 1},
						{X: 2, Y: 2},
						{X: 2, Y: 3},
					},
					// diagonal SE
					{
						{X: 0, Y: 0},
						{X: 1, Y: 1},
						{X: 2, Y: 2},
					},
					{
						{X: 0, Y: 1},
						{X: 1, Y: 2},
						{X: 2, Y: 3},
					},
					// diagonal SW
					{
						{X: 2, Y: 0},
						{X: 1, Y: 1},
						{X: 0, Y: 2},
					},
					{
						{X: 2, Y: 1},
						{X: 1, Y: 2},
						{X: 0, Y: 3},
					},
				},
			},
		}
		for i, testCase := range testCases {
			t.Run(fmt.Sprintf("Test case: %d", i), func(t *testing.T) {
				grid := day4.NewGrid(testCase.gridWidth, testCase.gridHeight)

				windows := collectWindows(grid.Windows(testCase.windowSize))

				assert.Equal(t, testCase.expectedWindows, windows)
			})
		}
	})
}

func TestGridHorizontalWindows(t *testing.T) {
	t.Run("generates no possible windows in a 0x0 grid.", func(t *testing.T) {
		grid := day4.NewGrid(0, 0)

		windows := collectWindows(grid.HorizontalWindows(3))

		assert.Equal(t, [][]day4.Coordinate{}, windows)
	})

	t.Run("generates one possible window of size 3 within a 3x1 grid.", func(t *testing.T) {
		grid := day4.NewGrid(3, 1)

		windows := collectWindows(grid.HorizontalWindows(3))

		assert.Equal(t, [][]day4.Coordinate{
			{
				{X: 0, Y: 0},
				{X: 1, Y: 0},
				{X: 2, Y: 0},
			},
		}, windows)
	})

	t.Run("generates 4 possible windows of size 4 within a 5x2 grid.", func(t *testing.T) {
		grid := day4.NewGrid(5, 2)

		windows := collectWindows(grid.HorizontalWindows(4))

		assert.Equal(t, [][]day4.Coordinate{
			{
				{X: 0, Y: 0},
				{X: 1, Y: 0},
				{X: 2, Y: 0},
				{X: 3, Y: 0},
			},
			{
				{X: 1, Y: 0},
				{X: 2, Y: 0},
				{X: 3, Y: 0},
				{X: 4, Y: 0},
			},
			{
				{X: 0, Y: 1},
				{X: 1, Y: 1},
				{X: 2, Y: 1},
				{X: 3, Y: 1},
			},
			{
				{X: 1, Y: 1},
				{X: 2, Y: 1},
				{X: 3, Y: 1},
				{X: 4, Y: 1},
			},
		}, windows)
	})
}

func TestGridVerticalWindows(t *testing.T) {
	t.Run("generates no possible windows in a 0x0 grid.", func(t *testing.T) {
		grid := day4.NewGrid(0, 0)

		windows := collectWindows(grid.VerticalWindows(3))

		assert.Equal(t, [][]day4.Coordinate{}, windows)
	})

	t.Run("generates one possible window of size 3 within a 1x3 grid.", func(t *testing.T) {
		grid := day4.NewGrid(1, 3)

		windows := collectWindows(grid.VerticalWindows(3))

		assert.Equal(t, [][]day4.Coordinate{
			{
				{X: 0, Y: 0},
				{X: 0, Y: 1},
				{X: 0, Y: 2},
			},
		}, windows)
	})

	t.Run("generates 4 possible windows of size 4 within a 2x5 grid.", func(t *testing.T) {
		grid := day4.NewGrid(2, 5)

		windows := collectWindows(grid.VerticalWindows(4))

		assert.Equal(t, [][]day4.Coordinate{
			{
				{X: 0, Y: 0},
				{X: 0, Y: 1},
				{X: 0, Y: 2},
				{X: 0, Y: 3},
			},
			{
				{X: 1, Y: 0},
				{X: 1, Y: 1},
				{X: 1, Y: 2},
				{X: 1, Y: 3},
			},
			{
				{X: 0, Y: 1},
				{X: 0, Y: 2},
				{X: 0, Y: 3},
				{X: 0, Y: 4},
			},
			{
				{X: 1, Y: 1},
				{X: 1, Y: 2},
				{X: 1, Y: 3},
				{X: 1, Y: 4},
			},
		}, windows)
	})
}

func TestGridDiagonalSEWindows(t *testing.T) {
	t.Run("generates no possible windows in a 0x0 grid.", func(t *testing.T) {
		grid := day4.NewGrid(0, 0)

		windows := collectWindows(grid.DiagonalSEWindows(3))

		assert.Equal(t, [][]day4.Coordinate{}, windows)
	})

	t.Run("generates one possible window of size 3 within a 3x3 grid.", func(t *testing.T) {
		grid := day4.NewGrid(3, 3)

		windows := collectWindows(grid.DiagonalSEWindows(3))

		assert.Equal(t, [][]day4.Coordinate{
			{
				{X: 0, Y: 0},
				{X: 1, Y: 1},
				{X: 2, Y: 2},
			},
		}, windows)
	})

	t.Run("generates 6 possible windows of size 3 within a 4x5 grid.", func(t *testing.T) {
		grid := day4.NewGrid(4, 5)

		windows := collectWindows(grid.DiagonalSEWindows(3))

		assert.Equal(t, [][]day4.Coordinate{
			{
				{X: 0, Y: 0},
				{X: 1, Y: 1},
				{X: 2, Y: 2},
			},
			{
				{X: 1, Y: 0},
				{X: 2, Y: 1},
				{X: 3, Y: 2},
			},
			{
				{X: 0, Y: 1},
				{X: 1, Y: 2},
				{X: 2, Y: 3},
			},
			{
				{X: 1, Y: 1},
				{X: 2, Y: 2},
				{X: 3, Y: 3},
			},
			{
				{X: 0, Y: 2},
				{X: 1, Y: 3},
				{X: 2, Y: 4},
			},
			{
				{X: 1, Y: 2},
				{X: 2, Y: 3},
				{X: 3, Y: 4},
			},
		}, windows)
	})
}

func TestGridDiagonalSWWindows(t *testing.T) {
	t.Run("generates no possible windows in a 0x0 grid.", func(t *testing.T) {
		grid := day4.NewGrid(0, 0)

		windows := collectWindows(grid.DiagonalSWWindows(3))

		assert.Equal(t, [][]day4.Coordinate{}, windows)
	})

	t.Run("generates one possible window of size 3 within a 3x3 grid.", func(t *testing.T) {
		grid := day4.NewGrid(3, 3)

		windows := collectWindows(grid.DiagonalSWWindows(3))

		assert.Equal(t, [][]day4.Coordinate{
			{
				{X: 2, Y: 0},
				{X: 1, Y: 1},
				{X: 0, Y: 2},
			},
		}, windows)
	})

	t.Run("generates 6 possible windows of size 3 within a 4x5 grid.", func(t *testing.T) {
		grid := day4.NewGrid(4, 5)

		windows := collectWindows(grid.DiagonalSWWindows(3))

		assert.Equal(t, [][]day4.Coordinate{
			{
				{X: 2, Y: 0},
				{X: 1, Y: 1},
				{X: 0, Y: 2},
			},
			{
				{X: 3, Y: 0},
				{X: 2, Y: 1},
				{X: 1, Y: 2},
			},
			{
				{X: 2, Y: 1},
				{X: 1, Y: 2},
				{X: 0, Y: 3},
			},
			{
				{X: 3, Y: 1},
				{X: 2, Y: 2},
				{X: 1, Y: 3},
			},
			{
				{X: 2, Y: 2},
				{X: 1, Y: 3},
				{X: 0, Y: 4},
			},
			{
				{X: 3, Y: 2},
				{X: 2, Y: 3},
				{X: 1, Y: 4},
			},
		}, windows)
	})
}
