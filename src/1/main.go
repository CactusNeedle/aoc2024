package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	leftNumbers, rightNumbers := readFileNumbers()
	leftNumbers = mergeSort(leftNumbers)
	rightNumbers = mergeSort(rightNumbers)

	sumOfDifferences := calculateSumOfDifferences(leftNumbers, rightNumbers)
	fmt.Println(fmt.Sprintf("Sum of differences %d", sumOfDifferences)) // 1151792

	similarityScore := calculateSimilarityScore(leftNumbers, rightNumbers)
	fmt.Println(fmt.Sprintf("Similarity score %d", similarityScore)) //21790168
}

func calculateSumOfDifferences(leftNumbers []int, rightNumbers []int) (sumOfDifferences int) {
	for i := 0; i < len(leftNumbers); i++ {
		leftNumber := leftNumbers[i]
		rightNumber := rightNumbers[i]
		if leftNumber > rightNumber {
			sumOfDifferences = sumOfDifferences + (leftNumber - rightNumber)
		} else {
			sumOfDifferences = sumOfDifferences + (rightNumber - leftNumber)
		}
	}
	return
}

func calculateSimilarityScore(leftNumbers []int, rightNumbers []int) (similarityScore int) {
	rightNumbersCount := make(map[int]int)
	for _, rightNumber := range rightNumbers {
		rightNumbersCount[rightNumber] = rightNumbersCount[rightNumber] + 1
	}

	for _, leftNumber := range leftNumbers {
		similarityScore = similarityScore + leftNumber*rightNumbersCount[leftNumber]
	}
	return
}

func readFileNumbers() (leftNumbers []int, rightNumbers []int) {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var lineNumber = 0
	for scanner.Scan() {
		lineNumber = lineNumber + 1
		line := scanner.Text()
		parts := strings.Split(line, "   ")
		if len(parts) != 2 {
			panic(fmt.Sprintf("Unexpected amount of numbers on line %d", lineNumber))
		}
		leftNumber, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		rightNumber, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		leftNumbers = append(leftNumbers, leftNumber)
		rightNumbers = append(rightNumbers, rightNumber)
	}
	return
}

func mergeSort(array []int) (sortedArray []int) {
	sortedArray = make([]int, len(array))
	midpoint := len(array) / 2
	subArray1 := array[:midpoint]
	if len(subArray1) != 1 {
		subArray1 = mergeSort(subArray1)
	}
	subArray2 := array[midpoint:]
	if len(subArray2) != 1 {
		subArray2 = mergeSort(subArray2)
	}
	subArray1Index := 0
	subArray2Index := 0
	for subArray1Index+subArray2Index < len(subArray1)+len(subArray2) {
		if subArray1Index == len(subArray1) {
			sortedArray[subArray1Index+subArray2Index] = subArray2[subArray2Index]
			subArray2Index++
		} else if subArray2Index == len(subArray2) {
			sortedArray[subArray1Index+subArray2Index] = subArray1[subArray1Index]
			subArray1Index++
		} else {
			subArray1Value := subArray1[subArray1Index]
			subArray2Value := subArray2[subArray2Index]
			if subArray1Value > subArray2Value {
				sortedArray[subArray1Index+subArray2Index] = subArray2Value
				subArray2Index++
			} else {
				sortedArray[subArray1Index+subArray2Index] = subArray1Value
				subArray1Index++
			}
		}
	}
	return
}
