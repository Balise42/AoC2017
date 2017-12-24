package main

import (
	"bufio"
	"day20"
	"fmt"
	"log"
	"os"
	"strings"
	"day24"
)

func main() {
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
	fmt.Println(day24.GetLongest(graph, 0))
}
