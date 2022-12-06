package six

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

func Part1() {
	fmt.Println("p1 ", findUnique(input, 4))
}

func Part2() {
	fmt.Println("p2 ", findUnique(input, 14))
}

func findUnique(s string, searchLength int) int {
	counter := map[string]bool{}
	for c := searchLength; c < len(input); c++ {
		for _, ch := range input[c-searchLength : c] {
			counter[string(ch)] = true
		}
		if len(counter) == searchLength {
			return c
		}
		counter = map[string]bool{}
	}
	return -1
}

func init() {
	Part1()
	Part2()
}
