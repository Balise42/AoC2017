package day7

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Name     string
	Weight   int
	Children []string
	Parent   []string
}

func CreateTree(filename string) map[string]*Node {
	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		log.Fatal("can't open file", err)
	}

	scanner := bufio.NewScanner(file)

	nodelist := make(map[string]*Node)

	for scanner.Scan() {
		line := scanner.Text()
		var left string
		comps := strings.Split(line, " -> ")
		left = comps[0]

		parts := strings.Split(left, " ")
		nodename := parts[0]
		nodeweight, err := strconv.Atoi(parts[1][1 : len(parts[1])-1])
		if err != nil {
			log.Fatal("can't convert to int", parts[1], err)
		}
		addNode(nodelist, nodename, nodeweight)

		if len(comps) > 1 {
			addNodeDependency(nodename, comps[1], nodelist)
		}
	}

	return nodelist
}

func GetHead(tree map[string]*Node) string {
	for _, v := range tree {
		if len(v.Parent) == 0 {
			return v.Name
		}
	}
	return ""
}

func addNodeDependency(node string, deps string, nodelist map[string]*Node) {
	deplist := strings.Split(deps, ", ")
	for _, dep := range deplist {
		(nodelist[node]).Children = append((nodelist[node]).Children, dep)
		_, ok := nodelist[dep]
		if ok {
			(nodelist[dep]).Parent = append((nodelist[dep]).Parent, node)
		} else {
			nodelist[dep] = &Node{dep, -1, make([]string, 0, 0), []string{node}}
		}
	}
}

func addNode(nodelist map[string]*Node, name string, weight int) {
	_, ok := nodelist[name]
	if !ok {
		nodelist[name] = &Node{name, weight, make([]string, 0, 0), make([]string, 0, 0)}
	} else {
		(nodelist[name]).Weight = weight
	}
}
