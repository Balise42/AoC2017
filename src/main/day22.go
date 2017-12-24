package main

import (
	"os"
	"log"
	"bufio"
	"strings"
	"day20"
	"fmt"
	"day22"
)

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		log.Fatal("can't open file", err)
	}

	scanner := bufio.NewScanner(file)

	grid := make(map[day20.Pair]rune)

	i := 0
	size := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		size = len(line)
		for j, c := range line {
			grid[day20.Pair{A: j, B: i}] = c
		}
		i++
	}

	pos := day20.Pair{A:i/2, B:size/2}
	dir := "up"
	a := 0
	sum := 0

	for i := 0; i<10000; i++ {
		a, pos, dir = day22.Evolve(grid, pos, dir)
		sum = sum + a
	}

	fmt.Println(sum)

}