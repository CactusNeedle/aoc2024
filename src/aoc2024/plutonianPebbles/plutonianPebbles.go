package plutonianPebbles

import (
	"strconv"
	"strings"
)

func RunPartOne(progressUpdater func(fraction float64, intermediaryResult int), input string) {
	stones := parseInput(input)
	for i := 0; i < 25; i++ {
		progressUpdater(float64(i)/float64(75), countStones(stones))
		stones = blink(stones)
	}
	progressUpdater(1, countStones(stones))
	return
}

func RunPartTwo(progressUpdater func(fraction float64, intermediaryResult int), input string) {
	stones := parseInput(input)
	for i := 0; i < 75; i++ {
		progressUpdater(float64(i)/float64(75), countStones(stones))
		stones = blink(stones)
	}
	progressUpdater(1, countStones(stones))
	return
}

func blink(stones map[string]int) (newStones map[string]int) {
	newStones = make(map[string]int)
	for stone, _ := range stones {
		if stone == "0" {
			newStones["1"] = newStones["1"] + stones[stone]
		} else if len(stone)%2 == 0 {
			newStoneLeft := strings.TrimLeft(stone[:len(stone)/2], "0")
			if len(newStoneLeft) == 0 {
				newStoneLeft = "0"
			}
			newStones[newStoneLeft] = newStones[newStoneLeft] + stones[stone]
			newStoneRight := strings.TrimLeft(stone[len(stone)/2:], "0")
			if len(newStoneRight) == 0 {
				newStoneRight = "0"
			}
			newStones[newStoneRight] = newStones[newStoneRight] + stones[stone]
		} else {
			stoneInt, err := strconv.Atoi(stone)
			if err != nil {
				panic(err)
			}
			newStone := strconv.Itoa(stoneInt * 2024)
			newStones[newStone] = newStones[newStone] + stones[stone]
		}
	}
	return
}

func countStones(stones map[string]int) (result int) {
	for _, count := range stones {
		result = result + count
	}
	return
}

func parseInput(input string) map[string]int {
	result := make(map[string]int)
	for _, number := range strings.Split(input, " ") {
		result[number] = result[number] + 1
	}
	return result
}
