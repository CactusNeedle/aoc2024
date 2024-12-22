package mullItOver

import (
	"strconv"
)

func RunPartOne(progressUpdater func(fraction float64, intermediaryResult int), input string) {
	characters := parseInput(input)

	sumOfMultiplications := 0
	for i := 0; i < len(characters); i++ {
		processMulStart(characters, &i, &sumOfMultiplications)
		if i%100 == 0 {
			progressUpdater(float64(i)/float64(len(input)), sumOfMultiplications)
		}
	}
	progressUpdater(1, sumOfMultiplications)
	return
}

func RunPartTwo(progressUpdater func(fraction float64, intermediaryResult int), input string) {
	characters := parseInput(input)

	do := true
	sumOfMultiplications := 0
	for i := 0; i < len(characters); i++ {
		if do && characters[i] == 'm' {
			processMulStart(characters, &i, &sumOfMultiplications)
		} else if characters[i] == 'd' {
			processDoOrDontStart(characters, &i, &do)
		}
		if i%100 == 0 {
			progressUpdater(float64(i)/float64(len(input)), sumOfMultiplications)
		}
	}
	progressUpdater(1, sumOfMultiplications)
	return
}

func processMulStart(characters []rune, index *int, sumOfMultiplications *int) {
	if characters[*index] != 'm' {
		return
	}
	*index++
	if characters[*index] != 'u' {
		return
	}
	*index++
	if characters[*index] != 'l' {
		return
	}
	*index++
	if characters[*index] != '(' {
		return
	}
	*index++
	firstNumber, err := findNumber(characters, index)
	if firstNumber > 999 || err != nil {
		return
	}
	if characters[*index] != ',' {
		return
	}
	*index++
	secondNumber, err := findNumber(characters, index)
	if secondNumber > 999 || err != nil {
		return
	}
	if characters[*index] != ')' {
		return
	}
	*sumOfMultiplications = *sumOfMultiplications + firstNumber*secondNumber
}

func processDoOrDontStart(characters []rune, index *int, do *bool) {
	if characters[*index] != 'd' {
		return
	}
	*index++
	if characters[*index] != 'o' {
		return
	}
	*index++
	if characters[*index] == '(' {
		*index++
		if characters[*index] != ')' {
			return
		}
		*do = true
	} else if characters[*index] == 'n' {
		*index++
		if characters[*index] != '\'' {
			return
		}
		*index++
		if characters[*index] != 't' {
			return
		}
		*index++
		if characters[*index] != '(' {
			return
		}
		*index++
		if characters[*index] != ')' {
			return
		}
		*do = false
	} else {
		return
	}
}

func findNumber(characters []rune, index *int) (number int, err any) {
	numberString := ""
	for *index < len(characters) {
		_, err = strconv.Atoi(string(characters[*index]))
		if err != nil {
			return strconv.Atoi(numberString)
		}
		numberString = numberString + string(characters[*index])
		*index++
	}
	return strconv.Atoi(numberString)
}

func parseInput(input string) []rune {
	return []rune(input)
}
