package ceresSearch

import (
	"strings"
)

func RunPartOne(progressUpdater func(fraction float64, intermediaryResult int), input string) {
	grid := parseInput(input)

	count := 0
	for row := 0; row < len(grid); row++ {
		currentPartialMatch := ""
		for col := 0; col < len(grid); col++ {
			processPotentialMatch(&currentPartialMatch, &count, grid[row][col])
		}
		progressUpdater(float64(row)/float64(len(grid)*6), count)
	}
	for col := 0; col < len(grid); col++ {
		currentPartialMatch := ""
		for row := 0; row < len(grid); row++ {
			processPotentialMatch(&currentPartialMatch, &count, grid[row][col])
		}
		progressUpdater((float64(len(grid))+float64(col))/float64(len(grid)*6), count)
	}
	for startRow := 0; startRow < len(grid); startRow++ {
		currentPartialMatch := ""
		for currentRow := startRow; currentRow >= 0; currentRow-- {
			processPotentialMatch(&currentPartialMatch, &count, grid[currentRow][startRow-currentRow])
		}
		progressUpdater((float64(2*len(grid))+float64(startRow))/float64(len(grid)*6), count)
	}
	for startRow := len(grid) - 1; startRow >= 0; startRow-- {
		currentPartialMatch := ""
		for currentRow := startRow; currentRow < len(grid); currentRow++ {
			processPotentialMatch(&currentPartialMatch, &count, grid[currentRow][currentRow-startRow])
		}
		progressUpdater((float64(3*len(grid))+float64(startRow))/float64(len(grid)*6), count)
	}
	for startCol := len(grid) - 1; startCol > 0; startCol-- {
		currentPartialMatch := ""
		for currentCol := startCol; currentCol < len(grid); currentCol++ {
			processPotentialMatch(&currentPartialMatch, &count, grid[currentCol-startCol][currentCol])
		}
		progressUpdater((float64(4*len(grid))+float64(startCol))/float64(len(grid)*6), count)
	}
	for startCol := len(grid) - 1; startCol > 0; startCol-- {
		currentPartialMatch := ""
		for currentCol := startCol; currentCol < len(grid); currentCol++ {
			processPotentialMatch(&currentPartialMatch, &count, grid[len(grid)+startCol-currentCol-1][currentCol])
		}
		progressUpdater((float64(5*len(grid))+float64(startCol))/float64(len(grid)*6), count)
	}
	progressUpdater(1, count)
	return
}

func RunPartTwo(progressUpdater func(fraction float64, intermediaryResult int), input string) {
	grid := parseInput(input)

	count := 0
	for startRow := 0; startRow < len(grid); startRow++ {
		currentPartialMatch := ""
		for currentRow := startRow; currentRow >= 0; currentRow-- {
			processPotentialCrossMatch(grid, currentRow, startRow-currentRow, &currentPartialMatch, &count)
		}
		progressUpdater(float64(startRow)/float64(len(grid)*2), count)
	}
	for startCol := 1; startCol <= len(grid); startCol++ {
		currentPartialMatch := ""
		for currentCol := startCol; currentCol < len(grid); currentCol++ {
			processPotentialCrossMatch(grid, len(grid)+startCol-currentCol-1, currentCol, &currentPartialMatch, &count)
		}
		progressUpdater((float64(len(grid))+float64(startCol))/float64(len(grid)*2), count)
	}
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
