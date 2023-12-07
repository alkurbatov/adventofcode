// --- Day 7: Camel Cards ---
// https://adventofcode.com/2023/day/7
package main

import (
	"bufio"
	"errors"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var errBadInput = errors.New("not enough input fields")

var strength = map[string]int{
	"five of a kind":  100000000,
	"four of a kind":  10000000,
	"full house":      1000000,
	"three of a kind": 100000,
	"two pair":        10000,
	"one pair":        1000,
	"A":               14,
	"K":               13,
	"Q":               12,
	"J":               11,
	"T":               10,
	"9":               9,
	"8":               8,
	"7":               7,
	"6":               6,
	"5":               5,
	"4":               4,
	"3":               3,
	"2":               2,
}

type Input struct {
	hand string
	bid  int
}

type State struct {
	Input
	handStrength  int
	cardsStrength []int
}

func calcCardsStrength(hand string, withJokers bool) []int {
	cardsStrength := make([]int, len(hand))

	for i, card := range hand {
		if card == 'J' && withJokers {
			cardsStrength[i] = 1
			continue
		}

		cardsStrength[i] = strength[string(card)]
	}

	return cardsStrength
}

func calcHandStrength(hand string, withJokers bool) int {
	jokers := 0
	cards := make(map[rune]int, 0)

	for _, card := range hand {
		if card == 'J' && withJokers {
			jokers++
		}

		if _, ok := cards[card]; ok {
			cards[card]++
			continue
		}

		cards[card] = 1
	}

	// 0 - nothing, 1 - single, 2 - pair, 3 - trinity, 4 - fours, 5 - fifth
	cmb := make([]int, 6)
	for _, count := range cards {
		cmb[count]++
	}

	switch jokers {
	case 4:
		return strength["five of a kind"]
	case 3:
		if cmb[2] == 1 {
			return strength["five of a kind"]
		}

		return strength["four of a kind"]
	case 2:
		if cmb[3] == 1 {
			return strength["five of a kind"]
		}

		if cmb[2] == 2 {
			return strength["four of a kind"]
		}

		return strength["three of a kind"]
	case 1:
		if cmb[4] == 1 {
			return strength["five of a kind"]
		}

		if cmb[3] == 1 {
			return strength["four of a kind"]
		}

		if cmb[2] == 2 {
			return strength["full house"]
		}

		if cmb[2] == 1 {
			return strength["three of a kind"]
		}

		return strength["one pair"]
	default:
		if cmb[5] == 1 {
			return strength["five of a kind"]
		}

		if cmb[4] == 1 {
			return strength["four of a kind"]
		}

		if cmb[3] == 1 && cmb[2] == 1 {
			return strength["full house"]
		}

		if cmb[3] == 1 && cmb[1] == 2 {
			return strength["three of a kind"]
		}

		if cmb[2] == 2 {
			return strength["two pair"]
		}

		if cmb[2] == 1 {
			return strength["one pair"]
		}

		return 0
	}
}

func newState(input Input, withJokers bool) State {
	handStrength := calcHandStrength(input.hand, withJokers)
	cardsStrength := calcCardsStrength(input.hand, withJokers)

	return State{input, handStrength, cardsStrength}
}

func process(inputs []Input, withJokers bool) []State {
	states := make([]State, 0, len(inputs))
	for _, input := range inputs {
		states = append(states, newState(input, withJokers))
	}

	sort.SliceStable(states, func(i, j int) bool {
		if states[i].handStrength != states[j].handStrength {
			return states[i].handStrength < states[j].handStrength
		}

		for k := 0; k < len(states[i].cardsStrength); k++ {
			if states[i].cardsStrength[k] == states[j].cardsStrength[k] {
				continue
			}

			return states[i].cardsStrength[k] < states[j].cardsStrength[k]
		}

		return states[i].bid < states[j].bid
	})

	return states
}

func readInput(src string) (Input, error) {
	fields := strings.Fields(src)

	if len(fields) < 2 {
		return Input{}, errBadInput
	}

	bid, err := strconv.Atoi(fields[1])
	if err != nil {
		return Input{}, err
	}

	return Input{fields[0], bid}, nil
}

func main() {
	input, err := os.Open("./2023/7/input.txt")
	if err != nil {
		log.Printf("Failed to open input.txt: %v", err)
		return
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	inputs := make([]Input, 0)

	for scanner.Scan() {
		src := scanner.Text()

		input, err := readInput(src)
		if err != nil {
			log.Printf("Faile to read input: %v", err)
			return
		}

		inputs = append(inputs, input)
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Failed to scan seeds from input.txt: %v", err)
		return
	}

	var firstPartResult, secondPartResult int

	states := process(inputs, false)
	for i, state := range states {
		firstPartResult += (i + 1) * state.bid
	}

	log.Printf("Part 1 result: %d", firstPartResult)

	states = process(inputs, true)
	for i, state := range states {
		secondPartResult += (i + 1) * state.bid
	}

	log.Printf("Part 2 result: %d", secondPartResult)
}
