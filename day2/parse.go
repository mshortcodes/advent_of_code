package main

import (
	"fmt"
	"strconv"
	"strings"
)

// parseInput parses the data into a slice of reports (slices).
func parseInput(input []byte) (reports, error) {
	var parsed reports

	inputStr := string(input)
	lines := strings.Split(inputStr, "\n")

	for _, line := range lines {
		var newLine []int
		nums := strings.Fields(line)
		for _, num := range nums {
			num, err := strconv.Atoi(num)
			if err != nil {
				return nil, fmt.Errorf("error converting to int: %s", err)
			}
			newLine = append(newLine, num)
		}
		parsed = append(parsed, newLine)
	}

	return parsed, nil
}
