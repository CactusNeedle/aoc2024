package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := readInput()

	matches := countMatches(input)
	fmt.Println(fmt.Sprintf("Amount of matches: %d", matches)) // 2543

	crossMatches := countCrossMatches(input)
	fmt.Println(fmt.Sprintf("Amount of cross matches: %d", crossMatches)) // 1930
}

func countMatches(input [][]rune) (count int) {
	for row := 0; row < len(input); row++ {
		currentPartialMatch := ""
		for col := 0; col < len(input); col++ {
			processPotentialMatch(&currentPartialMatch, &count, input[row][col])
		}
	}
	for col := 0; col < len(input); col++ {
		currentPartialMatch := ""
		for row := 0; row < len(input); row++ {
			processPotentialMatch(&currentPartialMatch, &count, input[row][col])
		}
	}
	for startRow := 0; startRow < len(input); startRow++ {
		currentPartialMatch := ""
		for currentRow := startRow; currentRow >= 0; currentRow-- {
			processPotentialMatch(&currentPartialMatch, &count, input[currentRow][startRow-currentRow])
		}
	}
	for startRow := len(input) - 1; startRow >= 0; startRow-- {
		currentPartialMatch := ""
		for currentRow := startRow; currentRow < len(input); currentRow++ {
			processPotentialMatch(&currentPartialMatch, &count, input[currentRow][currentRow-startRow])
		}
	}
	for startCol := len(input) - 1; startCol > 0; startCol-- {
		currentPartialMatch := ""
		for currentCol := startCol; currentCol < len(input); currentCol++ {
			processPotentialMatch(&currentPartialMatch, &count, input[currentCol-startCol][currentCol])
		}
	}
	for startCol := len(input) - 1; startCol > 0; startCol-- {
		currentPartialMatch := ""
		for currentCol := startCol; currentCol < len(input); currentCol++ {
			processPotentialMatch(&currentPartialMatch, &count, input[len(input)+startCol-currentCol-1][currentCol])
		}
	}
	return
}

func countCrossMatches(input [][]rune) (count int) {
	for startRow := 0; startRow < len(input); startRow++ {
		currentPartialMatch := ""
		for currentRow := startRow; currentRow >= 0; currentRow-- {
			processPotentialCrossMatch(input, currentRow, startRow-currentRow, &currentPartialMatch, &count)
		}
	}
	for startCol := 1; startCol <= len(input); startCol++ {
		currentPartialMatch := ""
		for currentCol := startCol; currentCol < len(input); currentCol++ {
			processPotentialCrossMatch(input, len(input)+startCol-currentCol-1, currentCol, &currentPartialMatch, &count)
		}
	}
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
