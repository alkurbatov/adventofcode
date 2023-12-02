package main

import (
	"bufio"
	"errors"
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

var ErrNoID = errors.New("game ID not found")

var (
	gameID = regexp.MustCompile(`Game (\d+):`)
	blues  = regexp.MustCompile(` (\d+) blue`)
	reds   = regexp.MustCompile(` (\d+) red`)
	greens = regexp.MustCompile(` (\d+) green`)
)

func extractGameID(src string) (int, error) {
	raw := gameID.FindStringSubmatch(src)
	if len(raw) == 0 {
		return 0, ErrNoID
	}

	return strconv.Atoi(raw[1])
}

func fewestPossible(src string, re *regexp.Regexp) int {
	counts := re.FindAllStringSubmatch(src, -1)
	if counts == nil {
		return 0
	}

	localMax := -1

	for _, count := range counts {
		for i := 1; i < len(count); i++ {
			num, err := strconv.Atoi(count[i])
			if err != nil {
				log.Printf("Failed to extract balls count: %v", err)
				return -1
			}

			if num > localMax {
				localMax = num
			}
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

	var firstPartResult, secondPartResult int

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		src := scanner.Text()

		fewerRed := fewestPossible(src, reds)
		fewerBlue := fewestPossible(src, blues)
		fewerGreen := fewestPossible(src, greens)

		secondPartResult += fewerRed * fewerBlue * fewerGreen

		if fewerRed > maxRed || fewerBlue > maxBlue || fewerGreen > maxGreen {
			continue
		}

		id, err := extractGameID(src)
		if err != nil {
			log.Printf("Failed to extract game ID: %v", err)
			return
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
