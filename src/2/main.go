package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reports := readReports()

	safeReportCount := calculateSafeReportCount(reports)
	fmt.Println(fmt.Sprintf("Number of safe reports: %d", safeReportCount)) // 670

	safeReportWithOneRemovalCount := calculateSafeReportCountWithOneRemoval(reports)
	fmt.Println(fmt.Sprintf("Number of safe reports with one removal: %d", safeReportWithOneRemovalCount)) // 700
}

func calculateSafeReportCountWithOneRemoval(reports [][]int) (count int) {
	for _, report := range reports {
		if isSafeReport(report) {
			count++
		} else {
			for i := 0; i < len(report); i++ {
				reportWithRemoval := append(append(make([]int, 0), report[:i]...), report[i+1:]...)
				if isSafeReport(reportWithRemoval) {
					count++
					break
				}
			}
		}
	}
	return
}

func calculateSafeReportCount(reports [][]int) (count int) {
	for _, report := range reports {
		if isSafeReport(report) {
			count++
		}
	}
	return
}

func isSafeReport(report []int) bool {
	safe := true
	ascending := false
	if report[0] < report[1] {
		ascending = true
	}
	for i := 0; i < len(report)-1 && safe; i++ {
		difference := report[i] - report[i+1]
		if ascending {
			if difference < -3 || difference > -1 {
				safe = false
			}
		} else {
			if difference > 3 || difference < 1 {
				safe = false
			}
		}
	}
	return safe
}

func readReports() (reports [][]int) {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		report := make([]int, len(parts))
		for partIndex, part := range parts {
			number, err := strconv.Atoi(part)
			if err != nil {
				panic(err)
			}
			report[partIndex] = number
		}
		reports = append(reports, report)
	}
	return
}
