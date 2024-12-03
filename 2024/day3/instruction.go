package day3

import (
	"regexp"
	"strconv"
)

type MulInstruction struct {
	X, Y int
}
type DoInstruction struct {
}
type DontInstruction struct {
}

var instructionRegExpr = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

func ParseInstructions(input string) []interface{} {
	matches := instructionRegExpr.FindAllStringSubmatch(input, -1)

	instructions := make([]interface{}, len(matches))

	for i, match := range matches {
		x, err := strconv.Atoi(match[1])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(match[2])
		if err != nil {
			panic(err)
		}
		instructions[i] = MulInstruction{
			X: x,
			Y: y,
		}
	}

	return instructions
}

func GetAllMultiplicationInstructions(instructions []interface{}) []MulInstruction {
	multiplicationInstructions := make([]MulInstruction, 0)

	for _, instruction := range instructions {
		switch v := instruction.(type) {
		case MulInstruction:
			multiplicationInstructions = append(multiplicationInstructions, v)
		default:
			continue
		}
	}

	return multiplicationInstructions
}

func GetAllEnableMultiplicationInstructions(istructions []interface{}) []MulInstruction {
	return []MulInstruction{}
}

func SumProducts(instructions []MulInstruction) int {
	sum := 0

	for _, instruction := range instructions {
		sum += instruction.X * instruction.Y
	}

	return sum
}
