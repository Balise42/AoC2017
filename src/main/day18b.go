package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

var q [][]int
var mux []sync.Mutex

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

	q = make([][]int, 2, 2)
	q[0] = make([]int, 0, 0)
	q[1] = make([]int, 0, 0)
	mux = make([]sync.Mutex, 2, 2)

	go run(program, 0)
	go run(program, 1)
	var input string
	fmt.Scanln(&input)
}

func run(program [][]string, prog int) {
	counter := 0
	registers := make(map[string]int)
	registers["p"] = prog
	count := 0
	for {
		if counter >= len(program) {
			break
		}
		instruction := program[counter]
		if instruction[0] == "snd" {
			val, err := strconv.Atoi(instruction[1])
			if err != nil {
				val = registers[instruction[1]]
			}
			if prog == 1 {
				count = count + 1
				fmt.Println(count)
			}
			mux[prog].Lock()
			q[prog] = append(q[prog], val)
			mux[prog].Unlock()
		} else if instruction[0] == "rcv" {
			if prog == 0 {
				for {
					mux[1].Lock()
					if len(q[1]) != 0 {
						mux[1].Unlock()
						break
					}
					mux[1].Unlock()
				}
				mux[1].Lock()
				registers[instruction[1]], q[1] = q[1][0], q[1][1:]
				mux[1].Unlock()
			} else {
				for {
					mux[0].Lock()
					if len(q[0]) != 0 {
						mux[0].Unlock()
						break
					}
					mux[0].Unlock()
				}
				mux[0].Lock()
				registers[instruction[1]], q[0] = q[0][0], q[0][1:]
				mux[0].Unlock()
			}
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
