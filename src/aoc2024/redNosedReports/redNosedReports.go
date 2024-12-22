package redNosedReports

import (
	"strconv"
	"strings"
)

func RunPartOne(progressUpdater func(fraction float64, intermediaryResult int), input string) {
	reports := parseInput(input)

	safeReportCount := 0
	for i, report := range reports {
		if isSafeReport(report) {
			safeReportCount++
		}
		progressUpdater(float64(i)/float64(len(input)), safeReportCount)
	}
	progressUpdater(1, safeReportCount)
	return
}

func RunPartTwo(progressUpdater func(fraction float64, intermediaryResult int), input string) {
	reports := parseInput(input)

	safeReportCount := 0
	for i, report := range reports {
		if isSafeReport(report) {
			safeReportCount++
		} else {
			for i := 0; i < len(report); i++ {
				reportWithRemoval := append(append(make([]int, 0), report[:i]...), report[i+1:]...)
				if isSafeReport(reportWithRemoval) {
					safeReportCount++
					break
				}
			}
		}
		progressUpdater(float64(i)/float64(len(input)), safeReportCount)
	}
	progressUpdater(1, safeReportCount)
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

func parseInput(input string) (reports [][]int) {
	for _, line := range strings.Split(input, "\r\n") {
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
