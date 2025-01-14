package hoofIt

import (
	"strconv"
	"strings"
)

func RunPartOne(progressUpdater func(fraction float64, intermediaryResult int), input string) {
	grid := parseInput(input)
	progressUpdater(1, calculateTrailheadScore(grid, true))
	return
}

func RunPartTwo(progressUpdater func(fraction float64, intermediaryResult int), input string) {
	grid := parseInput(input)
	progressUpdater(1, calculateTrailheadScore(grid, false))
	return
}

func calculateTrailheadScore(grid [][]int, unique bool) (score int) {
	for startRow := 0; startRow < len(grid); startRow++ {
		for startCol := 0; startCol < len(grid); startCol++ {
			if grid[startRow][startCol] == 0 {
				validPaths := [][]int{{startRow, startCol}}
				for len(validPaths) > 0 {
					currentRow := validPaths[0][0]
					currentCol := validPaths[0][1]
					currentHeight := grid[currentRow][currentCol]
					if currentHeight == 9 {
						score++
					} else {
						if currentRow-1 >= 0 && grid[currentRow-1][currentCol] == currentHeight+1 {
							validPaths = addUniqueCoordinate(validPaths, currentRow-1, currentCol, unique)
						}
						if currentCol+1 < len(grid) && grid[currentRow][currentCol+1] == currentHeight+1 {
							validPaths = addUniqueCoordinate(validPaths, currentRow, currentCol+1, unique)
						}
						if currentRow+1 < len(grid) && grid[currentRow+1][currentCol] == currentHeight+1 {
							validPaths = addUniqueCoordinate(validPaths, currentRow+1, currentCol, unique)
						}
						if currentCol-1 >= 0 && grid[currentRow][currentCol-1] == currentHeight+1 {
							validPaths = addUniqueCoordinate(validPaths, currentRow, currentCol-1, unique)
						}
					}
					validPaths = validPaths[1:]
				}
			}
		}
	}
	return
}

func addUniqueCoordinate(coordinates [][]int, row int, col int, unique bool) [][]int {
	if unique {
		for _, coordinate := range coordinates {
			if coordinate[0] == row && coordinate[1] == col {
				return coordinates
			}
		}
	}
	return append(coordinates, []int{row, col})
}

func parseInput(input string) (grid [][]int) {
	for _, line := range strings.Split(input, "\r\n") {
		currentRow := make([]int, 0)
		for _, heightChar := range []rune(line) {
			height, err := strconv.Atoi(string(heightChar))
			if err != nil {
				panic(err)
			}
			currentRow = append(currentRow, height)
		}
		grid = append(grid, currentRow)
	}
	return
}
