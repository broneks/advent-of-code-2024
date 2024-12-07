package day1

import (
	"aoc2024/shared"
	"log"
	"sort"
	"strconv"
	"strings"
)

func Solution() {
	var listA []int
	var listB []int
	var distance int

	shared.ScanFile("./solutions/day1/input.txt", func(line string) {
		parts := strings.Split(line, "   ")

		a, _ := strconv.Atoi(parts[0])
		b, _ := strconv.Atoi(parts[1])

		listA = append(listA, a)
		listB = append(listB, b)
	})

	sort.Ints(listA)
	sort.Ints(listB)

	for i := range listA {
		a, b := listA[i], listB[i]

		var diff int
		if a > b {
			diff = a - b
		} else {
			diff = b - a
		}

		distance += diff
	}

	log.Println(distance)
}
