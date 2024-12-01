package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"slices"
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
	lines, err := readLinesFromFile(filepath)
	handleError(err)
	numbers, err := parseListsFromLines(lines)
	handleError(err)
	sorted1stElements := obtainNthElementFromNumbers(numbers, 0)
	sorted2ndElements := obtainNthElementFromNumbers(numbers, 1)
	differences := make([]int, 0)
	for index := range sorted1stElements {
		differences = append(differences, int(math.Abs(float64(sorted1stElements[index]-sorted2ndElements[index]))))
	}
	sum := 0
	for _, difference := range differences {
		sum += difference
	}
	fmt.Printf("Sum of differences: %d", sum)
}

func readLinesFromFile(path string) ([]string, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	data := string(bytes)
	return strings.Split(data, "\n"), nil
}

func parseListsFromLines(lines []string) ([]([]int), error) {
	var rawNumbers []([]string)
	for _, line := range lines {
		numberRegexp := regexp.MustCompile(`\d+`)
		numbersOnLine := numberRegexp.FindAllString(line, -1)
		rawNumbers = append(rawNumbers, numbersOnLine)
	}

	var numbers []([]int)
	for _, record := range rawNumbers {
		newRecord := make([]int, 0)
		for _, rawNumber := range record {
			parsedInt, err := strconv.Atoi(rawNumber)
			if err != nil {
				return nil, err
			}
			newRecord = append(newRecord, parsedInt)
		}
		numbers = append(numbers, newRecord)
	}
	return numbers, nil
}

func obtainNthElementFromNumbers(numbers []([]int), index int16) []int {
	output := make([]int, 0)
	for _, number := range numbers {
		output = append(output, number[index])
	}
	slices.Sort(output)
	return output
}
