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

var instructionRegExpr = regexp.MustCompile(`(do)\(\)|(don't)\(\)|(mul)\((\d{1,3}),(\d{1,3})\)`)

func ParseInstructions(input string) []interface{} {
	matches := instructionRegExpr.FindAllStringSubmatch(input, -1)

	instructions := make([]interface{}, len(matches))

	for i, match := range matches {
		if match[1] == "do" {
			instructions[i] = DoInstruction{}
		} else if match[2] == "don't" {
			instructions[i] = DontInstruction{}
		} else if match[3] == "mul" {
			x, err := strconv.Atoi(match[4])
			if err != nil {
				panic(err)
			}
			y, err := strconv.Atoi(match[5])
			if err != nil {
				panic(err)
			}
			instructions[i] = MulInstruction{
				X: x,
				Y: y,
			}
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

func GetAllEnabledMultiplicationInstructions(instructions []interface{}) []MulInstruction {
	multiplicationInstructions := make([]MulInstruction, 0)

	enabled := true
	for _, instruction := range instructions {
		switch v := instruction.(type) {
		case DoInstruction:
			enabled = true
		case DontInstruction:
			enabled = false
		case MulInstruction:
			if enabled {
				multiplicationInstructions = append(multiplicationInstructions, v)
			}
		}
	}

	return multiplicationInstructions
}

func SumProducts(instructions []MulInstruction) int {
	sum := 0

	for _, instruction := range instructions {
		sum += instruction.X * instruction.Y
	}

	return sum
}
