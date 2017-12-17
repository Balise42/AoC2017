package main

import (
	"fmt"
	"main/day10"
	"main/day12"
	"strconv"
)

func main() {
	input := "vbqugkhl"

	grid := make([]string, 128, 128)

	for i := 0; i < 128; i++ {
		hash := day10.ComputeHash(input + "-" + strconv.FormatInt(int64(i), 10))
		grid[i] = day10.ToBinary(hash)
	}

	graph := createGraph(grid)
	components := day12.GetConnectedComponents(graph)
	max := 0
	for _, v := range components {
		if v > max {
			max = v
		}
	}
	fmt.Println(components)
	fmt.Println(max)
}

func createGraph(grid []string) map[string][]string {
	graph := make(map[string][]string)

	for i := 0; i < 128; i++ {
		for j := 0; j < 128; j++ {
			if grid[i][j] == '1' {
				if _, ok := graph[key(i, j)]; !ok {
					graph[key(i, j)] = make([]string, 0)
				}
				addNeighbors(grid, i, j, graph)
			}
		}
	}
	return graph
}

func addNeighbors(grid []string, i int, j int, graph map[string][]string) {
	if i > 0 && grid[i-1][j] == '1' {
		addEdge(i, j, i-1, j, graph)
	}
	if i < 127 && grid[i+1][j] == '1' {
		addEdge(i, j, i+1, j, graph)
	}
	if j > 0 && grid[i][j-1] == '1' {
		addEdge(i, j, i, j-1, graph)
	}
	if j < 127 && grid[i][j+1] == '1' {
		addEdge(i, j, i, j+1, graph)
	}
}

func addEdge(x1 int, y1 int, x2 int, y2 int, graph map[string][]string) {
	k1 := key(x1, y1)
	k2 := key(x2, y2)
	graph[k1] = append(graph[k1], k2)
}

func key(x int, y int) string {
	return strconv.FormatInt(int64(x), 10) + "--" + strconv.FormatInt(int64(y), 10)
}
