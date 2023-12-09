// --- Day 8: Haunted Wasteland ---
// https://adventofcode.com/2023/day/8
package main

import (
	"bufio"
	"errors"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/alkurbatov/adventofcode/internal/calculus"
)

var errNoNodes = errors.New("no nodes found")

var nodeRe = regexp.MustCompile(`[\dA-Z]{3}`)

type Node struct {
	id    string
	left  string
	right string
}

func parseRoute(scanner *bufio.Scanner) ([]rune, error) {
	scanner.Scan()

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return []rune(scanner.Text()), nil
}

func parseNodes(scanner *bufio.Scanner) (map[string]Node, error) {
	nodes := make(map[string]Node, 0)

	for scanner.Scan() {
		src := scanner.Text()

		if src == "" {
			continue
		}

		fields := nodeRe.FindAllString(src, -1)
		if fields == nil || len(fields) < 3 {
			return nil, errNoNodes
		}

		nodes[fields[0]] = Node{fields[0], fields[1], fields[2]}
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Failed to scan seeds from input.txt: %v", err)
		return nil, err
	}

	if len(nodes) == 0 {
		return nil, errNoNodes
	}

	return nodes, nil
}

func calcSteps(route []rune, from, to string, nodes map[string]Node) int {
	var steps, pos int

	for {
		if strings.HasSuffix(nodes[from].id, to) {
			return steps
		}

		direction := route[pos]
		if direction == 'L' {
			from = nodes[from].left
		} else {
			from = nodes[from].right
		}

		pos++
		if pos == len(route) {
			pos = 0
		}

		steps++
	}
}

func calcGhostSteps(route []rune, nodes map[string]Node) int {
	steps := make([]int, 0)

	for from := range nodes {
		if !strings.HasSuffix(from, "A") {
			continue
		}

		steps = append(steps, calcSteps(route, from, "Z", nodes))
	}

	return calculus.FindLCM(steps...)
}

func main() {
	input, err := os.Open("./2023/8/input.txt")
	if err != nil {
		log.Printf("Failed to open input.txt: %v", err)
		return
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	route, err := parseRoute(scanner)
	if err != nil {
		log.Printf("Failed to parse routes from input: %v", err)
	}

	nodes, err := parseNodes(scanner)
	if err != nil {
		log.Printf("Failed to parse nodes from input: %v", err)
	}

	firstPartResult := calcSteps(route, "AAA", "ZZZ", nodes)
	log.Printf("Part 1 result: %d", firstPartResult)

	secondPartResult := calcGhostSteps(route, nodes)
	log.Printf("Part 2 result: %d", secondPartResult)
}
