package day6

import "strconv"

func Hash(mem []int) string {
	hash := ""
	for _, m := range mem {
		hash = hash + strconv.FormatInt(int64(m), 10) + "#"
	}
	return hash
}

func Modify(mem []int) []int {
	maxi := MaxIndex(mem)
	toRedistribute := mem[maxi]
	mem[maxi] = 0
	for i := 1; i <= toRedistribute; i++ {
		mem[(maxi+i)%len(mem)] = mem[(maxi+i)%len(mem)] + 1
	}
	return mem
}

func MaxIndex(mem []int) int {
	max := -1
	index := 0
	for i, x := range mem {
		if x > max {
			index = i
			max = x
		}
	}
	return index
}
