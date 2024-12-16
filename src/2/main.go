package main

import (
	"aoc2024/shared"
	"fmt"
)

func main() {
	leftNumbers, rightNumbers := shared.ReadFileNumbers("../resources/input.txt")

	rightNumbersCount := make(map[int]int)
	for _, rightNumber := range rightNumbers {
		rightNumbersCount[rightNumber] = rightNumbersCount[rightNumber] + 1
	}

	result := 0
	for _, leftNumber := range leftNumbers {
		result = result + leftNumber*rightNumbersCount[leftNumber]
	}

	fmt.Println(result) //21790168
}
