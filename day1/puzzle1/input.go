package main

import (
	"fmt"
	"strconv"
	"strings"
)

func parseInput(input []byte) (lists, error) {
	var parsed lists

	inputStr := string(input)
	inputStr = strings.TrimSpace(inputStr)
	lines := strings.Split(inputStr, "\n")

	for i := 0; i < len(lines); i++ {
		line := strings.Fields(lines[i])
		firstNum, err := strconv.Atoi(line[0])
		if err != nil {
			return lists{}, fmt.Errorf("error cleaning input")
		}

		secondNum, err := strconv.Atoi(line[1])
		if err != nil {
			return lists{}, fmt.Errorf("error cleaning input")
		}

		parsed.left = append(parsed.left, firstNum)
		parsed.right = append(parsed.right, secondNum)
	}

	return parsed, nil
}
