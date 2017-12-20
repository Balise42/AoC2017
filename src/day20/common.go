package day20

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type Vector struct {
	X int
	Y int
	Z int
}

func ParseLine(line string) (Vector, Vector, Vector) {
	re := regexp.MustCompile("p=\\<(.+),(.+),(.+)>, v=\\<(.+),(.+),(.+)>, a=\\<(.+),(.+),(.+)>")
	matched := re.FindStringSubmatch(line)
	p := Vector{toInt(matched[1]), toInt(matched[2]), toInt(matched[3])}
	v := Vector{toInt(matched[4]), toInt(matched[5]), toInt(matched[6])}
	a := Vector{toInt(matched[7]), toInt(matched[8]), toInt(matched[9])}
	return p, v, a
}

func toInt(s string) int {
	i, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		log.Fatal("can't even atoi", err)
	}
	return i
}

func MinVector(vecs []Vector) int {
	min := -1
	val := -1
	for i, vec := range vecs {
		if length(vec) == min {
			fmt.Println("DUPES", min)
		}
		if val == -1 || length(vec) < min {
			min = length(vec)
			val = i
		}
	}
	fmt.Println(min, val)
	return val
}

func length(v Vector) int {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}
