package day16

import (
	"log"
	"strconv"
	"strings"
)

func ApplyMove(move string, line string) string {
	if move[0] == 'p' {
		return swap(line, string(move[1]), string(move[3]))
	} else if move[0] == 's' {
		rot, err := strconv.Atoi(move[1:])
		if err != nil {
			log.Fatal("s: can't convert int "+move[1:]+": ", err)
		}
		return rotate(line, rot)
	} else {
		toks := strings.Split(move, "/")
		i1, err := strconv.Atoi(toks[0][1:])
		if err != nil {
			log.Fatal("x: can't convert int "+toks[0][1:]+": ", err)
		}
		i2, err := strconv.Atoi(toks[1])
		if err != nil {
			log.Fatal("x: can't convert int "+toks[1]+": ", err)
		}
		return swap(line, string(line[i1]), string(line[i2]))
	}
}

func swap(line string, str1 string, str2 string) string {
	line = strings.Replace(line, str1, "W", -1)
	line = strings.Replace(line, str2, str1, -1)
	line = strings.Replace(line, "W", str2, -1)
	return line
}

func rotate(line string, rot int) string {
	for ; rot > 0; rot-- {
		line = string(line[len(line)-1]) + line[:len(line)-1]
	}
	return line
}
