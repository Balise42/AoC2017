package main

import (
	"main/day8"
	"os"
	"fmt"
)

func main() {
	instructions := day8.ParseFileToInstructions(os.Args[1])
	_, max := day8.ProcessInstructions(instructions)
	fmt.Println(max)
}