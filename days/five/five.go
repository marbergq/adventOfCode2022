package five

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Part func(columnAndRows [][]string, fromAmount, fromColumn, toColumn int) [][]string

func Run(part Part) string {
	arrangement := []string{}
	for _, l := range strings.Split(input, "\n") {
		if l == "" {
			break
		}

		arrangement = append(arrangement, l)
	}

	instructions := []string{}
	instructions = append(instructions, strings.Split(input, "\n")[len(arrangement)+1:]...)

	columnRow := [][]string{}
	for column := 1; column < len(arrangement[0])-1; column += 4 {
		columns := []string{}
		for row := len(arrangement) - 2; row >= 0; row-- {
			col := string(arrangement[row][column])
			if col != " " {
				columns = append(columns, col)
			}
		}
		columnRow = append(columnRow, columns)
	}

	for _, instruction := range instructions {
		instruction = strings.ReplaceAll(instruction, "move ", "")
		instruction = strings.ReplaceAll(instruction, "from ", "")
		instruction = strings.ReplaceAll(instruction, "to ", "")
		split := strings.Split(instruction, " ")
		fromAmount, fromColumn, toColumn := split[0], split[1], split[2]
		fromAmountInt, _ := strconv.Atoi(string(fromAmount))
		fromColumnInt, _ := strconv.Atoi(string(fromColumn))
		toColumnInt, _ := strconv.Atoi(string(toColumn))
		columnRow = part(columnRow, fromAmountInt, fromColumnInt, toColumnInt)

	}

	result := ""
	for _, column := range columnRow {
		result += column[len(column)-1]
	}

	return result
}

func init() {
	p1 := Run(func(columnRow [][]string, fromAmountInt, fromColumnInt, toColumnInt int) [][]string {
		for row := 0; row < fromAmountInt; row++ {
			columnRow[toColumnInt-1] = append(columnRow[toColumnInt-1], columnRow[fromColumnInt-1][len(columnRow[fromColumnInt-1])-1])
			columnRow[fromColumnInt-1] = columnRow[fromColumnInt-1][:len(columnRow[fromColumnInt-1])-1]
		}
		return columnRow
	})
	p2 := Run(func(columnRow [][]string, fromAmountInt, fromColumnInt, toColumnInt int) [][]string {
		columnRow[toColumnInt-1] = append(columnRow[toColumnInt-1], columnRow[fromColumnInt-1][len(columnRow[fromColumnInt-1])-fromAmountInt:len(columnRow[fromColumnInt-1])]...)
		columnRow[fromColumnInt-1] = columnRow[fromColumnInt-1][:len(columnRow[fromColumnInt-1])-fromAmountInt]
		return columnRow
	})

	println("Part 1:", p1)
	println("Part 2:", p2)

}
