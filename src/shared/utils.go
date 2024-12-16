package shared

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadFileNumbers(path string) (leftNumbers []int, rightNumbers []int) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
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
