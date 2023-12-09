// --- Day DAY: Title ---
// https://adventofcode.com/YEAR/day/DAY
package main

import (
	"bufio"
	"log"
	"os"
)

func readInput(path string) ([]int, error) {
	input, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer input.Close()

	output := make([]int, 0)
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		src := scanner.Text()
		log.Print(src)
	}

	return output, scanner.Err()
}

func main() {
	input, err := readInput("./YEAR/DAY/input.txt")
	if err != nil {
		log.Printf("Failed to read data from input.txt: %v", err)
		return
	}

	log.Print(input)

	var firstPartResult, secondPartResult int

	log.Printf("Part 1 result: %d", firstPartResult)
	log.Printf("Part 2 result: %d", secondPartResult)
}
