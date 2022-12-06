package three

import (
	_ "embed"
	"fmt"
	"sort"
	"strings"
)

//go:embed input.txt
var input string

var lookup map[string]int

func Part1() {
	priorities := []string{}

	for _, l := range strings.Split(input, "\n") {
		first, sec := l[:len(l)/2], l[len(l)/2:]
		// first, sec = sortString(first), sortString(sec)
		mem := map[string]bool{}
		for i := 0; i < len(first); i++ {
			for j := 0; j < len(sec); j++ {
				if first[i] == sec[j] {
					mem[string(first[i])] = true
				}
			}
		}
		for k := range mem {
			priorities = append(priorities, k)
		}
	}

	fmt.Printf("Priorities: %v", priorities)
	prioSum := 0
	for i := 0; i < len(priorities); i++ {
		p := priorities[i]
		prioValue := lookup[p]
		prioSum += prioValue
	}

	fmt.Printf("Part 1: %v", prioSum)
}

func Part2() {
	priorities := []string{}

	splitstr := strings.Split(input, "\n")
	for i := 2; i < len(splitstr); i += 3 {

		first, sec, thrid := string(splitstr[i-2]), string(splitstr[i-1]), string(splitstr[i])
		// first, sec = sortString(first), sortString(sec)
		mem := map[string]bool{}
		for i := 0; i < len(first); i++ {
			for j := 0; j < len(sec); j++ {
				if first[i] == sec[j] {
					for k := 0; k < len(thrid); k++ {
						if first[i] == thrid[k] {
							mem[string(first[i])] = true
						}
					}
				}
			}
		}
		for k := range mem {
			priorities = append(priorities, k)
		}
	}

	fmt.Printf("Priorities: %v", priorities)
	prioSum := 0
	for i := 0; i < len(priorities); i++ {
		p := priorities[i]
		prioValue := lookup[p]
		prioSum += prioValue
	}

	fmt.Printf("Part 2: %v", prioSum)
}

func sortString(s string) string {
	split := strings.Split(s, "")
	sort.Strings(split)
	return strings.Join(split, "")
}

func init() {
	lookup = map[string]int{}
	for i, j := 'a', 1; i <= 'z'; i, j = i+1, j+1 {
		lookup[string(i)] = j
	}
	for i, j := 'A', 27; i <= 'Z'; i, j = i+1, j+1 {
		lookup[string(i)] = j
	}
	Part1()
	Part2()
}
