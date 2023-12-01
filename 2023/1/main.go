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

	result := 0

	// NB (alkurbatov): Becase the input contains strings like oneight, eighthree or sevenine
	// and Golang does not provide negative lookahead, we have to try reversed approach.
	firstDigit := regexp.MustCompile(`\d|one|two|three|four|five|six|seven|eight|nine`)
	lastDigit := regexp.MustCompile(`\d|eno|owt|eerht|ruof|evif|xis|neves|thgie|enin`)

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		src := scanner.Text()

		first := firstDigit.FindString(src)
		if first == "" {
			log.Printf("Failed to find first digit in input string: %s", src)
			return
		}

		last := lastDigit.FindString(reverse(src))
		if last == "" {
			log.Printf("Failed to find last digit in input string: %s", src)
			return
		}

		hiddenNumber := numsMap[first]*10 + numsMap[reverse(last)]
		result += hiddenNumber

		log.Printf("%s -> %d", src, hiddenNumber)
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Failed to scan input.txt: %v", err)
		return
	}

	log.Printf("Result: %d", result)
}
