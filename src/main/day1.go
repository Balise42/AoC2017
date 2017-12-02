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
