package main

import (
	"os"
	"./day12"
	"fmt"
)

func main() {
	filename := os.Args[1]
	graph := day12.CreateGraphFromFile(filename)
	visited := make(map[string]int)
	day12.GetConnectedComponent(graph, "0", 1, visited)
	sum := 0
	for _, v := range visited {
		sum = sum + v
	}
	fmt.Println(sum)
}