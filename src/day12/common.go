package day12

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func CreateGraphFromFile(filename string) map[string][]string {
	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		log.Fatal("Can't open file", err)
	}

	graph := make(map[string][]string)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		elems := strings.Split(line, " <-> ")
		follow := strings.Split(elems[1], ", ")

		_, ok := graph[elems[0]]
		if !ok {
			graph[elems[0]] = make([]string, 0)
		}
		graph[elems[0]] = append(graph[elems[0]], follow...)
	}

	return graph
}

func GetConnectedComponent(graph map[string][]string, vertex string, label int, visited map[string]int) {
	queue := make([]string, 0)
	queue = append(queue, vertex)
	for len(queue) > 0 {
		elem := queue[0]
		if v, ok := visited[elem]; !ok || v == 0 {
			visited[elem] = label
			queue = append(queue, graph[elem]...)
		}
		queue = queue[1:]
	}
}

func GetConnectedComponents(graph map[string][]string) map[string]int {
	visited := make(map[string]int)
	for k, _ := range graph {
		visited[k] = 0
	}

	added := true

	for i := 1; added; i++ {
		added = false
		for k, v := range visited {
			if v == 0 {
				GetConnectedComponent(graph, k, i, visited)
				added = true
				break
			}
		}
	}

	return visited
}
