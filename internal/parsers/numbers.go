package parsers

import (
	"strconv"
	"strings"
)

// Read sequence of numbers from string.
func ReadNumbers(src string) ([]int, error) {
	result := make([]int, 0)

	for _, field := range strings.Fields(src) {
		num, err := strconv.Atoi(field)
		if err != nil {
			return nil, err
		}

		result = append(result, num)
	}

	return result, nil
}
