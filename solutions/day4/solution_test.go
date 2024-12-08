package day4

import "testing"

func TestGetGridWordCount(t *testing.T) {
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
	expected := 18

	result := getGridWordCount(grid, "XMAS")

	if result != expected {
		t.Errorf("Expected result to be %d, got %d", expected, result)
	}
}
