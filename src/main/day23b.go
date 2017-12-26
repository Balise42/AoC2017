package main

import "fmt"

func main() {
	b := 108100
	c := 125100
	h := 0

	for {
		f := 1
		d := 2
		g := 2

		for g != 0 {
			e := 2

			for g != 0 {
				g = d
				g = g * e
				g = g - b

				if g != 0 {
					f = 0
				}

				e = e + 1
				g = e
				g = g - b
			}

			d = d + 1
			g = d
			g = g - b
		}

		if f != 0 {
			h = h + 1
		}
		g = b
		g = g - c
		if g == 0 {
			break
		}
		b = b + 17
	}
	fmt.Println(h)
}