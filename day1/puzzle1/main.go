package main

import (
	"fmt"
	"log"
	"os"
	"sort"
)

type lists struct {
	left  []int
	right []int
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("error reading input file: %s", err)
	}

	lists, err := parseInput(input)
	if err != nil {
		log.Fatalf("error parsing input: %s", err)
	}

	sort.Ints(lists.left)
	sort.Ints(lists.right)

	distance := lists.getDistance()
	fmt.Printf("Total distance between left and right lists:\n%d\n", distance)
}

func (l *lists) getDistance() int {
	var distance int

	for i := 0; i < len(l.left); i++ {
		switch {
		case l.left[i] < l.right[i]:
			distance += l.right[i] - l.left[i]
		case l.left[i] > l.right[i]:
			distance += l.left[i] - l.right[i]
		default:
		}
	}

	return distance
}
