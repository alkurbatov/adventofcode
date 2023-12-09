// --- Day 1: The Tyranny of the Rocket Equation ---
// https://adventofcode.com/2019/day/1
package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
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
		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}

		output = append(output, value)
	}

	return output, scanner.Err()
}

func calcFuel(mass int) int {
	return max(mass/3-2, 0)
}

func calcFuelWithAdditionalFuel(mass int) int {
	result := 0

	for mass > 0 {
		fuel := calcFuel(mass)

		mass = fuel
		result += fuel
	}

	return result
}

func main() {
	input, err := readInput("./2019/1/input.txt")
	if err != nil {
		log.Printf("Failed to read data from input.txt: %v", err)
		return
	}

	var firstPartResult, secondPartResult int

	for _, mass := range input {
		firstPartResult += calcFuel(mass)
		secondPartResult += calcFuelWithAdditionalFuel(mass)
	}

	log.Printf("Part 1 result: %d", firstPartResult)
	log.Printf("Part 2 result: %d", secondPartResult)
}
