package main

import (
	"day12"
	"fmt"
	"os"
)

func main() {
	filename := os.Args[1]
	graph := day12.CreateGraphFromFile(filename)
	component := day12.GetConnectedComponents(graph)
	max := 0
	fmt.Println(component)
	for _, v := range component {
		if v > max {
			max = v
		}
	}
	fmt.Println(max)
}
