package main

import (
	"os"
	"./day12"
	"fmt"
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