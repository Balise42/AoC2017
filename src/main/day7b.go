package main

import (
	"day7"
	"fmt"
	"log"
	"os"
)

func main() {
	filename := os.Args[1]
	tree := day7.CreateTree(filename)
	head := day7.GetHead(tree)

	candidate := head
	for nextNode := head; nextNode != ""; nextNode = getOddChild(tree, nextNode) {
		candidate = nextNode
	}

	fmt.Println(candidate)

	fmt.Println(getNormalWeight(tree, candidate))
}

func getWeight(node string, tree map[string]*day7.Node) int {
	weight := tree[node].Weight
	for _, child := range tree[node].Children {
		weight = weight + getWeight(child, tree)
	}
	return weight
}

func getNormalWeight(tree map[string]*day7.Node, nodename string) int {
	node := tree[nodename]
	parent := node.Parent[0]
	var otherchild string
	if tree[parent].Children[0] != nodename {
		otherchild = tree[parent].Children[0]
	} else {
		otherchild = tree[parent].Children[1]
	}
	return getWeight(otherchild, tree) - childrenWeight(nodename, tree)
}

func childrenWeight(node string, tree map[string]*day7.Node) int {
	weight := 0
	for _, child := range tree[node].Children {
		weight = weight + getWeight(child, tree)
	}
	return weight
}

func getOddChild(tree map[string]*day7.Node, node string) string {
	if len(tree[node].Children) <= 2 {
		log.Fatal("nope.")
	}

	r1 := getWeight(tree[node].Children[0], tree)
	r2 := getWeight(tree[node].Children[1], tree)
	if r1 == r2 {
		for _, child := range tree[node].Children {
			if getWeight(child, tree) != r1 {
				return child
			}
		}
	} else {
		if r1 == getWeight(tree[node].Children[2], tree) {
			return tree[node].Children[1]
		} else {
			return tree[node].Children[0]
		}
	}
	return ""
}
