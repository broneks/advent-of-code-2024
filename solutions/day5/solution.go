package day5

import (
	"aoc2024/shared"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type PageOrderingRules map[string]*[]string

func isPageUpdateCorrect(pageUpdate []string, pageOrderingRules PageOrderingRules) bool {
	pagesSeen := make(map[string]bool)
	pagesMissing := make(map[string]bool)

	for _, page := range pageUpdate {
		if _, missing := pagesMissing[page]; missing {
			return false
		}

		pagesSeen[page] = true

		pagesBefore, ruleExists := pageOrderingRules[page]

		if !ruleExists {
			continue
		}

		for _, pageBefore := range *pagesBefore {
			if _, found := pagesSeen[pageBefore]; !found {
				pagesMissing[pageBefore] = true
			}
		}
	}

	return true
}

func getPageUpdateMiddleValue(pageUpdate []string, pageOrderingRules PageOrderingRules) int {
	if !isPageUpdateCorrect(pageUpdate, pageOrderingRules) {
		return 0
	}

	middle, err := strconv.Atoi(pageUpdate[len(pageUpdate)/2])
	if err != nil {
		log.Fatalln("Invalid middle number")
		return 0
	}

	return middle
}

func Solution() {
	pageOrderingRules := make(PageOrderingRules)
	var pageUpdates [][]string
	var hasReachedSeparator bool

	shared.ScanFile("./solutions/day5/input.txt", func(line string) {
		if hasReachedSeparator {
			pageUpdates = append(pageUpdates, strings.Split(line, ","))
		} else {
			if line == "" {
				hasReachedSeparator = true
				return
			}

			pageBefore, pageAfter, found := strings.Cut(line, "|")
			if !found {
				log.Fatalln("Invalid rule format")
				return
			}

			rule, exists := pageOrderingRules[pageAfter]
			if exists {
				*rule = append(*rule, pageBefore)
			} else {
				pageOrderingRules[pageAfter] = &[]string{pageBefore}
			}
		}
	})

	var result int

	for _, pageUpdate := range pageUpdates {
		result += getPageUpdateMiddleValue(pageUpdate, pageOrderingRules)
	}

	fmt.Println(result)
}
