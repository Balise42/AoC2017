package main

import (
	"os"
	"log"
	"io/ioutil"
	"strings"
	"strconv"
	"fmt"
	"main/day6"
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

	seen := make(map[string]bool)

	for i := 0; ; i++ {
		if hasSeen(mem, seen) {
			fmt.Println(i)
			break
		}
		add(mem, seen)
		mem = day6.Modify(mem)
	}
}

func hasSeen(mem []int, seen map[string]bool) bool {
	_, ok := seen[day6.Hash(mem)]
	return ok
}


func add(mem []int, seen map[string]bool) {
	seen[day6.Hash(mem)] = true
}
