// --- Day DAY: Title ---
// https://adventofcode.com/YEAR/day/DAY
package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	input, err := os.Open("./YEAR/DAY/input.txt")
	if err != nil {
		log.Printf("Failed to open input.txt: %v", err)
		return
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		src := scanner.Text()
		log.Print(src)
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Failed to scan seeds from input.txt: %v", err)
		return
	}

	var firstPartResult, secondPartResult int

	log.Printf("Part 1 result: %d", firstPartResult)
	log.Printf("Part 2 result: %d", secondPartResult)
}
