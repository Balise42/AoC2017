package main

import (
	"day6"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		log.Fatal("can't open file", err)
	}

	b, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatal("couldn't read file", err)
	}

	input := string(b)
	mems := strings.Split(strings.TrimSpace(input), "\t")

	mem := make([]int, len(mems))
	for i, m := range mems {
		mem[i], err = strconv.Atoi(m)
		if err != nil {
			log.Fatal("couldn't convert number", m, err)
		}
	}

	seen := make(map[string]int)

	for i := 0; ; i++ {
		x, ok := whereSeen(mem, seen)
		if ok {
			fmt.Println(i - x)
			break
		}
		addIndex(mem, seen, i)
		mem = day6.Modify(mem)
	}
}

func addIndex(mem []int, seen map[string]int, iteration int) {
	seen[day6.Hash(mem)] = iteration
}

func whereSeen(mem []int, seen map[string]int) (int, bool) {
	x, ok := seen[day6.Hash(mem)]
	return x, ok
}
