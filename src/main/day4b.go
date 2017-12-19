package main

import (
	"LineProcessor"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	filename := os.Args[1]
	res := LineProcessor.SumLines(filename, func(line string) int64 {
		wordmap := make(map[string]bool)
		words := strings.Split(line, " ")
		for _, word := range words {
			sortedWord := []byte(word)
			sort.Slice(sortedWord, func(i, j int) bool {
				return sortedWord[i] < sortedWord[j]
			})
			_, cont := wordmap[string(sortedWord)]
			if cont {
				return 0
			}
			wordmap[string(sortedWord)] = true
		}
		return 1
	})
	fmt.Println(res)
}
