package day21

import (
	"strings"
)

func AddRule(rules map[string]string, line string) {
	toks := strings.Split(line, " => ")
	rules[toks[0]] = toks[1]
}

func Enhance(grid []string, rules map[string]string) []string {
	splitGrid := Split(grid)
	enhanced := enhanceSplit(splitGrid, rules)
	return Rebuild(enhanced)
}

func Rebuild(split [][]string) []string {
	res := make([]string, 0, 0)
	size := strings.Count(split[0][0], "/")
	for i := 0; i< len(split); i++ {
		for j := 0; j <= size; j++ {
			str := ""
			for k := 0; k < len(split); k++ {
				str = str + strings.Split(split[i][k], "/")[j]
			}
			res = append(res, str)
		}
	}
	return res
}

func Split(grid[]string) [][][]string {
	var unitSize int
	if len(grid) % 2 == 0 {
		unitSize = 2
	} else {
		unitSize = 3
	}

	splitGrid := make([][][]string, len(grid)/unitSize, len(grid)/unitSize)
	for i := range splitGrid {
		splitGrid[i] = make([][]string, len(grid)/unitSize, len(grid)/unitSize)
		for j := range splitGrid[i] {
			splitGrid[i][j] = make([]string, unitSize, unitSize)
		}
	}

	for i := 0; i<len(grid) / unitSize; i++ {
		for j := 0; j<len(grid) / unitSize; j++ {
			for k := 0; k<unitSize; k++ {
				splitGrid[i][j][k] = grid[i*unitSize + k][unitSize*j:unitSize*(j+1)]
			}
		}
	}

	return splitGrid
}

func enhanceSplit(grid [][][]string, rules map[string]string) [][]string {
	enhanced := make([][]string, len(grid), len(grid))

	for i, col := range grid {
		enhanced[i] = make([]string, len(grid))
		for j, cell := range col {
			enhanced[i][j] = EnhanceCell(cell, rules)
		}
	}

	return enhanced
}

func EnhanceCell(cell []string, rules map[string]string) string {

	for rule, result := range rules {
		if Matches(cell, rule) {
			return result
		}
	}
	return "ERROR"
}

func Matches(cell []string, rule string) bool{
	flatCell := strings.Join(cell, "/")
	if len(flatCell) != len(rule) {
		return false
	}
	if strings.Count(flatCell, "#") != strings.Count(rule, "#") {
		return false
	}
	for i := 0; i<4; i++ {
		if MatchesExactly(cell, rule) {
			return true
		}
		if MatchesExactly(FlipH(cell), rule) {
			return true
		}
		if MatchesExactly(FlipV(cell), rule) {
			return true
		}
		cell = Rotate(cell)
	}
	return false
}

func Rotate(cell []string) []string {
	if len(cell) == 2 {
		return []string{
			string(cell[0][1]) + string(cell[1][1]),
			string(cell[0][0]) + string(cell[1][0]),
		}
	} else {
		return []string{
			string(cell[0][2]) + string(cell[1][2]) + string(cell[2][2]),
			string(cell[0][1]) + string(cell[1][1]) + string(cell[2][1]),
			string(cell[0][0]) + string(cell[1][0]) + string(cell[2][0]),
		}
	}
}

func MatchesExactly(cell []string, rule string) bool {
	return strings.Join(cell, "/") == rule
}

func FlipH(cell []string) []string {
	newCell := make([]string, len(cell), len(cell))
	for i, row := range cell {
		val := reverse(row)
		newCell[i] = val
	}
	return newCell
}

func reverse(str string) string {
	if len(str) == 2 {
		return string(str[1]) + string(str[0])
	} else {
		return string(str[2]) + string(str[1]) + string(str[0])
	}
}

func FlipV(cell []string) []string{
	newCell := make([]string, len(cell), len(cell))
	if len(cell) == 2 {
		newCell[0] = cell[1]
		newCell[1] = cell[0]
	} else {
		newCell[0] = cell[2]
		newCell[1] = cell[1]
		newCell[2] = cell[0]
	}
	return newCell
}


