package daythree

import (
	"aoc2024/shared"
	"fmt"
)

func Solution() {
	instructionProcessor := NewMulInstructionProcessor()

	shared.ScanFile("./solutions/daythree/input.txt", func(line string) {
		instructionProcessor.Collect(line)
	})

	result := instructionProcessor.CalculateAll()

	fmt.Println(result)
}
