package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"main/day9"
	"os"
)

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("can't open file ", err)
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal("can't read file ", err)
	}
	clean, garbage := day9.CleanData(data)
	score := day9.ComputeScore(clean)
	fmt.Println(score)
	fmt.Println(garbage)
}
