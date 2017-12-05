package main

import (
	"./LineProcessor"
	"os"
	"strings"
	"fmt"
)

func main() {
	filename := os.Args[1]
	res := LineProcessor.SumLines(filename, func (line string) int64 {
		wordmap := make(map[string]bool)
		words := strings.Split(line, " ")
		for _, word := range words {
			_, cont := wordmap[word]
			if cont {
				return 0
			}
			wordmap[word] = true
		}
		return 1
	})
	fmt.Println(res)
}