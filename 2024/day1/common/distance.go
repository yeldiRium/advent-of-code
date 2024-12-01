package common

import "slices"

func Distance(leftList, rightList []int) int {
	if len(leftList) != len(rightList) {
		panic("input lists must have the same length")
	}

	slices.Sort(leftList)
	slices.Sort(rightList)

	distance := 0
	for i, left := range leftList {
		right := rightList[i]

		difference := left - right
		if difference < 0 {
			difference = -difference
		}

		distance += difference
	}

	return distance
}

