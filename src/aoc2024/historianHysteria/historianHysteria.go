package historianHysteria

import (
	"strconv"
	"strings"
)

func RunPartOne(progressUpdater func(fraction float64, intermediaryResult int), input string) {
	leftNumbers, rightNumbers := parseInput(input)
	leftNumbers = mergeSort(leftNumbers)
	rightNumbers = mergeSort(rightNumbers)

	sumOfDifferences := 0
	for i := 0; i < len(leftNumbers); i++ {
		leftNumber := leftNumbers[i]
		rightNumber := rightNumbers[i]
		if leftNumber > rightNumber {
			sumOfDifferences = sumOfDifferences + (leftNumber - rightNumber)
		} else {
			sumOfDifferences = sumOfDifferences + (rightNumber - leftNumber)
		}
		progressUpdater(float64(i)/float64(len(input)), sumOfDifferences)
	}
	progressUpdater(1, sumOfDifferences)
	return
}

func RunPartTwo(progressUpdater func(fraction float64, intermediaryResult int), input string) {
	leftNumbers, rightNumbers := parseInput(input)
	leftNumbers = mergeSort(leftNumbers)
	rightNumbers = mergeSort(rightNumbers)

	rightNumbersCount := make(map[int]int)
	for i, rightNumber := range rightNumbers {
		rightNumbersCount[rightNumber] = rightNumbersCount[rightNumber] + 1
		progressUpdater(float64(i)/float64(len(input)*2), 0)
	}

	similarityScore := 0
	for i, leftNumber := range leftNumbers {
		similarityScore = similarityScore + leftNumber*rightNumbersCount[leftNumber]
		progressUpdater(float64(len(input)+i)/float64(len(input)*2), similarityScore)
	}
	progressUpdater(1, similarityScore)
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

func parseInput(input string) (leftNumbers []int, rightNumbers []int) {
	for _, line := range strings.Split(input, "\r\n") {
		parts := strings.Split(line, "   ")
		if len(parts) != 2 {
			panic("Invalid input")
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
