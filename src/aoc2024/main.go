package main

import (
	"aoc2024/ceresSearch"
	"aoc2024/guardianGallivant"
	"aoc2024/historianHysteria"
	"aoc2024/mullItOver"
	"aoc2024/printQueue"
	"aoc2024/redNosedReports"
	"bufio"
	"fmt"
	"math"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

var options = []string{
	"(Day 1 - Part 1) Historian Hysteria",
	"(Day 1 - Part 2) Historian Hysteria",
	"(Day 2 - Part 1) Red-Nosed Reports",
	"(Day 2 - Part 2) Red-Nosed Reports",
	"(Day 3 - Part 1) Mull It Over",
	"(Day 3 - Part 2) Mull It Over",
	"(Day 4 - Part 1) Ceres Search",
	"(Day 4 - Part 2) Ceres Search",
	"(Day 5 - Part 1) Print Queue",
	"(Day 5 - Part 2) Print Queue",
	"(Day 6 - Part 1) Guardian Gallivant",
	"(Day 6 - Part 2) Guardian Gallivant",
}

func main() {
	mainMenu()
}

func mainMenu() {
	for {
		// Display menu
		clearScreen()
		fmt.Println("Welcome to advent of code 2024!")
		fmt.Println("Which challenge would you like to run?")
		for i, option := range options {
			fmt.Printf("%d: %s\n", i, option)
		}

		// Process user input
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Selection: ")
		scanner.Scan()
		selectedOption, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			time.Sleep(2 * time.Second)
			continue
		}
		if selectedOption < 0 || selectedOption >= len(options) {
			fmt.Println("Invalid option. Please select a valid challenge.")
			time.Sleep(2 * time.Second)
			continue
		}
		clearScreen()
		runOption(selectedOption)
		fmt.Print("Press any button to continue...")
		scanner.Scan()
		if err != nil {
			panic(err)
		}
	}
}

func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

func runOption(option int) {
	start := time.Now()

	progressUpdater := func(fraction float64, intermediaryResult int) {
		timeSinceStart := time.Since(start)
		printProgress(timeSinceStart, fraction, intermediaryResult)
	}

	switch option {
	case 0:
		historianHysteria.RunPartOne(progressUpdater, readFile("./historianHysteria/input.txt")) // 1151792
	case 1:
		historianHysteria.RunPartTwo(progressUpdater, readFile("./historianHysteria/input.txt")) // 21790168
	case 2:
		redNosedReports.RunPartOne(progressUpdater, readFile("./redNosedReports/input.txt")) // 670
	case 3:
		redNosedReports.RunPartTwo(progressUpdater, readFile("./redNosedReports/input.txt")) // 700
	case 4:
		mullItOver.RunPartOne(progressUpdater, readFile("./mullItOver/input.txt")) // 184576302
	case 5:
		mullItOver.RunPartTwo(progressUpdater, readFile("./mullItOver/input.txt")) // 118173507
	case 6:
		ceresSearch.RunPartOne(progressUpdater, readFile("./ceresSearch/input.txt")) // 2543
	case 7:
		ceresSearch.RunPartTwo(progressUpdater, readFile("./ceresSearch/input.txt")) // 1930
	case 8:
		printQueue.RunPartOne(progressUpdater, readFile("./printQueue/input.txt")) // 5268
	case 9:
		printQueue.RunPartTwo(progressUpdater, readFile("./printQueue/input.txt")) // 5799
	case 10:
		guardianGallivant.RunPartOne(progressUpdater, readFile("./guardianGallivant/input.txt")) // 4454
	case 11:
		guardianGallivant.RunPartTwo(progressUpdater, readFile("./guardianGallivant/input.txt")) // 1503
	}
}

func readFile(path string) (content string) {
	b, err := os.ReadFile(path)
	if err != nil {
		fmt.Print(err)
	}
	return string(b)
}

func printProgress(timeSinceStart time.Duration, fraction float64, intermediaryResult int) {
	completed := fraction * float64(50)
	remaining := float64(50) - completed
	fmt.Printf("\033[F\033[K") // Move up and clear the line
	fmt.Printf("\033[F\033[K") // Move up and clear the line
	fmt.Println(fmt.Sprintf("Result: %d", intermediaryResult))
	fmt.Println(fmt.Sprintf("[%s%s] %ds - %.2f%%", strings.Repeat("=", int(completed)),
		strings.Repeat(" ", int(remaining)), int64(math.Round(timeSinceStart.Seconds())), fraction*100))
}