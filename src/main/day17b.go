package main

import (
	"fmt"
)

func main() {
	step := 369
	size := 50000000
	pos := 0

	for i := 1; i <= size; i++ {
		pos = (pos+step)%i + 1
		if pos == 1 {
			fmt.Println(i)
		}
	}

}
