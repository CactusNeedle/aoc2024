package bridgeRepair

import (
	"strconv"
	"strings"
)

const (
	ADD      int = 0
	MULTIPLY     = 1
	ELEPHANT     = 2
)

func RunPartOne(progressUpdater func(fraction float64, intermediaryResult int), input string) {
	results, arguments := parseInput(input)
	countPotentialValidCalculations(progressUpdater, results, arguments, []int{ADD, MULTIPLY})

}

func RunPartTwo(progressUpdater func(fraction float64, intermediaryResult int), input string) {
	results, arguments := parseInput(input)
	countPotentialValidCalculations(progressUpdater, results, arguments, []int{ADD, MULTIPLY, ELEPHANT})
}

func countPotentialValidCalculations(progressUpdater func(fraction float64, intermediaryResult int), results []int,
	arguments [][]int, validOperators []int) {
	count := 0
	for index := 0; index < len(results); index++ {
		operatorCombinations := getOperatorCombinations(arguments[index], validOperators)
		for _, operators := range operatorCombinations {
			if isValidCalculation(results[index], arguments[index], operators) {
				count = count + results[index]
				break
			}
		}
		if index%10 == 0 {
			progressUpdater(float64(index)/float64(len(results)), count)
		}
	}
	progressUpdater(1, count)
}

func getOperatorCombinations(arguments []int, validOperators []int) (combinations [][]int) {
	length := len(arguments) - 1
	totalCombinations := pow(len(validOperators), length)
	combinations = make([][]int, totalCombinations)
	for i := 0; i < totalCombinations; i++ {
		combination := make([]int, length)
		for j := 0; j < length; j++ {
			combination[j] = validOperators[(i/pow(len(validOperators), length-j-1))%len(validOperators)]
		}
		combinations[i] = combination
	}
	return
}

func pow(base, exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}

func isValidCalculation(result int, arguments []int, operators []int) bool {
	return getCalculationResult(arguments, operators) == result
}

func getCalculationResult(arguments []int, operators []int) (calculationResult int) {
	calculationResult = arguments[0]
	for index := 1; index < len(arguments); index++ {
		if operators[index-1] == ADD {
			calculationResult = calculationResult + arguments[index]
		} else if operators[index-1] == MULTIPLY {
			calculationResult = calculationResult * arguments[index]
		} else if operators[index-1] == ELEPHANT {
			newCalculationResult, err := strconv.Atoi(strconv.Itoa(calculationResult) + strconv.Itoa(arguments[index]))
			if err != nil {
				panic(err)
			}
			calculationResult = newCalculationResult
		} else {
			panic("This should never happen")
		}
	}
	return
}

func parseInput(input string) (calculationResults []int, calculationArguments [][]int) {
	for _, line := range strings.Split(input, "\r\n") {
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			panic("Incorrect input!")
		}
		calculationResult, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		calculationResults = append(calculationResults, calculationResult)
		subParts := strings.Split(strings.Trim(parts[1], " "), " ")
		arguments := make([]int, 0)
		for _, subPart := range subParts {
			calculationArgument, err := strconv.Atoi(subPart)
			if err != nil {
				panic(err)
			}
			arguments = append(arguments, calculationArgument)
		}
		calculationArguments = append(calculationArguments, arguments)
	}
	return
}
