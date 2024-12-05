package daythree

import (
	"aoc2024/shared"
	"fmt"
	"regexp"
	"strconv"
)

func collectInstructions(instructions *[]string, line string) {
	var instruction string

	for _, char := range line {
		instructionLen := len(instruction)

		if instructionLen > 0 {
			switch instruction[instructionLen-1] {
			case 'm':
				if char == 'u' {
					instruction += "u"
				}
			case 'u':
				if char == 'l' {
					instruction += "l"
				}
			case 'l':
				if char == '(' {
					instruction += "("
				}
			case '(', ',':
				switch char {
				case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
					instruction += string(char)
				}
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				switch instructionLen {
				// e.g. "mul(123"
				case 7:
					if char == ',' {
						instruction += ","
					} else {
						instruction = "" // invalid
					}
				// e.g. "mul(123,456"
				case 11:
					if char == ')' {
						instruction += ")"
					} else {
						instruction = "" // invalid
					}
				default:
					switch char {
					case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', ',':
						instruction += string(char)
					}
				}
			case ')':
				*instructions = append(*instructions, instruction) // instruction found!
				instruction = ""
			default:
				instruction = "" // invalid
			}
		} else if char == 'm' {
			instruction += "m"
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
