package day5

import "fmt"

type OrderRule struct {
	Earlier int
	Later   int
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
	brokenRule     OrderRule
	firstPosition  int
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
						brokenRule:     rule,
						firstPosition:  i,
						secondPosition: j,
					})
				}
			}
		}
	}

	return brokenRules
}

var ErrCantBeSorted = fmt.Errorf("the input can not be sorted according to the given order rules")

func Sort(elements []int, ruleBook RuleBook) ([]int, error) {
	elementGraph := make(graph)
	for _, element := range elements {
		elementGraph[element] = make(map[int]interface{})
	}
	for _, rule := range ruleBook {
		_, earlierContained := elementGraph[rule.Earlier]
		_, laterContained := elementGraph[rule.Later]
		if !(earlierContained && laterContained) {
			continue
		}

		mapping, ok := elementGraph[rule.Earlier]
		if !ok {
			mapping = make(map[int]interface{})
		}
		mapping[rule.Later] = struct{}{}
	}

	return elementGraph.topSort()
}

type graph map[int]map[int]interface{}

// topSort is a destructive operation on the graph. Only use it on graphs that
// were created for this purpose.
func (g graph) topSort() ([]int, error) {
	orderedNodes := make([]int, 0)
	nodesWithNoIncomingEdges := make([]int, 0)

	for node, neighbours := range g {
		if len(neighbours) == 0 {
			nodesWithNoIncomingEdges = append(nodesWithNoIncomingEdges, node)
			delete(g, node)
		}
	}

	for len(nodesWithNoIncomingEdges) > 0 {
		activeNode := nodesWithNoIncomingEdges[0]
		nodesWithNoIncomingEdges = nodesWithNoIncomingEdges[1:]

		orderedNodes = append([]int{activeNode}, orderedNodes...)

		for node, neighbours := range g {
			if _, ok := neighbours[activeNode]; !ok {
				continue
			}
			delete(neighbours, activeNode)
			if len(neighbours) == 0 {
				nodesWithNoIncomingEdges = append(nodesWithNoIncomingEdges, node)
				delete(g, node)
			}
		}
	}

	if len(g) > 0 {
		return nil, ErrCantBeSorted
	}

	return orderedNodes, nil
}
