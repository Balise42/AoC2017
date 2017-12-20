package main

import (
	"bufio"
	"day20"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("can't open file", err)
	}

	accelerations := make([]day20.Vector, 0, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		_, _, a := day20.ParseLine(strings.TrimSpace(line))
		accelerations = append(accelerations, a)
	}

	fmt.Println(day20.MinVector(accelerations))
}
