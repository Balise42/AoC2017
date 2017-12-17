package main

import (
	"./day3"
	"fmt"
	"log"
	"os"
	"strconv"
)

type grid struct {
	array [][]int
}

func newGrid() grid {
	g := grid{make([][]int, 100)}
	for i := range g.array {
		g.array[i] = make([]int, 100)
	}
	g.set(0, 0, 1)
	return g
}

func (g grid) set(x int, y int, value int) {
	g.array[x+50][y+50] = value
}

func (g grid) get(x int, y int) int {
	return g.array[x+50][y+50]
}

func main() {
	input, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal("Input must be an integer.")
	}

	g := newGrid()

	for i := 2; ; i++ {
		xf, yf := day3.GetCoordinates(i)
		x := int(xf)
		y := int(yf)
		res := g.get(x-1, y-1) + g.get(x-1, y) + g.get(x-1, y+1) + g.get(x, y-1) + g.get(x, y+1) + g.get(x+1, y-1) + g.get(x+1, y) + g.get(x+1, y+1)
		if res > input {
			fmt.Println(res)
			break
		}
		g.set(x, y, res)
	}
}
