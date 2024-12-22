package printQueue

import (
	"slices"
	"strconv"
	"strings"
)

func RunPartOne(progressUpdater func(fraction float64, intermediaryResult int), input string) {
	rules, updates := parseInput(input)

	sum := 0
	for i, update := range updates {
		if isValidUpdate(rules, update) {
			sum = sum + update[(len(update)-1)/2]
		}
		progressUpdater(float64(i)/float64(len(updates)), sum)
	}
	progressUpdater(1, sum)
	return
}

func RunPartTwo(progressUpdater func(fraction float64, intermediaryResult int), input string) {
	rules, updates := parseInput(input)

	sum := 0
	for i, update := range updates {
		if !isValidUpdate(rules, update) {
			sum = sum + createValidUpdate(rules, update)[(len(update)-1)/2]
		}
		progressUpdater(float64(i)/float64(len(updates)), sum)
	}
	progressUpdater(1, sum)
	return
}

func isValidUpdate(rules map[int][]int, update []int) bool {
	for currentIndex := 0; currentIndex < len(update); currentIndex++ {
		for comparingIndex := currentIndex; comparingIndex >= 0; comparingIndex-- {
			if slices.Contains(rules[update[currentIndex]], update[comparingIndex]) {
				return false
			}
		}
	}
	return true
}

func createValidUpdate(rules map[int][]int, update []int) (validUpdate []int) {
	for _, page := range update {
		for index := 0; index < len(update) || len(validUpdate) == 0; index++ {
			potentialValidUpdate := append([]int(nil), validUpdate...)
			potentialValidUpdate = slices.Insert(potentialValidUpdate, index, page)
			if isValidUpdate(rules, potentialValidUpdate) {
				validUpdate = potentialValidUpdate
				break
			}
		}
	}
	return
}

func parseInput(input string) (rules map[int][]int, updates [][]int) {
	rules = make(map[int][]int)
	readingRules := true
	for _, line := range strings.Split(input, "\r\n") {
		if line == "" {
			readingRules = false
		} else if readingRules {
			splitLine := strings.Split(line, "|")
			if len(splitLine) != 2 {
				panic("Invalid input!")
			}
			number1, err := strconv.Atoi(splitLine[0])
			if err != nil {
				panic(err)
			}
			number2, err := strconv.Atoi(splitLine[1])
			if err != nil {
				panic(err)
			}
			if rules[number1] == nil {
				rules[number1] = make([]int, 1)
			}
			rules[number1] = append(rules[number1], number2)
		} else {
			update := make([]int, 0)
			splitLine := strings.Split(line, ",")
			for _, page := range splitLine {
				pageNumber, err := strconv.Atoi(page)
				if err != nil {
					panic(err)
				}
				update = append(update, pageNumber)
			}
			updates = append(updates, update)
		}
	}
	return
}
