package gardenGroups

import (
	"aoc2024/lib"
	"strings"
)

func RunPartOne(progressUpdater func(fraction float64, intermediaryResult int), input string) {
	grid := parseInput(input)
	groups := findGroups(progressUpdater, grid)
	fencePrice := 0
	for _, group := range groups {
		fencePrice = fencePrice + len(group)*calculateGroupPerimeter(group)
	}
	progressUpdater(1, fencePrice)
}

func RunPartTwo(progressUpdater func(fraction float64, intermediaryResult int), input string) {

}

func findGroups(progressUpdater func(fraction float64, intermediaryResult int), grid [][]rune) [][][]int {
	var allCellIndexes [][]int
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid); col++ {
			allCellIndexes = append(allCellIndexes, []int{row, col})
		}
	}
	groups := make([][][]int, 0)
	for len(allCellIndexes) != 0 {
		groupIndexes := findGroup(grid, allCellIndexes[0][0], allCellIndexes[0][1])
		groups = append(groups, groupIndexes)
		allCellIndexes = lib.Remove(allCellIndexes, groupIndexes)
		progressUpdater(1-float64(len(allCellIndexes))/float64(len(grid)*len(grid)), 0)
	}
	return groups
}

func findGroup(grid [][]rune, row int, col int) [][]int {
	currentGroupIndexes := [][]int{{row, col}}
	uncheckedGroupMemberIndexes := [][]int{{row, col}}
	for len(uncheckedGroupMemberIndexes) != 0 {
		currentPotentialGroupMemberIndex := uncheckedGroupMemberIndexes[0]
		currentRow := currentPotentialGroupMemberIndex[0]
		currentCol := currentPotentialGroupMemberIndex[1]
		if currentRow-1 >= 0 && grid[currentRow-1][currentCol] == grid[currentRow][currentCol] &&
			!lib.Contains(currentGroupIndexes, currentRow-1, currentCol) {
			currentGroupIndexes = append(currentGroupIndexes, []int{currentRow - 1, currentCol})
			uncheckedGroupMemberIndexes = append(uncheckedGroupMemberIndexes, []int{currentRow - 1, currentCol})
		}
		if currentRow+1 < len(grid) && grid[currentRow+1][currentCol] == grid[currentRow][currentCol] &&
			!lib.Contains(currentGroupIndexes, currentRow+1, currentCol) {
			currentGroupIndexes = append(currentGroupIndexes, []int{currentRow + 1, currentCol})
			uncheckedGroupMemberIndexes = append(uncheckedGroupMemberIndexes, []int{currentRow + 1, currentCol})
		}
		if currentCol-1 >= 0 && grid[currentRow][currentCol-1] == grid[currentRow][currentCol] &&
			!lib.Contains(currentGroupIndexes, currentRow, currentCol-1) {
			currentGroupIndexes = append(currentGroupIndexes, []int{currentRow, currentCol - 1})
			uncheckedGroupMemberIndexes = append(uncheckedGroupMemberIndexes, []int{currentRow, currentCol - 1})
		}
		if currentCol+1 < len(grid) && grid[currentRow][currentCol+1] == grid[currentRow][currentCol] &&
			!lib.Contains(currentGroupIndexes, currentRow, currentCol+1) {
			currentGroupIndexes = append(currentGroupIndexes, []int{currentRow, currentCol + 1})
			uncheckedGroupMemberIndexes = append(uncheckedGroupMemberIndexes, []int{currentRow, currentCol + 1})
		}
		uncheckedGroupMemberIndexes = uncheckedGroupMemberIndexes[1:]
	}
	return currentGroupIndexes
}

func calculateGroupPerimeter(group [][]int) int {
	perimeter := 0
	for _, groupMember := range group {
		if !lib.Contains(group, groupMember[0]-1, groupMember[1]) {
			perimeter = perimeter + 1
		}
		if !lib.Contains(group, groupMember[0]+1, groupMember[1]) {
			perimeter = perimeter + 1
		}
		if !lib.Contains(group, groupMember[0], groupMember[1]-1) {
			perimeter = perimeter + 1
		}
		if !lib.Contains(group, groupMember[0], groupMember[1]+1) {
			perimeter = perimeter + 1
		}
	}
	return perimeter
}

func parseInput(input string) (grid [][]rune) {
	for _, line := range strings.Split(input, "\r\n") {
		grid = append(grid, []rune(line))
	}
	return
}
