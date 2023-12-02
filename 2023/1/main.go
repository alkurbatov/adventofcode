package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
)

var numsMap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
}

func reverse(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}

	return string(rns)
}

func main() {
	input, err := os.Open("./2023/1/input.txt")
	if err != nil {
		log.Printf("Failed to open input.txt: %v", err)
		return
	}
	defer input.Close()

	var firstPartResult, secondPartResult int

	// NB (alkurbatov): Becase the input contains strings like oneight, eighthree or sevenine
	// and Golang does not provide negative lookahead, we have to try 'reversed' approach.
	digits := regexp.MustCompile(`\d`)
	firstNumber := regexp.MustCompile(`\d|one|two|three|four|five|six|seven|eight|nine`)
	lastNumber := regexp.MustCompile(`\d|eno|owt|eerht|ruof|evif|xis|neves|thgie|enin`)

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		src := scanner.Text()

		allDigits := digits.FindAllString(src, -1)
		if len(allDigits) == 0 {
			log.Printf("Failed to find all digits in input string: %s", src)
			return
		}

		hiddenNumber := numsMap[allDigits[0]]*10 + numsMap[allDigits[len(allDigits)-1]]
		firstPartResult += hiddenNumber
		log.Printf("Part 1: %s -> %d", src, hiddenNumber)

		firstNum := firstNumber.FindString(src)
		if firstNum == "" {
			log.Printf("Failed to find first number in input string: %s", src)
			return
		}

		lastNum := lastNumber.FindString(reverse(src))
		if lastNum == "" {
			log.Printf("Failed to find last number in input string: %s", src)
			return
		}

		hiddenNumber = numsMap[firstNum]*10 + numsMap[reverse(lastNum)]
		secondPartResult += hiddenNumber
		log.Printf("Part 2: %s -> %d", src, hiddenNumber)
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Failed to scan input.txt: %v", err)
		return
	}

	log.Printf("Part 1 result: %d", firstPartResult)
	log.Printf("Part 2 result: %d", secondPartResult)
}
