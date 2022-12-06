package two

import (
	_ "embed"
	"fmt"
	"strings"
)

type Attack string

const (
	ROCK     = "X"
	PAPER    = "Y"
	SCISSORS = "Z"
)

type Winner int

const (
	OP   = 0
	ME   = 6
	DRAW = 3
)

//go:embed input.txt
var input string

func Score(cmd Attack) int {
	switch cmd {
	//ROCK
	case ROCK:
		return 1
	//Paper
	case PAPER:
		return 2
	//Scissors
	case SCISSORS:
		return 3
	}

	panic("invalid input")
}
func MapOpAttack(op string) Attack {
	switch op {
	case "A":
		return ROCK
	case "B":
		return PAPER
	case "C":
		return SCISSORS
	}
	panic("invalid input")
}
func Wins(op, me Attack) Winner {
	if op == me {
		return DRAW
	}

	switch op {
	case ROCK:
		if me == PAPER {
			return ME
		}
		return OP
	case PAPER:
		if me == SCISSORS {
			return ME
		}
		return OP
	case SCISSORS:
		if me == ROCK {
			return ME
		}
		return OP
	default:
		panic("invalid input")
	}
}

func Part1() []int {
	scores := []int{}
	for _, line := range strings.Split(input, "\n") {
		op, me := MapOpAttack(string(line[0])), Attack(line[2])
		winner := Wins(op, me)
		s := Score(me)
		scores = append(scores, s+int(winner))
	}

	return scores
}
func Play(op, me Attack) Attack {
	switch me {
	case ROCK:
		// loose
		switch op {
		case SCISSORS:
			return PAPER
		case ROCK:
			return SCISSORS
		}
		return me
	case PAPER:
		//draw
		return op
	case SCISSORS:
		switch op {
		case SCISSORS:
			return ROCK
		case ROCK:
			return PAPER
		}
		return me
	default:
		panic("invalid input")
	}
}
func Part2() int {
	lines := []string{}
	for _, line := range strings.Split(input, "\n") {
		op, me := MapOpAttack(string(line[0])), Attack(line[2])
		winner := Play(op, me)
		me = winner
		lines = append(lines, string(line[0])+" "+string(winner))
	}
	input = strings.Join(lines, "\n")
	p1 := Part1()
	sum := 0
	for _, score := range p1 {
		sum += score
	}
	return sum
}

func init() {
	p1 := Part1()
	sum := 0
	for _, score := range p1 {
		sum += score
	}

	p2 := Part2()
	fmt.Printf("Day 2: %d, %d", sum, p2)
}
