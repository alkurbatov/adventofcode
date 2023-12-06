// --- Day 6: Wait For It ---
// https://adventofcode.com/2023/day/6
package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput(src string) (races []int, total int, err error) {
	values := strings.Fields(src)
	races = make([]int, 0, len(values)-1)

	var sb strings.Builder
	var num int

	for i := 1; i < len(values); i++ {
		num, err = strconv.Atoi(values[i])
		if err != nil {
			return nil, 0, err
		}

		races = append(races, num)

		sb.WriteString(values[i])
	}

	total, err = strconv.Atoi(sb.String())
	if err != nil {
		return nil, 0, err
	}

	return races, total, nil
}

func findVictoryBorder(time, distance int, cmp func(int, int) bool) int {
	l := 0
	r := time

	for l <= r {
		pivot := (r-l)/2 + l
		if cmp((time-pivot)*pivot, distance) {
			r = pivot - 1
			continue
		}

		l = pivot + 1
	}

	return l
}

func findVictoriesCount(time, distance int) int {
	left := findVictoryBorder(time, distance, func(l, r int) bool {
		return l > r
	})
	right := findVictoryBorder(time, distance, func(l, r int) bool {
		return l <= r
	})

	return right - left
}

func main() {
	input, err := os.Open("./2023/6/input.txt")
	if err != nil {
		log.Printf("Failed to open input.txt: %v", err)
		return
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	i := 0
	races := make([][]int, 2)
	totals := make([]int, 2)

	for scanner.Scan() {
		src := scanner.Text()

		values, total, err := readInput(src)
		if err != nil {
			log.Printf("Failed to reed values from %s: %v", src, err)
			return
		}

		races[i] = values
		totals[i] = total
		i++
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Failed to scan seeds from input.txt: %v", err)
		return
	}

	firstPartResult := 1
	for i := 0; i < len(races[0]); i++ {
		firstPartResult *= findVictoriesCount(races[0][i], races[1][i])
	}
	log.Printf("Part 1 result: %d", firstPartResult)

	secondPartResult := findVictoriesCount(totals[0], totals[1])
	log.Printf("Part 2 result: %d", secondPartResult)
}
