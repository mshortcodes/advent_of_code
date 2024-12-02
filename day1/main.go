package main

import (
	"fmt"
	"log"
	"os"
	"sort"
)

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
	similarity := lists.getSimilarity()

	fmt.Printf("Total distance: %d\n", distance)
	fmt.Printf("Similarity score: %d\n", similarity)
}
