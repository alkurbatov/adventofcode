// --- Day 5: If You Give A Seed A Fertilizer ---
// https://adventofcode.com/2023/day/5
package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const maxInt = 9223372036854775807

type Range struct {
	Delta int
	Start int
	End   int
}

func newRange(description string) (Range, error) {
	fields := strings.Fields(description)

	dst, err := strconv.Atoi(fields[0])
	if err != nil {
		return Range{}, err
	}

	src, err := strconv.Atoi(fields[1])
	if err != nil {
		return Range{}, err
	}

	length, err := strconv.Atoi(fields[2])
	if err != nil {
		return Range{}, err
	}

	return Range{dst - src, src, src + length}, nil
}

type conversionMap []Range

func newConversionMap(ranges []Range) conversionMap {
	sort.SliceStable(ranges, func(i, j int) bool {
		return ranges[i].Start < ranges[j].Start
	})

	return conversionMap(ranges)
}

func (c conversionMap) find(value int) int {
	for _, rng := range c {
		if value < rng.Start {
			return value
		}

		if value >= rng.End {
			continue
		}

		return rng.Delta + value
	}

	return value
}

type Pipeline []conversionMap

func (p Pipeline) process(location int) int {
	for _, cMap := range p {
		location = cMap.find(location)
	}

	return location
}

func readSeeds(src string) ([]int, error) {
	seeds := make([]int, 0)

	for _, rawNum := range strings.Fields(src[len("seeds:"):]) {
		num, err := strconv.Atoi(rawNum)
		if err != nil {
			return nil, err
		}

		seeds = append(seeds, num)
	}

	return seeds, nil
}

func readPipeline(scanner *bufio.Scanner) (Pipeline, error) {
	pipeline := make(Pipeline, 0)
	ranges := make([]Range, 0)

	for scanner.Scan() {
		src := scanner.Text()

		if strings.HasSuffix(src, "map:") {
			continue
		}

		if src != "" {
			rng, err := newRange(src)
			if err != nil {
				return nil, err
			}

			ranges = append(ranges, rng)

			continue
		}

		if len(ranges) != 0 {
			pipeline = append(pipeline, newConversionMap(ranges))
			ranges = make([]Range, 0)
		}
	}

	if len(ranges) != 0 {
		pipeline = append(pipeline, newConversionMap(ranges))
	}

	return pipeline, nil
}

func solveFirstPart(seeds []int, pipeline Pipeline) int {
	result := maxInt
	for _, seed := range seeds {
		result = min(result, pipeline.process(seed))
	}

	return result
}

func solveSecondPart(seeds []int, pipeline Pipeline) int {
	result := maxInt

	for i := 1; i < len(seeds); i += 2 {
		for seed := seeds[i-1]; seed < seeds[i-1]+seeds[i]; seed++ {
			result = min(result, pipeline.process(seed))
		}
	}

	return result
}

func main() {
	input, err := os.Open("./2023/5/input.txt")
	if err != nil {
		log.Printf("Failed to open input.txt: %v", err)
		return
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	if !scanner.Scan() {
		log.Printf("Failed to scan seeds from input.txt: %v", scanner.Err())
		return
	}

	seeds, err := readSeeds(scanner.Text())
	if err != nil {
		log.Printf("Failed to convert seed number to int: %v", err)
		return
	}

	pipeline, err := readPipeline(scanner)
	if err != nil {
		log.Printf("Failed to read pipeline: %v", err)
		return
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Failed to scan input.txt: %v", err)
		return
	}

	firstPartResult := solveFirstPart(seeds, pipeline)
	log.Printf("Part 1 result: %d\n", firstPartResult)

	secondPartResult := solveSecondPart(seeds, pipeline)
	log.Printf("Part 2 result: %d", secondPartResult)
}
