package day4

import "iter"

// Grid represents a grid of characters in row-major order.
// (https://en.wikipedia.org/wiki/Row-_and_column-major_order)
type Grid [][]rune
type Coordinate struct {
	X, Y uint
}

func NewGrid(width, height uint) Grid {
	grid := make([][]rune, width)
	for i := uint(0); i < width; i++ {
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

func (grid Grid) RuneAt(x, y uint) rune {
	return grid[x][y]
}

func (grid Grid) Windows(size uint) iter.Seq[[]Coordinate] {
	return func(yield func(value []Coordinate) bool) {
		for window := range grid.HorizontalWindows(size) {
			if !yield(window) {
				return
			}
		}
		for window := range grid.VerticalWindows(size) {
			if !yield(window) {
				return
			}
		}
		for window := range grid.DiagonalSEWindows(size) {
			if !yield(window) {
				return
			}
		}
		for window := range grid.DiagonalSWWindows(size) {
			if !yield(window) {
				return
			}
		}
	}
}

func (grid Grid) HorizontalWindows(size uint) iter.Seq[[]Coordinate] {
	return func(yield func(value []Coordinate) bool) {
		if grid.Width() < size {
			return
		}

		for y := uint(0); y < grid.Height(); y++ {
			for x := uint(0); x < grid.Width()-size+1; x++ {
				window := make([]Coordinate, size)
				for i := uint(0); i < size; i++ {
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

func (grid Grid) VerticalWindows(size uint) iter.Seq[[]Coordinate] {
	return func(yield func(value []Coordinate) bool) {
		if grid.Height() < size {
			return
		}

		for y := uint(0); y < grid.Height()-size+1; y++ {
			for x := uint(0); x < grid.Width(); x++ {
				window := make([]Coordinate, size)
				for i := uint(0); i < size; i++ {
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
func (grid Grid) DiagonalSEWindows(size uint) iter.Seq[[]Coordinate] {
	return func(yield func(value []Coordinate) bool) {
		if grid.Height() < size || grid.Width() < size {
			return
		}

		for y := uint(0); y < grid.Height()-size+1; y++ {
			for x := uint(0); x < grid.Width()-size+1; x++ {
				window := make([]Coordinate, size)

				for i := uint(0); i < size; i++ {
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
func (grid Grid) DiagonalSWWindows(size uint) iter.Seq[[]Coordinate] {
	return func(yield func(value []Coordinate) bool) {
		if grid.Height() < size || grid.Width() < size {
			return
		}

		for y := uint(0); y < grid.Height()-size+1; y++ {
			for x := uint(0); x < grid.Width()-size+1; x++ {
				window := make([]Coordinate, size)

				for i := uint(0); i < size; i++ {
					window[i] = Coordinate{
						X: x - i + size - 1,
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

// X3 iterates X-es in the grid with a diagonal length of three.
// I.e. in a grid of size 3x3 the resulting X is:
//
// X.X
// .X.
// X.X
//
// The produced coordinates are each the top-left coordinate of
// an X in the grid.
func (grid Grid) X3() iter.Seq[Coordinate] {
	return func(yield func(value Coordinate) bool) {
		for diagonal := range grid.DiagonalSEWindows(3) {
			if !yield(diagonal[0]) {
				return
			}
		}
	}
}
