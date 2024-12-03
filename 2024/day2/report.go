package day2

type Report = []int

func CountSafeReports(reports []Report, dampen bool) int {
	safeReports := 0

	for _, report := range reports {
		if ReportSafety(report, dampen) {
			safeReports += 1
		}
	}

	return safeReports
}

func ReportSafety(report Report, dampen bool) (safe bool) {
	if len(report) <= 1 {
		return true
	}

	lowerBound := 1
	upperBound := 3
	if report[1]-report[0] < 0 {
		lowerBound = -3
		upperBound = -1
	}

	safeDifference := func(difference int) bool {
		return difference >= lowerBound && difference <= upperBound
	}

	for i := 1; i < len(report); i++ {
		if safeDifference(report[i] - report[i-1]) {
			continue
		}

		if !dampen {
			return false
		}

		// We have encountered an unsafe level. There are three possible causes for this:
		// 1) The first two levels lead us to falsely believe the elements should ascend/descend.
		//    Removing the first level fixes this.
		// 2) The level at i is unsafe. E.g. 1, 2, 6, 3, where i is 2.
		//    Removing the level at i fixes this.
		// 3) The level at i-1 is unsafe. E.g. 1, 4, 2, 3, where i is 2.
		//    Removing the level at i-1 fixes this.
		option1 := make([]int, len(report)-1)
		copy(option1, report[1:])
		if ReportSafety(option1, false) {
			return true
		}

		option2 := make([]int, len(report)-1)
		copy(option2[:i], report[:i])
		copy(option2[i:], report[i+1:])
		if ReportSafety(option2, false) {
			return true
		}

		option3 := make([]int, len(report)-1)
		copy(option3[:i-1], report[:i-1])
		copy(option3[i-1:], report[i:])
		if ReportSafety(option3, false) {
			return true
		}

		return false
	}

	return true
}
