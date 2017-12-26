package day25

import (
	"log"
	"os"
	"bufio"
	"regexp"
	"strings"
	"strconv"
)

type State struct {
	Towrite []int
	Dir []int
	Next []string
}

type Machine struct {
	Tape        map[int]int
	Cursor      int
	Transitions map[string]State
	CurrState       string
	NumIter int
}

func NewMachine() Machine {
	machine := Machine {make(map[int]int), 0, make(map[string]State), "A", 0}
	return machine
}

func NewState() State {
	state := State {make([]int, 2, 2), make([]int, 2, 2), make([]string, 2, 2)}
	return state
}

func CreateMachine(filename string) Machine {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		log.Fatal("can't open file", err)
	}

	machine := NewMachine()

	scanner := bufio.NewScanner(file)

	machine.CurrState = getStartState(scanner)

	machine.NumIter = getNumIter(scanner)

	for scanner.Scan() {
		stateName := getStateName(scanner)
		state := NewState()
		for i := 0; i<2; i++ {
			stateIdx := getStateIdx(scanner)
			stateWrite := getWrite(scanner)
			stateMove := getMove(scanner)
			stateNext := getNext(scanner)
			state.Towrite[stateIdx] = stateWrite
			state.Dir[stateIdx] = stateMove
			state.Next[stateIdx] = stateNext
		}
		machine.Transitions[stateName] = state
	}

	return machine
}

func (m Machine) Checksum() int {
	tot := 0
	for _, v := range m.Tape {
		tot = tot + v
	}
	return tot
}

func (m * Machine) Execute() {
	if m.NumIter <= 0 {
		return
	}
	val, ok := m.Tape[m.Cursor]
	if !ok {
		val = 0
	}
	m.Tape[m.Cursor] = m.Transitions[m.CurrState].Towrite[val]
	m.Cursor = m.Cursor + m.Transitions[m.CurrState].Dir[val]
	m.CurrState = m.Transitions[m.CurrState].Next[val]
	m.NumIter = m.NumIter - 1
}

func getWrite(scanner *bufio.Scanner) int {
	return getIntFromRegex(scanner,"- Write the value (.)\\.")
}

func getMove(scanner *bufio.Scanner) int {
	dir := getStringFromRegex(scanner, "- Move one slot to the (.+)\\.")
	if dir == "right" {
		return 1
	} else {
		return -1
	}
}

func getNext(scanner *bufio.Scanner) string {
	return getStringFromRegex(scanner, "- Continue with state (.)\\.")
}

func getLine(scanner *bufio.Scanner) string {
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func getIntFromRegex(scanner *bufio.Scanner, reg string) int {
	i, err := strconv.Atoi(getStringFromRegex(scanner, reg))
	if err != nil {
		log.Fatal("can't atoi", err)
	}
	return i
}

func getStringFromRegex(scanner *bufio.Scanner, reg string) string {
	line := getLine(scanner)
	re, err := regexp.Compile(reg)
	if err != nil {
		log.Fatal("can't regex", err)
	}
	res := re.FindStringSubmatch(line)
	return res[1]
}

func getStateIdx(scanner *bufio.Scanner) int {
	return getIntFromRegex(scanner, "If the current value is (.):")
}

func getStartState(scanner *bufio.Scanner) string {
	return getStringFromRegex(scanner,"Begin in state (.)\\.")
}

func getStateName(scanner *bufio.Scanner) string {
	return getStringFromRegex(scanner,"In state (.):")
}

func getNumIter(scanner *bufio.Scanner) int {
	return getIntFromRegex(scanner,"Perform a diagnostic checksum after (\\d+) steps\\.")
}