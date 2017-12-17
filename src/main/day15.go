package main

import (
	"fmt"
	"strconv"
)

func main() {
	gena := int64(16807)
	genb := int64(48271)
	mod := int64(2147483647)

	a := int64(289)
	b := int64(629)

	tot := 0

	for i := 0; i < 40000000; i++ {
		a = (gena * a) % mod
		b = (genb * b) % mod
		abin := strconv.FormatInt(a, 2)
		bbin := strconv.FormatInt(b, 2)
		diff := false
		for j := 1; j <= 16; j++ {
			if abin[len(abin)-j] != bbin[len(bbin)-j] {
				diff = true
				break
			}
		}
		if !diff {
			tot = tot + 1
		}
	}
	fmt.Println(tot)
}
