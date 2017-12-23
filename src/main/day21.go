package main

import (
	"os"
	"log"
	"bufio"
	"strings"
	"fmt"
	"day21"
)

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal("can't open file", err)
	}

	rules := make(map[string]string)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		day21.AddRule(rules, strings.TrimSpace(scanner.Text()))
	}

	grid := make([]string, 3, 3)
	grid[0] = ".#."
	grid[1] = "..#"
	grid[2] = "###"

	for i := 0; i<18; i++ {
		grid = day21.Enhance(grid, rules)
	}

	numOn := 0

	for _, row := range grid {
		numOn = numOn + strings.Count(row, "#")
	}

	fmt.Println(numOn)
}