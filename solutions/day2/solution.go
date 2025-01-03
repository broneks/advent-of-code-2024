package day2

import (
	"aoc2024/shared"
	"fmt"
	"log"
	"strconv"
	"strings"
)

const (
	minDiff = 1
	maxDiff = 3
)

func isGraduallyIncreasing(levels []int) bool {
	for i := 1; i < len(levels); i++ {
		prev, curr := levels[i-1], levels[i]

		if curr < prev || (curr-prev < minDiff) || (curr-prev > maxDiff) {
			return false
		}
	}

	return true
}

func isGraduallyDecreasing(levels []int) bool {
	for i := 1; i < len(levels); i++ {
		prev, curr := levels[i-1], levels[i]

		if prev < curr || (prev-curr < minDiff) || (prev-curr > maxDiff) {
			return false
		}
	}

	return true
}

func Solution() {
	safeCount := 0

	shared.ScanFile("./solutions/day2/input.txt", func(line string) {
		parts := strings.Split(line, " ")

		levels := make([]int, len(parts))

		for i, part := range parts {
			level, err := strconv.Atoi(part)
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				return
			}
			levels[i] = level
		}

		if isGraduallyIncreasing(levels) || isGraduallyDecreasing(levels) {
			safeCount += 1
		} else {
			for i := 0; i < len(levels); i++ {
				var reduced []int

				reduced = append(reduced, levels[:i]...)
				reduced = append(reduced, levels[i+1:]...)

				if isGraduallyIncreasing(reduced) || isGraduallyDecreasing(reduced) {
					safeCount += 1
					break
				}
			}
		}
	})

	log.Println(safeCount)
}
