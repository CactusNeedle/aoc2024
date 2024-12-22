package guardianGallivant

import (
	"strings"
)

const (
	NORTH         int = 0
	EAST              = 1
	SOUTH             = 2
	WEST              = 3
	OUT_OF_BOUNDS     = 4
)

func RunPartOne(progressUpdater func(fraction float64, intermediaryResult int), input string) {
	grid := parseInput(input)
	visitedTiles := make([][]int, 0)
	direction := NORTH
	row, col := findCurrentLocation(grid)
	for direction != OUT_OF_BOUNDS {
		if !containsTile(visitedTiles, row, col) {
			visitedTiles, _ = addVisitedTile(visitedTiles, nil, row, col, direction)
		}
		row, col, direction = move(grid, row, col, direction)
		if len(visitedTiles)%10 == 0 {
			progressUpdater(0, len(visitedTiles))
		}
	}
	progressUpdater(1, len(visitedTiles))
}

func RunPartTwo(progressUpdater func(fraction float64, intermediaryResult int), input string) {
	grid := parseInput(input)
	count := 0
	startRow, startCol := findCurrentLocation(grid)
	for obstructionRow := 0; obstructionRow < len(grid); obstructionRow++ {
		for obstructionCol := 0; obstructionCol < len(grid); obstructionCol++ {
			if grid[obstructionCol][obstructionRow] == '#' || grid[obstructionCol][obstructionRow] == '^' {
				continue
			}
			grid[obstructionCol][obstructionRow] = '#'
			direction := NORTH
			row, col := startRow, startCol
			visitedTiles := make([][]int, 0)
			visitedDirections := make([]int, 0)
			for direction != OUT_OF_BOUNDS {
				visitedTiles, visitedDirections = addVisitedTile(visitedTiles, visitedDirections, row, col, direction)
				row, col, direction = move(grid, row, col, direction)
				if hasLoop(visitedTiles, visitedDirections) {
					count++
					break
				}
			}
			grid[obstructionCol][obstructionRow] = '.'
			progressUpdater(float64(obstructionRow*len(grid)+obstructionCol)/float64(len(grid)*len(grid)), count)
		}
	}
	progressUpdater(1, count)
	return
}

func containsTile(visitedTiles [][]int, row int, col int) bool {
	for _, visitedTile := range visitedTiles {
		if visitedTile[0] == row && visitedTile[1] == col {
			return true
		}
	}
	return false
}

func addVisitedTile(visitedTiles [][]int, visitedDirections []int, row int, col int, direction int) ([][]int, []int) {
	visitedTile := make([]int, 2)
	visitedTile[0] = row
	visitedTile[1] = col
	visitedTiles = append(visitedTiles, visitedTile)
	if visitedDirections != nil {
		visitedDirections = append(visitedDirections, direction)
	}
	return visitedTiles, visitedDirections
}

func hasLoop(visitedTiles [][]int, visitedDirections []int) bool {
	if len(visitedTiles) < 2 {
		return false
	}
	for index := 0; index < len(visitedTiles)/2; index++ {
		if visitedTiles[index][0] == visitedTiles[len(visitedTiles)-1][0] &&
			visitedTiles[index][1] == visitedTiles[len(visitedTiles)-1][1] &&
			visitedDirections[index] == visitedDirections[len(visitedTiles)-1] {
			return true
		}
	}
	return false
}

func findCurrentLocation(input [][]rune) (row int, col int) {
	for row = 0; row < len(input); row++ {
		for col = 0; col < len(input); col++ {
			if input[row][col] == '^' {
				return
			}
		}
	}
	panic("No guard found!")
}

func move(input [][]rune, currentRow int, currentCol int, direction int) (newRow int, newCol int, newDirection int) {
	if direction == NORTH {
		if currentRow-1 < 0 {
			newDirection = OUT_OF_BOUNDS
		} else if input[currentRow-1][currentCol] == '#' {
			newDirection = EAST
			newRow, newCol, newDirection = move(input, currentRow, currentCol, newDirection)
		} else {
			newDirection = direction
			newRow = currentRow - 1
			newCol = currentCol
		}
	} else if direction == EAST {
		if currentCol+1 >= len(input) {
			newDirection = OUT_OF_BOUNDS
		} else if input[currentRow][currentCol+1] == '#' {
			newDirection = SOUTH
			newRow, newCol, newDirection = move(input, currentRow, currentCol, newDirection)
		} else {
			newDirection = direction
			newRow = currentRow
			newCol = currentCol + 1
		}
	} else if direction == SOUTH {
		if currentRow+1 >= len(input) {
			newDirection = OUT_OF_BOUNDS
		} else if input[currentRow+1][currentCol] == '#' {
			newDirection = WEST
			newRow, newCol, newDirection = move(input, currentRow, currentCol, newDirection)
		} else {
			newDirection = direction
			newRow = currentRow + 1
			newCol = currentCol
		}
	} else if direction == WEST {
		if currentCol-1 < 0 {
			newDirection = OUT_OF_BOUNDS
		} else if input[currentRow][currentCol-1] == '#' {
			newDirection = NORTH
			newRow, newCol, newDirection = move(input, currentRow, currentCol, newDirection)
		} else {
			newDirection = direction
			newRow = currentRow
			newCol = currentCol - 1
		}
	} else {
		panic("Move called out of bounds!")
	}
	return
}

func parseInput(input string) (grid [][]rune) {
	for _, line := range strings.Split(input, "\r\n") {
		grid = append(grid, []rune(line))
	}
	return
}
