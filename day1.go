package main

import (
	"os"
	"fmt"
)

func main() {
	input := os.Args[1]
	res := 0
	for i := 0; i<len(input) - 1; i++ {
		res = res + getLocal(input[i], input[i+1])
	}
	res = res + getLocal(input[0], input[len(input) - 1])
	fmt.Println(res)
}

func getLocal(c1 byte, c2 byte) int {
	if c1 != c2 {
		return 0
	}
	return (int) (c1 - 48)
}