package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func handleError(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var filepath string
	_, err := fmt.Scan(&filepath)
	handleError(err)
	lines, err := readLines(filepath)
	handleError(err)
	levels, err := parseLevels(lines)
	handleError(err)
	safeReportCount := countSafeReports(levels)
	fmt.Printf("Safe report count: %d", safeReportCount)

}

func readLines(filepath string) ([]string, error) {
	raw, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	data := string(raw)
	return strings.Split(data, "\n"), nil
}

func parseLevels(lines []string) ([]([]int), error) {
	reports := make([]([]string), 0)
	for _, line := range lines {
		reports = append(reports, strings.Split(line, " "))
	}
	parsedReports := make([]([]int), 0)
	for _, report := range reports {
		parsedReport := make([]int, 0)
		for _, level := range report {
			parsedLevel, err := strconv.Atoi(level)
			if err != nil {
				return nil, err
			}
			parsedReport = append(parsedReport, parsedLevel)
		}
		parsedReports = append(parsedReports, parsedReport)
	}
	return parsedReports, nil
}

func isReportSafe(report []int) bool {
	shouldBeAscending := report[0] < report[1]
	for index := range report {
		if index == len(report)-1 {
			break
		}
		nextIndex := index + 1
		if shouldBeAscending && report[index] > report[nextIndex] {
			return false
		} else if !shouldBeAscending && report[index] < report[nextIndex] {
			return false
		}
		difference := math.Abs(float64(report[index]) - float64(report[nextIndex]))
		if difference < 1 || difference > 3 {
			return false
		}
	}
	return true
}

func countSafeReports(reports []([]int)) int {
	count := 0
	for _, report := range reports {
		if isReportSafe(report) {
			count++
		}
	}
	return count
}
