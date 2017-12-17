package main

import (
	"bufio"
	"fmt"
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
		log.Fatal("can't open file")
	}

	scanner := bufio.NewScanner(file)
	firewall := make(map[int]int)
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(strings.TrimSpace(line), ": ")
		layer, _ := strconv.Atoi(tokens[0])
		level, _ := strconv.Atoi(tokens[1])
		firewall[layer] = level
	}

	severity := 0
	for layer, level := range firewall {
		if layer%((level-1)*2) == 0 {
			severity = severity + layer*level
		}
	}

	fmt.Println(severity)
}
