package main

import (
	"./day10"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	l := 256
	input := "130,126,1,11,140,2,255,207,18,254,246,164,29,104,0,224"
	lengths := make([]byte, 0)
	for _, length := range strings.Split(input, ",") {
		a, _ := strconv.Atoi(length)
		lengths = append(lengths, byte(a))
	}

	list := day10.CircularList{make([]byte, l, l)}
	for i := 0; i < l; i++ {
		list.Set(i, byte(i))
	}
	list, _, _ = day10.ComputeOneRound(list, lengths, 0, 0)
	fmt.Println(int(list.Get(0)) * int(list.Get(1)))
}
