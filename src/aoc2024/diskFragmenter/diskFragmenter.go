package diskFragmenter

import (
	"strconv"
)

func RunPartOne(progressUpdater func(fraction float64, intermediaryResult int), input string) {
	unwrappedDiskMap := unwrap(parseInput(input))

	lastUsedFreeSpaceIndex := 0
	for index := len(unwrappedDiskMap) - 1; index >= 0; index-- {
		for unwrappedDiskMap[lastUsedFreeSpaceIndex] != -1 {
			lastUsedFreeSpaceIndex++
		}
		if lastUsedFreeSpaceIndex < index && unwrappedDiskMap[index] >= 0 {
			unwrappedDiskMap[lastUsedFreeSpaceIndex] = unwrappedDiskMap[index]
			unwrappedDiskMap[index] = -1
		}
		if index%100 == 0 {
			progressUpdater(float64(len(unwrappedDiskMap)-index)/float64(len(unwrappedDiskMap)), -1)
		}
	}
	progressUpdater(1, calculateCheckSum(unwrappedDiskMap))
}

func RunPartTwo(progressUpdater func(fraction float64, intermediaryResult int), input string) {
	unwrappedDiskMap := unwrap(parseInput(input))

	currentFileNumber := -1
	currentFileLength := 0
	for index := len(unwrappedDiskMap) - 1; index >= 0; index-- {
		if unwrappedDiskMap[index] != currentFileNumber {
			if currentFileNumber != -1 {
				freeSpaceStartingIndex := findContinuousFreeSpace(unwrappedDiskMap, currentFileLength, index)
				if freeSpaceStartingIndex != -1 {
					for freeSpaceIndex := 0; freeSpaceIndex < currentFileLength; freeSpaceIndex++ {
						unwrappedDiskMap[freeSpaceStartingIndex+freeSpaceIndex] = unwrappedDiskMap[index+freeSpaceIndex+1]
						unwrappedDiskMap[index+freeSpaceIndex+1] = -1
					}
				}
			}
			currentFileNumber = unwrappedDiskMap[index]
			currentFileLength = 1
		} else {
			currentFileLength++
		}
		if index%100 == 0 {
			progressUpdater(float64(len(unwrappedDiskMap)-index)/float64(len(unwrappedDiskMap)), -1)
		}
	}
	progressUpdater(1, calculateCheckSum(unwrappedDiskMap))
}

func unwrap(diskMap []rune) (unwrappedDiskMap []int) {
	for index, value := range diskMap {
		repeatCount, err := strconv.Atoi(string(value))
		if err != nil {
			panic(err)
		}
		for repeat := 0; repeat < repeatCount; repeat++ {
			if index%2 == 0 {
				unwrappedDiskMap = append(unwrappedDiskMap, index/2)
			} else {
				unwrappedDiskMap = append(unwrappedDiskMap, -1)
			}
		}
	}
	return
}

func findContinuousFreeSpace(diskMap []int, length int, maxIndex int) int {
	currentFreeSpaceStartingIndex := -1
	currentFreeSpaceLength := 0
	for index, value := range diskMap {
		if index > maxIndex {
			return -1
		}
		if value == -1 {
			if currentFreeSpaceStartingIndex == -1 {
				currentFreeSpaceStartingIndex = index
			}
			currentFreeSpaceLength++
			if currentFreeSpaceLength == length {
				return currentFreeSpaceStartingIndex
			}
		} else {
			currentFreeSpaceStartingIndex = -1
			currentFreeSpaceLength = 0
		}
	}
	return -1
}

func calculateCheckSum(diskMap []int) (checkSum int) {
	for index, value := range diskMap {
		if value >= 0 {
			checkSum += index * value
		}
	}
	return
}

func parseInput(input string) []rune {
	return []rune(input)
}
