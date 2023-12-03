// --- Day 3: Gear Ratios ---
// https://adventofcode.com/2023/day/3
package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"unicode"
)

type Point struct {
	i int
	j int
}

func main() {
	input, err := os.Open("./2023/3/input.txt")
	if err != nil {
		log.Printf("Failed to open input.txt: %v", err)
		return
	}
	defer input.Close()

	var i int

	schematic := make([][]rune, 0)
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		schematic = append(schematic, make([]rune, 0))

		for _, char := range scanner.Text() {
			schematic[i] = append(schematic[i], char)
		}

		i++
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Failed to scan input.txt: %v", err)
		return
	}

	var firstPartResult, secondPartResult int

	var adjacent bool

	var gearPos *Point

	partNum := make([]rune, 0)
	gears := make(map[Point][]int)

	for i := range schematic {
		for j, char := range schematic[i] {
			if !unicode.IsDigit(char) {
				if len(partNum) == 0 {
					continue
				}

				num, err := strconv.Atoi(string(partNum))
				if err != nil {
					log.Printf("Failed to parse string to int: %v", err)
					return
				}

				if adjacent {
					firstPartResult += num
				}

				if gearPos != nil {
					if _, ok := gears[*gearPos]; !ok {
						gears[*gearPos] = make([]int, 0)
					}

					gears[*gearPos] = append(gears[*gearPos], num)
				}

				partNum = make([]rune, 0)
				adjacent = false
				gearPos = nil

				continue
			}

			partNum = append(partNum, char)

			neighbors := make([]Point, 0)

			if i > 0 { // up
				neighbors = append(neighbors, Point{i - 1, j})
			}

			if i < len(schematic)-1 { // down
				neighbors = append(neighbors, Point{i + 1, j})
			}

			if j > 0 { // left
				neighbors = append(neighbors, Point{i, j - 1})
			}

			if j < len(schematic[0])-1 { // right
				neighbors = append(neighbors, Point{i, j + 1})
			}

			if j > 0 {
				if i > 0 { // upper-left
					neighbors = append(neighbors, Point{i - 1, j - 1})
				}

				if i < len(schematic)-1 { // lower-left
					neighbors = append(neighbors, Point{i + 1, j - 1})
				}
			}

			if j < len(schematic[0])-1 {
				if i > 0 { // upper-right
					neighbors = append(neighbors, Point{i - 1, j + 1})
				}

				if i < len(schematic)-1 { // lower-right
					neighbors = append(neighbors, Point{i + 1, j + 1})
				}
			}

			for _, point := range neighbors {
				symbol := schematic[point.i][point.j]
				if unicode.IsDigit(symbol) || symbol == '.' {
					continue
				}

				adjacent = true

				if symbol == '*' {
					gearPos = &Point{point.i, point.j}
				}
			}
		}
	}

	for _, parts := range gears {
		if len(parts) != 2 {
			continue
		}

		secondPartResult += parts[0] * parts[1]
	}

	log.Printf("Part 1 result: %d", firstPartResult)
	log.Printf("Part 2 result: %d", secondPartResult)
}
