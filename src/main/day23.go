package main

import (
	"os"
	"log"
	"bufio"
	"strings"
	"strconv"
	"fmt"
)

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		log.Fatal("can't open file", err)
	}

	program := make([][]string, 0, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		program = append(program, strings.Split(strings.TrimSpace(scanner.Text()), " "))
	}

	registers := make(map[string]int)
	counter := 0
	nummul := 0

	for {
		if counter >= len(program) {
			break
		}
		instruction := program[counter]

		if instruction[0] == "set" {
			val, err := strconv.Atoi(instruction[2])
			if err != nil {
				registers[instruction[1]] = registers[instruction[2]]
			} else {
				registers[instruction[1]] = val
			}
		} else if instruction[0] == "mul" {
			nummul = nummul + 1
			val, err := strconv.Atoi(instruction[2])
			if err != nil {
				registers[instruction[1]] = registers[instruction[1]] * registers[instruction[2]]
			} else {
				registers[instruction[1]] = registers[instruction[1]] * val
			}
		} else if instruction[0] == "sub" {
			val, err := strconv.Atoi(instruction[2])
			if err != nil {
				registers[instruction[1]] = registers[instruction[1]] - registers[instruction[2]]
			} else {
				registers[instruction[1]] = registers[instruction[1]] - val
			}
		} else if instruction[0] == "jnz" {
			val, err := strconv.Atoi(instruction[1])
			if err != nil {
				val = registers[instruction[1]]
			}
			if val != 0 {
				skip, err := strconv.Atoi(instruction[2])
				if err != nil {
					skip = registers[instruction[2]]
				}
				counter = counter + skip
				continue
			}
		}
		counter = counter + 1
	}

	fmt.Println(nummul)
}