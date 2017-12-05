package main

import (
	"os"
	"log"
	"bufio"
	"fmt"
	"strconv"
)

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		log.Fatal("can't open file", err)
	}


	stat, err := file.Stat()
	if err != nil {
		log.Fatal("can't stat file", err)
	}

	scanner := bufio.NewScanner(file)
	jumps := make([]int, stat.Size())
	var i = 0
	for scanner.Scan() {

		jumpstr := scanner.Text()
		jump, err := strconv.Atoi(jumpstr)
		if err != nil {
			log.Fatal("not a number", jumpstr, err)
		}
		jumps[i] = jump
		i++
	}

	numelems := i
	fmt.Println(numelems)

	step := 0
	pos := 0
	for {
		if pos < 0 || pos >= numelems {
			break
		}
		step = step + 1
		tmp := jumps[pos]
		jumps[pos] = jumps[pos] + 1
		pos = pos + tmp
		//fmt.Println(jumps)
	}

	fmt.Println(step)
}