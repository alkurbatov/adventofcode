// --- Day 2: 1202 Program Alarm ---
// https://adventofcode.com/2019/day/2
package main

import (
	"bufio"
	"log"
	"os"

	"github.com/alkurbatov/adventofcode/internal/parsers"
)

func readInput(path string) ([]int, error) {
	input, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Scan()

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return parsers.ReadNumbers(scanner.Text(), ",")
}

func runIntcode(program []int) int {
	i := 0

	for i < len(program) {
		instruction := program[i]

		if instruction == 99 {
			break
		}

		p1 := program[i+1]
		p2 := program[i+2]
		p3 := program[i+3]

		switch instruction {
		case 1:
			program[p3] = program[p1] + program[p2]
		case 2:
			program[p3] = program[p1] * program[p2]
		}

		i += 4
	}

	return program[0]
}

func restoreGravityAssist(program []int) int {
	programCopy := append([]int(nil), program...)

	// Apply rules
	programCopy[1] = 12 // noun
	programCopy[2] = 2  // verb

	return runIntcode(programCopy)
}

func completeGravityAssist(program []int) (noun, verb int) {
	for noun = 0; noun < 100; noun++ {
		for verb = 0; verb < 100; verb++ {
			programCopy := append([]int(nil), program...)

			// Apply rules
			programCopy[1] = noun
			programCopy[2] = verb

			if runIntcode(programCopy) == 19690720 {
				return noun, verb
			}
		}
	}

	return noun, verb
}

func main() {
	program, err := readInput("./2019/2/input.txt")
	if err != nil {
		log.Printf("Failed to read data from input.txt: %v", err)
		return
	}

	firstPartResult := restoreGravityAssist(program)
	log.Printf("Part 1 result: %d", firstPartResult)

	noun, verb := completeGravityAssist(program)
	secondPartResult := 100*noun + verb
	log.Printf("Part 2 result: %d", secondPartResult)
}
