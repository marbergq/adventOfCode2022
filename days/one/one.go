package one

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func Part1() []int {
	results := []int{}
	current := 0
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			results = append(results, current)
			current = 0
		} else {
			v, _ := strconv.Atoi(line)
			current += v
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(results)))
	return results
}

func Part2() int {
	sum := 0
	for _, elf := range Part1()[:3] {
		sum += elf
	}

	return sum
}

func init() {
	fmt.Printf("Day 1: %d, %d", Part1()[0], Part2())
}
