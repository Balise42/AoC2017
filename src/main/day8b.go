package main

import (
	"day8"
	"fmt"
	"os"
)

func main() {
	instructions := day8.ParseFileToInstructions(os.Args[1])
	_, max := day8.ProcessInstructions(instructions)
	fmt.Println(max)
}
