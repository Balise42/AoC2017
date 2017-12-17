package main

import (
	"./day16"
	"fmt"
	"io/ioutil"
	"log"
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

	for _, move := range moves {
		line = day16.ApplyMove(move, line)
	}
	fmt.Println(line)
}
