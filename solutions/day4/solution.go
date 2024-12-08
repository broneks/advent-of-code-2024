package day4

import (
	"aoc2024/shared"
	"fmt"
)

func checkNorth(count *int, grid []string, word string, x, y int) {
	gridW := len(grid[0])
	lastIndex := len(word) - 1

	if y-lastIndex >= 0 {
		// N
		for i := range word {
			if grid[y-i][x] != word[i] {
				break
			}
			if i == lastIndex {
				*count++
			}
		}
		if x+lastIndex < gridW {
			// NE
			for i := range word {
				if grid[y-i][x+i] != word[i] {
					break
				}
				if i == lastIndex {
					*count++
				}
			}
		}
		if x-lastIndex >= 0 {
			// NW
			for i := range word {
				if grid[y-i][x-i] != word[i] {
					break
				}
				if i == lastIndex {
					*count++
				}
			}
		}
	}
}

func checkEast(count *int, grid []string, word string, x, y int) {
	gridW := len(grid[0])
	lastIndex := len(word) - 1

	if x+lastIndex < gridW {
		// E
		if grid[y][x:x+lastIndex+1] == word {
			*count++
		}
	}
}

func checkSouth(count *int, grid []string, word string, x, y int) {
	gridH := len(grid)
	gridW := len(grid[0])
	lastIndex := len(word) - 1

	if y+lastIndex < gridH {
		// S
		for i := range word {
			if grid[y+i][x] != word[i] {
				break
			}
			if i == lastIndex {
				*count++
			}
		}
		if x+lastIndex < gridW {
			// SE
			for i := range word {
				if grid[y+i][x+i] != word[i] {
					break
				}
				if i == lastIndex {
					*count++
				}
			}
		}
		if x-lastIndex >= 0 {
			// SW
			for i := range word {
				if grid[y+i][x-i] != word[i] {
					break
				}
				if i == lastIndex {
					*count++
				}
			}
		}
	}
}

func checkWest(count *int, grid []string, word string, x, y int) {
	lastIndex := len(word) - 1

	if x-lastIndex >= 0 {
		// W (backward)
		for i := lastIndex; i >= 0; i-- {
			if grid[y][x-i] != word[i] {
				break
			}
			if i == 0 {
				*count++
			}
		}
	}
}

func getSurroundingAreaWordCount(grid []string, word string, x, y int) (count int) {
	checkNorth(&count, grid, word, x, y)
	checkEast(&count, grid, word, x, y)
	checkSouth(&count, grid, word, x, y)
	checkWest(&count, grid, word, x, y)

	return
}

func getGridWordCount(grid []string, word string) (count int) {
	firstChar := rune(word[0])

	for y, row := range grid {
		for x, char := range row {
			if char == firstChar {
				count += getSurroundingAreaWordCount(grid, word, x, y)
			}
		}
	}

	return
}

func Solution() {
	var grid []string

	shared.ScanFile("./solutions/day4/input.txt", func(line string) {
		grid = append(grid, line)
	})

	result := getGridWordCount(grid, "XMAS")

	fmt.Println(result)
}
