package main

import (
	"os"
	"log"
	"bufio"
	"day20"
	"strings"
	"fmt"
	"day24"
)

func main () {
	filename := os.Args[1]
	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		log.Fatal("can't open file", err)
	}

	graph := make(map[day20.Pair]bool)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		day24.AddToGraph(graph, strings.TrimSpace(scanner.Text()))
	}
	fmt.Println(graph)
	fmt.Println(day24.GetLargestWeight(graph, 0))
}
