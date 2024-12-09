package day4

import "iter"

// Grid represents a grid of characters in row-major order.
// (https://en.wikipedia.org/wiki/Row-_and_column-major_order)
type Grid [][]rune
type Coordinate struct {
	X, Y uint
}

func NewGrid(width, height int) Grid {
	grid := make([][]rune, width)
	for i := 0; i < width; i++ {
		grid[i] = make([]rune, height)
	}
	return grid
}

func (grid Grid) Width() uint {
	return uint(len(grid))
}
func (grid Grid) Height() uint {
	if len(grid) == 0 {
		return 0
	}
	return uint(len(grid[0]))
}

func (grid Grid) RuneAt(x, y int) rune {
	return grid[x][y]
}

func (grid Grid) HorizontalWindows(length uint) iter.Seq[[]Coordinate] {
	return func(yield func(value []Coordinate) bool) {
		if grid.Width() < length {
			return
		}

		for y := uint(0); y < grid.Height(); y++ {
			for x := uint(0); x < grid.Width()-length+1; x++ {
				window := make([]Coordinate, length)
				for i := uint(0); i < length; i++ {
					window[i] = Coordinate{
						X: x + i,
						Y: y,
					}
				}

				if !yield(window) {
					return
				}
			}
		}
	}
}

func (grid Grid) VerticalWindows(length uint) iter.Seq[[]Coordinate] {
	return func(yield func(value []Coordinate) bool) {
		if grid.Height() < length {
			return
		}

		for y := uint(0); y < grid.Height()-length+1; y++ {
			for x := uint(0); x < grid.Width(); x++ {
				window := make([]Coordinate, length)
				for i := uint(0); i < length; i++ {
					window[i] = Coordinate{
						X: x,
						Y: y + i,
					}
				}

				if !yield(window) {
					return
				}
			}
		}
	}
}

// DiagonalSEWindows iterates diagonal windows that are oriented from
// north-west toward south-east.
func (grid Grid) DiagonalSEWindows(length uint) iter.Seq[[]Coordinate] {
	return func(yield func(value []Coordinate) bool) {
		if grid.Height() < length || grid.Width() < length {
			return
		}

		for y := uint(0); y < grid.Height()-length+1; y++ {
			for x := uint(0); x < grid.Width()-length+1; x++ {
				window := make([]Coordinate, length)

				for i := uint(0); i < length; i++ {
					window[i] = Coordinate{
						X: x + i,
						Y: y + i,
					}
				}

				if !yield(window) {
					return
				}
			}
		}
	}
}

// DiagonalSWWindows iterates diagonal windows that are oriented from
// north-east toward south-west.
func (grid Grid) DiagonalSWWindows(length uint) iter.Seq[[]Coordinate] {
	return func(yield func(value []Coordinate) bool) {
		if grid.Height() < length || grid.Width() < length {
			return
		}

		for y := uint(0); y < grid.Height()-length+1; y++ {
			for x := uint(0); x < grid.Width()-length+1; x++ {
				window := make([]Coordinate, length)

				for i := uint(0); i < length; i++ {
					window[i] = Coordinate{
						X: x - i + length - 1,
						Y: y + i,
					}
				}

				if !yield(window) {
					return
				}
			}
		}
	}
}
