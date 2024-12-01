package common

func CountOccurences(list []int) map[int]int {
	counts := make(map[int]int)

	for _, value := range list {
		count, _ := counts[value]
		counts[value] = count + 1
	}

	return counts
}

func Similarity(leftList, rightList []int) int {
	if len(leftList) != len(rightList) {
		panic("input lists must have the same length")
	}

	similarity := 0

	occurences := CountOccurences(rightList)

	for _, left := range leftList {
		occurence := occurences[left]
		similarity += left * occurence
	}

	return similarity
}
