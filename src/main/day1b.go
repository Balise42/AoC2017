package main

import (
	"os"
	"fmt"
)

func main() {
	input := os.Args[1]
	res := 0
	for i := 0; i<len(input); i++ {
		res = res + getLocal(input[i], input[(i+len(input)/2)%len(input)])
	}
	fmt.Println(res)
}
