package day4

import "testing"

func TestGetGridXmasCrossCount(t *testing.T) {
	grid := []string{
		"MMMSXXMASM",
		"MSAMXMSMSA",
		"AMXSXMAAMM",
		"MSAMASMSMX",
		"XMASAMXAMM",
		"XXAMMXXAMA",
		"SMSMSASXSS",
		"SAXAMASAAA",
		"MAMMMXMMMM",
		"MXMXAXMASX",
	}
	expected := 9

	result := getGridXmasCrossCount(grid)

	if result != expected {
		t.Errorf("Expected result to be %d, got %d", expected, result)
	}
}
