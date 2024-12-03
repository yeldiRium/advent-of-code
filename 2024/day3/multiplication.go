package day3

import (
	"regexp"
	"strconv"
)

type MulInstruction struct {
	X, Y int
}

var mulInstructionRegExp = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

func ParseMulInstructions(input string) []MulInstruction {
	matches := mulInstructionRegExp.FindAllStringSubmatch(input, -1)

	instructions := make([]MulInstruction, len(matches))

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

func SumProducts(instructions []MulInstruction) int {
	sum := 0

	for _, instruction := range instructions {
		sum += instruction.X * instruction.Y
	}

	return sum
}
