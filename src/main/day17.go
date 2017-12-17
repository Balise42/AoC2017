package main

import (
	"fmt"
)

func main() {
	step := 369

	buffer := make([]int, 1, 2017)
	buffer[0] = 0
	pos := 0

	for i := 1; i <= 2017; i++ {
		pos = (pos+step)%i + 1
		buffer = append(buffer[:pos], append([]int{i}, buffer[pos:]...)...)
	}

	for i, val := range buffer {
		if val == 2017 {
			fmt.Println(buffer[i+1])
			break
		}
	}
}
