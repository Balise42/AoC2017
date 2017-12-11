package main

import (
	"main/day10"
	"fmt"
)

func main() {
	input := "130,126,1,11,140,2,255,207,18,254,246,164,29,104,0,224"
	lengths := make([]byte, 0)
	for _, c := range input {
		lengths = append(lengths, byte(c))
	}
	lengths = append(lengths,17, 31, 73, 47, 23)

	skip := 0
	pos := 0
	l := 256
	list := day10.CircularList{ List: make([]byte, l, l)}
	for i := 0; i<l; i++ {
		list.Set(i, byte(i))
	}

	for i := 0; i<64; i++ {
		list, skip, pos = day10.ComputeOneRound(list, lengths, skip, pos)
	}

	hashval := make([]byte, 16, 16)

	for i := 0; i<16; i++ {
		hashval[i] = list.Get(i*16)
		for j := 1; j<16; j++ {
			hashval[i] = hashval[i] ^ list.Get(i*16+j)
		}
	}

	for i := 0; i<16; i++ {
		fmt.Printf("%02x", hashval[i])
	}
	fmt.Print("\n")
}