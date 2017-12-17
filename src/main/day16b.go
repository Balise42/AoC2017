package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"main/day16"
	"os"
	"strings"
)

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		log.Fatal("can't open file ", filename, ": ", err)
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal("can't read file ", filename, ": ", err)
	}

	dance := string(bytes)
	moves := strings.Split(strings.TrimSpace(dance), ",")

	line := "abcdefghijklmnop"

	seen := make(map[string]int)
	seen[line] = 0

	numit := -1
	for i := 1; numit == -1; i++ {
		for _, move := range moves {
			line = day16.ApplyMove(move, line)
		}
		if it, ok := seen[line]; ok {
			fmt.Println(line)
			fmt.Println(it, i)
			numit = (1000000000 - it) % i
		} else {
			seen[line] = i
		}
	}
	fmt.Println(numit)

	line = "abcdefghijklmnop"

	for i := 0; i < numit; i++ {
		for _, move := range moves {
			line = day16.ApplyMove(move, line)
		}
	}
	fmt.Println(line)
}
