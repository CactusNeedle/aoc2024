package lib

func IterateRows(grid [][]rune, process func(row int, col int, reset bool)) {
	for row, _ := range grid {
		reset := true
		for col, _ := range grid[row] {
			process(row, col, reset)
			reset = false
		}
	}
}

func IterateCols(grid [][]rune, process func(row int, col int, reset bool)) {
	for col, _ := range grid[0] {
		reset := true
		for row, _ := range grid[0] {
			process(row, col, reset)
			reset = false
		}
	}
}

func IterateDiagonalsSouthWestToNorthEast(grid [][]rune, process func(row int, col int, reset bool)) {
	for startRow := 0; startRow < len(grid); startRow++ {
		reset := true
		for currentCol := 0; startRow-currentCol >= 0; currentCol++ {
			process(startRow-currentCol, currentCol, reset)
			reset = false
		}
	}
	for startCol := 1; startCol < len(grid); startCol++ {
		reset := true
		for currentRow := len(grid) - 1; startCol+len(grid)-currentRow-1 < len(grid); currentRow-- {
			process(currentRow, startCol+len(grid)-currentRow-1, reset)
			reset = false
		}
	}
}

func IterateDiagonalsNorthWestToSouthEast(grid [][]rune, process func(row int, col int, reset bool)) {
	for startCol := len(grid) - 1; startCol >= 0; startCol-- {
		reset := true
		for currentRow := 0; startCol+currentRow < len(grid); currentRow++ {
			process(currentRow, startCol+currentRow, reset)
			reset = false
		}
	}
	for startRow := 1; startRow < len(grid); startRow++ {
		reset := true
		for currentCol := 0; startRow+currentCol < len(grid); currentCol++ {
			process(startRow+currentCol, currentCol, reset)
			reset = false
		}
	}
}
