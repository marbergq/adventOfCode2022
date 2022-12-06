package four

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func rangeNumber(r string) []int {
	split := strings.Split(r, "-")
	min, _ := strconv.Atoi(split[0])
	max, _ := strconv.Atoi(split[1])
	parts := []int{}
	for i := min; i <= max; i++ {
		parts = append(parts, i)
	}

	return parts
}

func Overlap(arr1, arr2 []int) []int {
	overlap := []int{}
	for _, v := range arr1 {
		for _, v2 := range arr2 {
			if v == v2 {
				overlap = append(overlap, v)
			}
		}
	}

	return overlap
}

func Part1() {
	overlaps := 0
	for _, l := range strings.Split(input, "\n") {
		split := strings.Split(l, ",")
		p1 := split[0]
		p2 := split[1]
		p1Parts := rangeNumber(p1)
		p2Parts := rangeNumber(p2)
		overlap := Overlap(p1Parts, p2Parts)
		if len(p1Parts) == len(overlap) || len(p2Parts) == len(overlap) {
			overlaps++
		}
	}
	fmt.Printf("Part 1: %v", overlaps)
}

func Part2() {
	overlaps := 0
	for _, l := range strings.Split(input, "\n") {
		split := strings.Split(l, ",")
		p1 := split[0]
		p2 := split[1]
		p1Parts := rangeNumber(p1)
		p2Parts := rangeNumber(p2)
		overlap := Overlap(p1Parts, p2Parts)
		if len(overlap) > 0 {
			overlaps++
		}
	}
	fmt.Printf("Part 2: %v", overlaps)
}
func init() {
	Part1()
	Part2()
}
