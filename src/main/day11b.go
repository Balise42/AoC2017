package main

import (
	"day11"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	filename := os.Args[1]
	file, _ := os.Open(filename)
	bytes, _ := ioutil.ReadAll(file)
	steps := strings.Split(strings.TrimSpace(string(bytes)), ",")
	tally := make(map[string]int)

	max := 0
	for _, step := range steps {
		if val, ok := tally[step]; ok {
			tally[step] = val + 1
		} else {
			tally[step] = 1
		}
		day11.CancelCompletely(tally, "nw", "se")
		day11.CancelCompletely(tally, "ne", "sw")
		day11.CancelToNS(tally, "ne", "nw", "n")
		day11.CancelToNS(tally, "se", "sw", "s")
		day11.CancelCompletely(tally, "n", "s")

		sum := 0
		for _, v := range tally {
			sum = sum + v
		}
		if sum > max {
			max = sum
		}

	}
	fmt.Println(max)
}
