package main

import (
	"LineProcessor"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	filename := os.Args[1]
	processor := func(line string) int64 {
		strs := strings.Split(line, "\t")
		nums := make([]int64, len(strs))
		for i, numstr := range strs {
			nums[i], _ = strconv.ParseInt(numstr, 10, 64)
		}

		for i, n1 := range nums {
			for _, n2 := range nums[i+1:] {
				if n1%n2 == 0 {
					return n1 / n2
				}
				if n2%n1 == 0 {
					return n2 / n1
				}
			}
		}
		return -1
	}
	fmt.Println(LineProcessor.SumLines(filename, processor))
}
