package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

const (
	maxRed   = 12
	maxGreen = 13
	maxBlue  = 14
)

var (
	blues  = regexp.MustCompile(` (\d+) blue`)
	reds   = regexp.MustCompile(` (\d+) red`)
	greens = regexp.MustCompile(` (\d+) green`)
)

func maxShown(src string, re *regexp.Regexp) int {
	counts := re.FindAllStringSubmatch(src, -1)
	if counts == nil {
		return 0
	}

	var localMax int

	for _, count := range counts {
		for i := 1; i < len(count); i++ {
			num, err := strconv.Atoi(count[i])
			if err != nil {
				log.Printf("Failed to extract balls count: %v", err)
				return -1
			}

			localMax = max(localMax, num)
		}
	}

	return localMax
}

func main() {
	input, err := os.Open("./2023/2/input.txt")
	if err != nil {
		log.Printf("Failed to open input.txt: %v", err)
		return
	}
	defer input.Close()

	var id, firstPartResult, secondPartResult int

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		src := scanner.Text()
		id++

		maxRedShown := maxShown(src, reds)
		maxBlueShown := maxShown(src, blues)
		maxGreenShown := maxShown(src, greens)

		secondPartResult += maxRedShown * maxBlueShown * maxGreenShown

		if maxRedShown > maxRed || maxBlueShown > maxBlue || maxGreenShown > maxGreen {
			continue
		}

		firstPartResult += id
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Failed to scan input.txt: %v", err)
		return
	}

	log.Printf("Part 1 result: %d", firstPartResult)
	log.Printf("Part 2 result: %d", secondPartResult)
}
