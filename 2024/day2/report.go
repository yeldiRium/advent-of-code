package day2

type Report = []int

func CountSafeReports(reports []Report) int {
	safeReports := 0

	for _, report := range reports {
		if ReportSafety(report) {
			safeReports += 1
		}
	}

	return safeReports
}

func ReportSafety(report Report) (safe bool) {
	unsafeLevels := CountUnsafeLevels(report)

	return unsafeLevels == 0
}

func CountUnsafeLevels(report Report) int {
	derived := derivation(report)

	if len(derived) == 0 {
		return 0
	}

	lowerBound := 1
	upperBound := 3
	if derived[0] < 0 {
		lowerBound = -3
		upperBound = -1
	}

	unsafeLevelCount := 0

	for _, value := range derived {
		if value < lowerBound || value > upperBound {
			unsafeLevelCount += 1
		}
	}

	return unsafeLevelCount
}

func derivation(values []int) []int {
	if len(values) == 1 {
		return []int{}
	}

	result := make([]int, len(values) - 1)

	for i := 1; i < len(values); i++ {
		result[i-1] = values[i] - values[i-1]
	}

	return result
}

