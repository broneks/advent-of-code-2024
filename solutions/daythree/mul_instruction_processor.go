package daythree

import (
	"fmt"
	"regexp"
	"strconv"
)

type MulInstructionProcessor struct {
	InstructionProcessor
}

func NewMulInstructionProcessor() *MulInstructionProcessor {
	return &MulInstructionProcessor{
		InstructionProcessor{
			Type:         "mul",
			Instructions: &[]string{},
		},
	}
}

func (p *MulInstructionProcessor) Collect(line string) {
	var instruction string
	var commas int

	resetInstruction := func() {
		instruction = ""
		commas = 0
	}

	saveInstruction := func() {
		*p.Instructions = append(*p.Instructions, instruction)
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

func (p *MulInstructionProcessor) CalculateOne(instruction string) int {
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

func (p *MulInstructionProcessor) CalculateAll() int {
	var result int

	for _, instruction := range *p.Instructions {
		result += p.CalculateOne(instruction)
	}

	return result
}
