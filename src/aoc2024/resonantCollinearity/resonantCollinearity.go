package resonantCollinearity

import (
	"aoc2024/lib"
	"strings"
)

func RunPartOne(progressUpdater func(fraction float64, intermediaryResult int), input string) {
	grid := parseInput(input)
	antennaMap := constructAntennaMap(grid)
	antiNodeLocations := make([][]int, 0)

	processor := func(row1 int, col1 int, row2 int, col2 int) {
		rowDistance := row1 - row2
		colDistance := col1 - col2
		_, antiNodeLocations = addPotentialAntiNode(grid, antiNodeLocations, row1+rowDistance, col1+colDistance)
		_, antiNodeLocations = addPotentialAntiNode(grid, antiNodeLocations, row2-rowDistance, col2-colDistance)
	}
	processAntennaMap(progressUpdater, antennaMap, processor)
	progressUpdater(1, len(antiNodeLocations))
}

func RunPartTwo(progressUpdater func(fraction float64, intermediaryResult int), input string) {
	grid := parseInput(input)
	antennaMap := constructAntennaMap(grid)
	antiNodeLocations := make([][]int, 0)

	processor := func(row1 int, col1 int, row2 int, col2 int) {
		rowDistance := row1 - row2
		colDistance := col1 - col2
		outOfBounds := false
		for index := 0; !outOfBounds; index++ {
			outOfBounds, antiNodeLocations = addPotentialAntiNode(grid, antiNodeLocations, row1+index*rowDistance, col1+index*colDistance)
		}
		outOfBounds = false
		for index := 0; !outOfBounds; index++ {
			outOfBounds, antiNodeLocations = addPotentialAntiNode(grid, antiNodeLocations, row2-index*rowDistance, col2-index*colDistance)
		}
	}
	processAntennaMap(progressUpdater, antennaMap, processor)
	progressUpdater(1, len(antiNodeLocations))
}

func processAntennaMap(progressUpdater func(fraction float64, intermediaryResult int), antennaMap map[rune][][]int,
	processor func(row1 int, col1 int, row2 int, col2 int)) {
	count := 0
	for _, antennas := range antennaMap {
		progressUpdater(float64(count/len(antennaMap)), len(antennaMap))
		for antennaIndex, antenna := range antennas {
			for comparingAntennaIndex := antennaIndex + 1; comparingAntennaIndex < len(antennas); comparingAntennaIndex++ {
				processor(antenna[0], antenna[1], antennas[comparingAntennaIndex][0], antennas[comparingAntennaIndex][1])
			}
		}
		count++
	}
}

func addPotentialAntiNode(grid [][]rune, antiNodeLocations [][]int, row int, col int) (bool, [][]int) {
	if row >= 0 && row < len(grid) && col >= 0 && col < len(grid) {
		add := true
		for _, location := range antiNodeLocations {
			if location[0] == row && location[1] == col {
				add = false
			}
		}
		if add {
			antiNodeLocations = append(antiNodeLocations, []int{row, col})
		}
		return false, antiNodeLocations
	} else {
		return true, antiNodeLocations
	}
}

func constructAntennaMap(grid [][]rune) (antennaMap map[rune][][]int) {
	antennaMap = make(map[rune][][]int)
	processor := func(row int, col int, reset bool) {
		currentRune := grid[row][col]
		if currentRune != '.' {
			if antennaMap[currentRune] == nil {
				antennaMap[currentRune] = make([][]int, 0)
			}
			antennaMap[currentRune] = append(antennaMap[currentRune], []int{row, col})
		}
	}
	lib.IterateRows(grid, processor)
	return
}

func parseInput(input string) (grid [][]rune) {
	for _, line := range strings.Split(input, "\r\n") {
		grid = append(grid, []rune(line))
	}
	return
}
