package main

import (
	"main/day8"
	"os"
	"fmt"
)

func main() {
	instructions := day8.ParseFileToInstructions(os.Args[1])
	mem, _ := day8.ProcessInstructions(instructions)
	fmt.Println(day8.MaxValue(mem))
}