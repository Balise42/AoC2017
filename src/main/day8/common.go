package day8

import (
	"os"
	"log"
	"bufio"
	"strings"
	"strconv"
)

type Instruction struct {
	Variable string
	Operator string
	Operand int
	Conditionvariable string
	Conditionoperator string
	Conditionoperand int
}

func NewInstruction(variable string, operator string, operand string, condvariable string, conditionoperator string, conditionoperand string) Instruction {
	operandi, err := strconv.Atoi(operand)
	if err != nil {
		log.Fatal("Couldn't convert to int", err)
	}
	conditionoperandi, err := strconv.Atoi(conditionoperand)
	if err != nil {
		log.Fatal("Couldn't convert to int", err)
	}
	return Instruction{variable, operator, operandi, condvariable, conditionoperator, conditionoperandi}
}

func ParseFileToInstructions(filename string) []Instruction {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		log.Fatal("Couldn't open file", err)
	}

	result := make([]Instruction, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, " ")
		instruction := NewInstruction(tokens[0], tokens[1], tokens[2], tokens[4], tokens[5], tokens[6])
		result = append(result, instruction)
	}

	return result
}

func ProcessInstructions(prog []Instruction) (map[string]int, int){
	max := -99999999
	state := make(map[string]int)
	for _, instruction := range prog {
		val := compute(instruction, state)
		state[instruction.Variable] = val
		if val > max {
			max = val
		}
	}
	return state, max
}

func compute(instruction Instruction, state map[string]int) int {
	val, ok := state[instruction.Variable]
	if !ok {
		val = 0
	}

	if condition(instruction, state) {
		if instruction.Operator == "inc" {
			return val + instruction.Operand
		} else {
			return val - instruction.Operand
		}
	}

	return val
}

func condition(instruction Instruction, state map[string]int) bool {
	val, ok := state[instruction.Conditionvariable]
	if !ok {
		val = 0
	}
	if instruction.Conditionoperator == "<" {
		return val < instruction.Conditionoperand
	}
	if instruction.Conditionoperator == ">" {
		return val > instruction.Conditionoperand
	}
	if instruction.Conditionoperator == "==" {
		return val == instruction.Conditionoperand
	}
	if instruction.Conditionoperator == "<=" {
		return val <= instruction.Conditionoperand
	}
	if instruction.Conditionoperator == ">=" {
		return val >= instruction.Conditionoperand
	}
	if instruction.Conditionoperator == "!=" {
		return val != instruction.Conditionoperand
	}
	return false
}

func MaxValue(mem map[string]int) int {
	max := -99999999
	for _, v := range mem {
		if v > max {
			max = v
		}
	}
	return max
}