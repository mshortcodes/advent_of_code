package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	inputBytes, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("error reading input file: %s", err)
	}

	reports, err := parseInput(inputBytes)
	if err != nil {
		log.Fatalf("error parsing data: %s", err)
	}

	safeReports := reports.countSafe()

	fmt.Printf("There are %d safe reports.\n", safeReports)
}
