// --- Day 4: Scratchcards ---
// https://adventofcode.com/2023/day/4
package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
)

type scratchcard map[string]struct{}

type cardResult struct {
	Score          int
	CountOfMatches int
	CountOfCopies  int
}

func parseInput(src string) scratchcard {
	result := make(scratchcard)

	for _, num := range strings.Fields(src) {
		result[num] = struct{}{}
	}

	return result
}

func calcScore(winningScratchcard, elfScratchcard scratchcard) (score, countOfMatches int) {
	for num := range elfScratchcard {
		if _, ok := winningScratchcard[num]; ok {
			countOfMatches++

			if score == 0 {
				score = 1
				continue
			}

			score *= 2
		}
	}

	return score, countOfMatches
}

func main() {
	input, err := os.Open("./2023/4/input.txt")
	if err != nil {
		log.Printf("Failed to open input.txt: %v", err)
		return
	}
	defer input.Close()

	var firstPartResult, secondPartResult int

	scanner := bufio.NewScanner(input)
	re := regexp.MustCompile(`^Card [ \d]+: ([ \d]+) ` + regexp.QuoteMeta(`|`) + ` ([ \d]+)`)
	results := make([]cardResult, 0)

	for scanner.Scan() {
		src := scanner.Text()
		groups := re.FindAllStringSubmatch(src, -1)

		if groups == nil || len(groups) != 1 || len(groups[0]) < 3 {
			log.Printf("Failed to extract cards from %s", src)
			return
		}

		winningScratchcard := parseInput(groups[0][1])
		elfScratchcard := parseInput(groups[0][2])

		score, countOfMatches := calcScore(winningScratchcard, elfScratchcard)

		firstPartResult += score
		results = append(results, cardResult{score, countOfMatches, 1})
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Failed to scan input.txt: %v", err)
		return
	}

	for cardNum, result := range results {
		for i := 1; i <= result.CountOfMatches; i++ {
			results[cardNum+i].CountOfCopies += results[cardNum].CountOfCopies
		}

		secondPartResult += result.CountOfCopies
	}

	log.Printf("Part 1 result: %d", firstPartResult)
	log.Printf("Part 2 result: %d", secondPartResult)
}
