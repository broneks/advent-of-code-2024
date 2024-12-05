package daythree

import (
	"aoc2024/shared"
	"fmt"
	"regexp"
	"strconv"
)

func collectInstructions(instructions *[]string, line string) {
	var instruction string
	var commas int

	resetInstruction := func() {
		instruction = ""
		commas = 0
	}

	saveInstruction := func() {
		*instructions = append(*instructions, instruction)
		resetInstruction()
	}

	for _, char := range line {
		instructionLen := len(instruction)

		if instructionLen > 0 {
			switch instruction[instructionLen-1] {
			case 'm':
				if char == 'u' {
					instruction += "u"
				} else {
					resetInstruction()
				}
			case 'u':
				if char == 'l' {
					instruction += "l"
				} else {
					resetInstruction()
				}
			case 'l':
				if char == '(' {
					instruction += "("
				} else {
					resetInstruction()
				}
			case '(', ',':
				switch char {
				case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
					instruction += string(char)
				default:
					resetInstruction()
				}
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				switch char {
				case ',':
					if commas == 0 {
						instruction += ","
						commas++
					} else {
						resetInstruction()
					}
				case ')':
					if commas == 1 {
						instruction += ")"
						saveInstruction()
					} else {
						resetInstruction()
					}
				case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
					instruction += string(char)
				default:
					resetInstruction()
				}
			}
		} else if char == 'm' {
			instruction += "m" // potential "mul" instruction
		}
	}
}

func mul(instruction string) int {
	matches := regexp.MustCompile(`mul\((\d+),(\d+)\)`).FindStringSubmatch(instruction)

	if len(matches) == 3 {
		a, err := strconv.Atoi(matches[1])
		b, err := strconv.Atoi(matches[2])
		if err != nil {
			fmt.Println("Error converting string to int:", err)
			return 0
		}

		return a * b
	}

	return 0
}

func calculateInstructions(instructions *[]string) int {
	var result int

	for _, instruction := range *instructions {
		switch instruction[:3] {
		case "mul":
			result += mul(instruction)
		}
	}

	return result
}

func Solution() {
	var instructions []string

	shared.ScanFile("./solutions/daythree/input.txt", func(line string) {
		collectInstructions(&instructions, line)
	})

	result := calculateInstructions(&instructions)

	fmt.Println(result)
}
