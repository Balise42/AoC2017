package main

import (
	"day7"
	"fmt"
	"os"
)

func main() {
	tree := day7.CreateTree(os.Args[1])
	fmt.Println(day7.GetHead(tree))
}
