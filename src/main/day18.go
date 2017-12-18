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
	if err != nil {
		log.Fatal("can't open file ", file, ": ", err)
	}

	program := make([][]string, 0, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		program = append(program, strings.Split(strings.TrimSpace(scanner.Text()), " "))
	}

	rcvd := 0
	counter := 0
	registers := make(map[string]int)

	for {
		if counter > len(program) {
			break
		}
		instruction := program[counter]
		if instruction[0] == "snd" {
			rcvd = registers[instruction[1]]
		} else if instruction[0] == "set" {
			val, err := strconv.Atoi(instruction[2])
			if err != nil {
				registers[instruction[1]] = registers[instruction[2]]
			} else {
				registers[instruction[1]] = val
			}
		} else if instruction[0] == "add" {
			init, ok := registers[instruction[1]]
			if !ok {
				init = 0
			}
			val, err := strconv.Atoi(instruction[2])
			if err != nil {
				registers[instruction[1]] = init + registers[instruction[2]]
			} else {
				registers[instruction[1]] = init + val
			}
		} else if instruction[0] == "mul" {
			init, ok := registers[instruction[1]]
			if !ok {
				init = 0
			}
			val, err := strconv.Atoi(instruction[2])
			if err != nil {
				registers[instruction[1]] = init * registers[instruction[2]]
			} else {
				registers[instruction[1]] = init * val
			}
		} else if instruction[0] == "mod" {
			init, ok := registers[instruction[1]]
			if !ok {
				init = 0
			}
			val, err := strconv.Atoi(instruction[2])
			if err != nil {
				registers[instruction[1]] = init % registers[instruction[2]]
			} else {
				registers[instruction[1]] = init % val
			}
		} else if instruction[0] == "rcv" {
			if rcvd != 0 {
				fmt.Println(rcvd)
				break
			}
		} else {
			val, err := strconv.Atoi(instruction[1])
			if err != nil {
				val = registers[instruction[1]]
			}
			if val > 0 {
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
}
