package ceresSearch

import (
	"aoc2024/lib"
	"strings"
)

func RunPartOne(progressUpdater func(fraction float64, intermediaryResult int), input string) {
	grid := parseInput(input)
	count := 0
	currentPartialMatch := ""

	processor := func(row int, col int, reset bool) {
		if reset {
			currentPartialMatch = ""
		}
		processPotentialMatch(&currentPartialMatch, &count, grid[row][col])
	}
	lib.IterateRows(grid, processor)
	lib.IterateCols(grid, processor)
	lib.IterateDiagonalsSouthWestToNorthEast(grid, processor)
	lib.IterateDiagonalsNorthWestToSouthEast(grid, processor)
	progressUpdater(1, count)
	return
}

func RunPartTwo(progressUpdater func(fraction float64, intermediaryResult int), input string) {
	grid := parseInput(input)
	count := 0
	currentPartialMatch := ""

	processor := func(row int, col int, reset bool) {
		if reset {
			currentPartialMatch = ""
		}
		processPotentialCrossMatch(grid, row, col, &currentPartialMatch, &count)
	}
	lib.IterateDiagonalsSouthWestToNorthEast(grid, processor)
	progressUpdater(1, count)
	return
}

func processPotentialMatch(currentPartialMatch *string, count *int, newChar rune) {
	newPartialMatchCandidate := *currentPartialMatch + string(newChar)
	if strings.HasPrefix("XMAS", newPartialMatchCandidate) || strings.HasPrefix("SAMX", newPartialMatchCandidate) {
		*currentPartialMatch = newPartialMatchCandidate
	} else {
		if newChar == 'X' {
			*currentPartialMatch = "X"
		} else if newChar == 'S' {
			*currentPartialMatch = "S"
		} else {
			*currentPartialMatch = ""
		}
	}
	if *currentPartialMatch == "XMAS" {
		*count++
		*currentPartialMatch = "S"
	}
	if *currentPartialMatch == "SAMX" {
		*count++
		*currentPartialMatch = "X"
	}
}

func processPotentialCrossMatch(input [][]rune, row int, col int, currentPartialMatch *string, count *int) {
	newPartialMatchCandidate := *currentPartialMatch + string(input[row][col])
	if strings.HasPrefix("MAS", newPartialMatchCandidate) || strings.HasPrefix("SAM", newPartialMatchCandidate) {
		*currentPartialMatch = newPartialMatchCandidate
	} else {
		if input[row][col] == 'M' {
			*currentPartialMatch = "M"
		} else if input[row][col] == 'S' {
			*currentPartialMatch = "S"
		} else {
			*currentPartialMatch = ""
		}
	}
	if (*currentPartialMatch == "MAS" || *currentPartialMatch == "SAM") && row+2 < len(input) && col-2 >= 0 {
		if input[row+2][col] == 'M' && input[row][col-2] == 'S' {
			*count++
		}
		if input[row+2][col] == 'S' && input[row][col-2] == 'M' {
			*count++
		}
		if *currentPartialMatch == "MAS" {
			*currentPartialMatch = "S"
		}
		if *currentPartialMatch == "SAM" {
			*currentPartialMatch = "M"
		}
	}
}

func parseInput(input string) (grid [][]rune) {
	for _, line := range strings.Split(input, "\r\n") {
		grid = append(grid, []rune(line))
	}
	return
}
