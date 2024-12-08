package day4

import (
	"aoc2024/shared"
	"fmt"
)

func getGridXmasCrossCount(grid []string) (count int) {
	gridW, gridH := len(grid[0]), len(grid)

	for y, row := range grid {
		for x, char := range row {
			if char == 'A' {
				// within bounds?
				if y-1 >= 0 && x+1 < gridW && y+1 < gridH && x-1 >= 0 {
					nw, ne := grid[y-1][x-1], grid[y-1][x+1]
					sw, se := grid[y+1][x-1], grid[y+1][x+1]

					if ((nw == 'M' && se == 'S') || (se == 'M' && nw == 'S')) &&
						((sw == 'M' && ne == 'S') || (ne == 'M' && sw == 'S')) {
						count++
					}
				}
			}
		}
	}

	return
}

func SolutionPart2() {
	var grid []string

	shared.ScanFile("./solutions/day4/input.txt", func(line string) {
		grid = append(grid, line)
	})

	result := getGridXmasCrossCount(grid)

	fmt.Println(result)
}
