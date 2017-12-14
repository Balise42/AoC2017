package main

import (
	"main/day10"
	"fmt"
)

func main() {
	input := "130,126,1,11,140,2,255,207,18,254,246,164,29,104,0,224"
	hash := day10.ComputeHash(input)
	fmt.Println(hash)
}