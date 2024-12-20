package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"time"
)

const (
	NORTH         int = 0
	EAST              = 1
	SOUTH             = 2
	WEST              = 3
	OUT_OF_BOUNDS     = 4
)

func main() {
	input := readInput()
	countDistinctVisitedTiles(input)        // 4454
	countObstructionsForInfiniteLoop(input) // 1503
}

func countDistinctVisitedTiles(input [][]rune) {
	fmt.Print("\n\n\n")
	start := time.Now()
	visitedTiles := make([][]int, 0)
	direction := NORTH
	row, col := findCurrentLocation(input)
	for direction != OUT_OF_BOUNDS {
		if !containsTile(visitedTiles, row, col) {
			visitedTiles, _ = addVisitedTile(visitedTiles, nil, row, col, direction)
		}
		row, col, direction = move(input, row, col, direction)
		printUpdate("Amount of distinct visited squares", len(visitedTiles), time.Since(start), 0)
	}
	printUpdate("Amount of distinct visited squares", len(visitedTiles), time.Since(start), 100)
}

func countObstructionsForInfiniteLoop(input [][]rune) {
	fmt.Print("\n\n\n")
	count := 0
	start := time.Now()
	startRow, startCol := findCurrentLocation(input)
	for obstructionRow := 0; obstructionRow < len(input); obstructionRow++ {
		for obstructionCol := 0; obstructionCol < len(input); obstructionCol++ {
			if input[obstructionCol][obstructionRow] == '#' || input[obstructionCol][obstructionRow] == '^' {
				continue
			}
			input[obstructionCol][obstructionRow] = '#'
			direction := NORTH
			row, col := startRow, startCol
			visitedTiles := make([][]int, 0)
			visitedDirections := make([]int, 0)
			for direction != OUT_OF_BOUNDS {
				visitedTiles, visitedDirections = addVisitedTile(visitedTiles, visitedDirections, row, col, direction)
				row, col, direction = move(input, row, col, direction)
				if hasLoop(visitedTiles, visitedDirections) {
					count++
					break
				}
			}
			input[obstructionCol][obstructionRow] = '.'
			printUpdate("Amount of obstructions resulting in infinite loop", count, time.Since(start),
				float64(obstructionRow*len(input)+obstructionCol)/float64(len(input)*len(input))*100)
		}
	}
	printUpdate("Amount of obstructions resulting in infinite loop", count, time.Since(start), 100)
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

func printUpdate(description string, count int, timeSinceStart time.Duration, percentage float64) {
	completed := (percentage * float64(50)) / 100
	remaining := float64(50) - completed
	fmt.Printf("\033[F\033[K") // Move up and clear the line
	fmt.Printf("\033[F\033[K") // Move up and clear the line
	fmt.Println(fmt.Sprintf("%s: %d", description, count))
	fmt.Println(fmt.Sprintf("[%s%s] %ds - %.2f%%", strings.Repeat("=", int(completed)),
		strings.Repeat(" ", int(remaining)), int64(math.Round(timeSinceStart.Seconds())), percentage))
}

func readInput() (input [][]rune) {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		input = append(input, []rune(scanner.Text()))
	}
	return
}
