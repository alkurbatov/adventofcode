// --- Day 9: Mirage Maintenance ---
// https://adventofcode.com/2023/day/9
package main

import (
	"bufio"
	"log"
	"os"

	"github.com/alkurbatov/adventofcode/internal/parsers"
)

func readInput(path string) ([][]int, error) {
	input, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	result := make([][]int, 0)

	for scanner.Scan() {
		sequence, err := parsers.ReadNumbers(scanner.Text(), " ")
		if err != nil {
			return nil, err
		}

		result = append(result, sequence)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func calcPredictions(values []int) (left, right int) {
	stepDiff := make([]int, 0)
	allZeroes := true

	for i := 1; i < len(values); i++ {
		nextValue := values[i] - values[i-1]

		if nextValue != 0 {
			allZeroes = false
		}

		stepDiff = append(stepDiff, nextValue)
	}

	if allZeroes {
		return values[0], values[len(values)-1]
	}

	left, right = calcPredictions(stepDiff)

	return values[0] - left, right + values[len(values)-1]
}

func main() {
	input, err := readInput("2023/9/input.txt")
	if err != nil {
		log.Printf("Failed to read data from input.txt: %v", err)
		return
	}

	var firstPartResult, secondPartResult int

	for _, history := range input {
		left, right := calcPredictions(history)

		firstPartResult += right
		secondPartResult += left
	}

	log.Printf("Part 1 result: %d", firstPartResult)
	log.Printf("Part 2 result: %d", secondPartResult)
}
