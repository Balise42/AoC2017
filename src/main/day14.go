package main

import (
	"fmt"
	"main/day10"
	"strconv"
)

func main() {
	input := "vbqugkhl"
	total := 0
	for i := 0; i < 128; i++ {
		hash := day10.ComputeHash(input + "-" + strconv.FormatInt(int64(i), 10))
		binhash := day10.ToBinary(hash)
		total = total + countOnes(binhash)
	}
	fmt.Println(total)
}

func countOnes(str string) int {
	res := 0
	for _, c := range str {
		if c == '1' {
			res = res + 1
		}
	}
	return res
}
