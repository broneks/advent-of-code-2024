package day3

import (
	"aoc2024/shared"
	"fmt"
	"regexp"
	"strconv"
)

const (
	mulCapture  = 1
	dontCapture = 4
	doCapture   = 5
)

func getMulResultForLine(line string) int {
	var result int

	reg := regexp.MustCompile(`(mul)\((\d{1,3}),(\d{1,3})\)|(don't)\(\)|(do)\(\)`)
	matches := reg.FindAllStringSubmatch(line, -1)

	var stop bool

	for _, match := range matches {
		instruction := match[mulCapture]

		if instruction == "" {
			instruction = match[dontCapture]
		}

		if instruction == "" {
			instruction = match[doCapture]
		}

		switch instruction {
		case "mul":
			if stop {
				continue
			}
			a, err := strconv.Atoi(match[2])
			b, err := strconv.Atoi(match[3])
			if err != nil {
				fmt.Printf("Error converting string to int: %v", err)
				return 0
			}
			result += a * b
		case "don't":
			stop = true
		case "do":
			stop = false
		}
	}

	return result
}

func Solution() {
	var line string

	shared.ScanFile("./solutions/day3/input.txt", func(l string) {
		line += l
	})

	result := getMulResultForLine(line)

	fmt.Println(result)
}
