package day22

import (
	"day20"
)

func Evolve(grid map[day20.Pair]rune, pos day20.Pair, dir string) (int, day20.Pair, string) {
	if _, ok := grid[pos]; !ok {
		grid[pos] = '.'
	}

	var inf int

	if grid[pos] == '.' {
		grid[pos] = '#'
		inf = 1
		if dir == "up" {
			dir = "left"
		} else if dir == "left" {
			dir = "down"
		} else if dir == "down" {
			dir = "right"
		} else {
			dir = "up"
		}
	} else {
		grid[pos] = '.'
		inf = 0
		if dir == "up" {
			dir = "right"
		} else if dir == "right" {
			dir = "down"
		} else if dir == "down" {
			dir = "left"
		} else {
			dir = "up"
		}
	}

	var newpos day20.Pair
	if dir == "up" {
		newpos = day20.Pair{pos.A, pos.B-1}
	} else if dir == "down" {
		newpos = day20.Pair{pos.A, pos.B+1}
	} else if dir == "left" {
		newpos = day20.Pair{pos.A - 1, pos.B}
	} else {
		newpos = day20.Pair{pos.A + 1, pos.B}
	}

	return inf, newpos, dir
}

func EvolveWeakened(grid map[day20.Pair]rune, pos day20.Pair, dir string) (int, day20.Pair, string) {
	if _, ok := grid[pos]; !ok {
		grid[pos] = '.'
	}

	inf := 0

	if grid[pos] == '.' {
		grid[pos] = 'W'
		if dir == "up" {
			dir = "left"
		} else if dir == "left" {
			dir = "down"
		} else if dir == "down" {
			dir = "right"
		} else {
			dir = "up"
		}
	} else if grid[pos] == '#' {
		grid[pos] = 'F'
		if dir == "up" {
			dir = "right"
		} else if dir == "right" {
			dir = "down"
		} else if dir == "down" {
			dir = "left"
		} else {
			dir = "up"
		}
	} else if grid[pos] == 'F' {
		grid[pos] = '.'
		if dir == "up" {
			dir = "down"
		} else if dir == "down" {
			dir = "up"
		} else if dir == "left" {
			dir = "right"
		} else {
			dir = "left"
		}
	} else {
		grid[pos] = '#'
		inf = 1
	}

	var newpos day20.Pair
	if dir == "up" {
		newpos = day20.Pair{pos.A, pos.B-1}
	} else if dir == "down" {
		newpos = day20.Pair{pos.A, pos.B+1}
	} else if dir == "left" {
		newpos = day20.Pair{pos.A - 1, pos.B}
	} else {
		newpos = day20.Pair{pos.A + 1, pos.B}
	}

	return inf, newpos, dir
}