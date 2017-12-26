package main

import (
	"os"
	"day25"
	"fmt"
)

func main() {
	machine := day25.CreateMachine(os.Args[1])
	for machine.NumIter > 0 {
		machine.Execute()
	}
	fmt.Println(machine.Checksum())
}