package main

import (
	"fmt"
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
