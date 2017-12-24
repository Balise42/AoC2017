package day24

import (
	"day20"
	"strings"
	"strconv"
	"log"
)

func GetLongest(graph map[day20.Pair]bool, v int) (int, int) {
	candidates := GetCandidates(graph, v)
	weight := 0
	length := 0

	for _, candidate := range candidates {
		graphCopy := Copy(graph)
		graphCopy[candidate] = true
		var newV int
		if candidate.A == v {
			newV = candidate.B
		} else {
			newV = candidate.A
		}

		newL, newW := GetLongest(graphCopy, newV)
		newL = newL + 1
		newW = newW + candidate.A + candidate.B

		if newL > length || (newL == length && newW > weight) {
			weight = newW
			length = newL
		}
	}
	return length, weight
}

func GetLargestWeight(graph map[day20.Pair]bool, v int) int {
	candidates := GetCandidates(graph, v)
	weight := 0
	for _, candidate := range candidates {
		graphCopy := Copy(graph)
		graphCopy[candidate] = true
		var newV int
		if candidate.A == v {
			newV = candidate.B
		} else {
			newV = candidate.A
		}
		newW := GetLargestWeight(graphCopy, newV) + candidate.A + candidate.B

		if newW > weight {
			weight = newW
		}
	}
	return weight
}

func GetCandidates(graph map[day20.Pair]bool, i int) []day20.Pair{
	candidates := make([]day20.Pair, 0, 0)
	for k, v := range graph {
		if !v && (k.A == i || k.B == i) {
			candidates = append(candidates, k)
		}
	}
	return candidates
}

func Copy(m map[day20.Pair]bool) map[day20.Pair]bool {
	mc := make(map[day20.Pair]bool)
	for k, v := range m {
		mc[k] = v
	}
	return mc
}

func AddToGraph(graph map[day20.Pair]bool, line string) {
	toks := strings.Split(line, "/")
	v1, err := strconv.Atoi(toks[0])
	if err != nil {
		log.Fatal("can't convert", err)
	}
	v2, err := strconv.Atoi(toks[1])
	if err != nil {
		log.Fatal("can't convert", err)
	}
	graph[day20.Pair{v1, v2}] = false
}
