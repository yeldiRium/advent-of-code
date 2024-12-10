package day5

type OrderRule struct {
	Earlier int
	Later int
}
type RuleBook = []OrderRule

func IsInOrder(elements []int, rules RuleBook) bool {
	brokenRules := brokenRules(elements, rules)

	if len(brokenRules) > 0 {
		return false
	}
	return true
}

type RuleBreak struct {
	brokenRule OrderRule
	firstPosition int
	secondPosition int
}
func brokenRules(elements []int, rules RuleBook) []RuleBreak {
	brokenRules := make([]RuleBreak, 0)

	for i := 0; i < len(elements); i++ {
		activeElement := elements[i]

		for j := i + 1; j < len(elements); j++ {
			compareElement := elements[j]

			for _, rule := range rules {
				if activeElement == rule.Later && compareElement == rule.Earlier {
					brokenRules = append(brokenRules, RuleBreak{
						brokenRule: rule,
						firstPosition: i,
						secondPosition: j,
					})
				}
			}
		}
	}

	return brokenRules
}
