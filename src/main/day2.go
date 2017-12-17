package main

import (
	"./LineProcessor"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	filename := os.Args[1]
	processor := func(line string) int64 {
		nums := strings.Split(line, "\t")
		var min, max int64
		min = 1000000
		max = 0
		for _, numstr := range nums {
			num, _ := strconv.ParseInt(numstr, 10, 64)
			if num > max {
				max = num
			}
			if num < min {
				min = num
			}
		}
		return max - min
	}
	fmt.Println(LineProcessor.SumLines(filename, processor))
}
