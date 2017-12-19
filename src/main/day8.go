package main

import (
	"day8"
	"fmt"
	"os"
)

func main() {
	instructions := day8.ParseFileToInstructions(os.Args[1])
	mem, _ := day8.ProcessInstructions(instructions)
	fmt.Println(day8.MaxValue(mem))
}
