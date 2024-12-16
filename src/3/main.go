package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	characters := readCharacters()

	sumOfMultiplications := calculateSumOfMultiplications(characters)
	fmt.Println(fmt.Sprintf("Sum of multiplications: %d", sumOfMultiplications)) // 184576302

	sumOfMultiplicationsWithDoAndDont := calculateSumOfMultiplicationsWithDoAndDont(characters)
	fmt.Println(fmt.Sprintf("Sum of multiplications with do and don't: %d", sumOfMultiplicationsWithDoAndDont)) //
}

func calculateSumOfMultiplications(characters []rune) (sumOfMultiplications int) {
	for i := 0; i < len(characters); i++ {
		processMulStart(characters, &i, &sumOfMultiplications)
	}
	return
}

func calculateSumOfMultiplicationsWithDoAndDont(characters []rune) (sumOfMultiplications int) {
	do := true
	for i := 0; i < len(characters); i++ {
		if do && characters[i] == 'm' {
			processMulStart(characters, &i, &sumOfMultiplications)
		} else if characters[i] == 'd' {
			processDoOrDontStart(characters, &i, &do)
		}
	}
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

func readCharacters() []rune {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	data, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	return []rune(string(data))
}
